class Solution {
public:
   
    vector<int> countBits(int num) {
        vector<int>vec;
        vector<int>tmp;
        vec.push_back(0);
        for(int i=1;i<=num;)
        {
            for(auto x: vec)
            {
                tmp.push_back(x+1);
                i++;
                if(i>num)
                    break;
            }
            vec.insert(vec.end(),tmp.begin(),tmp.end());
            tmp.clear();
        }
        return vec;
        
    }
};
