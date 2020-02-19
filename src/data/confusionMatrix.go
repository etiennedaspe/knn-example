package data

import "fmt"

type ConfusionMatrix struct {
	TestData    Samples
	Predictions Samples
}

func (cm ConfusionMatrix) Print() {
	if len(cm.TestData) != len(cm.Predictions) {
		panic("testData and predictions should have the same size")
	}

	var a [10][10]int

	for i, sample := range cm.TestData {
		a[sample.Class][cm.Predictions[i].Class]++
	}

	var str string
	for i := 0; i < 10; i++ {
		str = ""
		for j := 0; j < 10; j++ {
			str += "|" + fmt.Sprintf("%v", a[i][j])
		}
		fmt.Print(str + "\n")
	}
}
