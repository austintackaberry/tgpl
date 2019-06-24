package ch1

import (
	"os"
)

// Ex6 - Lissajous generates GIF animations of random Lissajous figures.
func Ex6() {
	Lissajous(os.Stdout, 5)
}
