func findMaxLength(nums []int) int {
    n:=len(nums)
    m:=make(map[int]int)
    m[0]=0;
    sum:=0;
    out:=0;
    for i:=0;i<n;i++{
        if nums[i]==0{
            sum+=-1
        }else{
            sum+=1
        }
        a,ok:=m[sum]
        if ok{
            if out<i-a+1{
                out=i-a+1
            }
        }else{
            m[sum]=i+1;
        }
    }
    return out;
}