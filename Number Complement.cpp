class Solution {
public:
    int findComplement(int num) {
        if(num==0)
            return 1;
        int n = floor(log2(num))+1; 
        return ((unsigned)(1 << n) - 1) ^ num;
        
    }
};
