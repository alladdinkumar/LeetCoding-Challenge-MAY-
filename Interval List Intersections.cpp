class Solution {
public:
    vector<vector<int>> intervalIntersection(vector<vector<int>>& A, vector<vector<int>>& B) {
        vector<vector<int>>vec;
        for(int i=0,j=0;i<A.size() && j<B.size();)
        {
            if(!(A[i][0] > B[j][1] || A[i][1] < B[j][0]))
            {
                int l=max(A[i][0],B[j][0]);
                int r=min(A[i][1],B[j][1]);
                vec.push_back({l,r});
            }
            if(i<A.size()-1 && B[j][1]>=A[i+1][0])
            {
                i++;
                continue;
            }
            if(j<B.size()-1 && A[i][1]>=B[j+1][0])
            {
                j++;
                continue;
            }
            i++;
            j++;
        
        }
        return vec;
        
    }
};
