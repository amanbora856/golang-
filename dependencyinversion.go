package main

import "fmt"

type socket interface {
	connect()
}

type connection struct {
	device socket
}
type laptopcharger struct{}

func (l laptopcharger) connect() {
	fmt.Println("laptop charger is connected")
}

type phonecharger struct{}

func (p phonecharger) connect() {
	fmt.Println("Phone charger is connected")
}

func (c connection) start() {
	c.device.connect()
}
func main() {

	laptop := laptopcharger{}
	phone := phonecharger{}
	laptopconnection := connection{device: laptop}
	phoneconnection := connection{device: phone}
	laptopconnection.start()
	phoneconnection.start()

}
