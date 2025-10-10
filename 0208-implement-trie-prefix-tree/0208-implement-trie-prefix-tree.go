type node struct {
    children map[rune]*node
    end      bool
}


type Trie struct {
    root *node
}


func Constructor() Trie {
    return Trie{
        root: &node{},
    }    
}


func (this *Trie) Insert(word string)  {
    curr := this.root
    for _, chr := range word {
        if curr.children == nil {
            curr.children = make(map[rune]*node)
        }
        if _, found := curr.children[chr]; !found {
            curr.children[chr] = &node{}
        }
        curr = curr.children[chr]
    }
    curr.end = true
}


func (this *Trie) Search(word string) bool {
    curr := this.root
    for _, chr := range word {
        if curr.children[chr] == nil {
            return false
        }
        curr = curr.children[chr]
    }
    return curr.end
}


func (this *Trie) StartsWith(prefix string) bool {
    curr := this.root
    for _, chr := range prefix {
        if curr.children[chr] == nil {
            return false
        }
        curr = curr.children[chr]
    }
    return true
}
