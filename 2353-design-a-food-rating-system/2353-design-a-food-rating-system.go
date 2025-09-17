/* 
Food attributes:
- Name
- Type (cuisine)
- rating

Operations:
- Modify rating for given food name
    - Can be done in O(1) if we can search by name in O(1)
    - Hash Table/ Map
- For a given cuisine type, return food name with highest rating
    - Store cuisine type to food mapping
        - For each cuisine type, calculate highest rating food
        - Takes O(n)
    - Instead of storing cuisine type -> food mapping, store cuisine type -> food with highest rating
        - Can be compared instantly during modification by doing lexicographic comparison of the current max food name with new food name
        - O(1) using hash table
*/

type food struct {
    name string
    cuisine string
    rating int
}

type FoodRatings struct {
    foodMap map[string]*food
    cuisineMap map[string]*food
    cuisineFoods map[string][]*food
}


func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
    foodMap := make(map[string]*food)
    cuisineMap :=  make(map[string]*food)
    cuisineFoods := make(map[string][]*food)

    for i := 0; i < len(foods); i++ {
        foodName := foods[i]
        cuisine := cuisines[i]

        givenFood := food{
            name: foodName,
            cuisine: cuisine,
            rating: ratings[i],
        }

        foodMap[foodName] = &givenFood
        // two cases for adding to ratings map. For a given cuisine:
        // If ratings exist, lexicographic comparison
        // If ratings don't exist, add this food as maximum
        if food, found := cuisineMap[cuisine]; !found {
            cuisineMap[cuisine] = &givenFood
        } else {
            if givenFood.rating > food.rating {
                cuisineMap[cuisine] = &givenFood
            } else if givenFood.rating == food.rating && givenFood.name < food.name {
                cuisineMap[cuisine] = &givenFood
            }
        }

        cuisineFoods[cuisine] = append(cuisineFoods[cuisine], &givenFood)
    }
    return FoodRatings{
        foodMap: foodMap,
        cuisineMap: cuisineMap,
        cuisineFoods: cuisineFoods,
    }
}


func (this *FoodRatings) updateMaxCuisine(givenFood *food, oldRating int) {
    cuisine := givenFood.cuisine

    // What happens when givenFood was already the one with highest rating?
    // recalculate maximum
    if givenFood.name == this.cuisineMap[cuisine].name {
        for _, food := range this.cuisineFoods[cuisine] {
            currMax := this.cuisineMap[cuisine]
            if food.rating > currMax.rating {
                this.cuisineMap[cuisine] = food
            } else if food.rating == currMax.rating && food.name < currMax.name {
                this.cuisineMap[cuisine] = food
            }
        }
        return
    }

    current := this.cuisineMap[cuisine]
    if givenFood.rating > current.rating {
        this.cuisineMap[givenFood.cuisine] = givenFood
    } else if givenFood.rating == current.rating && givenFood.name < current.name {
        this.cuisineMap[givenFood.cuisine] = givenFood
    }
}

func (this *FoodRatings) ChangeRating(food string, newRating int)  {
    oldRating := this.foodMap[food].rating
    this.foodMap[food].rating = newRating
    this.updateMaxCuisine(this.foodMap[food], oldRating)
}


func (this *FoodRatings) HighestRated(cuisine string) string {
    return this.cuisineMap[cuisine].name
}
