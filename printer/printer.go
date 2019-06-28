package printer

import "github.com/fatih/color"

func PrintMeta(msg string){
	color.Blue(msg)
}

func PrintMovement(msg string){
	color.Yellow(msg)
}

func PrintMerge(msg string){
	color.Red(msg)
}

func PrintGoRoutiens(msg string){
	color.Green(msg)
}