func firstUniqChar(s string) int {
    var p[26][2] bool
    for _,c := range s{
        if p[int(c)-int('a')][0]==false{
            p[int(c)-int('a')][0]=true
        }else if(p[int(c)-int('a')][0]==true){
            p[int(c)-int('a')][1]=true
        }
    }
    for i,c := range s{
        if p[int(c)-int('a')][1]==false{
            return i
        }
    }
    
    return -1
    
}