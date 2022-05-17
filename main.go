package main

import "container/list"

func main() {
	a := list.New()
	a.PushFront(1)
	b := a.Back().Value.(int)
	print(b)
}
