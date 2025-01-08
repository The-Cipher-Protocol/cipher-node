// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

contract EpochReward {
    // Hardcoded systemCaller address
    address public constant systemCaller = 0x000000fE0B5d325ae36f04667b75a22E29b94A08;

    uint256 public totalPaidLifetime;
    address public lastWinner;
    uint256 public lastRewardAmount;
    uint256 public lastRewardTimestamp;

    event RewardDistributed(address indexed recipient, uint256 amount);
    event FeeReceived(address indexed sender, uint256 amount);

    modifier onlySystemCaller() {
        require(msg.sender == systemCaller, "Not authorized: Only systemCaller");
        _;
    }

    function distributeReward(address _to, uint256 _amount) external onlySystemCaller {
        require(_amount > 0, "Reward must be > 0");
        require(_to != address(0), "Invalid recipient");
        require(address(this).balance >= _amount, "Insufficient balance");

        (bool sent, ) = _to.call{value: _amount}("");
        require(sent, "Transfer failed");

        lastWinner = _to;
        lastRewardAmount = _amount;
        lastRewardTimestamp = block.timestamp;
        totalPaidLifetime += _amount;

        emit RewardDistributed(_to, _amount);
    }

    receive() external payable {
        emit FeeReceived(msg.sender, msg.value);
    }

    fallback() external payable {
        emit FeeReceived(msg.sender, msg.value);
    }

    function totalAvailableRewards() external view returns (uint256) {
        return address(this).balance;
    }

    function getLastRewardInfo() external view returns (address, uint256, uint256) {
        return (lastWinner, lastRewardAmount, lastRewardTimestamp);
    }
}
