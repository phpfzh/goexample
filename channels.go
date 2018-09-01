package main

import ("fmt")


func foo(c chan int,someValue int){
	c <- someValue * 5
}

func main(){
	fmt.Println("H")

	fooVal := make(chan int)
	fmt.Println("fooVal",fooVal)

	go foo(fooVal,5)
	go foo(fooVal,3)

	//v1 := <-fooVal
	//v2 := <-fooVal

	v1,v2 := <-fooVal,<-fooVal

	fmt.Println("H",v1)
	fmt.Println("H",v2)
}