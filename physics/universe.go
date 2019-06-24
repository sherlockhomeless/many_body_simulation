package physics

import (
	"fmt"
	"github.com/sherlockhomeless/many_body_simulation/printer"
	"math/rand"
)

var Debug bool
var Worker int

type Star struct {
	ID   int
	X    float64
	Y    float64
	Mass uint32
}

type Universe struct {
	Stars        []Star
	StarCount    int
	Time         uint32
	X_Dim, Y_Dim uint32
	MaxMass      uint32
	HasFinished  bool
	TimeConvergence int

}

func (u *Universe) FillUniverse(){
	for i := 0; i < u.StarCount; i++{
		u.addStar(i)
	}
}

func (u *Universe) addStar(id int){
	newStar := Star{ID: id, X: float64(rand.Intn(int(u.X_Dim))), Y: float64(rand.Intn(int(u.Y_Dim))), Mass: uint32(rand.Intn(int(u.MaxMass)))}
	u.Stars = append(u.Stars, newStar)

	if (Debug){
		printer.PrintMeta(fmt.Sprintf("[CREATE] New Star %d with (%f,%f,%d)", newStar.ID, newStar.X, newStar.Y, newStar.Mass))
	}

}

func (u *Universe) RunSimulationRound(worker int) {
	for

	u.Time++

}