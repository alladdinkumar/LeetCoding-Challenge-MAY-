func checkStraightLine(coordinates [][]int) bool {
    if len(coordinates)==2{
            return true
        }
    var a int
    n:=len(coordinates)
    for i:=2;i<n;i++ {
      
        a = coordinates[0][0]*(coordinates[1][1]-coordinates[i][1]) + coordinates[1][0]*(coordinates[i][1]-coordinates[0][1]) + coordinates[i][0]*(coordinates[0][1]-coordinates[1][1])
        if a!=0{
            return false
        }
    }
    return true   
}