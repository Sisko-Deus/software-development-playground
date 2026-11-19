package main

import (
	"object-serializer/serializer"
)

type R struct {
	V2 int
}

type Test struct {
	V1 int
	S1 string
	R1 R
}

func main() {
	t := []Test{
		Test{V1: 10, S1: "Test_1", R1: R{V2: 101}},
		Test{V1: 11, S1: "Test_2", R1: R{V2: 102}},
	}
	srlz := serializer.NewJSONSerializer()
	srlz.IncludeTimestamp = true
	serializer.SerializePrint(srlz, t)
}
