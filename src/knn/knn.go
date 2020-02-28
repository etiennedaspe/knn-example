package knn

import (
	"knn_example/utils"
	"math"
	"sort"
)

type (
	Classifier struct {
		Samples utils.Samples
	}

	Prediction struct {
		Sample            utils.Sample
		NearestNeighbours utils.Samples
	}

	Predictions []Prediction

	Classes [10]int

	SampleDistance struct {
		Sample   utils.Sample
		Distance float64
	}

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
//						neighbours[s] = euclideanDistance(img, s)
//
//					sortByDistance(neighbours)
//
//					predictions[img] = majorityVoting(k-nearest neighbours)
//
//				return predictions
//
func (c Classifier) Predict(k int, images []utils.Features) Predictions {
	// bad value for k
	// fallback with a 1-nearest neighbor classifier
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

		// for k-nearest neighbours
		// count the number of representatives of each class
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
// TODO(ed) add weighted majority voting.
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
