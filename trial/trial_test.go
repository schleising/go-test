package trial

import "testing"

func TestHello(t *testing.T) {
	result := Trial()

	if result == false {
		t.Fatal("Testing Failed")
	} else {
		t.Log("Testing Passed")
	}
}
