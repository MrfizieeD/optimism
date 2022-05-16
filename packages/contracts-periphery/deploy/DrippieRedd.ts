/* Imports: External */
import { DeployFunction } from 'hardhat-deploy/dist/types'

/* Imports: Internal */
import { getDeployConfig } from '../src/deploy-config'

const deployFn: DeployFunction = async (hre) => {
  const { deterministic } = hre.deployments
  const { deployer } = await hre.getNamedAccounts()

  const deployConfig = getDeployConfig(hre.network.name)

  const { deploy } = await deterministic('DrippieRedd', {
    salt: hre.ethers.constants.HashZero,
    from: deployer,
    args: [deployConfig.initialDrippieReddOwner],
    log: true,
  })

  await deploy()
}

deployFn.tags = ['DrippieRedd']

export default deployFn
