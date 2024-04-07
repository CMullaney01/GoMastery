package pointers

import "fmt"

type Player struct {
	health int
}

func takeDamageFromExplosion(player *Player) {
	fmt.Println("player is taking damage from explosion")
	explosionDmg := 10

	player.health -= explosionDmg
}

func Pointer() {
	player := &Player{
		health: 100,
	}

	fmt.Printf("before explosion %v\n", player)
	takeDamageFromExplosion(player)
	fmt.Printf("after explosion %v", player)
}
