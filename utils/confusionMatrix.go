package utils

import "fmt"

// ConfusionMatrix shows the performance of the algorithm.
//
// Example with k=7:
//
//             | Actual class
//             |_________________________________________
//             | 0   1   2   3   4   5   6   7   8   9
// ____________|_________________________________________
//           0 | 87| 0 | 0 | 0 | 1 | 0 | 0 | 0 | 0 | 0  |
//           1 | 0 | 65| 1 | 1 | 1 | 0 | 1 | 1 | 11| 10 |
//           2 | 1 | 0 | 73| 5 | 0 | 0 | 0 | 0 | 2 | 5  |
//           3 | 0 | 0 | 0 | 68| 0 | 2 | 0 | 1 | 9 | 11 |
// Predicted 4 | 0 | 0 | 0 | 0 | 83| 0 | 0 | 0 | 0 | 9  |
// class     5 | 0 | 0 | 0 | 0 | 0 | 72| 9 | 0 | 0 | 10 |
//           6 | 0 | 0 | 0 | 0 | 0 | 0 | 91| 0 | 0 | 0  |
//           7 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 83| 1 | 5  |
//           8 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 1 | 78| 9  |
//           9 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 0 | 92 |
//
// More information:
// https://en.wikipedia.org/wiki/Confusion_matrix
//
type ConfusionMatrix struct {
	TestData    Samples
	Predictions Samples
}

const nbClasses = 10

// Print the confusion matrix in ASCII art.
func (cm ConfusionMatrix) Print() {
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
