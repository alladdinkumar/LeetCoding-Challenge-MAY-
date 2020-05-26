class Solution {
public:
    int findMaxLength(vector<int>& nums) {
        int n=nums.size();
        unordered_map<int,int>m;
        m[0]=0;
        int sum=0;
        int out=0;
        for(int i=0;i<n;i++)
        {
            int x=nums[i];
            sum+=x==0?-1:1;
            if(m.find(sum)!=m.end())
            {
                out=max(out,i-m[sum]+1);
            }
            else
            {
                m[sum]=i+1;
            }
        }
        return out;
    }
};
