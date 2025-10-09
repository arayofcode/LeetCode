func minTime(skill []int, mana []int) int64 {
    finish := make([]int64, len(skill))
    for j := range mana {
        var curTime int64 = 0
        for i := range skill {
            if curTime < finish[i] {
                curTime = finish[i]
            }
            curTime += int64(skill[i]) * int64(mana[j])
        }
        finish[len(skill) - 1] = curTime
        for i := len(skill) - 2; i >= 0; i-- {
            finish[i] = finish[i+1] - int64(skill[i+1]) * int64(mana[j])
        }
    }
    return finish[len(finish) - 1]
}