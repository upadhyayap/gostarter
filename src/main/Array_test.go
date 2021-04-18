package main

import (
	"log"
	"testing"
)

func TestArray(t *testing.T) {
	arr := testingArray()

	if len(arr) < 3 {
		t.FailNow()
		t.Error("Failed")
		log.Println("length is less than 3")
	}
}

func TestAvg(t *testing.T) {

	// table tests or data driven test
	type test struct {
		data []int
		avg  float64
	}

	tests := []test{
		test{[]int{1, 2}, 1.5},
		test{[]int{2, 3}, 3},
	}

	for _, test := range tests {
		avg := avgArray(test.data)
		if avg != test.avg {
			t.Fail()
			t.Error("Expected ", test.avg, "found ", avg)
		}
	}

}

func ExampleAvg() {
	avgArray([]int{1, 2})
	//Output
	//1.5
}

func BenchmarkAvg(b *testing.B) { // go test -bench Avg
	for i := 0; i < b.N; i++ {
		avgArray([]int{1, 2})
	}
}
