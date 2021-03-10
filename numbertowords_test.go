package numbertowords

import "testing"

func TestInvalidInput(t *testing.T) {
	result, err := Convert(-1)
	if result != "" || err == nil {
		t.Fatal("failed test for input -1")
	}

	result, err = Convert(100000)
	if result != "" || err == nil {
		t.Fatal("failed test for input -1")
	}

}

func TestUnits(t *testing.T) {
	testcases := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	for number, word := range testcases {
		result, err := Convert(number)
		if result != word || err != nil {
			t.Fail()
		}
	}
}
