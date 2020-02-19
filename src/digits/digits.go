package digits

import (
	"bufio"
	"knn_example/data"
	"os"
	"strconv"
	"strings"
)

const NbSamples = 1797

func LoadDigits() (samples data.Samples, err error) {
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
		var sp data.Sample
		a := strings.Split(sc.Text(), ",")
		sp.Class, err = strconv.Atoi(a[len(a)-1])
		if err != nil {
			return
		}

		var ft int
		for _, s := range a[:len(a)-1] {
			ft, err = strconv.Atoi(s)
			if err != nil {
				return
			}
			sp.Features = append(sp.Features, ft)
		}

		samples = append(samples, sp)
	}

	return
}
