package main

import "fmt"

func GenDisplaceFn(data ...float64) func (time float64) float64 {
	fn := func(time float64) float64 {
		return 1.0/2 * data[0] * time * time  + data[1]* time + data[2];
	}
	return fn;
}


func main() {
   var(
   	speedUp, speedZero, innitialDisplace, time float64
   )
   _, _ = fmt.Scan(&speedUp, &speedZero, &innitialDisplace)
   fn  := GenDisplaceFn(speedUp, speedZero, innitialDisplace);
   _, _ = fmt.Scan(&time)
   fmt.Print(fn(time))

}



