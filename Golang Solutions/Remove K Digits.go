func removeKdigits(num string, k int) string {
    if(k==len(num)){
        return "0"
    }
    var i int=0
    for ;k>0;k--{
        if i-1>0 {
            i=i-1  
        }else{
            i=0
        }
        for ;i <len(num)-1 && num[i] <= num[i+1];i++{}
        num=num[0:i]+num[i+1:]
    }
    
    for ;num[0]=='0';{
        
        num=num[1:]
        if num==""{
        return "0"
    }
    }
   
    return num
    
    
}