func findComplement(num int) int {
    if num==0{
        return 1
    }
    var n int= int(math.Log2(float64(num)))+1
    return ((1 << n) - 1) ^ num;
    
}