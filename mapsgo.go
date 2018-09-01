package main

import ("fmt")

func main(){
	fmt.Println("dfd")
	grades := make(map[string]float32)

	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 58

	fmt.Println(grades)

	TimGrades := grades["Timmy"]
	fmt.Println(TimGrades)

	delete(grades,"Timmy")
	fmt.Println(grades)

	for k,v := range grades {
		fmt.Println(k,":===:",v)
	}
} 