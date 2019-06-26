package main

import (
	"testing"

	"fmt"
)

func TestRobotWalk(t *testing.T) {
	x, y := RobotWalk(15, 3, 6, 2, 3, "EEWNSEEEENNWWWW")
	if x != 1 || y != 1 {
		t.Errorf(fmt.Sprintf("Actual : %v %v", x, y))
	}
}
