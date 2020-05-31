var d []float64
var p []int
func euclidean(x []int) float64{
    a:=math.Pow(float64(x[0]),2)
    b:=math.Pow(float64(x[1]),2)
    y:= math.Sqrt(a+b)
    return y
}
func insert(dist float64,index int){
    if d==nil{
        d=append(d,dist)
        p=append(p,index)
    }else{
        var i int
        for i=0;i<len(d) && d[i] < dist;{
            i++
        }
        d = append(d[:i],append([]float64{dist}, d[i:]...)... )
        p = append(p[:i],append([]int{index}, p[i:]...)... )
    }
}
func kClosest(points [][]int, K int) [][]int {
    var sol [][]int
    d=nil
    p=nil
    cnt:=0
    for i,x:=range points{
        dist:=euclidean(x)
        if cnt<K{
            insert(dist,i)
            cnt++
        }else if dist<d[len(d)-1]{
            insert(dist,i)
            d=d[:len(d)-1]
            p=p[:len(p)-1]
        }
    }
    for _,x:=range p{
        sol=append(sol,points[x])
    }
    return sol
}