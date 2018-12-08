package duck

import "fmt"

type IDuck interface {
	Quack() string
	Fly() string
}

type IFlyBehavior interface {
	Fly() string
}
type QuackBehavior func() string

type ConcreteDuck struct {
	QuackBehavior QuackBehavior
	FlyBehavior   IFlyBehavior
}

func (d ConcreteDuck) Quack() string {
	return d.QuackBehavior()
}

func (d ConcreteDuck) Fly() string {
	return d.FlyBehavior.Fly()
}

func NewDuck(behavior QuackBehavior, flyBehavior IFlyBehavior) *ConcreteDuck {
	var d = ConcreteDuck{behavior, flyBehavior}
	return &d
}

type noFlyBehavior struct {
}

func (f normalFlyBehavior) Fly() string {
	fmt.Println("Fly with wings")
	return "Fly with wings\n"
}

type normalFlyBehavior struct {
}

func (f noFlyBehavior) Fly() string {
	fmt.Println("I cant fly")
	return "I cant fly\n"
}

var squeakBehavior QuackBehavior = func() string {
	fmt.Println("squick")
	return "squick\n"
}
var crackBehavior QuackBehavior = func() string {
	fmt.Println("crack")
	return "crack\n"
}

type MallardDuck struct {
	ConcreteDuck
}

func NewMallardDuck() *MallardDuck {
	return &MallardDuck{ConcreteDuck{QuackBehavior: squeakBehavior, FlyBehavior: normalFlyBehavior{}}}
}
