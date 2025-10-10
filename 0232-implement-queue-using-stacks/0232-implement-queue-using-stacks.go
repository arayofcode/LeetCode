type stack struct {
    stack []int
}

func (s *stack) push(x int) {
    s.stack = append(s.stack, x)
}

func (s *stack) pop() int {
    val := s.stack[len(s.stack) - 1]
    s.stack = s.stack[:len(s.stack) - 1]
    return val
}

func (s *stack) peek() int {
    return s.stack[len(s.stack) - 1]
}

func (s *stack) empty() bool {
    return len(s.stack) == 0
}

type MyQueue struct {
    s1 stack
    s2 stack
}


func Constructor() MyQueue {
    return MyQueue{
        s1: stack{
            stack: []int{},
        },
        s2: stack{
            stack: []int{},
        },
    }
}


func (this *MyQueue) Push(x int)  {
    for !this.s1.empty() {
        val := this.s1.pop()
        this.s2.push(val)
    }
    this.s2.push(x)
    for !this.s2.empty() {
        val := this.s2.pop()
        this.s1.push(val)
    }
}


func (this *MyQueue) Pop() int {
    return this.s1.pop()
}


func (this *MyQueue) Peek() int {
    return this.s1.peek()
}


func (this *MyQueue) Empty() bool {
    return this.s1.empty()
}
