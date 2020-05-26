func frequencySort(s string) string {
    m:=make(map[byte]int)
    for i:=0;i<len(s);i++{
        _,ok:=m[s[i]]
        if !ok{
            m[s[i]]=1
        }else{
            m[s[i]]++
        }
    }
    m2:=make(map[int]string)
    for key,val:=range m{
        _,ok:=m2[val]
        if !ok{
            m2[val]=strings.Repeat(string(key),val)
        }else{
            m2[val]+=strings.Repeat(string(key),val)
        }
    }
    keys := make([]int, 0, len(m2))
    for k := range m2 {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    var str string=""
    for i:=len(keys)-1;i>=0;i--{
        str+=m2[keys[i]]
    }
    return str
}