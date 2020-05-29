func canFinish(numCourses int, prerequisites [][]int) bool {
    pre_count:=make([]int, numCourses)
    m:=make([][]int, numCourses)
    for _,x:= range prerequisites{
        pre_count[x[0]]++;
        m[x[1]]=append(m[x[1]],x[0]);
    }
    s:=[]int{}
    for i:=0;i<numCourses;i++{
        if pre_count[i]==0{
            s=append(s,i); 
        }          
    }
    var count int
    for count=0;len(s)!=0;count++{
        x:= s[0]
        s=s[1:]
        for _,next:=range m[x]{
            pre_count[next]--;
            if pre_count[next]==0{
                s=append(s,next)
            }       
        }
    }
    return count==numCourses;
   
}