// считаем сумму элементов массива, разбивая его на части и обрабатывая каждую часть в отдельной горутине
package main

import (
	"sync"
)

func sumArr(arr []int, start, end int, wg *sync.WaitGroup, res chan int) {
	defer wg.Done()

	sum := 0
	for i := start; i < end; i++ {
		sum += arr[i]
	}

	res <- sum
}

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	nG := 3 // кол-во горутин

// 	subarraySize := len(arr) / nG // размер подмассива для передача каждой горутине

// 	res := make(chan int, nG) // буфеизированный канал на nG элементов
// 	var wg sync.WaitGroup     //
// 	for i := 0; i < nG; i++ {
// 		start := i * subarraySize
// 		end := start + subarraySize

// 		if i == nG-1 {
// 			end = len(arr)
// 		}
// 		wg.Add(1)
// 		go sumArr(arr, start, end, &wg, res)
// 	}

// 	wg.Wait()
// 	close(res)

// 	total := 0
// 	for sum := range res {
// 		total += sum
// 	}
// 	fmt.Println("Total sum: ", total)
// }
