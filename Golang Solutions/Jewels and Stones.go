func numJewelsInStones(J string, S string) int {
    
    var j[128] bool
    var x int 
    var count int =0
    for i:=0;i<len(J);i++{
        x=int(J[i])
        j[x]=true
    }
    for i:=0;i<len(S);i++{
        x=int(S[i])
        if j[x]==true{
            count++
        }
            
    }
    return count
    
}