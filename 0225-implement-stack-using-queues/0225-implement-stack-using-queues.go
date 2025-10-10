type MyStack struct {
    q1 chan int
    q2 chan int
}


func Constructor() MyStack {
    return MyStack{
        q1: make(chan int, 100),
        q2: make(chan int, 100),
    }
}


func (this *MyStack) Push(x int)  {
    this.q2 <- x
    for _ = range len(this.q1) {
        val := <-this.q1
        this.q2 <- val
    }
    this.q1, this.q2 = this.q2, this.q1
}


func (this *MyStack) Pop() int {
    return <-this.q1
}


func (this *MyStack) Top() int {
    val := this.Pop()
    this.Push(val)
    return val
}


func (this *MyStack) Empty() bool {
    return len(this.q1) == 0
}
