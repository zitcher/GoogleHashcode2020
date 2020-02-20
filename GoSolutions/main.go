package main

import (
	"fmt"
	"github.com/GoogleHashcode2020/GoSolutions/pizza"
    "github.com/GoogleHashcode2020/GoSolutions/util"
	"github.com/GoogleHashcode2020/GoSolutions/qual"
	"log"
    "os"
)

//
func main() {
    args := os.Args[1:]
	content, err := util.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}

	solution := ""

    switch args[0] {
    case "pizza":
		solution, err = pizza.SolvePizza(content)
	    if err != nil {
			log.Fatal(err)
		}
	case "qual":
		solution, err = qual.Solution(content)
	    if err != nil {
			log.Fatal(err)
		}
    default :
        fmt.Printf("arg %s not accepted\n", args[0])
    }

	err = util.WriteFile(args[2], solution)
	if err != nil {
		log.Fatal(err)
	}
    os.Exit(0)
}
