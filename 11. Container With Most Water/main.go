import (
	"fmt"
	"math"
)

// Given n non-negative integers a1, a2, ..., an ,
// where each represents a point at coordinate (i, ai).
// n vertical lines are drawn such that the two endpoints of
// line i is at (i, ai) and (i, 0).
// Find two lines, which together with x-axis forms a container,
// such that the container contains the most water.
//
//Note: You may not slant the container and n is at least 2.

func maxArea(height []int) int {
	var max int
	for xID, x := range height {
		for yID, y := range height {
			distance := int(math.Abs(float64(xID-yID)) * math.Min(float64(x), float64(y)))
			if distance > max {
				max = distance
			}
		}
	}

	return max
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
