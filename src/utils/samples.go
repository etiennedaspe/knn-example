package utils

import "fmt"

type (
	Sample struct {
		Features Features
		Class    int
	}

	Samples  []Sample
	Features []int
)

// Print a sample in ASCII art
func (s Sample) Print() {
	if len(s.Features) != 64 {
		panic("not enough features")
	}

	fmt.Printf("\nClass : %v\n", s.Class)

	var str string
	for i := 0; i < 8; i++ {
		str = ""
		for j := 0; j < 8; j++ {
			str += " " + greyASCII(s.Features[8*i+j])
			str += " " + greyASCII(s.Features[8*i+j])
		}
		fmt.Print(str + "\n")
		fmt.Print(str + "\n")
	}
}

func (s Samples) Features() (fs []Features) {
	for _, sample := range s {
		fs = append(fs, sample.Features)
	}
	return
}

// greyASCII returns an ASCII character with a scale of ten shades of grey
//
// Source: http://paulbourke.net/dataformats/asciiart/
//
func greyASCII(i int) string {
	switch i {
	case 0:
		return " "
	case 1:
		return "."
	case 2:
		return ":"
	case 3:
		return "-"
	case 4:
		return "="
	case 5:
		return "+"
	case 6:
		return "*"
	case 7:
		return "#"
	case 8:
		return "%"
	default:
		return "@"
	}
}
