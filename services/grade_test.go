package services_test

import (
	"fmt"
	"gotest/services"
	"testing")

func TestCheckGrade(t *testing.T) {
	type CheckGradeTest struct {
		name_grade string
		expected   string
		score      int
	}

	check_result := []CheckGradeTest{
		{name_grade: "a", expected: "A", score: 100},
		{name_grade: "b", expected: "B", score: 80},
		{name_grade: "c", expected: "C", score: 70},
		{name_grade: "d", expected: "D", score: 60},
		{name_grade: "f", expected: "F", score: 50},
	}

	for _, val := range check_result {
  t.Run(val.name_grade,func (t *testing.T)  {
	grade := services.CheckGrade(val.score)
	expected := val.expected

	if grade != expected {
		t.Errorf("got %v, expected %v", grade, expected)
	}

  })
}
}


func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.CheckGrade(100)
	}
}


func ExampleCheckGrade (){
	grade := services.CheckGrade(100)
	fmt.Println(grade)
	// Output: A
}