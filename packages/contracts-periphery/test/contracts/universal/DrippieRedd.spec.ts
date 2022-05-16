import hre from 'hardhat'
import { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers'
import { Contract } from 'ethers'
import _ from 'lodash'
import { toRpcHexString } from '@eth-optimism/core-utils'

import { expect } from '../../setup'
import { deploy } from '../../helpers'

describe('DrippieRedd', () => {
  const DEFAULT_DRIP_NAME = 'drippity drip drip'
  const DEFAULT_DRIP_CONFIG = {
    recipient: '0x' + '11'.repeat(20),
    amount: hre.ethers.BigNumber.from(1),
    checkscript: hre.ethers.utils.hexConcat([
      '0x60', // PUSH1
      '0x01', // 0x01
    ]),
    interval: hre.ethers.BigNumber.from(100),
    data: '0x',
  }

  let signer1: SignerWithAddress
  let signer2: SignerWithAddress
  before('signer setup', async () => {
    ;[signer1, signer2] = await hre.ethers.getSigners()
  })

  let SimpleStorage: Contract
  let DrippieRedd: Contract
  beforeEach('deploy contracts', async () => {
    SimpleStorage = await deploy('SimpleStorage')
    DrippieRedd = await deploy('DrippieRedd', {
      signer: signer1,
      args: [signer1.address],
    })
  })

  beforeEach('balance setup', async () => {
    await hre.ethers.provider.send('hardhat_setBalance', [
      DrippieRedd.address,
      toRpcHexString(DEFAULT_DRIP_CONFIG.amount.mul(100000)),
    ])
    await hre.ethers.provider.send('hardhat_setBalance', [
      DEFAULT_DRIP_CONFIG.recipient,
      '0x0',
    ])
  })

  describe('constructor', () => {
    it('should set the owner', async () => {
      expect(await DrippieRedd.owner()).to.equal(signer1.address)
    })
  })

  describe('create', () => {
    describe('when called by the owner', () => {
      it('should create a drip with the given name', async () => {
        await expect(
          DrippieRedd.create(DEFAULT_DRIP_NAME, DEFAULT_DRIP_CONFIG)
        ).to.emit(DrippieRedd, 'DripCreated')

        const drip = await DrippieRedd.drips(DEFAULT_DRIP_NAME)
        expect(drip.status).to.equal(1)
        expect(drip.last).to.deep.equal(hre.ethers.BigNumber.from(0))
        expect(_.toPlainObject(drip.config)).to.deep.include(
          DEFAULT_DRIP_CONFIG
        )
      })

      it('should not be able to create the same drip twice', async () => {
        await DrippieRedd.create(DEFAULT_DRIP_NAME, DEFAULT_DRIP_CONFIG)

        await expect(
          DrippieRedd.create(DEFAULT_DRIP_NAME, DEFAULT_DRIP_CONFIG)
        ).to.be.revertedWith('DrippieRedd: drip with that name already exists')
      })
    })

    describe('when called by not the owner', () => {
      it('should revert', async () => {
        await expect(
          DrippieRedd.connect(signer2).create(
            DEFAULT_DRIP_NAME,
            DEFAULT_DRIP_CONFIG
          )
        ).to.be.revertedWith('UNAUTHORIZED')
      })
    })
  })

  describe('update', () => {
    describe('when called by the owner', () => {
      it('should update the config for the given drip', async () => {
        await DrippieRedd.create(DEFAULT_DRIP_NAME, DEFAULT_DRIP_CONFIG)
        await expect(
          DrippieRedd.update(DEFAULT_DRIP_NAME, {
            ...DEFAULT_DRIP_CONFIG,
            recipient: '0x' + '22'.repeat(20),
          })
        ).to.emit(DrippieRedd, 'DripConfigUpdated')

        const drip = await DrippieRedd.drips(DEFAULT_DRIP_NAME)
        expect(drip.status).to.equal(1)
        expect(drip.last).to.deep.equal(hre.ethers.BigNumber.from(0))
        expect(_.toPlainObject(drip.config)).to.deep.include({
          ...DEFAULT_DRIP_CONFIG,
          recipient: '0x' + '22'.repeat(20),
        })
      })

      it('should revert if the drip does not exist yet', async () => {
        await expect(
          DrippieRedd.update(DEFAULT_DRIP_NAME, DEFAULT_DRIP_CONFIG)
        ).to.be.revertedWith('DrippieRedd: drip with that name does not exist')
      })
    })

    describe('when called by not the owner', () => {
      it('should revert', async () => {
        await expect(
          DrippieRedd.connect(signer2).update(
            DEFAULT_DRIP_NAME,
            DEFAULT_DRIP_CONFIG
          )
        ).to.be.revertedWith('UNAUTHORIZED')
      })
    })
  })

  describe('toggle', () => {
    describe('when called by the owner', () => {
      it('should toggle the status of a drip', async () => {
        // On by default
        await DrippieRedd.create(DEFAULT_DRIP_NAME, DEFAULT_DRIP_CONFIG)
        expect((await DrippieRedd.drips(DEFAULT_DRIP_NAME)).status).to.equal(1)

        // Toggle off
        await expect(DrippieRedd.toggle(DEFAULT_DRIP_NAME)).to.emit(
          DrippieRedd,
          'DripStatusUpdated'
        )
        expect((await DrippieRedd.drips(DEFAULT_DRIP_NAME)).status).to.equal(2)

        // Toggle on
        await expect(DrippieRedd.toggle(DEFAULT_DRIP_NAME)).to.emit(
          DrippieRedd,
          'DripStatusUpdated'
        )
        expect((await DrippieRedd.drips(DEFAULT_DRIP_NAME)).status).to.equal(1)
      })

      it('should revert if the drip does not exist yet', async () => {
        await expect(DrippieRedd.toggle(DEFAULT_DRIP_NAME)).to.be.revertedWith(
          'DrippieRedd: drip with that name does not exist'
        )
      })
    })

    describe('when called by not the owner', () => {
      it('should revert', async () => {
        await expect(
          DrippieRedd.connect(signer2).toggle(DEFAULT_DRIP_NAME)
        ).to.be.revertedWith('UNAUTHORIZED')
      })
    })
  })

  describe('drip', () => {
    it('should drip the amount', async () => {
      await DrippieRedd.create(DEFAULT_DRIP_NAME, DEFAULT_DRIP_CONFIG)

      await expect(DrippieRedd.drip(DEFAULT_DRIP_NAME)).to.emit(
        DrippieRedd,
        'DripExecuted'
      )

      expect(
        await signer1.provider.getBalance(DEFAULT_DRIP_CONFIG.recipient)
      ).to.equal(DEFAULT_DRIP_CONFIG.amount)
    })

    it('should trigger a function if data is included', async () => {
      await DrippieRedd.create(DEFAULT_DRIP_NAME, {
        ...DEFAULT_DRIP_CONFIG,
        recipient: SimpleStorage.address,
        data: SimpleStorage.interface.encodeFunctionData('setValue', [
          '0x' + '33'.repeat(32),
        ]),
      })

      await DrippieRedd.drip(DEFAULT_DRIP_NAME)
      expect(await SimpleStorage.getValue()).to.equal('0x' + '33'.repeat(32))
    })

    it('should revert if dripping twice in one interval', async () => {
      await DrippieRedd.create(DEFAULT_DRIP_NAME, DEFAULT_DRIP_CONFIG)
      await DrippieRedd.drip(DEFAULT_DRIP_NAME)

      await expect(DrippieRedd.drip(DEFAULT_DRIP_NAME)).to.be.revertedWith(
        'DrippieRedd: drip interval has not elapsed'
      )

      await hre.ethers.provider.send('evm_increaseTime', [
        DEFAULT_DRIP_CONFIG.interval.add(1).toHexString(),
      ])

      await expect(DrippieRedd.drip(DEFAULT_DRIP_NAME)).to.not.be.reverted
    })

    it('should revert when the drip does not exist', async () => {
      await expect(DrippieRedd.drip(DEFAULT_DRIP_NAME)).to.be.revertedWith(
        'DrippieRedd: selected drip does not exist or is not currently active'
      )
    })

    it('should revert when the drip is not active', async () => {
      await DrippieRedd.create(DEFAULT_DRIP_NAME, DEFAULT_DRIP_CONFIG)
      await DrippieRedd.toggle(DEFAULT_DRIP_NAME)

      await expect(DrippieRedd.drip(DEFAULT_DRIP_NAME)).to.be.revertedWith(
        'DrippieRedd: selected drip does not exist or is not currently active'
      )
    })

    describe('checkscripts', () => {
      describe('ETH balance threshold', () => {
        it('should succeed if balance is below threshold', async () => {
          const threshold = hre.ethers.BigNumber.from(10)
          await hre.ethers.provider.send('hardhat_setBalance', [
            DEFAULT_DRIP_CONFIG.recipient,
            // Just below the threshold
            toRpcHexString(threshold.sub(1)),
          ])

          await DrippieRedd.create(DEFAULT_DRIP_NAME, {
            ...DEFAULT_DRIP_CONFIG,
            checkscript: hre.ethers.utils.hexConcat([
              '0x00',
              hre.ethers.utils.hexZeroPad(threshold.toHexString(), 32),
            ]),
          })

          await expect(DrippieRedd.drip(DEFAULT_DRIP_NAME)).to.not.be.reverted
        })

        it('should revert if balance is above threshold', async () => {
          const threshold = hre.ethers.BigNumber.from(10)
          await hre.ethers.provider.send('hardhat_setBalance', [
            DEFAULT_DRIP_CONFIG.recipient,
            // Just above the threshold
            toRpcHexString(threshold.add(1)),
          ])

          await DrippieRedd.create(DEFAULT_DRIP_NAME, {
            ...DEFAULT_DRIP_CONFIG,
            checkscript: hre.ethers.utils.hexConcat([
              '0x00',
              hre.ethers.utils.hexZeroPad(threshold.toHexString(), 32),
            ]),
          })

          await expect(DrippieRedd.drip(DEFAULT_DRIP_NAME)).to.be.revertedWith(
            'DrippieRedd: checkscript failed so drip is not yet ready to be triggered'
          )
        })
      })

      describe('contract call checkscript', () => {
        it('should be able to do a checkscript based on an external contract call', async () => {
          const desired = '0x' + '33'.repeat(32)

          // Checkscript that looks for a storage value to be a specific thing
          await DrippieRedd.create(DEFAULT_DRIP_NAME, {
            ...DEFAULT_DRIP_CONFIG,
            checkscript: hre.ethers.utils.hexConcat([
              '0x63', // PUSH4
              SimpleStorage.interface.getSighash('getValue'),
              '0x60', // PUSH1
              '0xE0', // 224
              '0x1B', // SHL
              '0x60', // PUSH1
              '0x00', // 0
              '0x52', // MSTORE
              '0x60', // PUSH1
              '0x20', // 32
              '0x60', // PUSH1
              '0x00', // 0
              '0x60', // PUSH1
              '0x20', // 32
              '0x60', // PUSH1
              '0x00', // 0
              '0x73', // 32
              SimpleStorage.address,
              '0x5A', // GAS
              '0xFA', // STATICCALL
              '0x60', // PUSH1
              '0x00', // 0
              '0x51', // MLOAD
              '0x7F', // PUSH32
              desired,
              '0x14', // EQ
            ]),
          })

          // Value is not yet set to desired value, should fail.
          await expect(DrippieRedd.drip(DEFAULT_DRIP_NAME)).to.be.revertedWith(
            'checkscript failed so drip is not yet ready to be triggered'
          )

          // Update value to desired value.
          await SimpleStorage.setValue(desired)

          // Value is now set to desired value, should succeed.
          await expect(DrippieRedd.drip(DEFAULT_DRIP_NAME)).to.not.be.reverted
        })
      })
    })
  })
})
