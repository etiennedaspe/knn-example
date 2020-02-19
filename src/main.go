package main

import (
	"fmt"
	"knn_example/data"
	"knn_example/digits"
	"knn_example/knn"
)

func main() {
	// load digits
	samples, err := digits.LoadDigits()
	if err != nil {
		panic(err)
	}

	// print first three digits
	if len(samples) != digits.NbSamples {
		panic("not enough samples")
	}
	fmt.Print("\n========= Samples =========\n")
	samples[0].Print()
	samples[1].Print()
	samples[2].Print()

	// split data in train and test samples
	var (
		trainingData = samples[:len(samples)/2]
		testData     = samples[len(samples)/2:]
	)

	// create KNN classifier
	classifier := knn.KNNC{Samples: trainingData}

	// predict
	predictions := classifier.Predict(7, testData.Features())

	// print first three predictions
	if len(predictions) != len(testData) {
		panic("not enough predictions")
	}
	fmt.Print("\n========= Predictions =========\n")
	predictions[0].Sample.Print()
	predictions[1].Sample.Print()
	predictions[2].Sample.Print()

	// print confusion matrix
	fmt.Print("\n========= Confusion Matrix =========\n")
	cm := data.ConfusionMatrix{
		TestData:    testData,
		Predictions: predictions.Samples(),
	}
	cm.Print()

	// first prediction
	fmt.Print("\n========= First Prediction =========\n")
	predictions[5].Sample.Print()

	// k nearest neighbours of the first prediction
	fmt.Print("\n========= K Nearest Neighbours =========\n")
	for _, n := range predictions[5].NearestNeighbours {
		n.Print()
	}
}
