// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

contract SkateGateway {
    address owner;
    mapping(uint256 => MessageData) public messages;
    mapping(address => bool) private operators;

    struct MessageData {
        string message;
        address signer;
    }

    constructor(address _owner) {
        owner = _owner;
    }

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    modifier onlyOperators() {
        require(operators[msg.sender] == true);
        _;
    }

    function registerOperator(address op) public onlyOwner {
        operators[op] = true;
    }

    function deregisterOperator(address op) public onlyOwner {
        operators[op] = false;
    }

    function postMsg(
        uint taskId,
        string memory message,
        address signer
    ) public onlyOperators {
        messages[taskId] = MessageData(message, signer);
    }

    function getMsg(uint taskId) public view returns (string memory) {
        return messages[taskId].message;
    }
}
