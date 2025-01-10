const FNV_OFFSET_BASIS_32 = 0x811c9dc5
const FNV_PRIME_32 = 0x01000193


type MyHashSet struct {
    data []*struct{}
    capacity int
    size int
}


func Constructor() MyHashSet {
    // If I'm seriously going to use 10^6 as capacity,
    // might as well use key as the index
    return MyHashSet{
        data: make([]*struct{}, 1000000),
        capacity: 1000000,
        size: 0,
    }
}


func (this *MyHashSet) Add(key int)  {
    index := hash(key) % this.capacity
    if this.data[index] != nil  {
        return
    }
    this.data[index] = &struct{}{}
    this.size++
}


func (this *MyHashSet) Remove(key int)  {
    index := hash(key) % this.capacity
    if this.data[index] != nil {
        this.data[index] = nil
        this.size--
    }
}


func (this *MyHashSet) Contains(key int) bool {
    index := hash(key) % this.capacity
    if this.data[index] != nil {
        return true
    }
    return false
}


func hash(key int) (hashResult int) {
    hashVal := FNV_OFFSET_BASIS_32
    prime := FNV_PRIME_32

    hashVal ^= key
    hashVal *= prime

    return hashVal
}
