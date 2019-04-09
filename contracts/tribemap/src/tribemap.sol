pragma solidity ^0.5.3;

contract Owned {

    /// `owner` is the only address that can call a function with this
    /// modifier
    modifier onlyOwner() {
        require(msg.sender == owner);
        _;
    }

    address public owner;

    /// @notice The Constructor assigns the message sender to be `owner`
    constructor() public {
        owner = msg.sender;
    }

    address newOwner = address(0);

    event OwnerUpdate(address _prevOwner, address _newOwner);

    ///change the owner
    function changeOwner(address _newOwner) public onlyOwner {
        require(_newOwner != owner);
        newOwner = _newOwner;
    }

    /// accept the ownership
    function acceptOwnership() public {
        require(msg.sender == newOwner);
        emit OwnerUpdate(owner, newOwner);
        owner = newOwner;
        newOwner = address(0);
    }
}

contract Tribemap is Owned {

    mapping(address => uint256) meshboxAddress;

    struct history {
        uint coin;
        uint num;
        address to;
    }

    mapping(uint => history) historyMap;
    uint[] historyList;

    function addAddress(address[] memory _owners, uint256 version) public onlyOwner() {
        uint len = _owners.length;
        for (uint i = 0; i < len; i ++) {
            meshboxAddress[_owners[i]] = version;
        }
    }

    function delAddress(address[] memory _owners) public onlyOwner() {
        uint len = _owners.length;
        for (uint i = 0; i < len; i ++) {
            meshboxAddress[_owners[i]] = 0;
        }
    }

    function existAddress(address _owner) view public returns (uint256){
        return meshboxAddress[_owner];
    }

    function makeCoin(address to , uint val) public onlyOwner {
        historyMap[block.number].to = to;
        historyMap[block.number].coin = val;
        historyMap[block.number].num = block.number;
        historyList.push(block.number);
    }

    function makeCoinLog() view public returns (
        uint[] memory numList,
        uint[] memory coinList,
        address[] memory toList
    ) {
        if (historyList.length > 0) {
            numList = historyList;
            coinList = new uint[](historyList.length);
            toList = new address[](historyList.length);
            for (uint i=0;i<historyList.length;i++) {
                uint k = historyList[i];
                toList[i] = historyMap[k].to;
                coinList[i] = historyMap[k].coin;
            }
        }
    }
}