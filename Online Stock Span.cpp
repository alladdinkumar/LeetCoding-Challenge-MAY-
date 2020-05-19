class StockSpanner {
    
public:
    
    vector<pair<int,int>>ss;
    StockSpanner() {
      
      ss.push_back({INT_MAX,1})  ;  
    }
    
    int next(int price) {
     
        int span=1;
        for(int i=ss.size()-1;i>0;)
        {
            if(ss[i].first<=price)
            {
                span+=ss[i].second;
                i-=ss[i].second;
            }
            else
            {
                break;
            }
            
        }
        ss.push_back({price,span});
        return span;
        
            
    }
};

/**
 * Your StockSpanner object will be instantiated and called as such:
 * StockSpanner* obj = new StockSpanner();
 * int param_1 = obj->next(price);
 */
