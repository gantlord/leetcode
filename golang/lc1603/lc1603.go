package lc1603

type ParkingSystem struct {
	bigSpaces, mediumSpaces, smallSpaces int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (p *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 && p.bigSpaces > 0 {
		p.bigSpaces--
		return true
	} else if carType == 2 && p.mediumSpaces > 0 {
		p.mediumSpaces--
		return true
	} else if carType == 3 && p.smallSpaces > 0 {
		p.smallSpaces--
		return true

	}
	return false
}
