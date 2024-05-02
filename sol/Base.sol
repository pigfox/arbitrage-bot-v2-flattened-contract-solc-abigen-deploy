// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

contract Somemath{
    uint private q;
    constructor(){}

    function setQ(uint _q)external{
        q = _q;
    }

    function getQ()external view returns (uint){
        return q;
    }

    function add(uint _x, uint _y) external pure returns(uint){
        return _x+_y;
    }
}

contract Stringer{
    string private x = "abc";
    constructor(){}
    function concat(string memory s)external view returns(string memory){
        return string(abi.encodePacked(x, s));
    }
}

contract Base is Somemath, Stringer{
    constructor(){}
}