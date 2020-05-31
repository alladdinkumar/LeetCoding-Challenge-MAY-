func minDistance(word1 string, word2 string) int {
    m:=len(word1)
    n:=len(word2)
    dp:=make([][]int ,m+1)
    for i:=0;i<m+1;i++{
        dp[i]=make([]int ,n+1)
        for j:=0;j<n+1;j++{
            if(i==0){
               dp[i][j]=j  
            }else if j==0{
               dp[i][j]=i 
            }else if word1[i-1]==word2[j-1]{
               dp[i][j]=dp[i-1][j-1]
            }else{
                dp[i][j]=dp[i-1][j]
                if dp[i][j]>dp[i][j-1]{
                    dp[i][j]=dp[i][j-1]
                }
                if dp[i][j]>dp[i-1][j-1]{
                    dp[i][j]=dp[i-1][j-1]
                }
                dp[i][j]++ 
            }
        }
    }
    return dp[m][n];
}