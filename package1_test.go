package package1

import "testing"

func TestSum(t *testing.T) {
  if Sum(1, 2) != 3 {
    t.Error("Expected 1 + 2 to equal 3")
  }
}

func TestSub(t *testing.T) {
  if Sub(2, 1) != 1 {
    t.Error("Expected 2 - 1 to equal 1")
  }
}
