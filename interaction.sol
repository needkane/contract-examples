pragma solidity ^0.4.0;

contract Test {
    
    uint val;
    
    function GetVal() returns (uint) {
        return val;
    }
    
    function SetVal(uint v) {
        val=v;
        return;
    }
}

contract FunctionInvoker{
    
    address _address = 0x7777e397863b17106fc9f0cb7144f7b536b3709e;
    function SetValue() public {
        Test(_address).SetVal(12345);
        return;
    }
    
    function GetValue() returns (uint){
        uint val=Test(_address).GetVal();
        return val; 
    }
}
