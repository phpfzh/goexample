package main

import "fmt"

const usixteenbitmax float64= 65535
const kmh_multiple float64 = 1.60934

type car struct{

	gas_pedal  uint16 //min 0 max 65535
	brake_pedal uint16
	steering_wheel  int16 //-32K - +32K
	top_speed_kmd float64
}

func (c car) kmh() float64{
	return float64(c.gas_pedal) * (c.top_speed_kmd/usixteenbitmax)
}


func main(){
	 a_car := car{gas_pedal:22341,
			brake_pedal:0,
			steering_wheel:12561,
			top_speed_kmd:225.0}

	 fmt.Println(a_car.gas_pedal)
	 fmt.Println(a_car.kmh())
	 fmt.Println(float64(a_car.gas_pedal))

	 fmt.Println(a_car.top_speed_kmd/usixteenbitmax)

}