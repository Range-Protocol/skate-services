// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.20;

library Errors {
    error InvalidPubKey();
    error OperatorNotAllowed();
    error AlreadyAnOperator();
    error MaxOperatorCountReached();
    error MinOperatorStakeNotSatisfied();
    error ZeroOperatorAddress();
    error OperatorAlreadyInAllowlist();
    error OperatorNotInAllowlist();
    error NotAnOperator();
    error AllowlistAlreadyEnabled();
    error AllowlistAlreadyDisabled();
    error ZeroStrategyAddress();
    error StrategyAlreadyAdded();
    error InvalidSignature();
    error DuplicateSignature();
    error SignaturesAreNotOrdered();
    error QuorumNotReached();
}
