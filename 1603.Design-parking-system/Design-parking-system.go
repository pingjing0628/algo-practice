package problem1603

type ParkingSystem struct {
	slot[3]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{[3]int{big, medium, small}}
}


func (this *ParkingSystem) AddCar(carType int) bool {
	if this.slot[carType - 1] == 0 {
		return false
	}

	this.slot[carType - 1]--

	return true
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */