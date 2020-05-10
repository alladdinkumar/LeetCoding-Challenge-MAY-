class Solution {
public:
    int findJudge(int N, vector<vector<int>>& trust) {
        if(N==1)
            return 1;
        
        bitset<1000> x[N+1];
        int j,cnt=0;
        bitset<1000>z;
        z.set();
        for(int i=0;i<trust.size();i++)
        {
            x[trust[i][0]][trust[i][1]]=1;      
        }
        
        for(int i=1;i<=N;i++)
        {
            
            if(!x[i].any())
            {
                j=i;
                cnt++;
            }  
            else
            {
                z = z & x[i];
            }
        }
        if(j==0 || cnt >1)
            return -1;
        if(j== z._Find_first() )
            return j;
        else
            return -1;
               
        
        
    }
};
