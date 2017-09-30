package main

import (
	"fmt"
	"os"
)

func main() {
	fout, err := os.Create("L:/squares.txt")
	if err != nil {
		fmt.Println("Could not open the file. Exiting...")
		fmt.Println("Press 'Enter' to continue")
		var input interface{}
        fmt.Scanf("%v",&input)
        os.Exit(1)
	}
	defer fout.Close()

    fout.WriteString("Square Numbers Report\n")
	fout.WriteString("NUMBER   SQUARE\n")

	fmt.Println("Enter a number: ")
	var nbr float64
	fmt.Scanf("%f\n", &nbr)
	for nbr != 0 {

		fmt.Fprintf(fout, "%6.2f   %6.2f\n", nbr, nbr*nbr)

		fmt.Println("Enter a number: ")
		fmt.Scanf("%f\n", &nbr)

	}

}
