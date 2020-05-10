func canConstruct(ransomNote string, magazine string) bool {
    m:= make(map[byte]int)
    for i:=0;i<len(magazine);i++{
        
        _, ok := m[magazine[i]]
        if !ok {
            m[magazine[i]]=1
        }else{
            m[magazine[i]]++
        }
        
    }
    for i:=0;i<len(ransomNote);i++{
        val, ok := m[ransomNote[i]]
        if !ok {
            return false;
        }else if val==0{
            return false;   
        }else{
            m[ransomNote[i]]--;
        }
        
    }
    return true;
    
}