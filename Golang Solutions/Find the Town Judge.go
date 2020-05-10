func findJudge(N int, trust [][]int) int {
    if N==1{
        return 1
    }
    m:= make(map[int]map[int]bool)
    j,cnt:=0,0
    for _,x:= range trust{
        inner,ok:=m[x[0]]
        if(!ok){
            inner=make(map[int]bool)
            m[x[0]]=inner
        }
        inner[x[1]]=true
        
    }
    for i:=1;i<=N;i++{
        _,ok:=m[i]
        if(!ok){
             j=i
            cnt++
        }
    }
    if j==0 && cnt>1{
        return -1
    }
    for i:=1;i<=N;i++{
        _,ok:=m[i]
        if(ok){
            if m[i][j]==false{
                return -1
            }
        }
    }
    return j
        
    
}