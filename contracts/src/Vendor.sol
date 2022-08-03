// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;


// Learn more about the ERC20 implementation 
// on OpenZeppelin docs: https://docs.openzeppelin.com/contracts/4.x/api/access#Ownable
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "hardhat/console.sol";
import "./bridge.sol";

contract Vendor is Ownable {

  // Our Token Contract
  uint8 public constant destinationChainID = 2;
  uint public constant substrateAddressLen = 32;

  BRIDGE _bridge;
  IERC20 token;
  address handler;
  bytes32 _resourceID;

  // token price for ETH
  uint256 public tokensPerEth = 1;

  // Event that log buy operation
  event BuyTokens(address buyer, uint256 amountOfETH, uint256 amountOfTokens);
  event CallData(bytes data);

  constructor(address bridge, address tokenAddress, address handlerAddress, bytes32 resourceID) {
    _bridge = BRIDGE(bridge);
    token = IERC20(tokenAddress);
    handler = handlerAddress;
    _resourceID = resourceID;
  }

  /**
  * @notice Allow users to buy token for ETH
  */
  function deposit (bytes32 to) external payable {
    require(msg.value > 0, "Send ETH to buy some tokens");

    uint256 amountToBuy = msg.value * tokensPerEth;

    // check if the Vendor Contract has enough amount of tokens for the transaction
    uint256 vendorBalance = token.balanceOf(address(this));
    require(vendorBalance >= amountToBuy, "Vendor contract has not enough tokens in its balance");

    // Transfer token to the msg.sender
    (bool approved) = token.approve(handler, amountToBuy);
    require(approved, "Failed to approve token transfer");

    bytes memory data = abi.encodePacked(bytes32(msg.value), bytes32(substrateAddressLen), to);
    emit CallData(data);

    _bridge.deposit(destinationChainID, _resourceID, data);

    (bool sent) = token.transfer(msg.sender, amountToBuy);
    require(sent, "Failed to transfer token to user");

    // emit the event
    emit BuyTokens(msg.sender, msg.value, amountToBuy);
  }

  /**
  * @notice Allow the owner of the contract to withdraw ETH
  */
  function withdraw() public onlyOwner {
    uint256 ownerBalance = address(this).balance;
    require(ownerBalance > 0, "Owner has not balance to withdraw");

    (bool sent,) = msg.sender.call{value: address(this).balance}("");
    require(sent, "Failed to send user balance back to the owner");
  }
}
