// SPDX-License-Identifier: MIT
// OpenZeppelin Contracts v4.3.2 (token/ERC20/IERC20.sol)

pragma solidity >=0.8.0;

/**
 * @dev Interface of the ERC20 standard as defined in the EIP.
 */
interface IERC20Holdable {

    struct Hold {
        address sender;
        address recipient;
        uint256 amount;
    }

    /**
     * @dev Moves `amount` tokens from the caller's account to account[0] (the token owner´s account)
     * until the hold is executed or returned.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {HoldCreated} event.
     */
    function hold(
      uint256 holdId,
      address recipient,
      uint256 amount
    ) external returns (uint256);

    /**
     * @dev Moves `amount` tokens from the caller's account to account[0] (the token owner´s account)
     * using the allowance mechanism until the hold is executed or returned. `amount` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {HoldCreated} event.
     */
    function holdFrom(
      uint256 holdId,
      address sender,
      address recipient,
      uint256 amount
    ) external returns (uint256);

    /**
    *  When the method is executed the money is transferred from the on hold balance to the wallet
    *  where the hold was instructed.
    *
    *  It can only be executed by the same address that created the hold or the contract owner.
    *
    *  Emits a {HoldExecuted} event.
    */
    function executeHold(
      uint256 holdId
    ) external returns (bool);

    /**
    *  Returns the money on hold and all the balance is available for transfers.
    *
    *  This function can only be called by the owner of the contract.
    *
    *  Emits a {HoldCanceled} event.
    */
    function cancelHold(
      uint256 holdId
    ) external returns (bool);
}
