class Solution {
public:
    bool checkInclusion(string s1, string s2) {
        map<char,int> m;
        vector<int> x;
        for(int i=0;i<s1.length();i++)
            m[p[i]]++;
        int l=0,r;
        int cnt=s1.length();
        for(int r=0;r<s2.length();r++)
        {
            if(m.find(s2[r])!=m.end() && m[s2[r]])
            {
               m[s2[r]]--;
                cnt--;
                if(cnt==0)
                    return true;
            }
            else{
                while(l<r)
                {
                    if(m.find(s2[l])!=m.end() )
                    {
                        m[s2[l]]++;
                        cnt++;
                    }
                    l++;
                    if(m.find(s2[r])!=m.end() && m[s2[r]])
                    {
                        r--;
                        break;
                    }
                }
                if(m.find(s2[l])==m.end())
                    l=r+1;
            }
           
            
        }
        return false;
        
    }
    
};
