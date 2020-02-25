package digits

import (
	"bufio"
	"knn_example/utils"
	"os"
	"strconv"
	"strings"
)

const NbSamples = 1797

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
		var sp utils.Sample
		l := strings.Split(sc.Text(), ",")

		// the class is the last column
		sp.Class, err = strconv.Atoi(l[len(l)-1])
		if err != nil {
			return
		}

		// parse features
		var ft int
		for _, sFt := range l[:len(l)-1] {
			ft, err = strconv.Atoi(sFt)
			if err != nil {
				return
			}
			sp.Features = append(sp.Features, ft)
		}

		samples = append(samples, sp)
	}

	return
}
