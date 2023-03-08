package main

import "fmt"

type transport interface {
	navigateToDestination()
}

type client struct {
}

func (c *client) startNavigation(t transport) {
	fmt.Println("client starteed navigation")
	t.navigateToDestination()
}

type boat struct {
}

func (b *boat) navigateToDestination() {
	fmt.Println("boat is navigating to island")
}
type car struct {
}

func (c *car) driveToDestination() {
	fmt.Println("car is driving to destination point")
}

// так как car способен исключительно ехать до точки предназначения,
// он не способен, примеру, плыть до неё, в случае, когда это  нужно.
// Допустим, нам требуется попасть на остров, в таком случае нам нужно 
// "адаптировать" car под плавание. Реализуем это

type carAdapter struct {
	c *car
}

func (car *carAdapter) navigateToDestination() {
	fmt.Println("adapting car to allow navigation")
	car.c.driveToDestination()
}

func main() {
	cl := new(client)
	b := new(boat)

	// b (boat struct) соответствует интерфейсу trasnport, 
	// проблем с навигации у cl (client struct) нет.
	cl.startNavigation(b)

	// но как видим, с (car struct) не способна реализовать навигацию
	// cl (client struct), так как не соответсвует интейрфейсу transport,
	// который способен navigateToDestination
	c := new(car)
	cl.startNavigation(c)

	// в этом случае, нам пригодится carAdapter, благодаря 
	// которому цель осуществима
	ca := carAdapter{c: c}
	cl.startNavigation(&ca)
}