package main

import (
	"fmt"
	"sync"
)

type arrs [2]int // an array for storing the philosopher's number and number of servings

type Chops struct {
	sync.Mutex
}

type Philo struct {
	leftCase, rightCase *Chops
	flag bool
	number *arrs
}


// eating philosophers
func (p Philo) eat() {
      if p.number[1] > 0 {
		  p.leftCase.Lock()
		  p.rightCase.Lock()

		  fmt.Println("starting to eat", p.number[0])
		  fmt.Println("eat")
		  p.number[1]--

		  fmt.Println("portion balance", p.number[1], "for philosoph number", p.number[0])
		  fmt.Println("finishing eating", p.number[0])
		  p.rightCase.Unlock()
		  p.leftCase.Unlock()
	  } else {
	  	fmt.Println("ate philosopher number", p.number[0])
		  return
	  }
}

// activation flag
func (p Philo) active() {
	if !p.flag {
	 p.flag = true
 }
}

// deactivation flag
func (p Philo) noactive() {
	if p.flag {
		p.flag = false
	}
}

// the choice of the right stick for the philosopher
func  checks(p *Philo, arrPhilos []*Philo) {
	if p.rightCase != nil  {
		return
	}
	 for _, values := range arrPhilos {
	    	if p == values {
	    		continue
			} else {
				if !values.flag {
					p.rightCase = values.leftCase
					break
	 			} else {
	 				continue
				}
			}
	 }
}

func main() {
	// creating an array of sticks
    CSticks := make([]*Chops, 5)
    for i := 0; i < 5; i++ {
    	CSticks[i] = new(Chops)
	}
    // creating philos
    philos := make([]*Philo, 5)


    for i := 0; i < 5; i++ {
        p :=  arrs{i+1, 3}
    	philos[i] = &Philo{CSticks[i], nil, false, &p}

	}

	// numbers of philosophers who are eating at a given time
	var numberPhilo1, numberPhilo2 int
	for {

		// Enter philosopher numbers
    	_, _ = fmt.Scan(&numberPhilo1, &numberPhilo2)
    	// Checks number philosopher
       if numberPhilo1 > 0 && numberPhilo1 < 6 && numberPhilo2 > 0 && numberPhilo2 < 6 {
		   philos[numberPhilo1-1].active()
		   philos[numberPhilo2-1].active()
		    checks(philos[numberPhilo1-1], philos)
		    checks(philos[numberPhilo2-1], philos)

		   go philos[numberPhilo1-1].eat()
		   go philos[numberPhilo2-1].eat()

		   philos[numberPhilo1-1].noactive()
		   philos[numberPhilo2-1].noactive()

	   } else {
	   	fmt.Println("Error numbers")
	   	break
	   }
	}


}
