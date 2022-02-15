package iteration

import "fmt"

const repeatCount = 5

func Repeat(repeatCount int, character string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated

}

func main() {
	fmt.Println(Repeat(3,"T"))
}
