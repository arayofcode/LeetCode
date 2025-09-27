type countNode struct {
    count  int
    // index from expiries
    index  int
}

type AuthenticationManager struct {
    ttl      int
    // token ID to expiry
    tokens   map[string]int
    // expiry time to count
    count    map[int]*countNode
    // maintaining all expiries
    expiries []int
}

func Constructor(timeToLive int) AuthenticationManager {
    return AuthenticationManager{
        ttl : timeToLive,
        tokens: make(map[string]int),
        count: make(map[int]*countNode),
        expiries: []int{},
    }
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int)  {
    expiry := currentTime + this.ttl
    this.tokens[tokenId] = expiry
    if _, found := this.count[expiry]; !found {
        // since generate will have strictly increasing currentTime
        // this is a sorted array
        this.expiries = append(this.expiries, expiry)
        this.count[expiry] = &countNode{
            count: 1,
            index: len(this.expiries) - 1,
        }
    } else {
        this.count[expiry].count++
    }
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int)  {
    if expiry, found := this.tokens[tokenId]; found && expiry > currentTime {
        this.count[expiry].count--
        newExpiry := currentTime + this.ttl
        this.tokens[tokenId] = newExpiry
        if _, found := this.count[newExpiry]; !found {
            this.expiries = append(this.expiries, newExpiry)
            this.count[newExpiry] = &countNode{
                count: 1,
                index: len(this.expiries) - 1,
            }
        } else {
            this.count[newExpiry].count++
        }
    }
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
    this.countCleanup()

    idx, _ := slices.BinarySearch(this.expiries, currentTime + 1)

    count := 0
    for _, expiry := range this.expiries[idx:] {
        count += this.count[expiry].count
    }
    return count
}

func (this *AuthenticationManager) countCleanup() {
    for i := 0; i < len(this.expiries); i++ {
        expiry := this.expiries[i]
        if this.count[expiry].count < 1 {
            this.expiries = append(this.expiries[:i], this.expiries[i+1:]...)
            for j := i; j < len(this.expiries); j++ {
                this.count[this.expiries[j]].index = j
            }
        }
    }
}