package main

import (
	"fmt"
	"knn_example/digits"
	"knn_example/knn"
	"knn_example/utils"
)

// The number of nearest neighbours used.
const k = 7

func main() {
	// Load digits.
	samples, err := digits.Load()
	if err != nil {
		panic(err)
	}

	// Print first three digits.
	fmt.Print("\n========= Samples =========\n")
	samples[0].Print()
	samples[1].Print()
	samples[2].Print()

	// Split samples in train and test dataset.
	var (
		trainingData = samples[:len(samples)/2]
		testData     = samples[len(samples)/2:]
	)

	// Create KNN classifier.
	classifier := knn.Classifier{Samples: trainingData}

	// Predict class for each image of the test dataset.
	predictions := classifier.Predict(k, testData.Features())

	// Print first three predictions.
	fmt.Print("\n========= Predictions =========\n")
	predictions[0].Sample.Print()
	predictions[1].Sample.Print()
	predictions[2].Sample.Print()

	// Print confusion matrix.
	fmt.Print("\n========= Confusion Matrix =========\n")
	cm := utils.ConfusionMatrix{
		TestData:    testData,
		Predictions: predictions.Samples(),
	}
	cm.Print()

	// First prediction.
	fmt.Print("\n========= First Prediction =========\n")
	predictions[0].Sample.Print()

	// The k-nearest neighbours of the first prediction.
	fmt.Print("\n========= K-Nearest Neighbours =========\n")
	for _, n := range predictions[0].NearestNeighbours {
		n.Print()
	}
}
