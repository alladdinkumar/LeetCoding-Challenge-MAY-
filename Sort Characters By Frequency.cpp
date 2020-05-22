class Solution {
    
public:
    string frequencySort(string s) {
        map<char,int> m;
        for(char c: s)
        {
            m[c]++;
        }
        vector<pair<int,char>> v;
        for(auto x:m)
        {
            v.push_back({x.second,x.first});
        }
        sort(v.begin(),v.end(),greater<pair<int,char>>());
        s="";
        for(auto x : v)
        {
            while(x.first--)
            {
                s+=x.second;
            }
            
        }
        return s;
        
    }
};
