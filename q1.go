package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
	"strings"
)


func takeValue(s string) []int{

	var myValue []int
	arrays := strings.Split(s, " ")

	for _, e := range arrays {
		i, _ := strconv.Atoi(e)
		myValue = append(myValue, i)
	}
	return myValue
}

type Node struct{
	index int
	priority int
}

type Queue struct{
	size int
	head *Node
	tail *Node
}

func (q *Queue) enqueue(index, priority int){
	q.size ++

	n := 

}

func (q *Queue) dequeue(){

}

func Solve(myValue, priority []int) int{

	total = myValue[0]
	myPriority_index := myValue[1]
	myPriority := priority[myPriority_index]

	// count := 1

	max := 0
	max_index := 0

	for i,e := range priority{
		if e > max{
			max = e
			max_index = i
		}
	}

	num_same := 0
	ind_in_same := 0
	bef_in_same := 0
	aft_in_same := 0
	for i, e := range priority{
		if e == myPriority{
			num_same ++
			if i < myPriority_index{
				ind_in_same ++
			} 
			if i < max_index{
				bef_in_same ++
			}
			if max_index < i {
				aft_in_same ++
			}	
		}
	}

	sort.Ints(priority)

	i := myValue[0]-1
	big := 0
	// fmt.Println(priority[i])
	for priority[i] > myPriority{
		big ++
		i --
	}

	if myPriority_index > max_index {
		ind_in_same = ind_in_same - bef_in_same
		aft_in_same = 0
	} 


	count := 1 + big + ind_in_same + aft_in_same
	// fmt.Println(priority, myPriority)
	// fmt.Println(count,big, ind_in_same, aft_in_same)
	if myPriority == max{
		return (ind_in_same+1)*2
	}else{

		return count*2
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var s string
	var myValue []int
	var priority []int

	scanner.Scan()
	s = scanner.Text()
	myValue = takeValue(s)

	scanner.Scan()
	s = scanner.Text()
	priority = takeValue(s)

	fmt.Print(Solve(myValue, priority))
}
