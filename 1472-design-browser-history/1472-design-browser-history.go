type Node struct {
    index   int
    url     string
    prev    *Node
    next    *Node
}

type BrowserHistory struct {
    head     *Node
    curr     *Node
    visited  []*Node
}


func Constructor(homepage string) BrowserHistory {
    node := &Node{
        index:   0,
        url:     homepage,
        prev:    nil,
        next:    nil,
    }

    head := node
    curr := node
    visited := []*Node{node}

    return BrowserHistory{
        head:     head,
        curr:     curr,
        visited:  visited,
    }
}


func (this *BrowserHistory) Visit(url string)  {
    node := &Node{
        index: this.curr.index + 1,
        url:   url,
        prev:  this.curr,
        next:  nil,
    }
    this.visited = append(this.visited[:this.curr.index + 1], node)
    this.curr.next = node
    this.curr = node
}


func (this *BrowserHistory) Back(steps int) string {
    newCurrIndex := this.curr.index - steps
    if newCurrIndex < 0 {
        newCurrIndex = 0
    }
    this.curr = this.visited[newCurrIndex]
    return this.curr.url
}


func (this *BrowserHistory) Forward(steps int) string {
    newCurrIndex := this.curr.index + steps
    if newCurrIndex >= len(this.visited) {
        newCurrIndex = len(this.visited) - 1
    }
    this.curr = this.visited[newCurrIndex]
    return this.curr.url
}
