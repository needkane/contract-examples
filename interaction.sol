pragma solidity ^0.4.2;                                                                                                            
 
contract Test {
    
    int public val ;
    
    function GetVal() public returns (int) {
        return val;
    }
    
    
    function SetVal(int v) public {
        val=v;
        return;
    }
}
 
contract FunctionInvoker{
    
    int val2;
    address addrVal;         
    address user_addr = msg.sender;
 
    constructor(address addr){
        addrVal=addr;
    }   
    
    function GetAddr() returns (address){
        return addrVal;
    }   
    
    function GetAddr2() returns (address){
        return msg.sender;
    }   
 
    
    function SetValue() returns (address){
        Test(addrVal).SetVal(12345);
        return addrVal;
    }   
    
    function GetValue() returns (int){
        val2=Test(addrVal).GetVal();
        if (val2 == 0){ 
            val2=-1;
        }   
        
        return val2; 
    }   
}
