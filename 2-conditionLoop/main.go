package main

import "fmt"

// "ali/methods"
// "ali/grade"

func main() {
	// guess.GeussMyNum()
	var script string = behaviour("behzad")
	println(script)
}

func behaviour(name string) string {
	fmt.Println("voroodi behaviour: " + name)
	var ret string = name + " khosh amadi"
	fmt.Println("khorooji: " + ret)
	return ret
}
