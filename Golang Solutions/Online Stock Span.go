type StockSpanner struct {
   ss1 []int
   ss2 []int
}


func Constructor() StockSpanner {
    var s StockSpanner
    s.ss1=append(s.ss1,9999999)
    s.ss2=append(s.ss2,1)
    return s
}


func (this *StockSpanner) Next(price int) int {
    span:=1
    for i:=len(this.ss1)-1;i>0;{
        if this.ss1[i]<=price{
            span+=this.ss2[i];
            i-=this.ss2[i]
        }else{
            break;
        }
        
    }
    this.ss1=append(this.ss1,price)
    this.ss2=append(this.ss2,span)
    return span
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */