// package main

// import (
// 	"fmt"
// 	"sync"
// )

// //deadlock example = each philosopher picks up left chopstick first

// type Chops struct{ sync.Mutex }

// type Philosopher struct {
// 	leftCS  *Chops
// 	rightCS *Chops
// }

// func (p Philosopher) eat() {
// 	for {
// 		p.leftCS.Lock()
// 		p.rightCS.Lock()

// 		fmt.Println("Philosopher is eating")

// 		p.leftCS.Unlock()
// 		p.rightCS.Unlock()

// 	}
// }

// func main() {

// 	CSticks := make([]*Chops, 5)

// 	for i := 0; i < 5; i++ {
// 		CSticks[i] = new(Chops)
// 	}

// 	philosophers := make([]*Philosopher, 5)

// 	for i := 0; i < 5; i++ {

// 		// deadlock occurs here
// 		philosophers[i] = &Philosopher{
// 			CSticks[i],
// 			CSticks[(i+1)%5],
// 		}

// 	}

// 	var wg sync.WaitGroup

// 	wg.Add(5)

// 	for i := 0; i < 5; i++ {
// 		go func(p *Philosopher) {
// 			defer wg.Done()
// 			p.eat()
// 		}(philosophers[i])
// 	}

// 	wg.Wait()
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

type Chops struct {
	sync.Mutex
}

type Philosopher struct {
	id      int
	leftCS  *Chops
	rightCS *Chops
}

func (p *Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 3; i++ { // eat 3 times (finite)

		// Break circular wait:
		// last philosopher picks RIGHT first
		if p.id == 4 {
			p.rightCS.Lock()
			p.leftCS.Lock()
		} else {
			p.leftCS.Lock()
			p.rightCS.Lock()
		}

		fmt.Printf("Philosopher %d is eating\n", p.id)
		time.Sleep(100 * time.Millisecond)

		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
}

func main() {
	const n = 5

	// create chopsticks
	CSticks := make([]*Chops, n)
	for i := 0; i < n; i++ {
		CSticks[i] = &Chops{}
	}

	// create philosophers
	philosophers := make([]*Philosopher, n)
	for i := 0; i < n; i++ {
		philosophers[i] = &Philosopher{
			id:      i,
			leftCS:  CSticks[i],
			rightCS: CSticks[(i+1)%n],
		}
	}

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go philosophers[i].eat(&wg)
	}

	wg.Wait()
	fmt.Println("All philosophers are done eating")
}
