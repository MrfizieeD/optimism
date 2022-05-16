// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

contract SimpleStorage {
    bytes32 value;

    function setValue(bytes32 _value) public payable {
        value = _value;
    }

    function getValue() public view returns (bytes32) {
        return value;
    }
}
