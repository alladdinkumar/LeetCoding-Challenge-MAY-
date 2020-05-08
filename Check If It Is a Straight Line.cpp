class Solution {
public:
    bool collinear(vector<int> x1,vector<int> x2,vector<int> x3)
    {
        int a=x1[0]*(x2[1]-x3[1]) + x2[0]*(x3[1]-x1[1]) + x3[0]*(x1[1]-x2[1]);
        return a==0?true:false;
    }
    bool checkStraightLine(vector<vector<int>>& coordinates) {
        if(coordinates.size()==2)
        {
            return true;
        }
        vector<int> x1=coordinates[0],x2=coordinates[1];
        for(int i=2;i<coordinates.size();i++)
        {
            if(!collinear(x1,x2,coordinates[i]))
                return false;
        }
        return true;
        
    }
};
