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


func TestCase2(t *testing.T){
	given := []int{4, 0}
	given2 := []int{1, 2, 3, 4}

	a := Solve(given, given2)

	if a != 8{
		t.Errorf("should be 8 but %d", a)
	}
}

func TestCase3(t *testing.T){
	given := []int{4, 1}
	given2 := []int{1, 2, 3, 4}

	a := Solve(given, given2)

	if a != 6{
		t.Errorf("should be 6 but %d", a)
	}
}

func TestCase4(t *testing.T){
	given := []int{4, 3}
	given2 := []int{1, 2, 3, 4}

	a := Solve(given, given2)

	if a != 2{
		t.Errorf("should be 2 but %d", a)
	}
}

func TestCase5(t *testing.T){
	given := []int{6, 0}
	given2 := []int{1,1,9,1,1,1}

	a := Solve(given, given2)

	if a != 10{
		t.Errorf("should be 10 but %d", a)
	}
}

func TestCase6(t *testing.T){
	given := []int{6, 1}
	given2 := []int{1,1,9,1,1,1}

	a := Solve(given, given2)

	if a != 12{
		t.Errorf("should be 12 but %d", a)
	}
}

func TestCase7(t *testing.T){
	given := []int{6, 2}
	given2 := []int{1,1,9,1,1,1}

	a := Solve(given, given2)

	if a != 2{
		t.Errorf("should be 2 but %d", a)
	}
}

func TestCase8(t *testing.T){
	given := []int{6, 3}
	given2 := []int{1,1,9,1,1,1}

	a := Solve(given, given2)

	if a != 4{
		t.Errorf("should be 4 but %d", a)
	}
}

func TestCase9(t *testing.T){
	given := []int{6, 4}
	given2 := []int{1,1,9,1,1,1}

	a := Solve(given, given2)

	if a != 6{
		t.Errorf("should be 6 but %d", a)
	}
}

func TestCase10(t *testing.T){
	given := []int{6, 5}
	given2 := []int{1,1,9,1,1,1}

	a := Solve(given, given2)

	if a != 8{
		t.Errorf("should be 8 but %d", a)
	}
}
func TestCase11(t *testing.T){
	given := []int{6, 0}
	given2 := []int{1,1,1,1,1,1}

	a := Solve(given, given2)

	if a != 2{
		t.Errorf("should be 2 but %d", a)
	}
}

func TestCase12(t *testing.T){
	given := []int{6, 1}
	given2 := []int{1,1,1,1,1,1}

	a := Solve(given, given2)

	if a != 4{
		t.Errorf("should be 4 but %d", a)
	}
}

func TestCase13(t *testing.T){
	given := []int{6, 2}
	given2 := []int{1,1,1,1,1,1}

	a := Solve(given, given2)

	if a != 6{
		t.Errorf("should be 6 but %d", a)
	}
}

func TestCase14(t *testing.T){
	given := []int{6, 5}
	given2 := []int{1,1,1,1,1,1}

	a := Solve(given, given2)

	if a != 12{
		t.Errorf("should be 12 but %d", a)
	}
}

func TestCase15(t *testing.T){
	given := []int{6, 0}
	given2 := []int{1,1,3,4,1,1}

	a := Solve(given, given2)

	if a != 10{
		t.Errorf("should be 10 but %d", a)
	}
}