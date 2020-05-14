const(
    a_size=26
)
type Node struct{
    child[a_size]*Node
    isend bool
}

type Trie struct {
    root *Node
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
        root: &Node{},
    }
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    head :=this.root;
    for i:=0;i<len(word);i++{
        var index int = int(word[i])-int('a')
        if head.child[index]==nil{
            head.child[index]= &Node{}
        }
                
        head=head.child[index]
    }
    head.isend=true;
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    head:=this.root
    for i:=0;i<len(word);i++{
        var index int = int(word[i])-int('a')
        if head.child[index]==nil{
            return false
        }      
        head=head.child[index]
    }
    return head!=nil && head.isend
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    head:=this.root
    for i:=0;i<len(prefix);i++{
        var index int = int(prefix[i])-int('a')
        if head.child[index]==nil{
            return false
        }      
        head=head.child[index]
    }
    return head!=nil 
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */