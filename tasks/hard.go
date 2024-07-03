package main

import "errors"

func DeleteDuplicates(arr []int) []int {
	ans := []int{}
	for i := range arr {
		flag := true
		for j := 0; j < i; j++ {
			if arr[i] == arr[j] {
				flag = false
			}
		}
		if flag {
			ans = append(ans, arr[i])
		}
	}
	return ans
}

func BubbleSort(arr []int32) {
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func Fibonacci(n int) []int {
	ans := make([]int, n)
	ans[0] = 1
	ans[1] = 1
	for i := 2; i < n; i++ {
		ans[i] = ans[i-1] + ans[i-2]
	}
	return ans
}

func CountInArr(arr []int, elem int) int {
	ans := 0
	for _, value := range arr {
		if value == elem {
			ans++
		}
	}
	return ans
}

func IntersectArr(a []int, b []int) []int {
	aWithoutDuplicates := DeleteDuplicates(a)
	bWithoutDuplicates := DeleteDuplicates(b)

	ans := []int{}

	for _, av := range aWithoutDuplicates { //конечно тут надо было посортить и воспользоваться идеей мержсорта,
		// но пока что мы "не умеем" сортить быстрее чем за квадрат(ну раз мы бабл сорт только что писали)
		for _, bv := range bWithoutDuplicates {
			if av == bv {
				ans = append(ans, av)
			}
		}
	}

	return ans
}

func IsAnagram(s1 string, s2 string) bool {
	s1Runes := []int32(s1)
	s2Runes := []int32(s2)
	BubbleSort(s1Runes)
	BubbleSort(s2Runes)

	if len(s1Runes) != len(s2Runes) {
		return false
	}
	for i := range s1Runes {
		if s1Runes[i] != s2Runes[i] {
			return false
		}
	}
	return true
}

func MergeSort(a []int, b []int) []int {
	ans := make([]int, len(a)+len(b))
	i := 0
	j := 0

	for i+j < len(a)+len(b) {
		if i == len(a) {
			ans[i+j] = b[j]
			j++
			continue
		}
		if j == len(b) {
			ans[i+j] = a[i]
			i++
			continue
		}
		if a[i] < b[j] {
			ans[i+j] = a[i]
			i++
		} else {
			ans[i+j] = b[j]
			j++
		}
	}
	return ans
}

func BinarySearch(arr []int, elem int) bool {
	l, r := 0, len(arr)
	for r-l > 1 {
		m := (l + r) / 2
		if arr[m] <= elem {
			l = m
		} else {
			r = m
		}
	}
	return arr[l] == elem
}

// HashTable realisation start
const (
	alphabetSize = 53
)

type HashTableCell struct {
	key   string
	value int
}

type HashTable struct { // я подумал что для простой реализации перестроение не нужно, очень жаль, если вдруг я ошибся!
	buckets    [][]HashTableCell
	size       int
	bucketsCnt int
}

func NewHashTable(bucketsCnt int) *HashTable {
	table := new(HashTable)
	table.size = 0
	table.bucketsCnt = bucketsCnt
	table.buckets = make([][]HashTableCell, bucketsCnt)
	return table
}

func (ht *HashTable) getHash(s string) int {
	ans := 0
	for _, sym := range s {
		ans *= alphabetSize
		ans += int(sym)
		ans %= ht.bucketsCnt
	}
	return ans
}

func (ht *HashTable) Add(key string, value int) {
	hash := ht.getHash(key)
	for i := range ht.buckets[hash] {
		if ht.buckets[hash][i].key == key {
			ht.buckets[hash][i].value = value
			return
		}
	}
	ht.buckets[hash] = append(ht.buckets[hash], HashTableCell{key: key, value: value})
	ht.size++
}

func (ht *HashTable) Delete(key string) {
	hash := ht.getHash(key)
	ind := -1
	for i := range ht.buckets[hash] {
		if ht.buckets[hash][i].key == key {
			ind = i
			break
		}
	}
	if ind != -1 {
		ht.buckets[hash] = append(ht.buckets[hash][:ind], ht.buckets[hash][ind+1:]...)
		ht.size--
	}
}

func (ht *HashTable) GetValue(key string) (int, error) {
	hash := ht.getHash(key)
	for i := range ht.buckets[hash] {
		if ht.buckets[hash][i].key == key {
			return ht.buckets[hash][i].value, nil
		}
	}
	return 0, errors.New("key doesn't exist")
}

func (ht *HashTable) GetSize() int {
	return ht.size
}

//end

// Queue realisation start
type Queue struct {
	beg []int
	end []int
}

func (q *Queue) tryMoveEndToBegin() {
	if len(q.beg) == 0 {
		n := len(q.end)
		for i := 0; i < n; i++ {
			q.beg = append(q.beg, q.end[n-i-1])
		}
		q.end = q.end[0:0]
	}
}

func (q *Queue) Push(x int) {
	q.end = append(q.end, x)
}

func (q *Queue) Pop() { // вообще по хорошему добавить случай, когда очередь пуста, но пока что просто будем не менять очередь
	q.tryMoveEndToBegin()
	if len(q.beg) != 0 {
		q.beg = q.beg[:len(q.beg)-1]
	}
}

func (q *Queue) Front() int {
	q.tryMoveEndToBegin()
	if len(q.beg) == 0 {
		return 0 // Очень плохой способ обрабатывать это, можно было бы чтобы функция возвращала (int, err) но пока что не умею хорошо кастомные ошибки выдавать
	}
	return q.beg[len(q.beg)-1]
}

func (q *Queue) Size() int {
	return len(q.beg) + len(q.end)
}

//end
