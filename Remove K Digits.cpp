class Solution {
public:
    string removeKdigits(string num, int k) {
        if(k==num.length())
            return "0";
        int i=0;
        while(k--)
        {
            i=i-1 > 0 ? i-1 : 0;
            while(i<num.length()-1 && num[i] <= num[i+1])
                i++;
            num.erase(num.begin()+i);
        }
        while(num[0]=='0')
            num.erase(num.begin());
        if (num=="")
            return "0";
        return num;
    }
};
