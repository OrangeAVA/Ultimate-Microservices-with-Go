package main

func main() {

	// for loop with initialization, condition and post
	for i := 0; i < 5; i++ {
		println(i)
	}

	// for loop with single condition
	i := 0
	for i < 5 {
		println(i)
		i++
	}

	// for loop with no condition
	i = 0
	for {
		println(i)
		i++
		if i >= 5 {
			break
		}
	}

	// for loop with continue
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		println(i)
	}

	// for loop with range
	s := []int{1, 2, 3}
	for k, v := range s {
		println(k, v)
	}
}
