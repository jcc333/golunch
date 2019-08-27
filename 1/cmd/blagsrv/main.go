package main

import (
	json1 "encoding/json"
	"fmt"
)

/*
// this is a type
// called 'error'
// which is an interface
// that requires Error with no args that returns a string
type error interface { Error() string }
*/
type MyCoolErrorStruct struct {
	ErrorMessage       string `json:"errorMessage"`
	PositionInTheFileX int    `json:"x"`
	PositionInTheFileY int    `json:"y"`
}

// This is a type called MyCoolerErrorStruct that's just anothername for MyCoolErrorStruct
type MyCoolerErrorStruct MyCoolErrorStruct

// This is a function on MyCoolErrorStruct called Error with no arguments that returns a string
func (mces MyCoolErrorStruct) Error() string {
	return fmt.Sprintf("error at %d:%d; '%s'", mces.PositionInTheFileX, mces.PositionInTheFileY, mces.ErrorMessage)
}

// This is a function on MyCoolerErrorStruct called Error with no arguments that returns a string
func (mces MyCoolerErrorStruct) Error() string {
	bytes, err := json1.Marshal(mces)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// func(MyCoolerErrorStruct) string

// This is a function that takes an error and returns nothing
func reportError(err error) {
	fmt.Println(err)
}

// This is a func
// called applyFunc
// from MyCoolerErrorStruct and
// a func(MyCoolerErrorStruct) string
// and returns a string
func applyFunc(s MyCoolerErrorStruct, fn func(MyCoolerErrorStruct) string) string {
	return s.fin(s)
}

// This is a function called main that has no arguments and returns nothing
func main() {
	val := MyCoolerErrorStruct{
		ErrorMessage:       "whoopsidaisy",
		PositionInTheFileX: 29,
		PositionInTheFileY: 30,
	}
	val1 := MyCoolErrorStruct(val)
	reportError(val)
	reportError(val1)
	fmt.Println(val.Error())
	fmt.Println(applyFunc(val, MyCoolerErrorStruct.Error))
}
