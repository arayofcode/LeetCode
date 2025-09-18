type TaskHeap []*Task

type Task struct {
	user     int
	task     int
	priority int
    deleted  bool
}

func (h TaskHeap) Len() int {
    return len(h)
}

func (h TaskHeap) Less(i, j int) bool {
    if h[i].priority != h[j].priority {
        return h[i].priority > h[j].priority
    }
    return h[i].task > h[j].task
}

func (h TaskHeap) Swap(i, j int) { 
    h[i], h[j] = h[j], h[i] 
}

func (h *TaskHeap) Push(x interface{}) {
    *h = append(*h, x.(*Task))
}

func (h *TaskHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type TaskManager struct {
    taskMap map[int]*Task
    maxHeap *TaskHeap
}

func Constructor(tasks [][]int) TaskManager {
    taskMap :=  make(map[int]*Task)
    maxHeap := make(TaskHeap, 0)

    for _, item := range tasks {
        userId := item[0]
        taskId := item[1]
        priority := item[2]

        newTask := &Task{
            user: userId,
            task: taskId,
            priority: priority,
            deleted: false,
        }
        taskMap[taskId] = newTask
        heap.Push(&maxHeap, newTask)
    }
    return TaskManager{
        taskMap: taskMap,
        maxHeap: &maxHeap,
    }
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
    task := &Task{
        user: userId,
        task: taskId,
        priority: priority,
        deleted: false,
    }
    this.taskMap[taskId] = task

    heap.Push(this.maxHeap, task)
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
    this.taskMap[taskId].deleted = true

    newTask := &Task{
        user: this.taskMap[taskId].user,
        task: taskId,
        priority: newPriority,
        deleted: false,
    }

    this.taskMap[taskId] = newTask

    heap.Push(this.maxHeap, newTask)
}

func (this *TaskManager) Rmv(taskId int) {
    // Lazy delete
    if _, found := this.taskMap[taskId]; found {
        this.taskMap[taskId].deleted = true
    }

    delete(this.taskMap, taskId)
}

func (this *TaskManager) ExecTop() int {
    for this.maxHeap.Len() > 0 {
        topTask := heap.Pop(this.maxHeap).(*Task)
        if !topTask.deleted {
            userId := topTask.user
            taskId := topTask.task
            delete(this.taskMap, taskId)
            return userId
        }
    }
    return -1
}
