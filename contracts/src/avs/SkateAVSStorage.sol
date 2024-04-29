// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.20;

import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {ISkateAVS} from "./interfaces/ISkateAVS.sol";
import {IAVSDirectory} from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import {IDelegationManager} from "./interfaces/IDelegationManager.sol";

abstract contract SkateAVSStorage is ISkateAVS {
    IStrategy[] internal _strategies;
    address[] internal _operators;
    mapping(address => bytes) internal _operatorPubkeys;
    mapping(address => bool) internal _allowlist;
    uint32 internal _maxOperatorCount;
    bool public _allowlistEnabled;
    uint96 public _minOperatorStake;

    function strategies() external view override returns (IStrategy[] memory) {
        return _strategies;
    }

    //    function operatorPubkeys(address operator) external view override returns () {}

    function isInAllowlist(address operator) external view override returns (bool) {
        return _allowlist[operator];
    }

    function maxOperatorCount() external view override returns (uint32) {
        return _maxOperatorCount;
    }

    function allowlistEnabled() external view override returns (bool) {
        return _allowlistEnabled;
    }

    function minOperatorStake() external view override returns (uint96) {
        return _minOperatorStake;
    }
}
