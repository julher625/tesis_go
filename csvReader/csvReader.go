package csvReader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func csvReader() {
	// 1. Open the file
	recordFile, err := os.Open("./Data_list_Baseline.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}
	// 2. Initialize the reader
	reader := csv.NewReader(recordFile)
	// 3. Read all the records
	// records, _ := reader.ReadAll()
	// 4. Iterate through the records as you wish
	// fmt.Println(records)

	for i := 0; ; i = i + 1 {
		record, err := reader.Read()
		if err == io.EOF {
			break // reached end of the file
		} else if err != nil {
			fmt.Println("An error encountered ::", err)
			return
		}
		if record[0] == "" {
			continue
		}
		fmt.Printf("Row %d : %v \n", i, record[0])

		// fmt.Printf("Row %d : %v \n", i, record)
	}

}
