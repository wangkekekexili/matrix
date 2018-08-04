package laff

import "fmt"

type ErrVectorSizeMismatch struct {
	left, right int
}

func (e ErrVectorSizeMismatch) Error() string {
	return fmt.Sprintf("vector size mismatch: %v and %v", e.left, e.right)
}
