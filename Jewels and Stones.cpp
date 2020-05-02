class Solution {
public:
    int numJewelsInStones(string J, string S) {
        bitset<128> j;
        int x;
        int count=0;
        for(int i=0;i<J.length();i++)
        {
            x=J[i];
            j[x]=1;
        }
        for(int i=0;i<S.length();i++)
        {
            x=S[i];
            if(j[x]==1)
                count++;
        }
        return count;
        
    }
};
