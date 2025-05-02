package main

// TWO POINTERS
func ReverseArray(nums []int) {
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		right--
		left++
	}
}

// срез это указатель на базовый массив, поэтому в данной ситуации
// происходит изменение nums без использования указателей в сигнатуре функции

// если бы nums := [5]int{1,2,3,4,5}, то тогда массив бы не изменился, так как
// внутри функции создалась бы локальная копия массива
