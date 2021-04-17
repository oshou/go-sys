package main

import (
	"encoding/csv"
	"os"
)

func main() {
	records := [][]string{
		[]string{"名前", "年齢", "身長", "体重"},
		[]string{"Tanaka", "31", "190cm", "97kg"},
		[]string{"Suzuki", "46", "180cm", "79kg"},
		[]string{"Matsui", "45", "188cm", "95kg"},
	}

	file, err := os.Create("decorator-csv-writeall.csv")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)

	err = writer.WriteAll(records)
	if err != nil {
		panic(err)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		panic(err)
	}
}
