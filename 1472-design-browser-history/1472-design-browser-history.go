type Node struct {
    url     string
    prev    *Node
    next    *Node
}

type BrowserHistory struct {
    head     *Node
    curr     *Node
}


func Constructor(homepage string) BrowserHistory {
    node := &Node{
        url:     homepage,
        prev:    nil,
        next:    nil,
    }

    head := node
    curr := node

    return BrowserHistory{
        head:     head,
        curr:     curr,
    }
}


func (this *BrowserHistory) Visit(url string)  {
    node := &Node{
        url:  url,
        prev: this.curr,
        next: nil,
    }
    this.curr.next = node
    this.curr = node
}


func (this *BrowserHistory) Back(steps int) string {
    for steps > 0 {
        if this.curr == this.head {
            break
        }
        this.curr = this.curr.prev
        steps--
    }
    return this.curr.url
}


func (this *BrowserHistory) Forward(steps int) string {
    for steps > 0 {
        if this.curr.next == nil {
            break
        }
        this.curr = this.curr.next
        steps--
    }
    return this.curr.url
}
