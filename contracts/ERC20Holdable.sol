// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import "../node_modules/@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "./interfaces/IERC20Holdable.sol";

contract ERC20Holdable is IERC20Holdable, ERC20 {
  address private tokenOwner;
  uint8 public constant DECIMALS = 18;
  uint256 public constant INITIAL_SUPPLY = 10000 * (10 ** uint256(DECIMALS));

  constructor() ERC20("ERC20Holdable", "ERC20H") {
    tokenOwner = msg.sender;
    _mint(tokenOwner, 10000 * (10 ** decimals()));
  }

  event HoldCreated(
    uint256 holdId,
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

  event HoldCanceled(
    uint256 holdId,
    address sender,
    address recipient,
    uint256 amount
  );

  mapping(address => mapping(address => uint256)) private _allowances;
  mapping(address => uint256) private _balances;
  mapping(uint256 => Hold) public holds;

  function hold(
    uint256 holdId,
    address recipient,
    uint256 amount
  ) public override returns (uint256)
  {
      _hold(_msgSender(), recipient, amount, holdId);

      emit HoldCreated(holdId, _msgSender(), recipient, amount);
      return holdId;
  }

  function holdFrom(
    uint256 holdId,
    address sender,
    address recipient,
    uint256 amount
  ) public override returns (uint256)
  {
      require(holds[holdId].amount <= 0, "Hold id must be unique");

      Hold storage newHold = holds[holdId];

      newHold.sender = sender;
      newHold.recipient = recipient;
      newHold.amount = amount;

      transferFrom(sender, tokenOwner, amount);

      emit HoldCreated(holdId, _msgSender(), recipient, amount);

      return holdId;
  }

  function executeHold(
    uint256 holdId
  ) public override returns (bool)
  {
      Hold storage executableHold = holds[holdId];
      require(executableHold.amount > 0, "Hold does not exist");
      require(_msgSender() == tokenOwner || executableHold.sender == _msgSender(), 'Only contract owner or hold creator can execute the hold.' );
      _transfer(tokenOwner, executableHold.recipient, executableHold.amount);
      _deleteHold(holdId);

      emit HoldExecuted(holdId, executableHold.sender, executableHold.recipient, executableHold.amount);
      return true;
  }

  function cancelHold(
    uint256 holdId
  ) public override returns (bool)
  {
      Hold storage removableHold = holds[holdId];
      require(_msgSender() == tokenOwner, 'Only contract owner can revert the hold.' );
      require(removableHold.amount > 0, "Hold does not exist");
      _transfer(tokenOwner, removableHold.sender, removableHold.amount);

      emit HoldCanceled(holdId, removableHold.sender, removableHold.recipient, removableHold.amount);
      _deleteHold(holdId);

      return true;
  }

  function _hold(
      address sender,
      address recipient,
      uint256 amount,
      uint256 holdId
  ) internal virtual {

      require(holds[holdId].amount <= 0, "Hold id must be unique");

      _transfer(sender, tokenOwner, amount);

      Hold storage newHold = holds[holdId];

      newHold.sender = sender;
      newHold.recipient = recipient;
      newHold.amount = amount;
  }

  function _deleteHold(
    uint256 holdId
  ) internal virtual {
    delete holds[holdId];
  }
}
