func earliestTime(tasks [][]int) int {
    minTaskEnd := tasks[0][1] + tasks[0][0]
    for i := 1; i < len(tasks); i++ {
        taskEnd := tasks[i][1] + tasks[i][0]
        if taskEnd < minTaskEnd {
            minTaskEnd = taskEnd
        }
    }
    return minTaskEnd
}