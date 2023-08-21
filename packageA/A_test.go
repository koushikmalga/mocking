package packageA

import (
	"mocking/packageB"
	"testing"
)

// Mock struct sor mocking C
// Create a mock struct which implements methods as packageC
type mock struct{}

func (m *mock) Sub(ele1, ele2 int) int {
	return 3
}

func TestA(t *testing.T) {
	mockc1 := &mock{}

	// Building NewB and NewA from the mock instead of calling packageC
	a := NewA(packageB.NewB(mockc1))

	var ele1, ele2 int
	ele1 = 5
	ele2 = 2
	result := a.func1(ele1, ele2)

	exp_result := 10

	if result != exp_result {
		t.Errorf("Expected result = %d, but got =%d", exp_result, result)
	}

}
