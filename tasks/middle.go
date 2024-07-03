package main

import "fmt"

func CelToF(a float32) float32 {
	return (a*9)/5 + 32
}

func FromNToOne(N int) {
	for i := N; i > 0; i-- {
		fmt.Print(i, " ")
	}
}

func StrLen(s string) int {
	ans := 0
	for i := range s {
		ans = i
	}
	return ans + 1
}

func InArray(arr []int, elem int) bool { // возможно тут хотим делать для любого массива, а не только интового, но я пока не умею в дженерики(
	for _, a := range arr {
		if a == elem {
			return true
		}
	}
	return false
}

func MeanValue(arr []int) float32 {
	sum := 0
	n := len(arr)
	for _, value := range arr {
		sum += value
	}
	return float32(sum) / float32(n)
}

func MultiTable(a int) {
	for i := 1; i < 10; i++ {
		fmt.Println(a, " * ", i, " = ", a*i)
	}
}

func IsPalindrome(s string) bool {
	n := len(s)
	for i := 0; i*2 < n; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

func MinMax(arr []int) (int, int) {
	mn := 1 << 31
	mx := -(1 << 31)

	for _, value := range arr {
		if value < mn {
			mn = value
		}
		if value > mx {
			mx = value
		}
	}
	return mn, mx
}

func DelFromSlice(arr []int, ind int) {
	arr = append(arr[:ind], arr[ind+1:]...)
}

func Index(arr []int, a int) int {
	for i, value := range arr {
		if a == value {
			return i
		}
	}
	return -1
}
