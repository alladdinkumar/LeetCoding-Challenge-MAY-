func intervalIntersection(A [][]int, B [][]int) [][]int {
    var vec [][]int
    var tmp []int
    for i,j:=0,0;i<len(A) && j<len(B); {
        if !(A[i][0] > B[j][1] || A[i][1] < B[j][0]){
            l:=int(math.Max(float64(A[i][0]),float64(B[j][0])))
            r:=int(math.Min(float64(A[i][1]),float64(B[j][1])))
            tmp=nil
            tmp=append(tmp,l)
            tmp=append(tmp,r)
            vec=append(vec,tmp)
        }
        if i<len(A)-1 && B[j][1]>=A[i+1][0]{
            i++
            continue
        }
        if j<len(B)-1 && A[i][1]>=B[j+1][0]{
            j++
            continue;
        }
        i++
        j++
        
    }
    return vec
}