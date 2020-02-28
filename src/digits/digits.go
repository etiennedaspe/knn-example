package digits

import (
	"bufio"
	"knn_example/utils"
	"os"
	"strconv"
	"strings"
)

// Load digit images and return an array of samples.
//
// The dataset used is the same as the one used in scikit-learn:
// https://scikit-learn.org/stable/modules/generated/sklearn.datasets.load_digits.html?highlight=dataset#sklearn.datasets.load_digits
//
// The dataset is the test set of the Optical Recognition of Handwritten Digits database.
//
// The file digits.csv contains 1797 lines.
// Each line contains 64 integers (elements of the 8x8 matrix mentioned below),
// plus one integer in the range 0..9 (the class, i.e. the value of the handwritten digit).
//
// Data preprocessing:
//
// "We used preprocessing programs made available by NIST to extract normalized bitmaps of handwritten digits from a preprinted form.
// From a total of 43 people, 30 contributed to the training set and different 13 to the test set.
// 32x32 bitmaps are divided into nonoverlapping blocks of 4x4 and the number of on pixels are counted in each block.
// This generates an input matrix of 8x8 where each element is an integer in the range 0..16.
// This reduces dimensionality and gives invariance to small distortions."
//
// Source : https://archive.ics.uci.edu/ml/datasets/Optical+Recognition+of+Handwritten+Digits
//
func Load() (samples utils.Samples, err error) {
	var f *os.File

	f, err = os.Open("./digits/digits.csv")
	if err != nil {
		return
	}

	defer func() {
		err = f.Close()
	}()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)

	for sc.Scan() {
		var s utils.Sample
		l := strings.Split(sc.Text(), ",")

		// the class is the last column in the csv file.
		s.Class, err = strconv.Atoi(l[len(l)-1])
		if err != nil {
			return
		}

		// parse features.
		var ft int
		for _, sFt := range l[:len(l)-1] {
			ft, err = strconv.Atoi(sFt)
			if err != nil {
				return
			}
			s.Features = append(s.Features, ft)
		}

		samples = append(samples, s)
	}

	return
}
