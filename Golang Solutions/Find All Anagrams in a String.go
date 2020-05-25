func findAnagrams(s string, p string) []int {
    m:=make(map[byte]int)
    x := []int{}
    for i:=0;i<len(p);i++{
        _,ok:=m[p[i]]
        if !ok{
            m[p[i]]=1
        }else{
            m[p[i]]++
        }
    }
    l,r:=0,0
    cnt:=len(p)
    for r=0;r<len(s);r++{
        a,ok:= m[s[r]]
        if ok && a>0{
            m[s[r]]--
            cnt--
            if cnt==0{
                x=append(x,l)
            }
                    
        }else{
            for ;l<r;{
                _,ok1:=m[s[l]]
                if ok1 {
                    m[s[l]]++
                    cnt++
                }
                l++
                b,ok2:=m[s[r]]
                if ok2 && b>0{
                    r--;
                    break;
                }
            }
            _,ok3:=m[s[l]]
            if !ok3{
                l=r+1
            }    
        }
    }
    return x
}