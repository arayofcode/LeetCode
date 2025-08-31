type travel struct {
	station string
	time    int
}

type travelStats struct {
	count   int
	average float64
}

type UndergroundSystem struct {
	activeTravels map[int]*travel
	travels       map[string]*travelStats
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		activeTravels: make(map[int]*travel),
		travels:       make(map[string]*travelStats),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.activeTravels[id] = &travel{
		station: stationName,
		time:    t,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
    travelDetails := this.activeTravels[id]
    key := fmt.Sprintf("%s-%s", travelDetails.station, stationName)
    duration := float64(t - travelDetails.time)

    if route, found := this.travels[key]; !found {
        this.travels[key] = &travelStats{
            count: 1,
            average: duration,
        }
    } else {
        count := float64(route.count)
        currAverage := route.average
        this.travels[key].count++
        this.travels[key].average = (currAverage * count + duration) / (count + 1.0)
    }
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    key := fmt.Sprintf("%s-%s", startStation, endStation)
    return this.travels[key].average
}
