package main
import (
				"fmt"
				"os"
				"bufio"
				"io"
				"strings"
				"strconv"
				"container/heap"
				"math"
				"sync"
)

type Item struct {
		value []int
		priority int
		index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	//item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value []int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}


func contains(slice []string, toFind string) bool {
		for _, value := range slice {
				if strings.EqualFold(value, toFind) {
								return true
				}
		}
		return false
}
func main() {

	reader := bufio.NewReader(os.Stdin)
	//points := make([]item, 0)
	px := make(PriorityQueue, 0)
	py := make(PriorityQueue, 0)
	//for i := 0; i < 6; i++ {
	//				_ , _ = reader.ReadString('\n')
	//}
	heap.Init(&px)
	heap.Init(&py)
	i := 0
	_ , _ = reader.ReadString('\n')
	for {
				text, err  := reader.ReadString('\n')
				if err == io.EOF || strings.EqualFold(text, "EOF\n") {
					break
				}
				myNums := strings.Split(text, " ")
				myNums[1] = strings.TrimSuffix(myNums[1], "\n")
				x, _:= strconv.Atoi(myNums[0])
				y, _:= strconv.Atoi(myNums[1])
				pointx := &Item{
										value: []int{x,y},
										priority: x,
										index: i,
				}
				pointy := &Item{
							value: []int{x,y},
							priority: y,
							index: i,
				}
				heap.Push(&px, pointx)
				heap.Push(&py, pointy)
				i++
		}
		//for px.Len() > 0 {
		//	item := heap.Pop(&px).(*Item)
		//	fmt.Print("Index: ", item.index, " ")
		//	fmt.Print("Prio: ", item.priority, " ")
		//	fmt.Print("Value: ", item.value, " ")
		//	fmt.Println("")
		//	//fmt.Printf("%.3d:%d \n", item.priority, item.value)
		//}
		//fmt.Println(px)
		firstHalf := px[0:(len(px)/2)]
		secondHalf := px[(len(px)/2):]
		firstHalfy := py[0:(len(py)/2)]
		secondHalfy := py[(len(py)/2):]

		var qx float64
		var qy float64
		var rx float64
		var ry float64
		var wg sync.WaitGroup

		go findClosest(firstHalf, &qx, 1, &wg)
		wg.Add(1)
		go findClosest(secondHalf, &rx, 2, &wg)
		wg.Add(1)
		go findClosest(firstHalfy, &qy, 3, &wg)
		wg.Add(1)
		go findClosest(secondHalfy, &ry, 4, &wg)
		wg.Add(1)
		wg.Wait()
		delta := math.Min(math.Min(qx, qy), math.Min(rx,ry))
		//fmt.Println("qx: ", qx, " rx: ", rx)
		//fmt.Println("qy: ", qy, " ry: ", ry)
		//fmt.Println(delta)
		if delta == math.Inf(1) {
						delta = 10000
		}
		xstar := firstHalf[len(firstHalf)-1]
		sy := make(PriorityQueue, 0)
		heap.Init(&sy)
		for _ , element := range py {
				if (element.value[0] - xstar.value[0]) < (int(delta)+1) {
								heap.Push(&sy, element)
				}
		}
		//fmt.Println(sy)
		//fmt.Println(xstar)
		min := math.Inf(1)
		//myPoints := make([]int, 2)
		length := len(sy)
		var innerLoop int
		if (length - 1) < 15 {
			innerLoop = length - 1
		} else {
			innerLoop = 15
		}
		for index, point := range sy {
				for i:=index +1; i <= innerLoop; i++ {
						currDist := distance(point.value, sy[i].value)
						if min > currDist{
			 			 		min = currDist
			 			 		//myPoints = []int{index, index+1}
			 			}
				}
		}
		//fmt.Println(min)
		trueMin := math.Min(min, delta)
		if trueMin == math.Inf(1) {
				trueMin = 0
		}
		fmt.Printf("%.15f\n",trueMin)
}
func distance(p1 []int, p2 []int) float64 {
	  var dist float64
		deltax := float64(p1[0] - p2[0])
		deltay := float64(p1[1] - p2[1])
		dist = math.Sqrt(deltax*deltax + deltay*deltay)
		return dist
}
func findClosest(q PriorityQueue, retVal *float64, goNo int, wg *sync.WaitGroup) {
		min := math.Inf(1)
		//myPoints := make([]int, 2)
		length := len(q)
		//fmt.Println("length ", length, "IN GOROUTINE: ", goNo)
		for index, point := range q {
			 //fmt.Println("INDEX ", index, "IN GOROUTINE: ", goNo)
			 if index > length-2{
					break
			 }
			 currDist := distance(point.value, q[index + 1].value)
			 if min > currDist{
						min = currDist
						//myPoints = []int{index, index+1}
			 }
		}
		//fmt.Println("GO:", goNo, " ",myPoints)
		//fmt.Println(min)
		*retVal = min
		wg.Done()
}
