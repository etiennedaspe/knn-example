package utils

import "fmt"

type ConfusionMatrix struct {
	TestData    Samples
	Predictions Samples
}

const NbClasses = 10

// Print the confusion matrix in ASCII art.
func (cm ConfusionMatrix) Print() {
	if len(cm.TestData) != len(cm.Predictions) {
		panic("testData and predictions should have the same size")
	}

	var a [NbClasses][NbClasses]int

	for i, sample := range cm.TestData {
		a[sample.Class][cm.Predictions[i].Class]++
	}

	var str string
	for i := 0; i < NbClasses; i++ {
		str = ""
		for j := 0; j < NbClasses; j++ {
			str += "| " + fmt.Sprintf("%v", a[i][j])
			if a[i][j] < NbClasses {
				str += " "
			}
			if j == NbClasses-1 {
				str += " |"
			}
		}
		fmt.Print(str + "\n")
	}
}
