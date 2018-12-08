package main

import (
		"github.com/bariskas/go-course/workshop1/task3/duck"
)

func playWithDuck(duck duck.IDuck) {
	duck.Quack()
	duck.Fly()
}

func main() {
	playWithDuck(duck.NewMallardDuck())
}
