type ParkingSystem struct {
    Big int
    Medium int
    Small int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{
        Big: big,
        Medium: medium,
        Small: small,
    }
}


func (this *ParkingSystem) AddCar(carType int) bool {
    if carType == 1 {
        if this.Big < 1 {
            return false
        }
        this.Big--
        return true
    }
    if carType == 2 {
        if this.Medium < 1 {
            return false
        }
        this.Medium--
        return true
    }
    if carType == 3 {
        if this.Small < 1 {
            return false
        }
        this.Small--
        return true
    }
    return true
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */