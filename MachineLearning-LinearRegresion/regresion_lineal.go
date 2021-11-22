package main

import (
	"fmt"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {
	//csv
	advertFile, err := os.Open("./import0921.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer advertFile.Close()

	//crear data frame
	advertDF := dataframe.ReadCSV(advertFile)
	//calcular regresion lineal min 25 50 75 max
	advertSummary := advertDF.Describe()
	fmt.Println(advertSummary)

	//0: mean
	//1: median
	//2: stddev
	//3: min
	//4: 25%
	//5: 50%
	//6: 75%
	//7: max
}
