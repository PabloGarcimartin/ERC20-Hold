// SPDX-License-Identifier: MIT
// OpenZeppelin Contracts v4.3.2 (token/ERC20/IERC20.sol)

pragma solidity ^0.8.0;

/**
 * @dev Interface of the ERC20 standard as defined in the EIP.
 */
interface IERC20Holdable {
    /**
     * @dev Moves `amount` tokens from the caller's account to account[0] (the token owner´s account)
     * until the hold is executed or returned.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function hold(
      uint256 holdId,
      address recipient,
      uint256 amount
    ) external returns (bool);

    /**
     * @dev Moves `amount` tokens from the caller's account to account[0] (the token owner´s account)
     * using the allowance mechanism until the hold is executed or returned. `amount` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function holdFrom(
      uint256 holdId,
      address sender,
      address recipient,
      uint256 amount
    ) external returns (bool);

    /**
    *  When the method is executed the money is transferred from the on hold balance to the wallet
    *  where the hold was instructed.
    *
    *  It can only be executed by the same address that created the hold.
    *
    *  Emits a {Transfer} event.
    */
    function executeHold(
      uint256 holdId
    ) external returns (bool);

    /**
    *  Returns the money on hold and all the balance is available for transfers.
    *
    *  This function can only be called by the owner of the contract.
    *
    *  Emits a {Transfer} event.
    */
    function removeHold(
      uint256 holdId
    ) external returns (bool);

}
