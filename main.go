package main

import "pokeapi/di"

func main() {
	err := di.Start()
	if err != nil {
		panic(err)
	}
}
