package main

import (
	"github.com/pebbe/novas"

	"fmt"
)

func main() {

	jpleph := "/my/opt/novas/share/JPLEPH"

	latitude, longitude := 53.21853, 6.5670 // Groningen, The Netherlands

	// END OF USER SETTINGS

	novas.Init(jpleph, false)

	t := novas.Now()
	fmt.Println(t)

	geo := novas.NewPlace(latitude, longitude, 0, 20, 1010)

	fmt.Println("\n            Distance   Altitude   Azimuth")
	for _, obj := range []*novas.Planet{
		novas.Mercury,
		novas.Venus,
		novas.Mars,
		novas.Jupiter,
		novas.Saturn,
		novas.Uranus,
		novas.Neptune,
		novas.Pluto,
	} {
		data := obj.Topo(t, geo, novas.REFR_NONE)
		fmt.Printf("%-8s%12.6f%11.2f%10.2f\n", obj.Name(), data.Dis, data.Alt, data.Az)
	}
}
