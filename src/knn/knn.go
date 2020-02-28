package knn

import (
	"knn_example/utils"
	"math"
	"sort"
)

type (
	// Classifiers contains training set.
	Classifier struct {
		Samples utils.Samples
	}

	// Prediction contains a sample with the features given and the class predicted,
	// with the list of the nearest neighbours.
	Prediction struct {
		Sample            utils.Sample
		NearestNeighbours utils.Samples
	}

	// Predictions stores a list of predictions.
	Predictions []Prediction

	// Classes is used to store the number of representatives for each class during the majority voting.
	Classes [10]int

	// SampleDistance is used to store a sample and its distance from the image to be predicted.
	SampleDistance struct {
		Sample   utils.Sample
		Distance float64
	}

	// ByDistance allows to sort by distance an array of SampleDistance.
	ByDistance []SampleDistance
)

// Predict the class of each image in images, using the k-nearest neighbours algorithm.
//
//			algorithm k-nearest neighbours is
// 				input: the number of nearest neighbours k,
//					   images for which we want to predict the class
//				output: a list of predictions where a prediction is an image with its predicted class
//
//				for each image img in images do
//					for each training sample s in the training set do
//						d = euclideanDistance(img, s)
//						neighbours += (s, d)
//
//					sortByDistance(neighbours)
//
//					predictions[img] = majorityVoting(first k neighbours)
//
//				return predictions
//
func (c Classifier) Predict(k int, images []utils.Features) Predictions {
	// bad value for k
	// fallback with a 1-nearest neighbour classifier
	if k <= 0 || k > len(c.Samples) {
		k = 1
	}

	var predictions Predictions

	// compute prediction for each image
	for _, img := range images {
		// compute euclidean distances between the features of the image to predict,
		// and the features of each sample of the model.
		var neighbours ByDistance
		for _, s := range c.Samples {
			d := euclideanDistance(img, s.Features)
			neighbours = append(neighbours, SampleDistance{
				Sample:   s,
				Distance: d,
			})
		}

		sort.Sort(neighbours)

		// for k-nearest neighbours, count the number of representatives of each class.
		var (
			candidates Classes       // for each class, the number of representatives
			nn         utils.Samples // nearest neighbours
		)
		for i := 0; i < k; i++ {
			candidates[neighbours[i].Sample.Class]++
			nn = append(nn, neighbours[i].Sample)
		}

		p := candidates.majorityVoting()

		// add the prediction found for the current image with its nearest neighbours.
		predictions = append(predictions, Prediction{
			Sample: utils.Sample{
				Features: img,
				Class:    p,
			},
			NearestNeighbours: nn,
		})
	}

	return predictions
}

// majorityVoting return the class with the most representatives.
//
// The index in the array of Classes cs matches the corresponding digit,
// e.g. the number of representatives for the class 0 will be stored at index 0.
//
// Example:
//
// 		Assume that for seven neighbours, we have two samples of 0, two samples of 4 and five samples of 6,
// 		an array c of classes will be filled like this:
//
// 		c = [2, 0, 0, 0, 2, 0, 5, 0, 0, 0]
//
//		Here the winner is the class 6 with five representatives.
//
func (cs Classes) majorityVoting() int {
	var (
		max    int
		winner int
	)
	for c, nbVotes := range cs {
		if nbVotes > max {
			winner = c
		}
	}
	return winner
}

// euclideanDistance computes the euclidean distance D between the two vectors x and y.
//
// 			D(x, y) = √ ∑ (xᵢ - yᵢ)²
//
func euclideanDistance(x, y utils.Features) (d float64) {
	for i := range x {
		d += float64((x[i] - y[i]) * (x[i] - y[i]))
	}
	math.Sqrt(d)
	return
}

func (b ByDistance) Len() int {
	return len(b)
}

func (b ByDistance) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByDistance) Less(i, j int) bool {
	return b[i].Distance < b[j].Distance
}

func (ps Predictions) Samples() (samples utils.Samples) {
	for _, p := range ps {
		samples = append(samples, p.Sample)
	}
	return
}
