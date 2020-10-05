package main

import (
	"islands/api"
	"log"
)

func main() {
	o, err := api.NewOcean(6)

	if err != nil {
		log.Fatalf("could not create ocean: +%v", err)
	}

	for _, land := range []struct{ x, y int }{
		{1, 0},
		{2, 1},
		{2, 2},
		{3, 2},
		{4, 2},
		{5, 1},
	} {
		log.Printf("adding land at x=%d,y=%d", land.x, land.y)

		if err := o.PutLand(land.x, land.y); err != nil {
			log.Fatalf("could not island %+v", err)
		}
	}

	log.Printf("island count %d", o.Islands())
}
