package digits

import (
	"bufio"
	"knn_example/utils"
	"os"
	"strconv"
	"strings"
)

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

		// the class is the last column
		s.Class, err = strconv.Atoi(l[len(l)-1])
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
			s.Features = append(s.Features, ft)
		}

		samples = append(samples, s)
	}

	return
}
