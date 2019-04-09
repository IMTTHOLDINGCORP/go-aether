pragma solidity ^0.4.19;

contract Register_0_0_1 {

    string vsn = "0.0.1";

    function getVsn() constant public returns(string) {
        return vsn;
    }

    address _owner;

    mapping(address => bool) table;
    mapping(string => bool) blacklist;
    mapping(string => address) idmap;

    function Register() public{
        _owner = msg.sender;
        table[_owner] = true;
    }

    modifier owner(address _addr) {
        require(_addr == _owner);
        _;
    }

    function append(address addr) public owner(msg.sender) {
        table[addr] = true;
    }

    function remove(address addr) public owner(msg.sender) {
        delete table[addr];
    }

    function verify(address addr) public constant returns(bool) {
        return table[addr];
    }

    function appendBlackList(string h) public owner(msg.sender) {
        blacklist[h] = true;
    }

    function removeBlackList(string h) public owner(msg.sender) {
        delete blacklist[h];
    }

    function verifyBlackList(string h) public constant returns(bool) {
        return blacklist[h];
    }


    function appendId(string id, string rawdata) public {
        assert(bytes(rawdata).length>0);
        if (idmap[id] == 0) {
            idmap[id] = msg.sender;
        }
    }

    function removeId(string id) public owner(msg.sender) {
        delete idmap[id];
    }

    function getId(string id) public constant returns(string _id,address _cb) {
        _id = id;
        _cb = idmap[id];
        return;
    }

    function hasId(string id) public constant returns(bool) {
        if (idmap[id]==0) {
            return false;
        }
        return true;
    }

}