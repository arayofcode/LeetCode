type ongoingJourney struct {
	id           int
	startStation string
	startTime    int
}

type routeDetails struct {
    journeyCount int
    averageTime float64
}

type UndergroundSystem struct {
	ongoingJourneys map[int]ongoingJourney
    routeStats map[string]routeDetails
}

func Constructor() UndergroundSystem {
    return UndergroundSystem{
	    ongoingJourneys: make(map[int]ongoingJourney),
        routeStats: make(map[string]routeDetails),
    }
}

func getRouteKey(startStation string, endStation string) string {
    return fmt.Sprintf("%s---%s", startStation, endStation)
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
    this.ongoingJourneys[id] = ongoingJourney{
        id: id,
        startStation: stationName,
        startTime: t,
    }
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
    journey := this.ongoingJourneys[id]
    delete(this.ongoingJourneys, id)

    duration := float64(t - journey.startTime)
    key := getRouteKey(journey.startStation, stationName)
    if route, found := this.routeStats[key]; !found {
        this.routeStats[key] = routeDetails{
            journeyCount: 1,
            averageTime: float64(duration),
        }
    } else {
        journeyCount := float64(route.journeyCount)
        currentAverage := route.averageTime
        newAverage := ((currentAverage * journeyCount) + duration) / (journeyCount + 1.0)
        route.journeyCount += 1
        route.averageTime = newAverage
        this.routeStats[key] = route
    }
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    key := getRouteKey(startStation, endStation)
    return this.routeStats[key].averageTime
}
