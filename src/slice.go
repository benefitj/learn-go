package main

func PrintSlice() {

	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	println("0~5: ", arr[0:5])
	println("0~5: ", arr[:5])
	println("5~9: ", arr[5:9])
	println("5~9: ", arr[5:])

}
