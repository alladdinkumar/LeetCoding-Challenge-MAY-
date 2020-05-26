func subsqaure(mat [][]int, k int) int{ 
    n:=len(mat)
    m:=len(mat[0])
    cnt:=0
    var stripSum [][]int
    var i,j int
    for i=0;i<n;i++{
        var temp []int
        for j=0;j<m;j++{
            temp=append(temp,0)
        }
        stripSum=append(stripSum,temp)
    }
    for j=0; j<m; j++{  
        sum:= 0
        for i=0; i<k; i++{
           sum += mat[i][j] 
        }   
        stripSum[0][j] = sum
        for i:=1; i<n-k+1; i++{ 
            sum += (mat[i+k-1][j] - mat[i-1][j])
            stripSum[i][j] = sum
       } 
   } 
    for i:=0; i<n-k+1; i++{ 
        sum := 0; 
        for j:= 0; j<k; j++{
            sum += stripSum[i][j]
        }
        if sum==k*k{
            cnt++ 
        }		
        for j:=1; j<m-k+1; j++{ 
            sum += (stripSum[i][j+k-1] - stripSum[i][j-1])
            if sum==k*k{
                cnt++
            }  		
      } 
   } 
   return cnt 
} 

func countSquares(matrix [][]int) int {
    cnt:=0
    n:=len(matrix)
    m:=len(matrix[0])
    var k int
    if m>n{
        k=n
    }else{
        k=m
    }
    for i:=1;i<=k;i++{
        	cnt+=subsqaure(matrix, i)
		}
	return cnt
}