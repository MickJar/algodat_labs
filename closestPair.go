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
		value []float64
		priority float64
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
	// item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value []float64, priority float64) {
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
				if err == io.EOF {
					break
				}
				myNums := strings.Split(text, " ")
				myNums[1] = strings.TrimSuffix(myNums[1], "\n")
				x, _:= strconv.ParseFloat(myNums[0], 64)
				y, _:= strconv.ParseFloat(myNums[1], 64)
				pointx := &Item{
										value: []float64{x,y},
										priority: x,
										index: i,
				}
				pointy := &Item{
							value: []float64{x,y},
							priority: y,
							index: i,
				}
				heap.Push(&px, pointx)
				heap.Push(&py, pointy)
				i++
		}
		// for _, element := range px{
		// 	fmt.Println(*element)
		// }
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
		// if delta == math.Inf(1) {
		// 				delta = 10000
		// }
		// fmt.Println(delta)
		// for _, element := range secondHalf{
		// 	fmt.Println(element)
		// }
		xstar := heap.Pop(&firstHalf).(*Item)
		sy := make(PriorityQueue, 0)
		heap.Init(&sy)
		// fmt.Println(xstar)
		for py.Len() > 0 {
				element := heap.Pop(&py).(*Item)
				if math.Abs(element.value[0] - xstar.value[0]) < delta {
								heap.Push(&sy, element)
				}
		}
		// fmt.Println(sy.Len())
		// for _, element := range sy{
		// 	fmt.Print(element.value[0], " ")
		// 	fmt.Println(element.value[1])
		// }
		//fmt.Println(sy)
		//fmt.Println(xstar)
		min := math.Inf(1)
		//myPoints := make([]int, 2)
		length := len(sy)
		// var innerLoop int
		// if (length - 1) < 1000 {
		// 	innerLoop = length
		// } else {
		// 	innerLoop = 1000
		// }
		for index, point := range sy {
				for i:=index +1; i < length; i++ {
					if math.Abs(point.value[1] - sy[i].value[1]) > delta{
						break
					}
						currDist := distance(point.value, sy[i].value)
						if min > currDist{
			 			 		min = currDist
			 			 		// fmt.Println(point.value , " ", sy[i].value)
			 			 		//myPoints = []int{index, index+1}
			 			}
				}
		}
		// for _, element := range sy{
		// 	fmt.Println(*element)
		// }
		// fmt.Println(min)
		trueMin := math.Min(min, delta)
		if trueMin == math.Inf(1) {
				trueMin = 0
		}
		fmt.Printf("%.6f\n",trueMin)
}	
func distance(p1 []float64, p2 []float64) float64 {
	  var dist float64
		deltax := float64(p1[0] - p2[0])
		deltay := float64(p1[1] - p2[1])
		dist = math.Sqrt(deltax*deltax + deltay*deltay)
		return dist
}
func findClosest(q PriorityQueue, retVal *float64, goNo int, wg *sync.WaitGroup) {
		min := math.Inf(1)
        //myPoints := make([]int, 2)
        // length := q.Len()
        //fmt.Println("length ", length, "IN GOROUTINE: ", goNo)
        prevElement := heap.Pop(&q).(*Item)
        for q.Len() > 0 {
            currElement := heap.Pop(&q).(*Item)
             //fmt.Println("INDEX ", index, "IN GOROUTINE: ", goNo)
             if currElement == nil{
                    break
             }
             currDist := distance(prevElement.value, currElement.value)
             if min > currDist{
                if goNo == 2{
                // fmt.Println(prevElement.value, " ", currElement.value)
                }
                        min = currDist
                        //myPoints = []int{index, index+1}
             }
             prevElement = currElement
        }
        //fmt.Println("GO:", goNo, " ",myPoints)
        //fmt.Println(min)
        *retVal = min
        wg.Done()
}
