Package unitconv

import "fmt"


type Feet float64
type Meters float64

func(m Meters) String() string{
    return fmt.Sprintf("%gM",m) 
}
