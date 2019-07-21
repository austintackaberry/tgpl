// unitconv converts C<>F, ft<>m, lbs<>kg
package unitconv

import (
	"fmt"
	"os"
	"strconv"
	"tgpl/ch2/lenconv"
	"tgpl/ch2/tempconv"
)

func Unitconv() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		m := lenconv.Meters(t)
		ft := lenconv.Feet(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c), m, lenconv.MToFt(m), ft, lenconv.FtToM(ft))
	}
}
