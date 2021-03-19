package main

type ParkingSystem struct {
	lot [3]int
}

func Constructor(big, medium, small int) ParkingSystem {
	return ParkingSystem{
		lot: [3]int{
			big, medium, small,
		},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.lot[carType-1] != 0 {
		this.lot[carType-1]--
		return true
	}
	return false
}
