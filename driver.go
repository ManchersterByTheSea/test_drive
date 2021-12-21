package test_drive

import (
	"fmt"
	"github.com/ManchersterByTheSea/test_common/register"
)

func init() {
	register.Register("test", &Test{})
}

type Test struct {
}

func (t Test) Create() {
	fmt.Println("Create test")

}
func (t Test) Update(a string, b int) error {

	fmt.Println("Update test", a, b)
	return nil
}
