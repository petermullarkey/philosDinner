package main

/* Implement the dining philosopher’s problem with the following constraints/modifications.

	There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
	
	Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
	
	The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
	
	In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
	
	The host allows no more than 2 philosophers to eat concurrently.
	
	Each philosopher is numbered, 1 through 5.
	
	When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” 
	on a line by itself, where <number> is the number of the philosopher.

	When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” 
	on a line by itself, where <number> is the number of the philosopher.
*/
   
import (
	"fmt"
	"sync"
	"time"
)
type ChopS struct{ sync.Mutex }
type Philo struct {
	dinerId int
	leftCS, rightCS *ChopS
}

func coordinateChowing(requestToEat <-chan bool, doneEating <-chan bool, wg *sync.WaitGroup){
	defer wg.Done()
	var PhilosFed = 0
	var numEating = 0
	for {
		if (numEating < 3) {
		<-requestToEat
		numEating++
		PhilosFed++
		fmt.Println("Have fed", PhilosFed, "hungry philosophers")
	
	if (<-doneEating) {
		numEating--
	}
	if (PhilosFed == 15){
		break
	}
		}

	}
}
func (p Philo) chowDown(requestToEat chan<- bool, doneEating chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for coursesEaten := 1; coursesEaten < 4; coursesEaten++ {
		requestToEat<-true
		p.leftCS.Lock()
		p.rightCS.Lock()
  
		fmt.Println("Philosopher",  p.dinerId, "starting to eat course", coursesEaten)
		fmt.Println("Philosopher",  p.dinerId, "done eating course", coursesEaten)
		doneEating<-true
		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
} 
func main() {

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
	   CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	okToEat := make(chan bool)
	doneEating := make(chan bool)
	for i := 0; i < 5; i++ {
	   philos[i] = &Philo{i, CSticks[i], CSticks[(i+1)%5]}
	}
	fmt.Println("start the feast")
	var wg sync.WaitGroup
	wg.Add(6)
	go coordinateChowing(okToEat, doneEating, &wg)
	for i := 0; i < 5; i++ {
		go philos[i].chowDown(okToEat, doneEating, &wg)
	 }
	wg.Wait()
	}
