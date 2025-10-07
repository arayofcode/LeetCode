type node struct {
    val int
    prev *node
    next *node
}

type MyLinkedList struct {
    head  *node
    tail  *node
    nodes []*node
}


func Constructor() MyLinkedList {
    return MyLinkedList{
        head:  nil,
        tail:  nil,
        nodes: []*node{},
    }
}


func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= len(this.nodes) {
        return -1
    }
    return this.nodes[index].val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    newNode := &node{
        val: val,
        prev: nil,
        next: this.head,
    }

    if this.head != nil {
        this.head.prev = newNode
    } else {
        this. tail = newNode
    }
    this.head = newNode
    this.nodes = append([]*node{newNode}, this.nodes...)
}


func (this *MyLinkedList) AddAtTail(val int)  {
    newNode := &node{
        val:  val,
        prev: this.tail,
        next: nil,
    }

    if this.tail != nil {
        this.tail.next = newNode
    } else {
        this.head = newNode
    }

    this.tail = newNode
    this.nodes = append(this.nodes, newNode)
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    n := len(this.nodes)
    if index < 0 || index > n {
        return
    }

    newNode := &node{
        val: val,
        prev: nil,
        next: nil,
    }
    
    if index == 0 {
        this.AddAtHead(val)
    } else if index == n {
        this.AddAtTail(val)
    } else {
        this.nodes[index].prev = newNode
        newNode.next = this.nodes[index]
        this.nodes[index - 1].next = newNode
        newNode.prev = this.nodes[index - 1]
        this.nodes = append(this.nodes[:index], append([]*node{newNode}, this.nodes[index:]...)...)
    }
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= len(this.nodes) {
        return
    }

    if index == 0 {
        this.head = this.head.next
        if this.head != nil {
            this.head.prev = nil
        } else {
            this.tail = nil
        }
        this.nodes = this.nodes[1:]
    } else if index == len(this.nodes) - 1 {
        this.tail = this.tail.prev
        if this.tail != nil {
            this.tail.next = nil
        } else {
            this.head = nil
        }
        this.nodes = this.nodes[:len(this.nodes)-1]
    } else {
        this.nodes[index-1].next = this.nodes[index].next
        this.nodes[index+1].prev = this.nodes[index].prev
        this.nodes = append(this.nodes[:index], this.nodes[index+1:]...)
    }
}
