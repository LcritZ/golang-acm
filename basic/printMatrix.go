package basic

import "fmt"

func printWidth(matrix [][]int, i int, j int, width int, direction bool) {
	var k int = 0
	if direction {
		for k < width {
			fmt.Println(matrix[i][j])
			k++
			j++
		}
	} else {
		for k < width {
			fmt.Println(matrix[i][j])
			k++
			j--
		}
	}
}

func printHeight(matrix [][]int, i int, j int, height int, direction bool) {
	var k int = 0
	if direction {
		for k < height {
			fmt.Println(matrix[i][j])
			k++
			i++
		}
	} else {
		for k < height {
			fmt.Println(matrix[i][j])
			k++
			i--
		}
	}
}

func PrintMatrix( matrix [][]int ) []int {
	// write code here
	if len(matrix) == 0 {
		return []int{}
	}

	var height int = len(matrix)
	var width int = len(matrix[0])

	fmt.Println("width: ", width)
	fmt.Println("height: ", height)
	var i int = 0
	var j int = 0
	for true {
		fmt.Println("width-true: ", width, "height-true: ", height)
		printWidth(matrix, i, j, width, true)
		height--
		if height == 0 {
			break
		}
		i++
		j = (width+len(matrix[0]))/2 -1
		fmt.Println("width-true: ", width, "height-true: ", height)
		printHeight(matrix, i, j, height, true)
		width--
		if width == 0 {
			break
		}
		i = height
		j--
		fmt.Println("width-false: ", width, "height-true: ", height)
		printWidth(matrix, i, j, width, false)
		height--
		if height == 0 {
			break
		}
		i--
		j = len(matrix[0]) - width -1
		fmt.Println("width-false: ", width, "height-false: ", height)
		fmt.Println("i : ", i, "j: ", j)
		printHeight(matrix, i, j, height, false)
		width--
		if width == 0 {
			break
		}
		fmt.Println("width-false: ", width, "height-false: ", height)
		i = (len(matrix)-height)/2
		j = (len(matrix[0])-width)/2
		fmt.Println("i : ", i, "j: ", j)

		//break
	}
	return []int{}

}
