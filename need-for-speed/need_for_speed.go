package speed

// TODO: define the 'Car' type struct
type Car struct{
    battery,batteryDrain,speed,distance int
}

type Track struct{
    distance int
    }

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
    return Car{
               battery:100,
               batteryDrain:batteryDrain,
               speed:speed,
               distance:0,
              }    
	panic("Please implement the NewCar function")
}

// TODO: define the 'Track' type struct

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{ distance:distance}
	panic("Please implement the NewTrack function")
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	if(car.batteryDrain >car.battery){
		return car
	}
    car.distance = car.speed
    car.battery= 100 - car.batteryDrain
    return car 
	panic("Please implement the Drive function")
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    if (track.distance/car.speed)*car.batteryDrain>car.battery{
        return false
    }else{
        return true
    }
	panic("Please implement the CanFinish function")
}
