package main

import "fmt"

func main() {
	const n = int64(1_000_000_000)
	const start = int64(2)

	l := int64(0) // 素数ではない数の個数
	k := start
	marks := [n]bool{false}
	// fmt.Println(marks)
	for {
		// fmt.Printf("Start at %d\n", k)

		if k*k > n {
			break
		}

		i := k
		for {
			j := k * i
			if j >= n {
				break
			}
			// fmt.Println("foo", j)
			if !marks[j] {
				marks[j] = true
				l++
			}
			// fmt.Printf("%d\n", j)
			i++
		}

		// Find min having false
		for i := k + 1; i < n; i++ {
			if !marks[i] {
				k = i
				break
			}
		}
	}

	// s := 0
	// for i := start; i < n; i++ {
	// 	if !marks[i] {
	// 		s++
	// 	}
	// }
	// fmt.Printf("%d\n", s)

	fmt.Println(n - l - 2)
}
