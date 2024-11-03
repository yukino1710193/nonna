package main

import (
	"context"

	"github.com/bonavadeur/nonna/pkg/bonalib"
	_ "github.com/bonavadeur/nonna/pkg/nonna"
)

func main() {
	bonalib.Log("Konnichiwa, Nonna desu")
	ctx := context.Background()

	// do something here ...

	<-ctx.Done() // hangout forever
}
