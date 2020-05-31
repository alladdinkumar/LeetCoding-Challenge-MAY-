class Solution {
public:
    float euclidean(vector<int> x)
    {
        return sqrt(pow(x[0],2)+pow(x[1],2));
    }
    vector<vector<int>> kClosest(vector<vector<int>>& points, int K) {
        int cnt=0;
        map<float,vector<int>> m;
        map<float,vector<int>>::iterator itr;
        for(int i=0;i<points.size();i++)
        {
            auto vec=points[i];
            float dist=euclidean(vec);
            if(cnt<K)
            {
                m[dist].push_back(i);
                cnt++;
            }
            else
            {
                itr = m.end();
                itr--;
                if(dist<itr->first)
                {
                    if(itr->second.size()>1)
                        itr->second.erase(itr->second.begin());
                    else
                        m.erase(itr);
                    m[dist].push_back(i); 
                }
            }
        }
        vector<vector<int>> sol;
        for(auto i:m)
        {
            for(auto x:i.second)
            {
                sol.push_back(points[x]);
            }
        }
        return sol;
        
    }
};
