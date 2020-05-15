func kadane(A[] int) int{
    answer:=A[0]
    sum:=0
    for _,x := range A{
        if x>sum+x{
            sum=x;
        }else{
            sum+=x
        }
        if answer<sum{
            answer=sum
        }
    }
    return answer
}
func maxSubarraySumCircular(A []int) int {
    wrap:=0
    ans1:=kadane(A)
    for i,x := range A{
        wrap+=x
        A[i]=-A[i]
    }
    ans2:=wrap+kadane(A)
    if ans2==0 || ans1>ans2{
        return ans1
    }else{
        return ans2
    }
}
   