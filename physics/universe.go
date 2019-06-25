package physics

import (
	"fmt"
	"github.com/sherlockhomeless/many_body_simulation/printer"
	"math"
	"math/rand"
)

var Debug bool
var Worker int

type Star struct {
	ID   int
	X    float64
	Y    float64
	Mass int
}

type Universe struct {
	Stars        []Star
	StarCount    int
	Time         uint32
	X_Dim, Y_Dim uint32
	MaxMass      uint32
	TimeConvergence int
	MinimalDistance float64
	G float64
	hasFinished  bool
	mergedStarsChan chan int

}

func (u *Universe) FillUniverse(){
	u.mergedStarsChan = make(chan int, u.StarCount)

	for i := 0; i < u.StarCount; i++{
		u.addStar(i)
	}
}

func (u *Universe) addStar(id int){
	newStar := Star{ID: id, X: float64(rand.Intn(int(u.X_Dim))), Y: float64(rand.Intn(int(u.Y_Dim))), Mass: rand.Intn(int(u.MaxMass))}
	u.Stars = append(u.Stars, newStar)

	if (Debug){
		printer.PrintMeta(fmt.Sprintf("[CREATE] New Star %d with (%f,%f,%d)", newStar.ID, newStar.X, newStar.Y, newStar.Mass))
	}

}

func (u *Universe) RunSimulation(){

	// build communication-channel for all workers
	var chanArr []chan Star = make([]chan Star, Worker)
	for i := 0; i < Worker; i++{
		chanArr[i] = make(chan Star, u.StarCount)
	}


	for true{
		// fill channels with current universe stars
		for _, channel := range chanArr{
			for _, star := range u.Stars{
				channel <- star
			}
		}

		// run simulation for workers
		for i := 0; i < Worker; i++ {
			go u.runSimulationRound(i, chanArr[i])
		}

		//TODO: just continous executing ==> some kind of sync must be here
		// read results from workers
		for i := 0; i < len(u.Stars); i++{
			u.Stars[i] = <- chanArr[i%Worker]
		}

		for merged_index := range u.mergedStarsChan{
			u.Stars = append(u.Stars[:merged_index], u.Stars[merged_index+1:]...)
		}

		if(Debug){
			printer.PrintMovement(fmt.Sprintf("[ROUND%d] %d Stars left", u.Time, u.StarCount))
		}
		u.Time++
	}


}


func (u *Universe) runSimulationRound(worker int, starChan chan Star) {

	for x := 0; x < len(u.Stars); x += Worker{
			starChan <- u.calculateMovement(x)
	}
}

func (u *Universe) calculateMovement(indexForStar int) Star{
	var distanceBtwStars float64
	var gravitationalForce float64
	var directions [2]float64
	var x1, x2, y1, y2 float64
	var m1, m2 int

	x1 = u.Stars[indexForStar].X
	y1 = u.Stars[indexForStar].Y
	m1 = u.Stars[indexForStar].Mass

	for i := 0; i < len(u.Stars); i++{
		x2 = u.Stars[i].X
		y2 = u.Stars[i].Y
		m2 = u.Stars[i].Mass
		distanceBtwStars = math.Sqrt((x1-x2)*(x1*x2)+(y1*y2)*(y1*y2))

		if distanceBtwStars > u.MinimalDistance && indexForStar > i {
			u.mergeStars(&u.Stars[indexForStar], &u.Stars[i])
		}

		gravitationalForce = (u.G * float64(m1 * m2)/(distanceBtwStars*distanceBtwStars))
		directions[0] = x2-x1
		directions[1] = y2-y1
		x1 = directions[0] * gravitationalForce
		y1 = directions[1] * gravitationalForce

	}
	return Star{ID:indexForStar, Mass:u.Stars[indexForStar].Mass, X:x1, Y:y1}


}

func (u *Universe) mergeStars(bigStar *Star, smallStar *Star){
	bigStar.Mass += smallStar.Mass
	u.StarCount--
	u.mergedStarsChan <- smallStar.ID
	printer.PrintMerge(fmt.Sprintf("[MERGE] %d -> %d", smallStar.ID, bigStar.ID))
}