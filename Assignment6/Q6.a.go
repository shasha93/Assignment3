package main

import "fmt"

type Enum int

func (e Enum) String() string {
	switch e {
	case 1:
		return "Alpha"
	case 2:
		return "Beta"
	}
	return ""
}

const (
	Alpha Enum = iota
	Beta
        
)



func main() {
     var t Enum=1
	fmt.Printf("%s\n", t)
	
}

