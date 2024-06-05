// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @dev Interface of the ERC20 standard as defined in the EIP.
 */
interface IERC20 {
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function transfer(address to, uint256 value) external returns (bool);
    function allowance(address owner, address spender) external view returns (uint256);
    function approve(address spender, uint256 value) external returns (bool);
    function transferFrom(address from, address to, uint256 value) external returns (bool);
}

interface IDex {
    function deposit(address token, uint256 amount) external;
    function withdraw(address token, uint256 amount) external;
    function getBalance(address token, address user) external view returns (uint256);
}

contract Trade {
    address private owner;
    address private profitDestination;

    event SwapExecuted(address indexed user, uint256 amount, address indexed sourceDex, address indexed destinationDex);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    constructor(address _profitDestination) {
        owner = msg.sender;
        profitDestination = _profitDestination;
    }

    function setOwner(address _owner) external onlyOwner {
        owner = _owner;
    }

    function setProfitDestination(address _profitDestination) external onlyOwner {
        profitDestination = _profitDestination;
    }

    function swap(
        address _sourceDexAddress,
        address _destinationDexAddress,
        address _tokenAddress,
        uint256 _amount
    ) external onlyOwner {
        require(_amount > 0, "Amount must be greater than 0");

        IDex sourceDex = IDex(_sourceDexAddress);
        IDex destinationDex = IDex(_destinationDexAddress);
        IERC20 token = IERC20(_tokenAddress);

        // Transfer tokens from user to this contract
        bool success = token.transferFrom(msg.sender, address(this), _amount);
        require(success, "Token transfer failed");

        // Deposit tokens to the source DEX
        token.approve(_sourceDexAddress, _amount);
        sourceDex.deposit(_tokenAddress, _amount);

        // Withdraw tokens from the source DEX
        sourceDex.withdraw(_tokenAddress, _amount);

        // Get initial balance of the destination DEX
        uint256 initialDestinationBalance = destinationDex.getBalance(_tokenAddress, address(this));

        // Conditionally deposit tokens to the destination DEX
        token.approve(_destinationDexAddress, _amount);
        destinationDex.deposit(_tokenAddress, _amount);

        // Withdraw tokens from the destination DEX
        destinationDex.withdraw(_tokenAddress, _amount);

        // Get final balance of the destination DEX
        uint256 finalDestinationBalance = destinationDex.getBalance(_tokenAddress, address(this));

        // Calculate profit
        uint256 profit = finalDestinationBalance - initialDestinationBalance;

        // Transfer profit to profitDestination
        if (profit > 0) {
            token.transfer(profitDestination, profit);
        }

        emit SwapExecuted(msg.sender, _amount, _sourceDexAddress, _destinationDexAddress);
    }

    function getProfitDestination() external view returns (address) {
        return profitDestination;
    }

    function getOwner() external view returns (address) {
        return owner;
    }
}
