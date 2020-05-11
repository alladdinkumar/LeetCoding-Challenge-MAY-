class Solution {
public:
    vector<vector<int>> floodFill(vector<vector<int>>& image, int sr, int sc, int newColor) {
        vector<vector<int>> directions={{0,1},{1,0},{0,-1},{-1,0}};
        queue<pair<int,int>> q;
        bitset<51>visited[51];
        pair<int,int>p,k;
        p.first=sr;
        p.second=sc;
        q.push(p);
        int val=image[sr][sc];
        while(!q.empty())
        {
            p=q.front();
            q.pop();
            
            image[p.first][p.second]=newColor;
            visited[p.first][p.second]=1;
            for(auto x : directions)
            {
                if(p.first+x[0]>=0 && p.first+x[0]<image.size() && p.second+x[1]>=0 && p.second+x[1]<image[0].size() ){
                   if(image[p.first+x[0]][p.second+x[1]]==val && visited[p.first+x[0]][p.second+x[1]]==0)
                    {
                        k.first=p.first+x[0];
                        k.second=p.second+x[1];
                        q.push(k);
                    } 
                }
            } 
        }
        return image;
    }
};
