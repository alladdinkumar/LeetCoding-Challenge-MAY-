var g [][]int
var colors []int
func dfs( cur int,color int )bool{
    colors[cur]=color
    for _, next:= range g[cur]{
        if colors[next]==color{
            return false
        }
        if colors[next]==0 && !dfs(next,-color){
            return false
        }
    
    }
    return true
}
func possibleBipartition(N int, dislikes [][]int) bool {
    g= make([][]int,N)
    for _,vec:= range dislikes{
        a:=vec[0]
        b:=vec[1]
        g[a-1]=append(g[a-1],b-1)
        g[b-1]=append(g[b-1],a-1)       
    }
    colors=make([]int,N)
    for i:=0;i<N;i++{
        if colors[i]==0 && !dfs(i,1){
            return false
        }     
    }
    return true;
}