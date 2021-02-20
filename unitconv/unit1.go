package unitconv

import(
   "fmt"
   )

//type Feets float64
//type Meters float64
//type Kilograms float64
//type Pounds float64

func FtoM(f Feet) Meters{
  Met:=Meters(f*0.3048)
  return Met
}

