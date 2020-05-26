package main

import (
	"fmt"
	"testing"
)

func TestDisJoin_Find(t *testing.T) {

	dJ := InitDisJoin(5)

	dJ.Merge(0, 1)
	dJ.Merge(1, 2)
	dJ.Merge(3, 2)

	fmt.Println(dJ.Find(4, 3))

}
