package main

func main() {
	sl := []int{}
	for i := range 10 {
		sl = append(sl, i)
	}
	for _, val := range sl {
		if isEven(val) {
			println("Even")
		} else {
			println("Odd")
		}
	}
}

func isEven(num int) bool {
	return num%2 == 0
}
