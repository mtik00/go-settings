package go_settings

import (
	"testing"
)

var getTests = []struct {
	key   string
	value bool
}{
	{"one.two", true},
	{"three", false},
}

var TEST_STRING = `
    {
        "one.two": true, // hanging comment

        // line comment
        "three": false
    }
`

func TestMain(t *testing.T) {
	s := NewSettings([]byte(TEST_STRING))

	for _, test := range getTests {
		rvalue := s.Get(test.key)
		if rvalue != test.value {
			t.Errorf("key [%v] was [%v]; wanted [%v]", test.key, rvalue, test.value)
		}
	}
}

func TestFile(t *testing.T) {

}
