package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var reg = regexp.MustCompile(`^([+\-]?)(0*)([0-9]+)`)

func myAtoi(str string) int {
	fields := strings.Fields(str)

	if len(fields) > 0 {
		field := fields[0]
		matches := reg.FindStringSubmatch(field)
		if len(matches) < 2 {
			return 0
		}

		neg := false
		var text string

		switch len(matches) {
		case 4:
			text = matches[3]
			neg = strings.HasPrefix(matches[1], "-")
		case 3:
			text = matches[2]
			neg = strings.HasPrefix(matches[1], "-")
		default:
			text = matches[1]
		}

		if len(text) >= 20 {
			if neg {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}

		if len(text) == 19 {
			for _, c := range "9223372036854775807" {
				x, err := strconv.Atoi(string(c))
				if err != nil {
					return 0
				}

				for _, target := range text {
					y, err := strconv.Atoi(string(target))
					if err != nil {
						return 0
					}
					if y > x {
						if neg {
							return math.MinInt32
						} else {
							return math.MaxInt32
						}
					}
				}
			}
		}
		n, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(err)
			return 0
		}

		if neg {
			n = -n
		}

		if n > math.MaxInt32 {
			return math.MaxInt32
		}
		if n < math.MinInt32 {
			return math.MinInt32
		}
		return n
	}

	return 0
}

func main() {
	//fmt.Println(myAtoi("-100000000"))
	//fmt.Println(myAtoi("-000000000000000000000000000000000000000000000000001"))
	fmt.Println(myAtoi("+-2"))

	// 9223372036854775807
	fmt.Println("9223372036854775808")
}
