type packet struct {
	source      int
	destination int
	timestamp   int
}

type Router struct {
	packets       chan *packet
	queuedPackets map[string]struct{}
    packetsByDest map[int][]int
}

func Constructor(memoryLimit int) Router {
	return Router{
		packets:       make(chan *packet, memoryLimit),
		queuedPackets: make(map[string]struct{}, memoryLimit),
		packetsByDest: make(map[int][]int, memoryLimit),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	pkt := &packet{
		source:      source,
		destination: destination,
		timestamp:   timestamp,
	}

	key := getKey(pkt)
    if _, found := this.queuedPackets[key]; found {
		return false
	}

    if len(this.packets) == cap(this.packets) {
        this.removeHead()
    }

    this.packets <- pkt
    this.queuedPackets[key] = struct{}{}
    this.packetsByDest[destination] = append(this.packetsByDest[destination], timestamp)
    return true
}

func (this *Router) ForwardPacket() []int {
    if len(this.packets) == 0 {
        return []int{}
    }
    node := this.removeHead()
    return []int{node.source, node.destination, node.timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
    nodes := this.packetsByDest[destination]
    left := sort.SearchInts(nodes, startTime)
    right := sort.SearchInts(nodes, endTime+1)
    return right - left
}

func getKey(pkt *packet) string {
    return fmt.Sprintf("%d-%d-%d", pkt.source, pkt.destination, pkt.timestamp)
}

func (this *Router) removeHead() *packet {
    node := <-this.packets
    key := getKey(node)
    delete(this.queuedPackets, key)
    nodes := this.packetsByDest[node.destination]
    i := sort.SearchInts(nodes, node.timestamp)
    this.packetsByDest[node.destination] = append(nodes[:i], nodes[i+1:]...)
    return node
}