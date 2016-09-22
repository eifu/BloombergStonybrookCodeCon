package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var s string
	var i,  myPriority, myPriority_index int
	var myValue []int

	var priority []int

	scanner.Scan()
	s = scanner.Text()
	arrays := strings.Split(s, " ")

	for _, e := range arrays {
		i, _ = strconv.Atoi(e)
		myValue = append(myValue, i)
	}

	scanner.Scan()
	s = scanner.Text()
	arrays = strings.Split(s, " ")	


	for _, e := range arrays {
		i, _ = strconv.Atoi(e)
		priority = append(priority, i)
	}

	// total = myValue[0]
	myPriority_index = myValue[1]
	myPriority = priority[myPriority_index]

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

	i = myValue[0]-1
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
		fmt.Print((ind_in_same+1)*2)
	}else{

		fmt.Print(count*2)
	}
}
