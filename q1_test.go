package main

import(
	"testing"
)


func testEqualSlice(a, b []int) bool{
	if a == nil && b == nil{
		return true 
	}else if a == nil || b == nil{
		return false
	}else if len(a) != len(b){
		return false
	}else {
		for i, a_element := range a {
			if a_element != b[i]{
				return false
			}
		}
		return true
	}
}



func TestTakeValue(t *testing.T){

	given := "1 0"
	given2 := "5"

	result := []int{1, 0}
	result2 := []int{5}

	a := takeValue(given)
	if !testEqualSlice(a, result){
		t.Errorf("should be [1 0], but %v", a)
	}

	a = takeValue(given2)
	if !testEqualSlice(a, result2){
		t.Errorf("should be [5], but %v", a)
	}

}

func TestCase1(t *testing.T){
	given := []int{1, 0}
	given2 := []int{5}

	a := Solve(given, given2)

	if a != 2{
		t.Errorf("should be 2 but %d", a)
	}
}