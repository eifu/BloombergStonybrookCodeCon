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



	for i = 0; i < len(priority); i++{
		if priority[i] > max{
			max = priority[i]
			max_index = i
		}
	}

	dub := 0
	for i = 0; i < myPriority_index; i ++{
		if i > max_index && priority[i] == myPriority{
			dub ++
		}
	}
	

	pos := 0
	if max != myPriority{
		if myPriority_index > max_index{
			for i = max_index+1; i < myPriority_index; i++{
				if i != myPriority_index && priority[i] == myPriority{
					pos ++
				}
			} 
		}else {
			for i = max_index+1; i < len(priority); i++{
				if i != myPriority_index && priority[i] == myPriority{
					pos ++
				}
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




	count := 1 + big + dub + pos
	fmt.Println(priority, myPriority)
	fmt.Println(count,big, dub, pos)

	fmt.Print(count*2)

}
