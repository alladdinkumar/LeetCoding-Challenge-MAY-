class Solution {
public:
int subsqaure(vector<vector<int>>& mat, int k) 
{ 
	int n=mat.size();
	int m=mat[0].size();
	int cnt=0;
   int stripSum[n][m]; 
   for (int j=0; j<m; j++) 
   {  
       int sum = 0; 
       for (int i=0; i<k; i++) 
          sum += mat[i][j]; 
       stripSum[0][j] = sum; 
       for (int i=1; i<n-k+1; i++) 
       { 
            sum += (mat[i+k-1][j] - mat[i-1][j]); 
            stripSum[i][j] = sum; 
       } 
   } 
   for (int i=0; i<n-k+1; i++) 
   { 
      int sum = 0; 
      for (int j = 0; j<k; j++) 
           sum += stripSum[i][j]; 
    
      if(sum==k*k)
        		cnt++;
      for (int j=1; j<m-k+1; j++) 
      { 
         sum += (stripSum[i][j+k-1] - stripSum[i][j-1]); 
       
         if(sum==k*k)
        		cnt++;
      } 
  
    
	  
   } 
   return cnt; 
} 
int countSquares(vector<vector<int>>& matrix) {
	int cnt=0;
	int n=matrix.size(),m=matrix[0].size();
	int k= m>n?n:m;
	
        for(int i=1;i<=k;i++)
        {
        	cnt+=subsqaure(matrix, i);
		}
		return cnt;
}
};
