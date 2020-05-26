func maxUncrossedLines(A []int, B []int) int {
    m:=len(A)
    n:= len(B)
    var arr [][]int
    for i:=0;i<=m;i++{
        var temp []int
        for j:=0;j<=n;j++{
            temp=append(temp,0)
        }
        arr=append(arr,temp)
    }
     for i:=0;i<=m;i++{
        for j:=0;j<=n;j++{
            if i==0 || j==0{
                arr[i][j]=0;
            }else if A[i-1]==B[j-1]{
                arr[i][j]=1+arr[i-1][j-1];
            }else{
                if arr[i-1][j]>arr[i][j-1]{
                    arr[i][j]=arr[i-1][j]
                }else{
                    arr[i][j]=arr[i][j-1]
                }
            }
        }
    }
    return arr[m][n];
}