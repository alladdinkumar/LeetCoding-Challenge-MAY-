func checkInclusion(s1 string, s2 string) bool {
    m:=make(map[byte]int)
    for i:=0;i<len(s1);i++{
        _,ok:=m[s1[i]]
        if !ok{
            m[s1[i]]=1
        }else{
            m[s1[i]]++
        }
    }
    l,r:=0,0
    cnt:=len(s1)
    for r=0;r<len(s2);r++{
        a,ok:= m[s2[r]]
        if ok && a>0{
            m[s2[r]]--
            cnt--
            if cnt==0{
                return true
            }
                    
        }else{
            for ;l<r;{
                _,ok1:=m[s2[l]]
                if ok1 {
                    m[s2[l]]++
                    cnt++
                }
                l++
                b,ok2:=m[s2[r]]
                if ok2 && b>0{
                    r--;
                    break;
                }
            }
            _,ok3:=m[s2[l]]
            if !ok3{
                l=r+1
            }    
        }
    }
    return false
    
}