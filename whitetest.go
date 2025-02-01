// package main


func IsEven(n int) bool {
	if n % 2 == 0 {
		return true
	}
	return false
}

func TestIsEven(t *testing.T) {
	if !IsEven(2) {
		t.Error("2 should be even")
	}
	if IsEven(3) {
		t.Error("3 should not be even")
	}
}