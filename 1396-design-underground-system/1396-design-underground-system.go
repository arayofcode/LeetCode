type (
	card      int
	station   string
	timestamp int
    routeKey  string
)

type ongoingJourney struct {
	id           card
	startStation station
	startTime    timestamp
}

type routeDetails struct {
    journeyCount int
    averageTime float64
}

type UndergroundSystem struct {
	ongoingJourneys map[card]ongoingJourney
    routeStats map[routeKey]routeDetails
}

func Constructor() UndergroundSystem {
    return UndergroundSystem{
	    ongoingJourneys: make(map[card]ongoingJourney),
        routeStats: make(map[routeKey]routeDetails),
    }
}

func getRouteKey(startStation station, endStation station) routeKey {
    return routeKey(fmt.Sprintf("%s---%s", startStation, endStation))
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
    cardId := card(id)
    startStation := station(stationName)
    startTime := timestamp(t)

    this.ongoingJourneys[cardId] = ongoingJourney{
        id: cardId,
        startStation: startStation,
        startTime: startTime,
    }
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
    cardId := card(id)
    endStation := station(stationName)
    endTime := timestamp(t)

    journey := this.ongoingJourneys[cardId]
    delete(this.ongoingJourneys, cardId)

    duration := float64(endTime - journey.startTime)
    key := getRouteKey(journey.startStation, endStation)
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
    key := getRouteKey(station(startStation), station(endStation))
    return this.routeStats[key].averageTime
}
