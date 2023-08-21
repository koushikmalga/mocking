package packageC

// C Interface
type CIfc interface {
	Sub(int, int) int
}
type C struct{}

// Sub method to subtract to elements
func (c *C) Sub(ele1, ele2 int) int {

	return ele1 - ele2
}
