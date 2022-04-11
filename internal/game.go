package internal

import (
	"fmt"
	"strconv"
)

type mine struct {
	x int
	y int
	z int
}

type Sector struct {
	Id        int  `json:"id"`
	X         int  `json:"x"`
	Y         int  `json:"y"`
	Z         int  `json:"z"`
	Radiation int  `json:"radiation"`
	Marked    bool `json:"marked"`
}

type Game struct {
	Id      string   `json:"id"`
	Sectors []Sector `json:"sectors"`
	State   string   `json:"state"`
	mines   []mine
}

func (sector Sector) print() {
	println(fmt.Sprintf("Sector[id=%d,marked=%v]", sector.Id, sector.Marked))
}

func NewGame(id string, scale int) Game {
	game := Game{Id: id, State: "PLAY", Sectors: make([]Sector, 0)}
	nextId := 0
	for x := 0; x < scale; x++ {
		for y := 0; y < scale; y++ {
			for z := 0; z < scale; z++ {
				sector := Sector{Id: nextId, X: x, Y: y, Z: z, Radiation: -1, Marked: false}
				game.Sectors = append(game.Sectors, sector)
				nextId++
			}
		}
	}
	game.mines = []mine{{x: 1, y: 1, z: 1}}
	return game
}

func (game *Game) Reveal(sectorId string) {
	index, _ := strconv.Atoi(sectorId)
	sector := &game.Sectors[index]
	sector.Radiation = 1
	sector.print()
}

func (game *Game) Mark(sectorId string) {
	index, _ := strconv.Atoi(sectorId)
	sector := &game.Sectors[index]
	sector.Marked = true
	sector.print()
}