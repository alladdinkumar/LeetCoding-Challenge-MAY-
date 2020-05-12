func singleNonDuplicate(nums []int) int {
    x:=nums[0]
    for i:=1;i<len(nums);i++{
        x=x^nums[i]
    }
    return x
    
}