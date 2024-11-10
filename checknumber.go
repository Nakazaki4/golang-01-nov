package main

/*import "fmt"

func main(){

	fmt.Println(CheckNumber("AHJDG"))

}*/

func CheckNumber(s string) bool {

	for i := 0; i < len(s); i++ {
		for j := '0'; j < '9'; j++ {
			if s[i] == byte(j) {
				return false
			}
		}		
	}

	return true
}