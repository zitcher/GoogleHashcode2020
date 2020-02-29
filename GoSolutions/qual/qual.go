package qual

import (
	"fmt"
	"sort"
	"log"
	"time"
	"math/rand"
	"strings"
	"strconv"
	"bytes"
	"encoding/gob"
)

func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func max(a int, b int) int {
	if (a > b) {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func valueOf(i int, j int, arr [][]int) int {
	if (i < 0 || j < 0) {
		return 0
	}
	return arr[i][j]
}

func insertSort(data []int, el int) []int {
	index := sort.Search(len(data), func(i int) bool { return data[i] > el })
	data = append(data, 0)
	copy(data[index+1:], data[index:])
	data[index] = el
	return data
}

/*
Book represents a book id and score
*/
type Book struct {
	ID int
	Score int
}

/*
LibSolution represents a solution
*/
type LibSolution struct {
	Score int
	AlreadyScannedBooks map[int]bool
	Solution []Library
	DaysRemaining int
}

/*
Library represents all lib propertiess
*/
type Library struct {
	ID int
	NumBooks int
	SignupTime int
	ShipPerDay int
	Books []Book
	ScoredBooks []Book
}

func addAllBooks(books []Book, alreadyScannedBooks map[int]bool) {
	for _, b := range books {
		alreadyScannedBooks[b.ID] = true
	}
}

func scoreLib(lib Library, remaining int, alreadyScannedBooks map[int]bool) (score int, scoredBooks []Book){
	scoredBooks = make([]Book, 0)

	for _, book := range lib.Books {
		// if book already scanned skip
		if _, ok := alreadyScannedBooks[book.ID]; ok {
		    continue
		}

		// lib.Books is already sorted greatest to least
		scoredBooks = append(scoredBooks, book)
	}

	score = 0
	maxBooks := min(lib.ShipPerDay * remaining, len(scoredBooks))
	// fmt.Printf("%v %v %v\n", maxBooks, lib.ShipPerDay * remaining, len(scoredBooks))
	scoredBooks = scoredBooks[:maxBooks]
	for _, book := range scoredBooks {
		score += book.Score
	}

	return score, scoredBooks
}

func scoreLibBySignUpTime(lib Library, remaining int, alreadyScannedBooks map[int]bool, libraries []Library) (score int, heurscore int, scoredBooks []Book){
	scoredBooks = make([]Book, 0)

	for _, book := range lib.Books {
		// if book already scanned skip
		if _, ok := alreadyScannedBooks[book.ID]; ok {
		    continue
		}

		// lib.Books is already sorted greatest to least
		scoredBooks = append(scoredBooks, book)
	}

	score = 0
	maxBooks := min(lib.ShipPerDay * remaining, len(scoredBooks))
	// fmt.Printf("%v %v %v\n", maxBooks, lib.ShipPerDay * remaining, len(scoredBooks))
	scoredBooks = scoredBooks[:maxBooks]
	for _, book := range scoredBooks {
		score += book.Score
	}
	heurscore = max(int(float64(score)/float64(lib.SignupTime)), 1)
	return score, heurscore, scoredBooks
}

func solutionToString(solution []Library) string {
	solutionString := strconv.Itoa(len(solution)) + "\n"
	for _, lib := range solution {
		solutionString += strconv.Itoa(lib.ID) + " " + strconv.Itoa(len(lib.ScoredBooks)) + "\n"

		bookIndicies := make([]int, 0)
		for _, book := range lib.ScoredBooks {
			bookIndicies = append(bookIndicies, book.ID)
		}
		booksString := strings.Trim(fmt.Sprintf("%d", bookIndicies),"[]")
		solutionString += booksString + "\n"
	}
	return solutionString
}

// CopySolution performs a deep copy of a solution
func CopySolution(s LibSolution) (LibSolution, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(s)
	if err != nil {
		log.Fatalf("Failed to encode %v\n", err)
	}
	var copy LibSolution
	err = dec.Decode(&copy)
	if err != nil {
		log.Fatalf("Failed to decode %v\n", err)
	}
	return copy, nil
}

// CopyLib performs a deep copy of a library
func CopyLib(orig Library) (Library, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(orig)
	if err != nil {
		log.Fatalf("Failed to encode %v\n", err)
	}
	var copy Library
	err = dec.Decode(&copy)
	if err != nil {
		log.Fatalf("Failed to decode %v\n", err)
	}
	return copy, nil
}

/*
Solution solves the hashcode qualification problem
*/
func Solution(input string) (string, error) {
	inputSplit := strings.Split(input, "\n")

	firstLine := strings.Split(inputSplit[0], " ")
	numBooks, _ := strconv.Atoi(firstLine[0])
	numLibs, _ := strconv.Atoi(firstLine[1])
	maxDays, _ := strconv.Atoi(firstLine[2])

	allBooks := make([]Book, 0)
	for index, score := range strings.Split(inputSplit[1], " "){
		iscore, _ := strconv.Atoi(score)
		allBooks = append(allBooks, Book{index, iscore})
	}

	fmt.Printf("%v %v %v \n", numBooks, numLibs, maxDays)
	// fmt.Printf("%#v\n", bookScores)

	libraries := make([]Library, 0)

	for i := 2; i < len(inputSplit); i+=2 {
		fLine := strings.Split(inputSplit[i], " ")
		if len(fLine) == 0 ||  fLine[0] == "" {
			break
		}

		books := make([]Book, 0)

		lib := Library{}
		lib.ID = len(libraries)
		lib.NumBooks, _ = strconv.Atoi(fLine[0])
		lib.SignupTime, _ = strconv.Atoi(fLine[1])
		lib.ShipPerDay, _ = strconv.Atoi(fLine[2])

		for _, bookid := range strings.Split(inputSplit[i + 1], " ") {
			ibookid, _ := strconv.Atoi(bookid)
			books = append(books, Book{ibookid, allBooks[ibookid].Score})
		}

		sort.SliceStable(books, func(i, j int) bool {
		    return books[i].Score > books[j].Score
		})

		lib.Books = books
		// fmt.Printf("books %#v\n", books)

		if len(lib.Books) != lib.NumBooks {
			log.Fatalf("lib.NumBooks %v, lib.Books %#v\n", lib.NumBooks, lib.Books)
		}

		libraries = append(libraries, lib)
	}

	// fmt.Printf("Libs %#v\n", libraries)

	if len(libraries) != numLibs {
		log.Fatalf("Len libraries %v, numLibs %v, libraries %#v\n", len(libraries), numLibs, libraries)
	}

	if len(libraries) != numLibs {
		log.Fatalf("numLibs %v, libraries %#v\n", numLibs, libraries)
	}

	fmt.Println("Starting Algorithm")
	solution := greedy(libraries, maxDays)
	// solution := buildSolutionExtraMemEfficient(libraries, maxDays)
	// solution := reverseGreedy(libraries, maxDays)

	sol := solutionToString(solution)
	// fmt.Printf("%s\n", sol)

	return sol, nil
}

func buildSolutionExtraMemEfficient(libraries []Library, maxDays int) []Library {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(libraries), func(i, j int) { libraries[i], libraries[j] = libraries[j], libraries[i] })

	greatestGroup := 0
	greatestScore := 0

	// maps score to solution with score
	solutionMap := make(map[int]LibSolution)

	// list of days remaining values
	solutionList := make([]int, 0)

	// intialize variables
	solutionList = append(solutionList, 0)
	solutionMap[0] = LibSolution{
		Score: 0,
		AlreadyScannedBooks: make(map[int]bool),
		Solution: make([]Library, 0),
		DaysRemaining: maxDays,
	}
	maxLen := 100
	radius := 50

	for i := 0; i < len(libraries); i++ {
		// radius = 10 * i
		fmt.Printf("prog: %f %v %v %v\n", float64(i) / float64(len(libraries)), len(solutionMap), len(solutionList), greatestScore)
		if libraries[i].SignupTime > maxDays {
			continue
		}

		temp := make([]int, len(solutionList))
		copy(temp, solutionList)

		for j := 0; j < len(solutionList); j++ {
			lib, _ := CopyLib(libraries[i])
			if solution, ok := solutionMap[solutionList[j]]; ok {
				// skip if greater than max slices
				if solution.DaysRemaining - 1  < lib.SignupTime {
					continue
				}

				newSolution, _ := CopySolution(solution)
				libScore, scoredBooks := scoreLib(lib, newSolution.DaysRemaining - lib.SignupTime, newSolution.AlreadyScannedBooks)
				lib.ScoredBooks = scoredBooks
				newSolution.Score += libScore
				addAllBooks(scoredBooks, newSolution.AlreadyScannedBooks)
				newSolution.DaysRemaining -= lib.SignupTime
				newSolution.Solution = append(newSolution.Solution, lib)


				// if already exists
				if existingSolution, ok := solutionMap[newSolution.DaysRemaining]; ok {
					// if existing is worse
					if existingSolution.Score < newSolution.Score {
						solutionMap[newSolution.DaysRemaining] = newSolution
						if newSolution.Score > greatestScore {
							greatestScore = newSolution.Score
							greatestGroup = newSolution.DaysRemaining
						}
					}
					continue
				}

				// skip or replace if items within radius
				if i, ok := within(temp, newSolution.DaysRemaining, radius); ok {
					existstingScore := temp[i]
					if existingSolution, ok := solutionMap[existstingScore]; ok {
						// replace if better
						if existingSolution.Score < newSolution.Score {
							// remove existing
							delete(solutionMap, existstingScore)
							temp = append(temp[:i], temp[i+1:]...)

							// add new
							solutionMap[newSolution.DaysRemaining] = newSolution
							temp = insertSort(temp, newSolution.DaysRemaining)
							if newSolution.Score > greatestScore {
								greatestScore = newSolution.Score
								greatestGroup = newSolution.DaysRemaining
							}
						}
						continue
					} else {
						log.Fatal("Score not found")
					}
				} else {
					// add new
					solutionMap[newSolution.DaysRemaining] = newSolution
					temp = insertSort(temp, newSolution.DaysRemaining)
					if newSolution.Score > greatestScore {
						greatestScore = newSolution.Score
						greatestGroup = newSolution.DaysRemaining
					}
				}
			}
		}
		solutionList = temp

		if len(solutionList) > maxLen {
			for k := 0; k < len(solutionList) - maxLen; k++ {
				delete(solutionMap, solutionList[k])
			}
			solutionList = solutionList[len(solutionList) - maxLen:]
		}

	}

	if sol, ok := solutionMap[greatestGroup]; ok {
		fmt.Printf("sol: %v\n", greatestScore)
		return sol.Solution
	}
	log.Fatalf("No solution %v", greatestGroup)
	return []Library{}
}

func within(data []int, el int, radius int) (int, bool) {
	index := sort.SearchInts(data, el)
	if index == 0 {
		if data[index] - el < radius {
			return index, true
		}
		return -1, false
	}
	if index == len(data) {
		if el - data[index - 1] < radius {
			return index - 1, true
		}
		return  -1, false
	}

	if data[index] - el < radius {
		return index, true
	}

	if el - data[index - 1] < radius {
		return index - 1, true
	}

	return  -1, false
}

func greedy(libraries []Library, maxDays int) []Library {
	// init problem
	rand.Seed(time.Now().UnixNano())
	alreadyScannedBooks := make(map[int]bool)
	remaining := maxDays

	solution := make([]Library, 0)
	score := 0

	maxLen := float32(len(libraries))
	for {
		fmt.Printf("prog %v\n", 1 - float32(len(libraries))/ maxLen)
		bestIndex, bestScore, bestScoredBooks := findBest(libraries, remaining, alreadyScannedBooks)
		if bestScore < 1 {
			break
		}

		// get library
		lib := libraries[bestIndex]
		lib.ScoredBooks = bestScoredBooks

		// adjust values
		remaining -= lib.SignupTime
		addAllBooks(bestScoredBooks, alreadyScannedBooks)
		libraries = append(libraries[:bestIndex], libraries[bestIndex+1:]...)

		// add to solution
		solution = append(solution, lib)
		score += bestScore

		if remaining < 1 {
			break
		}

	}

	fmt.Printf("score %v\n", score)
	return solution

}

func scoreRemoval(library Library, booksToCount map[int]int) (int, int) {
	score := 0
	heurscore := float64(0)

	for _, book := range library.Books {
		if booksToCount[book.ID] == 1 {
			score += book.Score
			heurscore += float64(book.Score)/float64(library.SignupTime)
		}
	}

	return score, score - library.SignupTime + int(heurscore) + len(library.Books)
}

func findWorst(libraries []Library, booksToCount map[int]int) (int) {
	worstIndex := 0
	bestHeurscore := 1000000000
	for index, lib := range libraries {
		// score, scoredBooks := scoreLib(lib, remaining - lib.SignupTime, alreadyScannedBooks)
		_, heurscore := scoreRemoval(lib, booksToCount)
		if heurscore < bestHeurscore {
			worstIndex = index
			bestHeurscore = heurscore
		}
	}
	return worstIndex
}

func removeAllBooks(books []Book, booksToCount map[int]int) int {
	score := 0

	for _, book := range books {
		if booksToCount[book.ID] == 0 {
			continue
		}
		if booksToCount[book.ID] == 1 {
			score += book.Score
			booksToCount[book.ID] = 0
		}
		booksToCount[book.ID]--
	}

	return score
}

func reverseGreedy(libraries []Library, maxDays int) []Library {
	// init problem
	booksToCount := make(map[int]int)

	score := 0
	for _, lib := range libraries {
		for _, book := range lib.Books {
			if _, ok := booksToCount[book.ID]; ok {
				booksToCount[book.ID]++
			} else {
				booksToCount[book.ID] = 1
				score += book.Score
			}
		}
	}

	days := 0
	for _, lib := range libraries {
		days += lib.SignupTime
	}

	fmt.Printf("len start %v\n", len(libraries))
	fmt.Printf("score %v\n", score)

	maxLibs := len(libraries)

	for {
		if days < maxDays || float64(len(libraries))/float64(maxLibs) < float64(.8) {
			break
		}

		worstIndex := findWorst(libraries, booksToCount)

		// get library
		lib := libraries[worstIndex]
		days -= lib.SignupTime

		// adjust values
		lostScore := removeAllBooks(lib.Books, booksToCount)
		fmt.Printf("score %v\n", lostScore)
		score -= lostScore
		libraries = append(libraries[:worstIndex], libraries[worstIndex+1:]...)
	}

	fmt.Printf("score %v\n", score)
	fmt.Printf("len lib %v\n", len(libraries))
	solution := greedy(libraries, maxDays)
	fmt.Printf("len sol %v\n", len(solution))
	return solution
}


func copyLibs(libraries []Library) []Library{
	libscopy := make([]Library, 0)
	for _, lib := range libraries {
		libscopy = append(libscopy, lib)
	}
	return libscopy
}


func findBest(libraries []Library, remaining int, alreadyScannedBooks map[int]bool) (int, int, []Book) {
	bestIndex := 0
	bestScore := 0
	bestHeurscore := 0
	bestScoredBooks := make([]Book, 0)
	for index, lib := range libraries {
		if remaining - lib.SignupTime <= 0 {
			continue
		}
		// score, scoredBooks := scoreLib(lib, remaining - lib.SignupTime, alreadyScannedBooks)
		score, heurscore, scoredBooks := scoreLibBySignUpTime(lib, remaining - lib.SignupTime, alreadyScannedBooks, libraries)
		if heurscore > bestHeurscore {
			bestIndex = index
			bestScore = score
			bestScoredBooks = scoredBooks
			bestHeurscore = heurscore
		}
	}
	return bestIndex, bestScore, bestScoredBooks
}

/*
	This reduces to set cover

	Library score:
		If read, now how many unique books could we get out of it
*/
