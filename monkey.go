package main

import "fmt"

func main() {
	array1 := [6]int{1, 2, 3, 4, 5, 6}
	var array2 = [6][6]int{array1, array1, array1, array1, array1, array1}
	array3 := [6][6][6]int{array2, array2, array2, array2, array2, array2}

	for i3, v3 := range array3 {
		for i2, v2 := range array2 {
			for i1, v1 := range array1 {
				fmt.Print(v3[i3][i3], v2[i2], v1, "\t")
				if i1 == 5 {
					fmt.Println()
				}
			}
		}
	}
}
