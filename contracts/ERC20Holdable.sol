// SPDX-License-Identifier: MIT
pragma solidity 0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
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
  mapping(uint256 => Hold) public holds;
  mapping(address => uint256) public heldBalance;
  uint256[] public holdIds;

  function hold(
    uint256 holdId,
    address recipient,
    uint256 amount
  ) public override returns (uint256)
  {
      _hold(_msgSender(), recipient, amount, holdId);

      emit HoldCreated(_msgSender(), recipient, amount);
      return holdId;
  }

  function holdFrom(
    uint256 holdId,
    address recipient,
    uint256 amount
  ) public override returns (uint256)
  {
      _hold(_msgSender(), recipient, amount, holdId);

      uint256 currentAllowance = _allowances[_msgSender()][_msgSender()];
      require(currentAllowance >= amount, "ERC20: transfer amount exceeds allowance");
      unchecked {
          _approve(_msgSender(), _msgSender(), currentAllowance - amount);
      }

      emit HoldCreated(_msgSender(), recipient, amount);

      return holdId;
  }

  function executeHold(
    uint256 holdId
  ) public override returns (bool)
  {
      Hold storage executableHold = holds[holdId];
      require(_msgSender() == tokenOwner || executableHold.sender == _msgSender(), 'Only contract owner or hold creator can execute the hold.' );
      require(executableHold.amount > 0, "Hold does not exist");
      _transfer(tokenOwner, executableHold.recipient, executableHold.amount);
      heldBalance[executableHold.sender] -= executableHold.amount;
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

      heldBalance[removableHold.sender] -= removableHold.amount;

      emit HoldRemoved(holdId, removableHold.sender, removableHold.recipient, removableHold.amount);
      _deleteHold(holdId);

      return true;
  }

  function getUserHolds(
    address userAddress
  )  public view override returns (Hold[] memory)
  {
    uint countUserHolds = 0;
    for (uint i=0; i <holdIds.length; i = i+1) {
      if( holds[holdIds[i]].sender == userAddress ){
        countUserHolds++;
      }
    }

    Hold[] memory userHolds = new Hold[](countUserHolds);
    uint count = 0;
    for (uint j=0; j <holdIds.length; j = j+1) {
      if( holds[holdIds[j]].sender == userAddress ){
        userHolds[count] = holds[holdIds[j]];
        count++;
      }
    }

    return userHolds;
  }

  function getHolds(
  )  public view override returns (Hold[] memory)
  {

    Hold[] memory allHolds = new Hold[](holdIds.length);
    for (uint j=0; j <holdIds.length; j = j+1) {
        allHolds[j] = holds[holdIds[j]];
    }

    return allHolds;
  }

  function getHold(
    uint256 holdId
  )  public view override returns (Hold memory)
  {
    Hold memory holdRequested = holds[holdId];

    return holdRequested;
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
