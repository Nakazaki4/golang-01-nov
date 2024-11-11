package main

//import "fmt"

func CanJump(arr []uint) bool {
	length := len(arr) - 1
	var position uint = 0

	for i := 0; i < length; i++ {
		position =+ arr[position]+1

		if position > uint(length) {
			return false
		}
	}

	return true
}

/*func main() {
	input := []uint{2, 3, 1, 1, 4}
	input1 := []uint{3, 2, 1, 0, 4}
	input2 := []uint{0}
	fmt.Println(CanJump(input))
	fmt.Println(CanJump(input1))
	fmt.Println(CanJump(input2))
}*/
