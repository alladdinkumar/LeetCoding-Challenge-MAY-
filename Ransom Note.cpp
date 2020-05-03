class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        map<char,int> m;
        for(int i=0;i<magazine.length();i++)
        {
            if(m.find(magazine[i])==m.end())
            {
                m[magazine[i]]=1;
            }
            else
            {
                m[magazine[i]]++;
            }
        }
        for(int i=0;i<ransomNote.length();i++)
        {
            if(m.find(ransomNote[i])==m.end())
            {
                return false;
            }
            else if(m[ransomNote[i]]==0)
            {
                return false;
            }
            else
            {
                m[ransomNote[i]]--;
            }
        }
        return true;
    }
};
