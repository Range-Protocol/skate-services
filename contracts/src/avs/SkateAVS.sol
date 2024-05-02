// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.20;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

import {IAVSDirectory} from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {ISignatureUtils} from "eigenlayer-contracts/src/contracts/interfaces/ISignatureUtils.sol";
import {IDelegationManager} from "./interfaces/IDelegationManager.sol";

import {SkateAVSStorage} from "./SkateAVSStorage.sol";
import {Errors} from "./Errors.sol";
//import {console2} from "forge-std/Test.sol";

contract SkateAVS is Initializable, UUPSUpgradeable, OwnableUpgradeable, PausableUpgradeable, SkateAVSStorage {
    using ECDSA for bytes32;
    using MessageHashUtils for bytes32;

    IAVSDirectory internal immutable _avsDirectory;
    IDelegationManager internal immutable _delegationManager;

    constructor(IAVSDirectory avsDirectory_, IDelegationManager delegationManager_) {
        _avsDirectory = avsDirectory_;
        _delegationManager = delegationManager_;
        _disableInitializers();
    }

    function initialize(
        address owner_,
        address[] calldata strategies_,
        string calldata metadataURI_,
        bool allowlistEnabled_
    ) external initializer {
        if (allowlistEnabled_) {
            _enableAllowlist();
        }

        if (bytes(metadataURI_).length != 0) {
            _avsDirectory.updateAVSMetadataURI(metadataURI_);
        }
        _setMaxOperatorCount(5);
        _setMinOperatorStake(1);
        _setStrategies(strategies_);
        _transferOwnership(owner_);
    }

    function registerOperator(
        bytes calldata pubkey,
        ISignatureUtils.SignatureWithSaltAndExpiry memory operatorSignature
    ) external override whenNotPaused {
        address operator = msg.sender;
        if (operator != address(uint160(uint256(keccak256(pubkey))))) revert Errors.InvalidPubKey();
        if (_allowlistEnabled && !_allowlist[operator]) revert Errors.OperatorNotAllowed();
        if (_isOperator(operator)) revert Errors.AlreadyAnOperator();
        if (_operators.length == _maxOperatorCount) revert Errors.MaxOperatorCountReached();
        if (_getTotalDelegations(operator) < _minOperatorStake) revert Errors.MinOperatorStakeNotSatisfied();

        _addOperator(operator, pubkey);
        _avsDirectory.registerOperatorToAVS(operator, operatorSignature);

        emit OperatorAdded(operator);
    }

    /**
     * @notice Returns the EigenLayer AVSDirectory contract.
     * @dev Implemented to match IServiceManager interface - required for compatibility with
     *      eigenlayer frontend.
     */
    function avsDirectory() external view override returns (address) {
        return address(_avsDirectory);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              AVS Views                                   //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Returns the currrent list of operator registered as OmniAVS.
     *         Operator.addr        = The operator's ethereum address
     *         Operator.pubkey      = The operator's 64 byte uncompressed secp256k1 public key
     *         Operator.staked      = The total amount staked by the operator, not including delegations
     *         Operator.delegated   = The total amount delegated, not including operator stake
     */
    function operators() external view override returns (Operator[] memory) {
        return _getOperators();
    }

    /**
     * @notice Returns the list of strategies that the AVS supports for restaking.
     * @dev Implemented to match IServiceManager interface - required for compatibility with
     *      eigenlayer frontend.
     */
    function getRestakeableStrategies() external view override returns (address[] memory) {
        return _getRestakeableStrategies();
    }

    /**
     * @notice Returns the list of strategies that the operator has potentially restaked on the AVS
     * @dev Implemented to match IServiceManager interface - required for compatibility with
     *      eigenlayer frontend.
     *
     *      This function is intended to be called off-chain
     *
     *      No guarantee is made on whether the operator has shares for a strategy. The off-chain
     *      service should do that validation separately. This matches the behavior defined in
     *      eigenlayer-middleware's ServiceManagerBase.
     *
     * @param operator The address of the operator to get restaked strategies for
     */
    function getOperatorRestakedStrategies(address operator) external view override returns (address[] memory) {
        if (!_isOperator(operator)) return new address[](0);
        return _getRestakeableStrategies();
    }

    /**
     * @notice Check if an operator can register to the AVS.
     *         Returns true, with no reason, if the operator can register to the AVS.
     *         Returns false, with a reason, if the operator cannot register to the AVS.
     * @dev This function is intented to be called off-chain.
     * @param operator The operator to check
     * @return canRegister True if the operator can register, false otherwise
     */
    function canRegister(address operator) external view override returns (bool) {
        if (!_delegationManager.isOperator(operator)) revert Errors.NotAnOperator();
        if (_allowlistEnabled && !_allowlist[operator]) revert Errors.OperatorNotAllowed();
        if (_isOperator(operator)) revert Errors.AlreadyAnOperator();
        if (_operators.length >= _maxOperatorCount) revert Errors.MaxOperatorCountReached();
        if (_getTotalDelegations(operator) < _minOperatorStake) revert Errors.MinOperatorStakeNotSatisfied();
        return true;
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Admin functions                             //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Sets AVS metadata URI with the AVSDirectory.
     */
    function setMetadataURI(string memory metadataURI) external override onlyOwner {
        _avsDirectory.updateAVSMetadataURI(metadataURI);
    }

    /**
     * @notice sets the new stratgies.
     * @param strategies_ the list of new strategies to set
     */
    function setStrategies(address[] calldata strategies_) external override onlyOwner {
        _setStrategies(strategies_);
    }

    /**
     * @notice Set the minimum operator stake.
     * @param stake The minimum operator stake, not including delegations
     */
    function setMinOperatorStake(uint96 stake) external override onlyOwner {
        _setMinOperatorStake(stake);
    }

    /**
     * @notice Set the maximum operator count.
     * @param count The maximum operator count
     */
    function setMaxOperatorCount(uint32 count) external override onlyOwner {
        _setMaxOperatorCount(count);
    }

    /**
     * @notice Add an operator to the allowlist.
     * @param operator The operator to add
     */
    function addToAllowlist(address operator) external override onlyOwner {
        if (operator == address(0x0)) revert Errors.ZeroOperatorAddress();
        if (_allowlist[operator]) revert Errors.OperatorAlreadyInAllowlist();
        _allowlist[operator] = true;
        emit OperatorAllowed(operator);
    }

    /**
     * @notice Remove an operator from the allowlist.
     * @param operator The operator to remove
     */
    function removeFromAllowlist(address operator) external override onlyOwner {
        if (!_allowlist[operator]) revert Errors.OperatorNotInAllowlist();
        _allowlist[operator] = false;
        emit OperatorDisallowed(operator);
    }

    /**
     * @notice Enable the allowlist.
     */
    function enableAllowlist() external override onlyOwner {
        _enableAllowlist();
    }

    /**
     * @notice Disable the allowlist.
     */
    function disableAllowlist() external override onlyOwner {
        _disableAllowlist();
    }

    /**
     * @notice Eject an operator from the AVS.
     */
    function ejectOperator(address operator) external override onlyOwner {
        _deregisterOperator(operator);
    }

    /**
     * @notice Pause the contract.
     * @dev This pauses registerOperatorToAVS, deregisterOperatorFromAVS, and syncWithOmni.
     */
    function pause() external override onlyOwner {
        _pause();
    }

    /**
     * @notice Unpause the contract.
     */
    function unpause() external override onlyOwner {
        _unpause();
    }

    function submitData(uint256 taskId, bytes calldata messageData, SignatureTuple[] calldata signatureTuples)
        public
        override
    {
        bytes32 digest = keccak256(abi.encodePacked(taskId, messageData)).toEthSignedMessageHash();
        uint256 sigsVerified;
        bool quorumSuccessful;
        for (uint256 i = 0; i < signatureTuples.length; i++) {
            SignatureTuple memory sigTuple = signatureTuples[i];
            if (!_isOperator(sigTuple.operator)) revert Errors.NotAnOperator();

            if (i > 0) {
                SignatureTuple memory prevSigTuple = signatureTuples[i - 1];
                if (sigTuple.operator == prevSigTuple.operator) revert Errors.DuplicateSignature();
                if (sigTuple.operator < prevSigTuple.operator) revert Errors.SignaturesAreNotOrdered();
            }

            if (ECDSA.recover(digest, sigTuple.signature) != sigTuple.operator) revert Errors.InvalidSignature();
            sigsVerified++;

            // 2/3 of operators must submit the data.
            // NOTE: future considerations
            // 1. Weighted by operator stake amount.
            // 2. to ensure strict BFT, should check using `(2 + amount * 2) / 3 > totalAmount`
            if (sigsVerified * 10_000 >= _getOperators().length * 6666) {
                quorumSuccessful = true;
                break;
            }
        }

        if (!quorumSuccessful) revert Errors.QuorumNotReached();
        emit DataSubmitted(taskId, messageData);
    }

    function batchSubmitData(
        uint256[] calldata taskIds,
        bytes[] calldata messageDatas,
        SignatureTuple[][] calldata signaturesTuples
    ) external override {
        for (uint256 i = 0; i < taskIds.length; i++) {
            submitData(taskIds[i], messageDatas[i], signaturesTuples[i]);
        }
    }

    /**
     * @notice Deregister an operator from the AVS. Forwards a call to EigenLayer's AVSDirectory.
     */
    function _deregisterOperator(address operator) private {
        if (!_isOperator(operator)) revert Errors.NotAnOperator();

        _removeOperator(operator);
        _avsDirectory.deregisterOperatorFromAVS(operator);

        emit OperatorRemoved(operator);
    }

    /**
     * @notice Add an operator to internal AVS state (_operators, _operatorPubkeys)
     * @dev Does not check if operator already exists
     */
    function _addOperator(address operator, bytes calldata pubkey) private {
        _operators.push(operator);
        _operatorPubkeys[operator] = pubkey;
    }

    /**
     * @notice Removes an operator from internal AVS state (_operators, _operatorPubkeys)
     * @dev Does not check if operator exists
     */
    function _removeOperator(address operator) private {
        for (uint256 i = 0; i < _operators.length;) {
            if (_operators[i] == operator) {
                _operators[i] = _operators[_operators.length - 1];
                _operators.pop();
                break;
            }
            unchecked {
                i++;
            }
        }
        delete _operatorPubkeys[operator];
    }

    /**
     * @notice Set the minimum operator stake.
     * @param stake The minimum operator stake, not including delegations
     */
    function _setMinOperatorStake(uint96 stake) private {
        _minOperatorStake = stake;
        emit MinOperatorStakeSet(stake);
    }

    /**
     * @notice Set the maximum operator count.
     * @param count The maximum operator count
     */
    function _setMaxOperatorCount(uint32 count) private {
        _maxOperatorCount = count;
        emit MaxOperatorCountSet(count);
    }

    /**
     * @notice Enable the allowlist.
     */
    function _enableAllowlist() private {
        if (_allowlistEnabled) revert Errors.AllowlistAlreadyEnabled();
        _allowlistEnabled = true;
        emit AllowlistEnabled();
    }

    /**
     * @notice Disable the allowlist.
     */
    function _disableAllowlist() private {
        if (!_allowlistEnabled) revert Errors.AllowlistAlreadyDisabled();
        _allowlistEnabled = false;
        emit AllowlistDisabled();
    }

    /**
     * @notice sets the new strategies.
     * @param strategies_ the new strategies
     */
    function _setStrategies(address[] calldata strategies_) private {
        delete _strategies;

        for (uint256 i = 0; i < strategies_.length;) {
            if (strategies_[i] == address(0x0)) revert Errors.ZeroStrategyAddress();

            // ensure no duplicates
            for (uint256 j = i + 1; j < strategies_.length;) {
                if (strategies_[i] == address(strategies_[j])) revert Errors.StrategyAlreadyAdded();
                unchecked {
                    j++;
                }
            }

            _strategies.push(IStrategy(strategies_[i]));
            unchecked {
                i++;
            }
        }

        emit StrategiesSet(strategies_);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Internal views                              //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Returns true if the operator is in the list of operators
     */
    function _isOperator(address operator) private view returns (bool) {
        return _operatorPubkeys[operator].length > 0;
    }

    /**
     * @notice Return current list of Operators, including their personal stake and delegated stake
     */
    function _getOperators() internal view returns (Operator[] memory) {
        Operator[] memory ops = new Operator[](_operators.length);

        for (uint256 i = 0; i < ops.length;) {
            address operator = _operators[i];

            uint96 total = _getTotalDelegations(operator);
            uint96 staked = _getSelfDelegations(operator);

            // this should never happen, but just in case
            uint96 delegated = total > staked ? total - staked : 0;
            bytes memory pubkey = _operatorPubkeys[operator];

            ops[i] = Operator(operator, pubkey, delegated, staked);
            unchecked {
                i++;
            }
        }

        return ops;
    }

    /**
     * @notice Returns the operator's self-delegations
     * @param operator The operator address
     */
    function _getSelfDelegations(address operator) internal view returns (uint96) {
        (IStrategy[] memory strategies_, uint256[] memory shares) = _delegationManager.getDelegatableShares(operator);

        uint96 staked;
        for (uint256 i = 0; i < strategies_.length; i++) {
            IStrategy strategyToCheck = strategies_[i];

            // find the strategy params for the strategy
            IStrategy strategy;
            for (uint256 j = 0; j < _strategies.length;) {
                if (address(_strategies[j]) == address(strategyToCheck)) {
                    strategy = _strategies[j];
                    break;
                }
                unchecked {
                    j++;
                }
            }

            // if strategy is not found, do not consider it in stake
            if (address(strategy) == address(0)) continue;

            staked += uint96(shares[i]);
        }

        return staked;
    }

    /**
     * @notice Returns total delegations to the operator, including self delegations
     * @param operator The operator address
     */
    function _getTotalDelegations(address operator) internal view returns (uint96) {
        uint96 total;
        IStrategy strategy;

        for (uint256 i = 0; i < _strategies.length;) {
            strategy = _strategies[i];
            uint256 shares = _delegationManager.operatorShares(operator, strategy);
            total += uint96(shares);
            unchecked {
                i++;
            }
        }

        return total;
    }

    /**
     * @notice Returns the list of restakeable strategy addresses
     */
    function _getRestakeableStrategies() internal view returns (address[] memory) {
        address[] memory strategies_ = new address[](_strategies.length);
        for (uint256 i = 0; i < _strategies.length;) {
            strategies_[i] = address(_strategies[i]);
            unchecked {
                i++;
            }
        }
        return strategies_;
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
