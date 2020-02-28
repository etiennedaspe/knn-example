package utils

import "fmt"

type ConfusionMatrix struct {
	TestData    Samples
	Predictions Samples
}

const nbClasses = 10

// Print the confusion matrix in ASCII art.
func (cm ConfusionMatrix) Print() {
	if len(cm.TestData) != len(cm.Predictions) {
		panic("testData and predictions should have the same size")
	}

	var a [nbClasses][nbClasses]int

	for i, sample := range cm.TestData {
		a[sample.Class][cm.Predictions[i].Class]++
	}

	var str string
	for i := 0; i < nbClasses; i++ {
		str = ""
		for j := 0; j < nbClasses; j++ {
			str += "| " + fmt.Sprintf("%v", a[i][j])
			if a[i][j] < nbClasses {
				str += " "
			}
			if j == nbClasses-1 {
				str += " |"
			}
		}
		fmt.Print(str + "\n")
	}
}
