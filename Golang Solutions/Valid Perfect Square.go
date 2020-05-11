func isPerfectSquare(num int) bool {
    if num==1{
        return true
    }       
    var l,r int
    l=1
    r=num/2;
    for ;l<=r;{
    var mid int =l+(r-l)/2;
    if mid*mid==num{
        return true
    }else if mid*mid<num {
        l=mid+1
    }else{
        r=mid-1
        }
    }
    return false;
    
}