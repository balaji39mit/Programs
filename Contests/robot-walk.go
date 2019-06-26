package main

import "fmt"

func RobotWalk(n, r, c, x, y int64, direction string) (int64, int64) {
	grid := make([][]bool, r)
	for k := range grid {
		grid[k] = make([]bool, c)
	}
	x, y = x-1, y-1
	grid[x][y] = true
	for i := int64(0); i < n; i++ {
		move := direction[i]
		switch move {
		case 'E':
			col := y + 1
			for col < c {
				if !grid[x][col] {
					grid[x][col] = true
					break
				}
				col++
			}
			y = col
		case 'W':
			col := y - 1
			for col >= 0 {
				if !grid[x][col] {
					grid[x][col] = true
					break
				}
				col--
			}
			y = col
		case 'N':
			row := x - 1
			for row >= 0 {
				if !grid[row][y] {
					grid[row][y] = true
					break
				}
				row--
			}
			x = row
		case 'S':
			row := x + 1
			for row < r {
				if !grid[row][y] {
					grid[row][y] = true
					break
				}
				row++
			}
			x = row
		default:
			return 0, 0
		}
	}
	return x + 1, y + 1
}

func main() {
	var t uint8
	fmt.Scanf("%d", &t)
	for i := uint8(1); i <= t; i++ {
		var n, r, c, x, y int64
		var direction string
		fmt.Scanf("%d %d %d %d %d", &n, &r, &c, &x, &y)
		fmt.Scanf("%s", &direction)
		xDest, yDest := RobotWalk(n, r, c, x, y, direction)
		fmt.Println(fmt.Sprintf("Case #%d: %d %d", i, xDest, yDest))
	}
}
