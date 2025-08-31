type station string
type card int
type timestamp int
type journeyKey card

// of format "{startStation}-{endStation}"
type journeysMapKey string

type UndergroundSystem struct {
	journeys        map[journeysMapKey][]journey
	ongoingJourneys map[journeyKey]journey
}

type journey struct {
	id           card
	startStation station
	startTime    timestamp
	endStation   station
	endTime      timestamp
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		journeys:        make(map[journeysMapKey][]journey),
		ongoingJourneys: make(map[journeyKey]journey),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.ongoingJourneys[journeyKey(id)] = journey{
		id:           card(id),
		startStation: station(stationName),
		startTime:    timestamp(t),
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	customerJourney := this.ongoingJourneys[journeyKey(id)]
	customerJourney.endStation = station(stationName)
	customerJourney.endTime = timestamp(t)

    mapKey := journeysMapKey(fmt.Sprintf("%s-%s", customerJourney.startStation, customerJourney.endStation))
    this.journeys[mapKey] = append(this.journeys[mapKey], customerJourney)
    delete(this.ongoingJourneys, journeyKey(id))
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    key := journeysMapKey(fmt.Sprintf("%s-%s", startStation, endStation))
    journeysFound := this.journeys[key]
    average := float64(0)
    for i, journey := range journeysFound {
        travelTime := journey.endTime - journey.startTime
        average = ((average * float64(i)) + float64(travelTime)) / float64(i+1)
    }
    return average
}
