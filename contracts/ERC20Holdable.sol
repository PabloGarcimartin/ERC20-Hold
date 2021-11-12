// SPDX-License-Identifier: MIT
pragma solidity 0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./interfaces/IERC20Holdable.sol";

contract ERC20Holdable is IERC20Holdable, ERC20 {
    address private owner;

    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

  constructor(uint256 _supply) ERC20("ERC20Holdable", "ERC20H") {
    owner = msg.sender;
    _mint(msg.sender, _supply * (10 ** decimals()));
  }

  struct Hold {
      address sender;
      address recipient;
      uint256 amount;
  }

  event HoldCreated(
      address sender,
      address recipient,
      uint256 amount
  );

  event HoldExecuted(
    uint256 holdId,
    address sender,
    address recipient,
    uint256 amount
  );

  event HoldRemoved(
    uint256 holdId,
    address sender,
    address recipient,
    uint256 amount
  );

  mapping(address => mapping(address => uint256)) private _allowances;
  mapping(address => uint256) private _balances;
  mapping(uint256 => Hold) internal holds;
  mapping(address => uint256) internal heldBalance;
  uint256[] public holdIds;

  function hold(
    uint256 holdId,
    address recipient,
    uint256 amount
  ) public override returns (bool)
  {
      _hold(_msgSender(), recipient, amount, holdId);

      emit HoldCreated(_msgSender(), recipient, amount);
      return true;
  }

  function holdFrom(
    uint256 holdId,
    address sender,
    address recipient,
    uint256 amount
  ) public override returns (bool)
  {
      _hold(sender, recipient, amount, holdId);

      uint256 currentAllowance = _allowances[sender][_msgSender()];
      require(currentAllowance >= amount, "ERC20: transfer amount exceeds allowance");
      unchecked {
          _approve(sender, _msgSender(), currentAllowance - amount);
      }

      emit HoldCreated(sender, recipient, amount);

      return true;
  }

  function executeHold(
    uint256 holdId
  ) public override returns (bool)
  {
      Hold storage executableHold = holds[holdId];

      require(executableHold.amount > 0, "Hold does not exist");
      _transfer(owner, executableHold.recipient, executableHold.amount);
      heldBalance[executableHold.sender] -= executableHold.amount;
      _deleteHold(holdId);

      emit HoldExecuted(holdId, executableHold.sender, executableHold.recipient, executableHold.amount);
      return true;
  }

  function removeHold(
    uint256 holdId
  ) public override returns (bool)
  {
      Hold storage removableHold = holds[holdId];

      require(removableHold.amount > 0, "Hold does not exist");
      _transfer(owner, removableHold.sender, removableHold.amount);

      heldBalance[removableHold.sender] -= removableHold.amount;

      emit HoldRemoved(holdId, removableHold.sender, removableHold.recipient, removableHold.amount);
      _deleteHold(holdId);

      return true;
  }


  /**
   * @dev Moves `amount` of tokens from `sender` to owner to set it on hold.
   *
   * This internal function is equivalent to {hold}, and can be used to
   * e.g. implement automatic token fees, slashing mechanisms, etc.
   *
   *
   * Requirements:
   *
   * - `sender` cannot be the zero address.
   * - `recipient` cannot be the zero address.
   * - `sender` must have a balance of at least `amount`.
   * - `holdId` must be unique.
   */
  function _hold(
      address sender,
      address recipient,
      uint256 amount,
      uint256 holdId
  ) internal virtual {

      require(holds[holdId].amount <= 0, "Hold id must be unique");

      _transfer(sender, owner, amount);

      Hold storage newHold = holds[holdId];

      newHold.sender = sender;
      newHold.recipient = recipient;
      newHold.amount = amount;
      holdIds.push(holdId);

      heldBalance[sender]+= amount;
  }

  function _deleteHold(
    uint256 holdId
  ) internal virtual {
    delete holds[holdId];
    uint i = 0;
    for (i; i <holdIds.length; i = i+1) {  //for loop example
       if( holdIds[i] == holdId){
         delete holdIds[i];
       }
    }
  }
}
