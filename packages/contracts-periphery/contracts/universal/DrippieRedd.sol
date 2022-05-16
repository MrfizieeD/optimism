// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { RetroReceiver } from "./RetroReceiver.sol";

/**
 * @title DrippieRedd
 * @notice DrippieRedd goes brrr.
 */
contract DrippieRedd is RetroReceiver {
    /**
     * Enum representing different status options for a given drip.
     */
    enum DripStatus {
        NONE,
        ACTIVE,
        PAUSED
    }

    /**
     * Represents the configuration for a given drip.
     */
    struct DripConfig {
        address payable recipient;
        bytes data;
        bytes checkscript;
        uint256 amount;
        uint256 interval;
    }

    /**
     * Represents the state of an active drip.
     */
    struct DripState {
        DripStatus status;
        DripConfig config;
        uint256 last;
    }

    /**
     * Emitted when a new drip is created.
     */
    event DripCreated(string indexed name, DripConfig config);

    /**
     * Emitted when a drip config is updated.
     */
    event DripConfigUpdated(string indexed name, DripConfig config);

    /**
     * Emitted when a drip status is updated.
     */
    event DripStatusUpdated(string indexed name, DripStatus status);

    /**
     * Emitted when a drip is executed.
     */
    event DripExecuted(string indexed name, address indexed executor, uint256 timestamp);

    /**
     * Maps from drip names to drip states.
     */
    mapping(string => DripState) public drips;

    /**
     * @param _owner Initial owner address.
     */
    constructor(address _owner) RetroReceiver(_owner) {}

    /**
     * Creates a new drip with the given name and configuration.
     *
     * @param _name Name of the drip.
     * @param _config Configuration for the drip.
     */
    function create(string memory _name, DripConfig memory _config) public onlyOwner {
        require(
            drips[_name].status == DripStatus.NONE,
            "DrippieRedd: drip with that name already exists"
        );

        // Note that "last" timestamp is being set to zero, which means the first drip can be
        // executed immediately after the drip is created.
        drips[_name] = DripState({ status: DripStatus.ACTIVE, config: _config, last: 0 });

        emit DripCreated(_name, _config);
    }

    /**
     * Configures a drip by name.
     *
     * @param _name Name of the drip to configure.
     * @param _config Drip configuration.
     */
    function update(string memory _name, DripConfig memory _config) public onlyOwner {
        require(
            drips[_name].status != DripStatus.NONE,
            "DrippieRedd: drip with that name does not exist"
        );

        drips[_name].config = _config;

        emit DripConfigUpdated(_name, _config);
    }

    /**
     * Toggles the status of a given drip.
     *
     * @param _name Name of the drip to toggle.
     */
    function toggle(string memory _name) public onlyOwner {
        require(
            drips[_name].status != DripStatus.NONE,
            "DrippieRedd: drip with that name does not exist"
        );

        if (drips[_name].status == DripStatus.ACTIVE) {
            drips[_name].status = DripStatus.PAUSED;
        } else {
            drips[_name].status = DripStatus.ACTIVE;
        }

        emit DripStatusUpdated(_name, drips[_name].status);
    }

    /**
     * Triggers a drip.
     *
     * @param _name Name of the drip to trigger.
     */
    function drip(string memory _name) public {
        DripState storage state = drips[_name];

        require(
            state.status == DripStatus.ACTIVE,
            "DrippieRedd: selected drip does not exist or is not currently active"
        );

        // Don't drip if the drip interval has not yet elapsed since the last time we dripped. This
        // is a safety measure that prevents a malicious recipient from, e.g., spending all of
        // their funds and repeatedly requesting new drips. Limits the potential impact of a
        // compromised recipient to just a single drip interval, after which the drip can be paused
        // by the owner address.
        require(
            state.last + state.config.interval <= block.timestamp,
            "DrippieRedd: drip interval has not elapsed since last drip"
        );

        // Checkscript is a special system for allowing drips to execute arbitrary EVM bytecode to
        // determine whether or not to execute the drip. A checkscript is a simply EVM bytecode
        // snippet that operates with the following requirements:
        // 1. Stack is initialized a single value, address of the drip recipient.
        // 2. Script can do any logic it wants.
        // 3. Script can signal a successful check by leaving a 1 at the top of the stack.
        // 4. Any value other than a 1 on the stack will signal a failed check.
        bytes memory checkscript = state.config.checkscript;
        address payable recipient = state.config.recipient;

        // Balance threshold checks are a common use case for this contract. Using the checkscript
        // system would be unnecessarily expensive for this, so we designate a special bytecode
        // string to be used for balance threshold checks. Specifically, we look for a 33 byte
        // string starting with the 0x00 (STOP) opcode, followed by the uint256 threshold amount.
        // Since a leading STOP opcode would normally halt execution (and be a useless checkscript)
        // we can safely treat this as a special string. Saves ~30k gas.
        bool executable;
        if (checkscript[0] == hex"00" && checkscript.length == 33) {
            assembly {
                let threshold := mload(add(checkscript, 33))
                executable := lt(balance(recipient), threshold)
            }
        } else {
            // Checkscript is only part of the EVM bytecode that actually gets executed on-chain.
            // We prepend a bytecode snippet that pushes the recipient address onto the stack and
            // allows the checkscript to operate on it. We then also append a snippet that takes
            // the final value on the stack and stores it in memory at 0..32 before reverting with
            // that value.
            bytes memory script = abi.encodePacked(
                // Snippet for pushing the recipient address to the stack.
                hex"73",
                recipient,
                // Actual user checkscript.
                checkscript,
                // Checkscript must leave a value on the stack, this cleanup segment will store
                // that value into memory at position 0..32 and then revert with that value which
                // allows us to access that value via returndatacopy below. This is a convenience
                // since checkscripts would otherwise have to do this manually.
                hex"60005260206000FD"
            );

            assembly {
                // Create a contract using the checkscript as initcode. This will execute the EVM
                // instructions included within the checkscript and, hopefully, leave a single 32
                // byte returndata value. We don't actually care about the returned contract
                // address since the script is intended to revert.
                pop(create(0, add(script, 32), mload(script)))

                // We expect the returned data to be exactly 32 bytes. Anything other than 32 bytes
                // will be ignored, which means "executable" will remain false.
                if eq(returndatasize(), 32) {
                    let ret := mload(0x40)
                    mstore(0x40, add(ret, 32))
                    returndatacopy(ret, 0, 32)
                    executable := eq(mload(ret), 1)
                }
            }
        }

        require(
            executable == true,
            "DrippieRedd: checkscript failed so drip is not yet ready to be triggered"
        );

        state.last = block.timestamp;
        (bool success, ) = recipient.call{ value: state.config.amount }(state.config.data);

        // Generally should not happen, but could if there's a misconfiguration (e.g., passing the
        // wrong data to the target contract), the recipient is not payable, or insufficient gas
        // was supplied to this transaction. We revert so the drip can be fixed and triggered again
        // later. Means we cannot emit an event to alert of the failure, but can reasonably be
        // detected by off-chain services even without an event.
        require(
            success == true,
            "DrippieRedd: drip was unsuccessful, check your configuration for mistakes"
        );

        emit DripExecuted(_name, msg.sender, block.timestamp);
    }
}
