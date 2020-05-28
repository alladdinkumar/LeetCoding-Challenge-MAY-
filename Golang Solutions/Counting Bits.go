func countBits(num int) []int {
    var vec []int
    var tmp []int
    vec=append(vec,0);
    for i:=1;i<=num;{
        for _,x:=range vec{
            tmp=append(tmp,x+1)
            i++
            if i>num{
                break
            }
        }
        vec=append(vec,tmp...)
        tmp=nil
    }
    return vec
}