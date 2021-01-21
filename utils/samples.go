package utils

import (
	"fmt"
)

const (
	sampleHeight = 8
	sampleWidth  = 8
)

type (
	// Sample represents an image of a handwritten digit (features) with its value (class).
	Sample struct {
		Features Features
		Class    int
	}

	Samples  []Sample
	Features []int
)

// Print a sample in ASCII art.
func (s Sample) Print() {
	if len(s.Features) != sampleHeight*sampleWidth {
		fmt.Print("can't print sample - not enough features")
		return
	}

	fmt.Printf("\nClass : %v\n", s.Class)

	var str string
	for i := 0; i < sampleHeight; i++ {
		str = ""
		for j := 0; j < sampleWidth; j++ {
			str += " " + greyASCII(s.Features[8*i+j])
			str += " " + greyASCII(s.Features[8*i+j])
		}
		fmt.Print(str + "\n")
		fmt.Print(str + "\n")
	}
}

// Features returns arrays of feature.
func (s Samples) Features() (fs []Features) {
	for _, sample := range s {
		fs = append(fs, sample.Features)
	}
	return
}

// greyASCII returns an ASCII character with a scale of ten shades of grey.
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
