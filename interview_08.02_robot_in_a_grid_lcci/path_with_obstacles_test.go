package lost_robot

import "testing"

func TestPathWithObstacles(t *testing.T) {

	grids := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}

	paths := pathWithObstacles(grids)

	t.Logf("paths=%+v", paths)
}
