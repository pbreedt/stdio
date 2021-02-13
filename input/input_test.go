package input_test

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"testing"

	"github.com/pbreedt/stdio/input"
)

var read, write, origStdIn *os.File

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	fmt.Println("SETUP")
	r, w, err := os.Pipe()
	read = r
	write = w
	if err != nil {
		panic(err)
	}

	origStdIn = os.Stdin
	os.Stdin = read
}

func teardown() {
	fmt.Println("TEARDOWN")
	// Restore stdin right after the test.
	write.Close()
	os.Stdin = origStdIn
}

func writeToStdIO(testData string) {
	_, err := write.Write([]byte(testData + "\n"))
	if err != nil {
		panic(err)
	}
	// write.Close()
}

type IntTestCases struct {
	input    string
	expValue int
	expError bool
}

func TestReadInt(t *testing.T) {
	testInputs := []IntTestCases{
		{"0", 0, false},
		{"1", 1, false},
		{"10000000", 10000000, false},
		{fmt.Sprintf("%d", math.MinInt64), math.MinInt64, false},
		{fmt.Sprintf("%d", math.MaxInt64), math.MaxInt64, false},
		{"1.5", 0, true}, // from here: expect errors
		{"alpha", 0, true},
	}

	for _, test := range testInputs {
		writeToStdIO(test.input)

		getint, err := input.ReadInt("read test int (" + test.input + ")\n")
		if test.expError && err != nil {
			fmt.Printf("Error expected, error occured (%v)\n", err)
		} else if !test.expError && err != nil {
			t.Errorf("Input error: %v", err)
		} else {
			if getint != test.expValue {
				t.Error() // to indicate test failed
			}
		}

	}
}

type FloatTestCases struct {
	input    string
	expValue float64
	expError bool
}

func TestReadFloat(t *testing.T) {
	testInputs := []FloatTestCases{
		{"0.0", 0, false},
		{"1", 1.0, false},
		{"1.5", 1.5, false},
		{"11111111.22222222", 11111111.22222222, false},
		{fmt.Sprintf(strconv.FormatFloat(math.SmallestNonzeroFloat64, 'f', -1, 64)), math.SmallestNonzeroFloat64, false},
		{fmt.Sprintf("%f", math.MaxFloat64), math.MaxFloat64, false},
		{"alpha", 0, true}, // from here: expect errors
	}

	for _, test := range testInputs {
		writeToStdIO(test.input)

		getfloat, err := input.ReadFloat("read test float (" + test.input + ")\n")

		if test.expError && err != nil {
			fmt.Printf("Error expected, error occured (%v)\n", err)
		} else if !test.expError && err != nil {
			t.Errorf("Input error: %v", err)
		} else {
			if getfloat != test.expValue {
				t.Errorf("Input %f did not match expected value %f", getfloat, test.expValue)
			}
		}
	}
}

type StringTestCases struct {
	input string
}

func TestReadString(t *testing.T) {
	testInputs := []StringTestCases{
		{"100"},
		{"This is test input"},
		{"abcdef"},
		{"±!@#$%^&*()_+"},
		{"àäæęūį"},
		{"aaa\tbbb"},
	}

	for _, test := range testInputs {
		writeToStdIO(test.input)

		getstr, err := input.ReadString("read test string (" + test.input + ")\n")
		if err != nil {
			t.Fatalf("Input error: %v", err)
		}

		if getstr != test.input {
			t.Error("Input", getstr, "did not match expected value", test.input)
		}
	}

}

type BoolTestCases struct {
	input    string
	expValue bool
	expError bool
}

func TestReadBool(t *testing.T) {
	testInputs := []BoolTestCases{
		{"1", true, false},
		{"T", true, false},
		{"True", true, false},
		{"true", true, false},
		{"TRUE", true, false},
		{"0", false, false},
		{"F", false, false},
		{"False", false, false},
		{"false", false, false},
		{"FALSE", false, false},
		{"Yes", false, true}, // from here: expect error (Y/N not allowed)
		{"YES", false, true},
		{"Y", false, true},
		{"No", false, true},
		{"NO", false, true},
		{"N", false, true},
	}

	for _, test := range testInputs {
		writeToStdIO(test.input)

		getbool, err := input.ReadBool("read test boolean ("+test.input+")\n", false) // allowYN = false
		if !test.expError && err != nil {
			t.Fatalf("Input error: %v", err)
		}

		if getbool != test.expValue {
			t.Error("Input", getbool, "did not match expected value", test.expValue)
		}
	}

}

func TestReadBoolWithYN(t *testing.T) {
	testInputs := []BoolTestCases{
		{"1", true, false},
		{"T", true, false},
		{"True", true, false},
		{"true", true, false},
		{"TRUE", true, false},
		{"0", false, false},
		{"F", false, false},
		{"False", false, false},
		{"false", false, false},
		{"FALSE", false, false},
		{"Yes", true, false}, // this time, from here do not expect error (Y/N is allowed)
		{"YES", true, false},
		{"Y", true, false},
		{"No", false, false},
		{"NO", false, false},
		{"N", false, false},
	}

	for _, test := range testInputs {
		writeToStdIO(test.input)

		getbool, err := input.ReadBool("read test boolean ("+test.input+")\n", true) // allowYN = true
		if !test.expError && err != nil {
			t.Fatalf("Input error: %v", err)
		}

		if getbool != test.expValue {
			t.Error("Input", getbool, "did not match expected value", test.expValue)
		}
	}

}
