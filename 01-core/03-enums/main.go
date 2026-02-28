package main

import "fmt"

const (
	//iota identifier is used in const declarations to simplify definitions of incrementing numbers.
	monday = iota + 1
	tuesday
	wednesday
	thursday
	friday
	saturday
	sunday
)

type ByteSize float64

const (
	_           = iota // ignore the first value by assigning to blank identifier to avoid multiplying by 0
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type LogLevel int

const (
	LogError LogLevel = iota
	LogWarning
	LogInfo
	LogDebug
)

func main() {
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(wednesday)
	fmt.Println(thursday)
	fmt.Println(friday)
	fmt.Println(saturday)
	fmt.Println(sunday)
}
