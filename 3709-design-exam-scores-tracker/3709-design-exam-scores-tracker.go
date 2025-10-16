type ExamTracker struct {
    // Exam times
    timestamps []int
    // Cumulative score until time i
    cumulative []int64
}


func Constructor() ExamTracker {
    return ExamTracker{
        timestamps: []int{},
        cumulative: []int64{},
    }
}


func (this *ExamTracker) Record(time int, score int)  {
    this.timestamps = append(this.timestamps, time)
    total := int64(score)
    if len(this.cumulative) > 0 {
        total += this.cumulative[len(this.cumulative) - 1]
    }
    this.cumulative = append(this.cumulative, total)
}


func (this *ExamTracker) TotalScore(startTime int, endTime int) int64 {
    i, _ := slices.BinarySearch(this.timestamps, startTime)
    j, found := slices.BinarySearch(this.timestamps, endTime)
    if found {
        j++
    }

    if i == j || i >= len(this.timestamps) {
        return 0
    }

    result := this.cumulative[j-1]
    if i > 0 {
        result -= this.cumulative[i-1]
    }
    return result
}
