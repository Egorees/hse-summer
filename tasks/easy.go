package main

import "fmt"

func HelloWorld() {
	fmt.Print("Hello, World!")
}

func Sum() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		panic(err)
	}
	fmt.Print(a + b)
}

func EvenOrOdd() {
	var a int
	_, err := fmt.Scan(&a)
	if err != nil {
		panic(err)
	}
	if a&1 == 0 {
		fmt.Print("Even")
	} else {
		fmt.Print("Odd")
	}
}

func MaxOfThree() {
	var a, b, c int
	_, err := fmt.Scan(&a, &b, &c)
	if err != nil {
		panic(err)
	}
	switch {
	case a >= b && a >= c:
		fmt.Print(a)
	case b >= a && b >= c:
		fmt.Print(b)
	case c >= a && c >= b:
		fmt.Print(c)
	default:
		fmt.Print("Something went wrong :(")
	}
}

func Factorial() {
	var a int
	_, err := fmt.Scan(&a)
	if err != nil {
		panic(err)
	}
	ans := 1
	for i := 1; i <= a; i++ {
		ans *= i
	}
	fmt.Print(ans)
}

func CharCheck() {
	alph := "aeuoiyAEUOIY"
	var c string
	_, err := fmt.Scan(&c)
	if err != nil {
		panic(err)
	}
	sym := rune(c[0])
	flag := false
	for _, char := range alph {
		if sym == char {
			flag = true
			break
		}
	}
	if flag {
		fmt.Print("YEES")
	} else {
		fmt.Print("No.")
	}
}

func PrimeNumbers() {
	var a int
	_, err := fmt.Scan(&a)
	if err != nil {
		return
	}

	notPrime := make([]bool, a+1)
	notPrime[0] = true
	notPrime[1] = true

	for i := 2; i <= a; i++ {
		if !notPrime[i] {
			for j := i * 2; j <= a; j += i {
				notPrime[j] = true
			}
		}
	}

	for num, isP := range notPrime {
		if !isP {
			fmt.Print(num, " ")
		}
	}
}

func ReverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i*2 < n; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes)
}

func SumOfArray(arr []int) int { //правда пока принимает не массив, а слайс, пока что не знаю как исправить TODO: подумать, как исправить
	ans := 0
	for _, value := range arr {
		ans += value
	}
	return ans
}

// Rectangle things
type Rectangle struct {
	width  int
	height int
}

func (rect Rectangle) Area() int {
	return rect.height * rect.width
}
