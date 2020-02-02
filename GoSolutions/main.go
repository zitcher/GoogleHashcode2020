package main

import (
	"fmt"
	"github.com/GoogleHashcode2020/GoSolutions/pizza"
    "github.com/GoogleHashcode2020/GoSolutions/util"
	"log"
    "os"
)

//
func main() {
    args := os.Args[1:]

    switch args[0] {
    case "pizza":
        err := solvePizza(args[1], args[2])
        if err != nil {
    		log.Fatal(err)
    	}
    default :
        fmt.Printf("arg %s not accepted\n", args[0])
    }

    os.Exit(0)
}

func solvePizza(in string, out string) error {
    content, err := util.ReadFile(in)
    if err != nil {
		log.Fatal(err)
        return err
	}

    solution, err := pizza.SolvePizza(content)
    if err != nil {
		log.Fatal(err)
        return err
	}

    err = util.WriteFile(out, solution)

    return err
}
