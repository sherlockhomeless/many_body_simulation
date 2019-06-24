package main

import "github.com/sherlockhomeless/many_body_simulation/physics"

// TODO: Zeitmessung
const (
	 MaxMass uint32 = 1000;
	 TimeConvergence int = 1000;
	 X uint32 = 500;
	 Y uint32 = 500;
	 Bodys int = 100;
	 MaxRounds int = 10000;
	 G float64 = 0.00000067;
	 worker_num int = 2;
)

var universe physics.Universe


func main() {
	physics.Debug = true
	physics.Worker = worker_num

	universe := new(physics.Universe)

	universe.StarCount = Bodys
	universe.Time = 0
	universe.X_Dim = X
	universe.Y_Dim = Y
	universe.MaxMass = MaxMass

	universe.FillUniverse()

	for true{


		if(universe.HasFinished) {
			return
		}
	}

}
