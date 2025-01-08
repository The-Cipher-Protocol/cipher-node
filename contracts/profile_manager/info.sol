// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

// Interface for the Staking contract
interface IStaking {
    function validators() external view returns (address[] memory);
    function accountStake(address addr) external view returns (uint256);
}

contract ValidatorInfo {
    // Hardcoded Staking contract address
    IStaking stakingContract = IStaking(0x000000000000000000000000000000000000FFff);

    struct ValidatorDetails {
        string name;
        string description;
        string website;
        string contactEmail;
        uint256 totalStaked;
    }

    mapping(address => ValidatorDetails) private validatorDetails;

    // Events
    event ValidatorInfoUpdated(
        address indexed validator,
        string name,
        string description,
        string website,
        string contactEmail,
        uint256 totalStaked
    );

    // Modifier to check if msg.sender is a validator
    modifier onlyValidator() {
        address[] memory validators = stakingContract.validators();
        bool isValidator = false;
        for (uint256 i = 0; i < validators.length; i++) {
            if (validators[i] == msg.sender) {
                isValidator = true;
                break;
            }
        }
        require(isValidator, "Only validators can call this function");
        _;
    }

    // Function for validators to update their own information
    function updateValidatorInfo(
        string memory _name,
        string memory _description,
        string memory _website,
        string memory _contactEmail
    ) public onlyValidator {
        validatorDetails[msg.sender].name = _name;
        validatorDetails[msg.sender].description = _description;
        validatorDetails[msg.sender].website = _website;
        validatorDetails[msg.sender].contactEmail = _contactEmail;
        validatorDetails[msg.sender].totalStaked = stakingContract.accountStake(msg.sender);

        emit ValidatorInfoUpdated(
            msg.sender,
            _name,
            _description,
            _website,
            _contactEmail,
            validatorDetails[msg.sender].totalStaked
        );
    }

    // Function to get validator information by address
    function getValidatorInfo(address validator) public view returns (
        string memory name,
        string memory description,
        string memory website,
        string memory contactEmail,
        uint256 totalStaked
    ) {
        ValidatorDetails memory details = validatorDetails[validator];
        return (
            details.name,
            details.description,
            details.website,
            details.contactEmail,
            details.totalStaked
        );
    }
}



// {
// version 8.0.26
// 	"language": "Solidity",
// 	"settings": {
// 		"optimizer": {
// 			"enabled": true,
// 			"runs": 200
// 		},
// 		"evmVersion": "london",
// 		"outputSelection": {
// 			"*": {
// 			"": ["ast"],
// 			"*": ["abi", "metadata", "devdoc", "userdoc", "storageLayout", "evm.legacyAssembly", "evm.bytecode", "evm.deployedBytecode", "evm.methodIdentifiers", "evm.gasEstimates", "evm.assembly"]
// 			}
// 		}
// 	}
// }
