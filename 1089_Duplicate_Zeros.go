package main

/*
Duplicate Zeros: https://leetcode.com/problems/duplicate-zeros/

kiểm tra nếu phẩn tử đó == 0 và sau i+1 phải nhỏ hơn độ dài của mảng
append
*/
func duplicateZeros(arr []int) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 && i+1 < len(arr) {
			arr = append(arr[:i+1], arr[i:len(arr)-1]...)
			i++
		}
	}
}
