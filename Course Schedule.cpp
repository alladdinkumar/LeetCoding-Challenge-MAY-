class Solution {
public:
    bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {
        vector<int>pre_count(numCourses,0);
        vector<vector<int>>m(numCourses);
        for(auto x: prerequisites){
           pre_count[x[0]]++;
           m[x[1]].push_back(x[0]);
        }
        stack<int>s;
        for(int i=0;i<numCourses;i++)
        {
            if(pre_count[i]==0)
                s.push(i);
        }
        int count=0;
        while(!s.empty())
        {
            count++;
            int x= s.top();
            s.pop();
            for(auto next: m[x]){
                pre_count[next]--;
                if(pre_count[next]==0)
                    s.push(next); 
            }
        }
        return count==numCourses;
    }
};
