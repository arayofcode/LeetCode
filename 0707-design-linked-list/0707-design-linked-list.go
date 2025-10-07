type node struct {
    val int
    next *node
}

type MyLinkedList struct {
    head *node
    size int
}


func Constructor() MyLinkedList {
    return MyLinkedList{
        head: nil,
        size: 0,
    }
}


func (this *MyLinkedList) Get(index int) int {
    if index < 0 || index >= this.size {
        return -1
    }
    
    curr := this.head
    for index > 0 {
        index--
        curr = curr.next
    }
    return curr.val
}


func (this *MyLinkedList) AddAtHead(val int)  {
    this.AddAtIndex(0, val)
}


func (this *MyLinkedList) AddAtTail(val int)  {
    this.AddAtIndex(this.size, val)
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
    if index < 0 || index > this.size {
        return
    }

    newNode := &node{
        val:  val,
        next: nil,
    }
    if index == 0 {
        newNode.next = this.head
        this.head = newNode
    } else {
        curr := this.head
        for index > 1 {
            index--
            curr = curr.next
        }

        newNode.next = curr.next
        curr.next = newNode
    }
    this.size++
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
    if index < 0 || index >= this.size {
        return
    }
    if index == 0 {
        this.head = this.head.next
    } else {
        curr := this.head
        for index > 1 {
            index--
            curr = curr.next
        }

        if curr.next != nil {
            curr.next = curr.next.next
        }
    }
    this.size--
}
