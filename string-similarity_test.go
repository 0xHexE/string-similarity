package string_similarity

import "testing"

func TestAbsValues(t *testing.T) {
	cases := []struct {
		input, result string
	}{
		{"test", "test"},
		{"test query", "testquery"},
		{"", ""},
	}

	for _, c := range cases {
		got := AbsValues(c.input)
		if got != c.result {
			t.Errorf("Input: %v %v", c.input, got)
		}
	}

}

func TestCompareString(t *testing.T) {
	cases := []struct {
		input1, input2 string
		result         float64
	}{
		{"french", "quebec", 0.0},
		//{ "france", "france", 1 },
		//{ "fRaNce", "france", 0.2 },
		//{ "healed", "sealed", 0.8 },
		//{ "web applications", "applications of the web", 0.7878787878787878 },
		//{ "this will have a typo somewhere", "this will huve a typo somewhere", 0.92 },
		//{ "Olive-green table for sale, in extremely good condition.", "For sale: table in very good  condition, olive green in colour.", 0.6060606060606061 },
		//{ "Olive-green table for sale, in extremely good condition.", "For sale: green Subaru Impreza, 210,000 miles", 0.2558139534883721 },
		//{ "Olive-green table for sale, in extremely good condition.", "Wanted: mountain bike with at least 21 gears.", 0.1411764705882353 },
		//{ "this has one extra word", "this has one word", 0.7741935483870968 },
		//{ "a", "a", 1 },
		//{ "a", "b", 0 },
		//{ "\"\"", "\"\"", 1 },
		//{ "a", "\"\"", 0 },
		//{ "\"\"", "a", 0 },
		//{ "apple event", "apple    event", 1 },
		//{ "iphone", "iphone x", 0.9090909090909091 },
	}

	for _, c := range cases {
		got := CompareString(c.input1, c.input2)
		if got != c.result {
			t.Errorf("Input1: %q Input2: %q Got result: %v Expected Result: %v", c.input1, c.input2, got, c.result)
		}
	}

}
