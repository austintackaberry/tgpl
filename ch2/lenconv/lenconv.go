// Package lenconv performs Feet and Meters conversions.
package lenconv

import "fmt"

type Feet float64
type Meters float64

func (m Meters) String() string { return fmt.Sprintf("%gm", m) }
func (ft Feet) String() string  { return fmt.Sprintf("%gft", ft) }
