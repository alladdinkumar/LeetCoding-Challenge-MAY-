class Solution {
public:
    vector<int> findAnagrams(string s, string p) {
        map<char,int> m;
        vector<int> x;
        for(int i=0;i<p.length();i++)
            m[p[i]]++;
        int l=0,r;
        int cnt=p.length();
        for(int r=0;r<s.length();r++)
        {
            if(m.find(s[r])!=m.end() && m[s[r]])
            {
               m[s[r]]--;
                cnt--;
                if(cnt==0)
                    x.push_back(l);
            }
            else{
                while(l<r)
                {
                    if(m.find(s[l])!=m.end() )
                    {
                        m[s[l]]++;
                        cnt++;
                    }
                    l++;
                    if(m.find(s[r])!=m.end() && m[s[r]])
                    {
                        r--;
                        break;
                    }
                }
                if(m.find(s[l])==m.end())
                    l=r+1;
            }
           
            
        }
        return x;
        
    }
};
