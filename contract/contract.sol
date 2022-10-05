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
        bool refund;
        uint256 amount;
    }

    // struct Transaction {

    //     address sender;
    //     address receiver;
    //     uint256 amount;
    // }

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
        subcribers[msg.sender].refund = false;
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

            require(msg.value >= 0.0001 ether, "You need to pay at least 0.0001 ether to sign the contract !!!");
            addSubciber();
            return "You subcriber successfully !!!";
        }
    }

    function removeElement(uint index) private {

        if (index >= arrSubcirbers.length) {

            revert("Error subcriber is not survive !!!");
        }
        for (uint i = index; i < arrSubcirbers.length - 1; i++ ) {

            arrSubcirbers[i] = arrSubcirbers[i + 1];
        }
        arrSubcirbers.pop();
    }

    function unSubcribe() public Authentication returns(string memory) {

        subcribers[msg.sender].isUnSubcirber = true;
        for (uint i = 0; i < arrSubcirbers.length; i++) {

            if (msg.sender == arrSubcirbers[i].addr) {

                removeElement(i);
                break;
            }
        }
        return "You unsubcriber successfully !!!";
    }

    function Refund() public returns(string memory) {

        require(subcribers[msg.sender].isSubciber == true  
                && subcribers[msg.sender].isUnSubcirber == true,
                "You have not unsubscribed !!!");
        require(subcribers[msg.sender].refund == false, "You got your money back !!!");
        (bool sent,) = msg.sender.call{value: subcribers[msg.sender].amount / 2 wei}("");
        require(sent, "Failed to send Ether");
        subcribers[msg.sender].refund = true;
        return "Refund successfully !!!";
    }

    function List() public view returns(Subcriber[] memory) {

        return arrSubcirbers;
    }
}