package templater

import "testing"

const wrongAnswer = "Wrong answer"

func TestEndWithTrue(t *testing.T) {
	str := "abcde"
	end := "de"
	if endWith(str, end) != true {
		t.Error(wrongAnswer)
	}
}

func TestEndWithFalse(t *testing.T) {
	str := "abcce"
	end := "de"
	if endWith(str, end) != false {
		t.Error(wrongAnswer)
	}
}

func TestEndWithEndBigger(t *testing.T) {
	str := "abcde"
	end := "asdabcde"
	if endWith(str, end) != false {
		t.Error(wrongAnswer)
	}
}
