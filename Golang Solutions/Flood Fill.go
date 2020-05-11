func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    var directions = [4][2]int { {0,1},{1,0},{0,-1},{-1,0} }
    q := [][]int{}
    var visited[51][51] bool
    q=append(q,[]int{sr,sc})
    val:=image[sr][sc]
    for ;len(q)!=0;{
        image[q[0][0]][q[0][1]]=newColor
        visited[q[0][0]][q[0][1]]=true
        for _,x := range directions{
            if(q[0][0]+x[0]>=0 && q[0][0]+x[0]<len(image) && q[0][1]+x[1]>=0 && q[0][1]+x[1]<len(image[0]) ){
                   if(image[q[0][0]+x[0]][q[0][1]+x[1]]==val && visited[q[0][0]+x[0]][q[0][1]+x[1]]==false){
                       q=append(q,[]int{q[0][0]+x[0],q[0][1]+x[1]})
                    } 
            }
        } 
        q=q[1:]
    }
    return image
    
}