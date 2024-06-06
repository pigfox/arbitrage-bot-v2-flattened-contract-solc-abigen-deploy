// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
// File: @openzeppelin/contracts/token/ERC20/IERC20.sol


// OpenZeppelin Contracts (last updated v5.0.0) (token/ERC20/IERC20.sol)

/**
 * @dev Interface of the ERC20 standard as defined in the EIP.
 */
interface IERC20 {
    /**
     * @dev Emitted when `value` tokens are moved from one account (`from`) to
     * another (`to`).
     *
     * Note that `value` may be zero.
     */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /**
     * @dev Emitted when the allowance of a `spender` for an `owner` is set by
     * a call to {approve}. `value` is the new allowance.
     */
    event Approval(address indexed owner, address indexed spender, uint256 value);

    /**
     * @dev Returns the value of tokens in existence.
     */
    function totalSupply() external view returns (uint256);

    /**
     * @dev Returns the value of tokens owned by `account`.
     */
    function balanceOf(address account) external view returns (uint256);

    /**
     * @dev Moves a `value` amount of tokens from the caller's account to `to`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transfer(address to, uint256 value) external returns (bool);

    /**
     * @dev Returns the remaining number of tokens that `spender` will be
     * allowed to spend on behalf of `owner` through {transferFrom}. This is
     * zero by default.
     *
     * This value changes when {approve} or {transferFrom} are called.
     */
    function allowance(address owner, address spender) external view returns (uint256);

    /**
     * @dev Sets a `value` amount of tokens as the allowance of `spender` over the
     * caller's tokens.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * IMPORTANT: Beware that changing an allowance with this method brings the risk
     * that someone may use both the old and the new allowance by unfortunate
     * transaction ordering. One possible solution to mitigate this race
     * condition is to first reduce the spender's allowance to 0 and set the
     * desired value afterwards:
     * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
     *
     * Emits an {Approval} event.
     */
    function approve(address spender, uint256 value) external returns (bool);

    /**
     * @dev Moves a `value` amount of tokens from `from` to `to` using the
     * allowance mechanism. `value` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(address from, address to, uint256 value) external returns (bool);
}

// File: Dex.sol


contract Dex {
    // Mapping from token address to user address to balance
    mapping(address => mapping(address => uint256)) public balances;

    event Deposit(address indexed token, address indexed user, uint256 amount);
    event Withdraw(address indexed token, address indexed user, uint256 amount);
    event Trade(address indexed token, address indexed user, uint256 amount);

    function deposit(address token, uint256 amount) external {
        require(amount > 0, "Amount must be greater than 0");

        IERC20 erc20Token = IERC20(token);
        bool success = erc20Token.transferFrom(msg.sender, address(this), amount);
        require(success, "Token transfer failed");

        balances[token][msg.sender] += amount;
        emit Deposit(token, msg.sender, amount);
    }

    function withdraw(address token, uint256 amount) external {
        require(amount > 0, "Amount must be greater than 0");
        require(balances[token][msg.sender] >= amount, "Insufficient balance");

        balances[token][msg.sender] -= amount;
        IERC20 erc20Token = IERC20(token);
        bool success = erc20Token.transfer(msg.sender, amount);
        require(success, "Token transfer failed");

        emit Withdraw(token, msg.sender, amount);
    }

    function trade(address token, uint256 amount) external {
        require(amount > 0, "Amount must be greater than 0");
        require(balances[token][msg.sender] >= amount, "Insufficient balance");

        balances[token][msg.sender] -= amount;
        balances[token][msg.sender] += amount;

        emit Trade(token, msg.sender, amount);
    }

    function getBalance(address token, address user) external view returns (uint256) {
        return balances[token][user];
    }
}
