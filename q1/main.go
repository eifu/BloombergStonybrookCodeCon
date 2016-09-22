package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var s string
	var i, total, myPriority, myPriority_index int
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

	total = myValue[0]
	myPriority_index = myValue[1]
	myPriority = priority[myPriority_index]

	count := 1
	max := 0
	max_index := 0
	for  i, e := range priority{
		if max < e{
			max = e
			max_index = i
		}
	}

	for i = total -1; i >= max_index; i-- {
		if i == myPriority_index{
			continue
		}else if priority[i] >= myPriority{
			count ++
		}
	}

	for i = max_index-1; i >= 0; i--{
		if priority[i] == myPriority{
			if i < myPriority_index{
				count ++
			}
		}else if priority[i] > myPriority{
			count ++
		}
	}

	fmt.Print(count*2)

}
