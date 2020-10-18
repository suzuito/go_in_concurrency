package main

import "fmt"

const N = int64(1_000_000)

// const N = int64(10_000_000_000) // 限界

func main() {
	const start = int64(2)

	nNonPrimes := int64(0) // 素数ではない数の個数
	k := start
	marks := [N]bool{false}

	for {
		// fmt.Printf("Start at %d\n", k)
		if k*k > N {
			break
		}

		// Marking
		doMark(k, &nNonPrimes, &marks)

		// Find next k
		for i := k + 1; i < N; i++ {
			if !marks[i] {
				k = i
				break
			}
		}
	}

	fmt.Println(N - nNonPrimes - 2)
}

// doMarkはk^2以上n以下のkの倍数を「素数ではない」とマークする
func doMark(k int64, nNonPrimes *int64, marks *[N]bool) {
	i := k
	for {
		j := k * i
		if j >= N {
			break
		}
		// fmt.Println("foo", j)
		if !(*marks)[j] {
			(*marks)[j] = true
			(*nNonPrimes)++
		}
		// fmt.Printf("%d\n", j)
		i++
	}
}
