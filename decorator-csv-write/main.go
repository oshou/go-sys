package main

import (
	"encoding/csv"
	"os"
)

func main() {
	record1 := []string{"名前", "年齢", "身長", "体重"}
	record2 := []string{"Tanaka", "31", "190cm", "97kg"}
	record3 := []string{"Suzuki", "46", "180cm", "79kg"}
	record4 := []string{"Matsui", "45", "188cm", "95kg"}

	file, err := os.Create("decorator-csv-write.csv")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)

	err = writer.Write(record1)
	if err != nil {
		panic(err)
	}

	err = writer.Write(record2)
	if err != nil {
		panic(err)
	}

	err = writer.Write(record3)
	if err != nil {
		panic(err)
	}

	err = writer.Write(record4)
	if err != nil {
		panic(err)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		panic(err)
	}
}
