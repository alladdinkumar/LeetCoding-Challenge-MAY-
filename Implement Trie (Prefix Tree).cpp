const int a_size=26;
struct Node{
    Node* child[a_size];
    bool isend;
    Node(){
        memset(child,NULL,sizeof(child));
        isend=false;
    }
};
class Trie {
    Node* root;
public:
    /** Initialize your data structure here. */
    Trie() {
       root=new Node();  
    }
    /** Inserts a word into the trie. */
    void insert(string word) {
        Node *head=root;
        for(int i=0;i<word.length();i++)
        {
            int index= word[i]-'a';
            if(head->child[index]==NULL)
                head->child[index]=new Node();
            head=head->child[index];
        }
        head->isend=true;
    }
    
    /** Returns if the word is in the trie. */
    bool search(string word) {
       Node *head=root;
        for(int i=0;i<word.length();i++)
        {
            int index= word[i]-'a';
            if(head->child[index]==NULL)
                return false;
            head=head->child[index];
        }
        return (head!=NULL && head->isend);         
    }
   
    /** Returns if there is any word in the trie that starts with the given prefix. */
   
    bool startsWith(string prefix) {
        Node *head=root;
        for(int i=0;i<prefix.length();i++)
        {
            int index= prefix[i]-'a';
            if(head->child[index]==NULL)
                return false;
            head=head->child[index];
        }
        return (head!=NULL);  
    }
};

/**
 * Your Trie object will be instantiated and called as such:
 * Trie* obj = new Trie();
 * obj->insert(word);
 * bool param_2 = obj->search(word);
 * bool param_3 = obj->startsWith(prefix);
 */
