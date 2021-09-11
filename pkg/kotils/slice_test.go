package kotils

import (
	"testing"
)

type Operation struct {
	item   int
	items  []int
	exists bool
}
type Operations []Operation

func TestMany(t *testing.T) {
	ops := Operations{
		Operation{
			item:   0,
			items:  []int{0, 3, 6},
			exists: true,
		},
		Operation{
			item:   -1,
			items:  []int{0, 3, 6},
			exists: false,
		},
		Operation{
			item:   3,
			items:  []int{0, 3, 6},
			exists: true,
		},
		Operation{
			item:   5,
			items:  []int{0, 3, 6},
			exists: false,
		},
		Operation{
			item:   6,
			items:  []int{0, 3, 6},
			exists: true,
		},
		Operation{
			item:   100,
			items:  []int{0, 3, 6},
			exists: false,
		},
		Operation{
			item:   100,
			items:  []int{},
			exists: false,
		},
		Operation{
			item:   5,
			items:  []int{5},
			exists: true,
		},
	}
	for _, op := range ops {
		if IntInSlice(op.item, op.items) != op.exists {
			t.Errorf("the value %v in %v gives %v", op.item, op.items, op.exists)
		}
	}
}
