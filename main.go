package main

import "fmt"

type footSolider struct {
	speed  int
	sRange int
	weapon string
}

type upgrade interface {
	action()
}

func (f footSolider) action() {

	fmt.Println("Solider has a shooting range of", f.sRange)
}

func grenade(u upgrade) {
	u.action()
}

func main() {

	p1 := footSolider{
		speed:  2,
		sRange: 2,
		weapon: "Gun",
	}

	p2 := footSolider{
		speed:  1,
		sRange: 5,
		weapon: "Rocket",
	}

	p3 := footSolider{
		speed:  5,
		sRange: 3,
		weapon: "Magic",
	}
	grenade(p1)
	grenade(p2)
	grenade(p3)
}
