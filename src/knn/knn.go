package knn

import (
	"knn_example/data"
	"math"
	"sort"
)

type KNNC struct {
	Samples data.Samples
}

type Prediction struct {
	Sample            data.Sample
	NearestNeighbours data.Samples
}

type Predictions []Prediction

func (p Predictions) Samples() (samples data.Samples) {
	for _, prediction := range p {
		samples = append(samples, prediction.Sample)
	}
	return
}

func (knnc KNNC) Predict(k int, d []data.Features) (predictions Predictions) {
	if k > len(knnc.Samples) {
		panic("not enough samples in classifier")
	}

	// compute prediction for each image
	for _, features := range d {
		// compute euclidean distances between the features of the image to predict,
		// and the features of each sample of the model.
		var distances ByDistance
		for _, sample := range knnc.Samples {
			distance := euclideanDistance(features, sample.Features)
			sd := SampleDistance{
				Sample:   sample,
				Distance: distance,
			}
			distances = append(distances, sd)
		}

		sort.Sort(distances)

		// for k-nearest neighbours
		// count the number of representatives of each class
		var (
			candidates Classes      // for each class, the number of representatives
			nn         data.Samples // nearest neighbours
		)
		for i := 0; i < k; i++ {
			candidates[distances[i].Sample.Class]++
			nn = append(nn, distances[i].Sample)
		}

		// add the prediction found for the current image
		// with its nearest neighbours.
		predictions = append(predictions, Prediction{
			Sample: data.Sample{
				Features: features,
				Class:    candidates.majorityVoting(),
			},
			NearestNeighbours: nn,
		})
	}
	return
}

type Classes [10]int

func (l Classes) majorityVoting() (winner int) {
	var max int
	for label, nbVotes := range l {
		if nbVotes > max {
			winner = label
		}
	}
	return
}

func euclideanDistance(x, y data.Features) (distance float64) {
	for i := range x {
		distance += float64((x[i] - y[i]) * (x[i] - y[i]))
	}
	math.Sqrt(distance)
	return
}

type SampleDistance struct {
	Sample   data.Sample
	Distance float64
}

type ByDistance []SampleDistance

func (b ByDistance) Len() int {
	return len(b)
}

func (b ByDistance) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByDistance) Less(i, j int) bool {
	return b[i].Distance < b[j].Distance
}
