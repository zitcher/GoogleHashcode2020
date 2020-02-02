package pizza

import (
	"fmt"
	"strings"
	"log"
	"strconv"
	"os"
)

func solutionToString(path []int) string {
	numPizzas := strconv.Itoa(len(path))
	pathString := strings.Trim(fmt.Sprintf("%d", path),"[]")
	sol :=  numPizzas + "\n" + pathString
	return sol
}

func sumList(arr []int) int{
	sum := 0

    for i := range arr {
        sum += arr[i]
    }

	return sum
}

func valueOf(i int, j int, arr [][]int) int {
	if (i < 0 || j < 0) {
		return 0
	}
	return arr[i][j]
}

func max(a int, b int) int {
	if (a > b) {
		return a
	}
	return b
}

func buildSolutionTable(numPizzas int, maxSlices int, pizzas []int) (solutionTable [][]int) {
	solutionTable = make([][]int, numPizzas)
	for i := range solutionTable {
		fmt.Println("built row")
	    solutionTable[i] = make([]int, maxSlices)
	}

	fmt.Println("built table")
	progress := 0.0
	maxProg := float64(maxSlices * numPizzas)
	for i := 0; i < numPizzas; i++ {
		for j := 0; j < maxSlices; j++ {
			progress++
			fmt.Printf("table prog: %d %d\n", progress, progress / maxProg)
			slice := pizzas[i]

			include := valueOf(i - 1, j - slice, solutionTable) + slice
			if (include > j) {
				include = 0
			}
			notInclude := valueOf(i - 1, j, solutionTable)
			solutionTable[i][j] = max(include, notInclude)
		}
	}
	return solutionTable
}

func tracePath(solutionTable [][]int, startR int, startC int, pizzas []int) ([]int, error) {
	var path []int

	r := startR
	c := startC
	for {
		// base case
		if (r == 0) {
			if (solutionTable[r][c] > 0) {
				pizza := pizzas[r]
				path = append(path, pizza)
			}
			return path, nil
		}

		pizza := pizzas[r]

		// not include
		if (solutionTable[r-1][c] == solutionTable[r][c]) {
			r = r-1
			continue
		}

		//include
		path = append(path, pizza)
		if (solutionTable[r - 1][c-pizza] != solutionTable[r][c] - pizza) {
			err := fmt.Errorf(
				"path does not match pizza cur %d prev %d",
				solutionTable[r - 1][c-pizza],
				solutionTable[r][c] - pizza,
			)
			log.Fatal(err)
			os.Exit(1)
			return path, err
		}

		r = r - 1
		c = c - pizza
	}
	return path, nil
}


func buildSolutionMemEfficient(numPizzas int, maxSlices int, pizzas []int) (path []int, opt int) {
	// row - pizzas
	// column - maxslices
	solutionTable := make([][]int, 2)
	for i := range solutionTable {
	    solutionTable[i] = make([]int, maxSlices)
	}
	fmt.Println("sol table done")

	// row - maxslices
	// column - path
	paths := make([][][]int, 2)
	for i := range paths {
		paths[i] = make([][]int, maxSlices)
		for j := range paths[i] {
			fmt.Printf("%d %d %d\n", i, j, len(paths[i]))
	    	paths[i][j] = []int{}
		}
	}
	fmt.Println("paths table done")

	curRow := 0
	altRow := 1
	progress := 0.0
	maxProg := maxSlices * numPizzas
	for i := 0; i < numPizzas; i++ {
		for j := 0; j < maxSlices; j++ {
			progress++
			fmt.Printf("table prog: %f %f\n", progress, progress / float64(maxProg))

			curRow = i % 2
			altRow = (curRow + 1) % 2

			slice := pizzas[i]

			include := valueOf(altRow, j - slice, solutionTable) + slice
			if (include > j) {
				include = 0
			}
			notInclude := valueOf(altRow, j, solutionTable)

			if (include > notInclude) {
				solutionTable[curRow][j] = include
				cp := make([]int, len(paths[altRow][j - slice]))
				copy(cp, paths[altRow][j - slice])
				paths[curRow][j] = append(cp, i)
			} else {
				solutionTable[curRow][j] = notInclude
				cp := make([]int, len(paths[altRow][j]))
				copy(cp, paths[altRow][j])
				paths[curRow][j] = cp
			}
		}
	}
	// fmt.Printf("%#v\n", solutionTable)
	// fmt.Printf("%#v\n", paths[curRow][maxSlices - 1])
	return paths[curRow][maxSlices - 1], solutionTable[curRow][maxSlices - 1]
}

/*
SolvePizza solves the hashcode pizza problem when given a path to the data
*/
func SolvePizza(input string) (string, error) {
	inputSplit := strings.Split(input, "\n")
	firstLine := strings.Split(inputSplit[0], " ")
	pizzaStrings := strings.Split(inputSplit[1], " ")
	maxSlices, _ := strconv.Atoi(firstLine[0])
	numPizzas, _ := strconv.Atoi(firstLine[1])
	pizzas := make([]int, numPizzas)

	if (len(pizzaStrings) != numPizzas) {
		err := fmt.Errorf("length of pizzas %d doesn't equal numPizzas %d", len(pizzaStrings), numPizzas)
		log.Fatal(err)
		return "", err
	}

	for idx, i := range pizzaStrings {
        j, err := strconv.Atoi(i)
        if err != nil {
			log.Fatal(err)
			return "", err
        }
        pizzas[idx] = j
    }

	fmt.Println("number of pizzas", len(pizzas))


	// solutionTable := buildSolutionTable(numPizzas, maxSlices + 1, pizzas)
	// solutionValue := solutionTable[numPizzas-1][maxSlices]
	//path, _ := tracePath(solutionTable, numPizzas-1, maxSlices, pizzas)

	path, solutionValue := treeSearchSolution(numPizzas, maxSlices + 1, pizzas)
	// sumPath := sumList(path)
	// if (sumPath != solutionValue) {
	// 	err := fmt.Errorf("sum of list %d not equal to solution %d", sumPath, solutionValue)
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%#v\n", solutionTable)
	fmt.Printf("%#v\n", solutionValue)
	// fmt.Printf("%#v\n", pizzas)
	// fmt.Printf("%#v\n", path)
	sol := solutionToString(path)
	// fmt.Printf("%s\n", sol)

	return sol, nil
}
