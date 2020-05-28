class Solution {
public:
    bool possibleBipartition(int N, vector<vector<int>>& dislikes)
    {
        g_=vector<vector<int>>(N);
        for(auto vec:dislikes)
        {
            int a=vec[0];
            int b=vec[1];
            g_[a-1].push_back(b-1);
            g_[b-1].push_back(a-1);       
        }
        colors_=vector<int>(N,0);
        for(int i=0;i<N;++i)
        {
            if(colors_[i]==0 && !dfs(i,1))
                return false;
        }
        return true;
    }
    private:
    vector<vector<int>>g_;
    vector<int>colors_;
    bool dfs( int cur,int color)
    {
        colors_[cur]=color;
        for(int next: g_[cur]){
            if(colors_[next]==color)
                return false;
            if(colors_[next]==0 && !dfs(next,-color))
                return false;
        }
        return true;
    }
};
