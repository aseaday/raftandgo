package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func swap(values []int32, i int32, j int32) {
	swap := values[i]
	values[i] = values[j]
	values[j] = swap
}

func quicksort_tail(values []int32, start int32, end int32) int32 {
	head := start
	tail := end
	key := values[head]
	istart := head + 1
	for head < tail {
		if values[istart] > key {
			values[istart], values[tail] = values[tail], values[istart]
			tail--
		} else {
			values[istart], values[head] = values[head], values[istart]
			head++
			istart++
		}
	}
	if end-start == 1 {
		return -1
	}
	return head
}

func quicksort3(values []int32, start int32, end int32) {
	stack := make([][2]int32, 100000)
	si := 0
	stack[si][0] = start
	stack[si][1] = end
	calls := 0
	for si >= 0 {
		istart := stack[si][0]
		iend := stack[si][1]
		si--
		calls++
		head := quicksort_tail(values, istart, iend)
		if head < 0 {
			continue
		} else {
			if head > istart {
				si++
				stack[si][0] = istart
				stack[si][1] = head
			}
			if head+1 < iend {
				si++
				stack[si][0] = head + 1
				stack[si][1] = iend
			}
		}
	}
}

func quicksort2(values []int32, start int32, end int32) {
	if end-start == 0 {
		return
	}
	head := start
	tail := end
	key := values[head]
	istart := head + 1
	for head < tail {
		if values[istart] > key {
			values[istart], values[tail] = values[tail], values[istart]
			tail--
		} else {
			values[istart], values[head] = values[head], values[istart]
			head++
			istart++
		}
	}
	quicksort2(values, start, head)
	if head+1 < end {
		quicksort2(values, head+1, end)
	}
}

func quicksort(values []int32, start int32, end int32) {
	size := end - start + 1
	if size <= 1 {
		return
	}
	if size == 2 {
		if values[end-1] < values[start] {
			swap(values, start, end-1)
		}
		return
	}
	iend := end
	istart := start
	key := values[istart]
	istart = istart + 1
	for istart < iend {
		for values[istart] < key {
			istart++
			if istart >= iend {
				break
			}
		}
		for values[iend] > key {
			iend--
			if istart >= iend {
				break
			}
		}
		if istart < iend {
			swap(values, istart, iend)
		}
	}
	var mid int32
	if values[istart] < key {
		mid = istart
	} else {
		mid = istart - 1
	}
	swap(values, start, mid)
	if size == 2 {
		return
	}
	quicksort(values, start, mid)
	quicksort(values, mid+1, end)
	return
}

func main() {
	data, err := ioutil.ReadFile("data")
	if err != nil {
		os.Exit(0)
	}
	data_string := string(data)
	numbers_string := strings.Split(data_string, "\n")[0:20000]
	numbers := make([]int32, len(numbers_string))
	for i, number_string := range numbers_string {
		if number_string != "" {
			number, _ := strconv.Atoi(number_string)
			numbers[i] = int32(number)
		}
	}
	quicksort3(numbers, 0, int32(len(numbers)-1))
	prevNumber := int32(0)
	for _, number := range numbers {
		if number < prevNumber {
			panic("Sort is Error")
		} else {
			prevNumber = number
		}
		number_s := strconv.FormatInt(int64(number), 10)
		fmt.Println(number_s)
	}
}
