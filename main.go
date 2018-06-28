package main

import (
	life2 "github.com/adewaleafolabi/game-of-life/life"
	"log"
)

func main() {
	life,err:= life2.NewGame(4,5)
	if err!=nil{
		log.Fatal(err)
	}
	life.RunSimulation()
}
