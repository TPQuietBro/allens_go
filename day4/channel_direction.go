package main

import (
	"fmt"
)

func main (){
	ch :=make(chan int , 10)
	sendData(ch, 12)
	recieveData(ch)

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case task1 := <-ch:
	// 		fmt.Println(task1)
	// 	case task2 := <-ch1:
	// 		fmt.Println(task2)
	// 	}
	// }
}
// 控制通道方向, 只发送
func sendData(ch chan<-  int, data int)  {
	fmt.Printf("data send: %v \n", data)
	ch <- data
}
// 控制通道方向, 只接收
func recieveData(ch <-chan int) {
	fmt.Printf("data recieve: %v", <-ch)
}