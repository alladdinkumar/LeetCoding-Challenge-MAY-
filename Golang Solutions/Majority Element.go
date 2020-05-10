func majorityElement(nums []int) int {
    m := make(map[int]int)
    for _,n := range nums{
        _,ok:=m[n]
        if !ok {
            m[n]=1
        }else{
            m[n]++
        }
        if m[n]>len(nums)/2{
            return n
        }
    }
    return -1
}