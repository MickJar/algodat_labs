package main

import (
				"fmt"
				"os"
				"bufio"
				"io"
				"strings"
				"strconv"
				"container/heap"
)
type Item struct {
	value    []string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
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
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value []string, priority int) {
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
	i := 0
	text, _ := reader.ReadString('\n')
	inputArgs := strings.Split(text, " ")
	noCities, _ := strconv.Atoi(inputArgs[0])
	inputArgs[1] = strings.TrimSuffix(inputArgs[1], "\n")
	noConnections , _ := strconv.Atoi(inputArgs[1])
	pq := make(PriorityQueue, noConnections)
	for {
				text, err  := reader.ReadString('\n')
				if err == io.EOF {
					break
				}
				keyValue := strings.Split(text, " ")
				keyValue[0] = strings.TrimSuffix(keyValue[0], "\n")
				if len(keyValue) > 1 {
						keyValue[2] = strings.TrimSuffix(keyValue[2], "\n")

						myInt, _ := strconv.Atoi(keyValue[2])
						pq[i] = &Item{
										value: []string{keyValue[0],keyValue[1]},
										priority: myInt,
										index: i,
						}
						i++
				}else {
						//unvisited = append(unvisited, keyValue[0])
				}
		}
		//fmt.Println("------------UNVISITED------------")
		//fmt.Println(unvisited )
		//fmt.Println("------------DISTS------------")
		//fmt.Println(pq)
		heap.Init(&pq)
		sum := 0
		visited := make([]string, 0)
		item := heap.Pop(&pq).(*Item)
		loc1 := item.value[0]
		loc2 := item.value[1]
		visited = append(visited, loc1)
		visited = append(visited, loc2)
		sum += item.priority
		//fmt.Println(unvisited)
		tempq := make([]*Item, 0)
		for len(visited) < noCities {
						if(pq.Len() == 0) {
								fmt.Println(-1)
								os.Exit(0)
						}
						currItem := heap.Pop(&pq).(*Item)
						if !contains(visited, currItem.value[0]) && contains(visited, currItem.value[1]) {
								visited = append(visited, currItem.value[0])
								sum += currItem.priority
								for _, i := range tempq {
												heap.Push(&pq, i)
								}
								tempq = tempq[:0]
						} else if !contains(visited, currItem.value[1]) && contains(visited, currItem.value[0]) {
								sum += currItem.priority
								visited = append(visited, currItem.value[1])
								for _, i := range tempq {
												heap.Push(&pq, i)
								}
								tempq = tempq[:0]
						} else {
							tempq = append(tempq, currItem)
						}
		}
		fmt.Println(sum)
		//for pq.Len() > 0 {
		//	item := heap.Pop(&pq).(*Item)
		//	fmt.Printf("%.2d:%s ", item.priority, item.value)
		//}

}
