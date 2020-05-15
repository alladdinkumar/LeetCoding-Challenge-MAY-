class Solution {
public:
    int kadane(vector<int> A)
    {
        int answer=A[0],sum=0;
        for(int x:A)
        {
            sum=max(x,sum+x);
            answer=max(answer,sum);
        }
      return answer;  
    }
    int maxSubarraySumCircular(vector<int>& A) {
        int wrap=0;
        int ans1=kadane(A);
        for(int i=0;i<A.size();i++ )
        {
            wrap+=A[i];
            A[i]=-A[i];
        }
        int ans2= wrap+kadane(A);
        if (ans2==0)
            return ans1;
        return ans1>ans2?ans1:ans2;
    }
    
};
