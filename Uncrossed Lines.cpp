class Solution {
public:
    
    int maxUncrossedLines(vector<int>& A, vector<int>& B) {
        int m=A.size();
        int n= B.size();
        int arr[m+1][n+1];
        for(int i=0;i<=m;i++)
        {
            for(int j=0;j<=n;j++)
            {
                if(i==0 || j==0)
                    arr[i][j]=0;
                else if(A[i-1]==B[j-1])
                {
                    arr[i][j]=1+arr[i-1][j-1];
                }
                else
                {
                    arr[i][j]=max(arr[i-1][j],arr[i][j-1]);
                }
            }
        }
        return arr[m][n];
    }
};
