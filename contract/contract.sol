// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

contract MyContract {

    
    address public creator;

    mapping (address => Subcriber) subcribers;

    Subcriber[] arrSubcirbers;

    struct Subcriber {

        address addr;
        bool isCreator;
        bool isSubciber;
        bool isUnSubcirber;
        uint256 amount;
    }


    constructor() {

        creator = msg.sender;
        subcribers[creator].isCreator = true;
        subcribers[creator].isSubciber = true;
        subcribers[creator].isUnSubcirber = false;
        subcribers[creator].addr = creator;
        subcribers[creator].amount = 0;
        arrSubcirbers.push(subcribers[creator]);
    }

    modifier Authentication {

        require((subcribers[msg.sender].isCreator == true 
                || subcribers[msg.sender].isSubciber == true)
                && subcribers[msg.sender].isUnSubcirber == false,
                "Not authentication !!!");
        _;
        
    }

    function addSubciber() private {

        subcribers[msg.sender].isCreator = false;
        subcribers[msg.sender].isSubciber = true;
        subcribers[msg.sender].isUnSubcirber = false;
        subcribers[msg.sender].addr = msg.sender;
        subcribers[msg.sender].amount = msg.value;
        arrSubcirbers.push(subcribers[msg.sender]);
    }

    function subcribe() public payable returns(string memory) {

        if (subcribers[msg.sender].isUnSubcirber == true) {

            revert("You have unsubscribed, you cannot subscribe again !!!");
        } else if (subcribers[msg.sender].isSubciber == true || subcribers[msg.sender].isCreator == true) {

            revert("You are already registered !!!");
        } else {

            require(msg.value >= 1 finney, "You need to pay at least 1 finney to sign the contract !!!");
            addSubciber();
            return "You subcriber successfully !!!";
        }
    }

    function unSubcribe() public return(string memory)
}
