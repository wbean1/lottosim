package draws

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"time"

	"github.com/jinzhu/copier"
)

func SimpleRandomDraw(powerball bool) {
	var maxWhite, maxGold int
	if powerball {
		maxWhite = 69
		maxGold = 26
	} else {
		maxWhite = 70
		maxGold = 25
	}
	ticket := drawTicket(maxWhite, maxGold)
	ticket.OrderTicket()
	ticket.PrintTicket()
}

func DoubleDraw(powerball bool) {
	var maxWhite, maxGold int
	if powerball {
		maxWhite = 69
		maxGold = 26
	} else {
		maxWhite = 70
		maxGold = 25
	}
	ticket := doubleDrawTicket(maxWhite, maxGold)
	ticket.OrderTicket()
	ticket.PrintTicket()
}

func FullTicketDoubleDraw(powerball bool) {
	var maxWhite, maxGold int
	if powerball {
		maxWhite = 69
		maxGold = 26
	} else {
		maxWhite = 70
		maxGold = 25
	}
	ticket := fullTicketDoubleDraw(maxWhite, maxGold)
	ticket.OrderTicket()
	ticket.PrintTicket()
}

func FirstOrderedDraw(powerball bool) {
	var maxWhite, maxGold int
	if powerball {
		maxWhite = 69
		maxGold = 26
	} else {
		maxWhite = 70
		maxGold = 25
	}
	ticket := firstOrderedTicketDraw(maxWhite, maxGold)
	ticket.PrintTicket()
}

func HighestFrequencyDraw(powerball bool, draws int) {
	var maxWhite, maxGold int
	if powerball {
		maxWhite = 69
		maxGold = 26
	} else {
		maxWhite = 70
		maxGold = 25
	}
	ticket := highestFrequencyTicketDraw(maxWhite, maxGold, draws)
	ticket.OrderTicket()
	ticket.PrintTicket()
}

func drawTicket(maxWhite, maxGold int) Ticket {
	var ticket Ticket
	ticket.WhiteBalls = append(ticket.WhiteBalls, drawBall(maxWhite, []int{}))
	ticket.WhiteBalls = append(ticket.WhiteBalls, drawBall(maxWhite, ticket.WhiteBalls))
	ticket.WhiteBalls = append(ticket.WhiteBalls, drawBall(maxWhite, ticket.WhiteBalls))
	ticket.WhiteBalls = append(ticket.WhiteBalls, drawBall(maxWhite, ticket.WhiteBalls))
	ticket.WhiteBalls = append(ticket.WhiteBalls, drawBall(maxWhite, ticket.WhiteBalls))
	ticket.GoldBall = drawBall(maxGold, []int{})
	return ticket
}

func doubleDrawTicket(maxWhite, maxGold int) Ticket {
	var ticket Ticket
	ticket.WhiteBalls = doubleDrawBalls(maxWhite, 5)
	ticket.GoldBall = doubleDrawBalls(maxGold, 1)[0]
	return ticket
}

func fullTicketDoubleDraw(maxWhite, maxGold int) Ticket {
	var tickets []Ticket
	numDrawn := 0
	for true {
		ticket := drawTicket(maxWhite, maxGold)
		numDrawn++
		ticket.OrderTicket()
		for _, i := range tickets {
			if reflect.DeepEqual(i, ticket) {
				fmt.Printf("numDrawn: %d\n", numDrawn)
				return ticket
			}
		}
		tickets = append(tickets, ticket)
	}
	return Ticket{}
}

func firstOrderedTicketDraw(maxWhite, maxGold int) Ticket {
	numDrawn := 0
	for true {
		t := drawTicket(maxWhite, maxGold)
		numDrawn++
		o := Ticket{}
		o.GoldBall = t.GoldBall
		copier.Copy(&o.WhiteBalls, &t.WhiteBalls)
		o.OrderTicket()
		if reflect.DeepEqual(t, o) {
			fmt.Printf("numDrawn: %d\n", numDrawn)
			return t
		}
	}
	return Ticket{}
}

func highestFrequencyTicketDraw(maxWhite, maxGold, draws int) Ticket {
	whites := make(map[int]int)
	golds := make(map[int]int)
	type kv struct {
		Key   int
		Value int
	}
	for draws > 0 {
		w := drawBall(maxWhite, []int{})
		g := drawBall(maxGold, []int{})
		whites[w] = whites[w] + 1
		golds[g] = golds[g] + 1
		draws--
	}
	var sortedWhites, sortedGolds []kv
	for k, v := range whites {
		sortedWhites = append(sortedWhites, kv{k, v})
	}
	for k, v := range golds {
		sortedGolds = append(sortedGolds, kv{k, v})
	}
	sort.Slice(sortedWhites, func(i, j int) bool {
		return sortedWhites[i].Value > sortedWhites[j].Value
	})
	sort.Slice(sortedGolds, func(i, j int) bool {
		return sortedGolds[i].Value > sortedGolds[j].Value
	})
	var t Ticket
	t.WhiteBalls = []int{sortedWhites[0].Key, sortedWhites[1].Key, sortedWhites[2].Key, sortedWhites[3].Key, sortedWhites[4].Key}
	t.GoldBall = sortedGolds[0].Key
	return t
}

func doubleDrawBalls(max, num int) []int {
	var winners []int
	drawn := make(map[int]int)
	for num > 0 {
		d := drawBall(max, winners)
		if drawn[d] > 0 {
			winners = append(winners, d)
			num = num - 1
		}
		drawn[d] = drawn[d] + 1
	}
	return winners
}

func drawBall(max int, excludes []int) int {
	x := 0
	for x < 1 {
		rand.Seed(time.Now().UnixNano())
		x = rand.Intn(max) + 1
		if intInSlice(x, excludes) {
			x = 0
		}
	}
	return x
}

func (ticket *Ticket) PrintTicket() {
	fmt.Printf("%+v\n", ticket)
}

func (ticket *Ticket) OrderTicket() {
	sort.Ints(ticket.WhiteBalls)
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
