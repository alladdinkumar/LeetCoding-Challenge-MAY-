class Solution {
public:
    int firstUniqChar(string s) {
        bitset<26>p,q;
        for(int i=0;i<s.length();i++)
        {
            if(p[s[i]-'a']==0)
            {
                p[s[i]-'a']=1;
            }
            else if(p[s[i]-'a']==1)
            {
                q[s[i]-'a']=1;
            }
        }
        for(int i=0;i<s.length();i++)
        {
            if(q[s[i]-'a']==0)
            {
                return i;
            }
            
        }
        return -1;
    }
};
