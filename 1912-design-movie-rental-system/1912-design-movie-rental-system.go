type entry struct {
    shop   int
    movie  int
    price  int
    rented bool
}

type MovieRentingSystem struct {
    entries []*entry          // report()
    unrented map[int][]*entry // search(), indexed by movie ID
    indexes map[[2]int]*entry // rent() & drop(), indexed by [shop, movie]
}


func Constructor(n int, entries [][]int) MovieRentingSystem {
    system := MovieRentingSystem{
        entries:  make([]*entry, 0, n),
        unrented: make(map[int][]*entry),
        indexes:  make(map[[2]int]*entry),
    }
    for _, e := range entries {
        entry := &entry{
            shop:  e[0],
            movie: e[1],
            price: e[2],
        }
        system.entries = append(system.entries, entry)
        system.indexes[[2]int{e[0], e[1]}] = entry
        system.unrented[e[1]] = append(system.unrented[e[1]], entry)
    }

    // Sort unrented movies for search
    for movie := range system.unrented {
        slices.SortStableFunc(system.unrented[movie], func(a, b *entry) int {
            if a.price != b.price {
                return a.price - b.price
            }
            return a.shop - b.shop
        })
    }
    
    // Sort all entries for report
    slices.SortStableFunc(system.entries, func(a, b *entry) int {
        if a.price != b.price {
            return a.price - b.price
        }
        if a.shop != b.shop {
            return a.shop - b.shop
        }
        return a.movie - b.movie
    })

    return system
}


func (this *MovieRentingSystem) Search(movie int) []int {
    movies := this.unrented[movie]
    results := make([]int, 0, 5)
    for _, m := range movies {
        if len(results) == 5 {
            break
        }
        if !m.rented {
            results = append(results, m.shop)
        }
    }
    return results
}


func (this *MovieRentingSystem) Rent(shop int, movie int)  {
    this.indexes[[2]int{shop, movie}].rented = true
}


func (this *MovieRentingSystem) Drop(shop int, movie int)  {
    this.indexes[[2]int{shop, movie}].rented = false
}


func (this *MovieRentingSystem) Report() [][]int {
    results := make([][]int, 0, 5)
    for _, entry := range this.entries {
        if len(results) == 5 {
            break
        }
        if entry.rented {
            results = append(results, []int{entry.shop, entry.movie})
        }
    }
    return results
}
