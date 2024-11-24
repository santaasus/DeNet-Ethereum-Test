// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MultiToken {
    struct Token {
        string name;
        string symbol;
        uint256 totalSupply;
    }

    mapping(string => Token) public tokens;
    mapping(string => mapping(address => uint256)) public balances;

    function createToken(string memory name, string memory symbol, uint256 totalSupply) public {
        // map key is short form token's name
        tokens[symbol] = Token(name, symbol, totalSupply);
        balances[symbol][msg.sender] = totalSupply; 
    }

    function balanceOf(string memory symbol, address account) public view returns (uint256) {
        return balances[symbol][account];
    }
}

