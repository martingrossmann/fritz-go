package writer

import (
	"encoding/csv"
	"github.com/martingrossmann/fritz-go/fritz"
	"log"
	"os"
)

const csvFileName = "counterdata.csv"

var csvHeader = []string{"date", "received", "sent"}

func WriteData(counter fritz.OnlineCounter) {
	checkErr(prepareFile(), "Error on file")
	checkErr(addData(counter), "Error on file")
}

// Add OnlineCounter data to existing CSV file
func addData(counter fritz.OnlineCounter) error {
	dateString := counter.DayOfData.Format("2006-01-02")
	record := []string{dateString, counter.DataReceived, counter.DataSent}

	csvFile, err := os.OpenFile(csvFileName, os.O_APPEND|os.O_RDWR, 0644)
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	checkErr(err, "Cannot read the CSV file.")
	log.Println((len(records) - 1), "data in CSV file.")

	// Check for the data if already exist.
	if records[len(records)-1][0] != record[0] {
		// Write new data to file
		w := csv.NewWriter(csvFile)
		if err := w.Write(record); err != nil {
			log.Fatalln("Error writing record to csv:", err)
		}
		w.Flush()
		log.Println("New data stored in file.")
	} else {
		log.Println("Data already stored in CSV.", record)
	}

	return csvFile.Close()
}

// Open the destination CSV file
// If not exist --> create new
// If already exist --> check if empty
// If empty --> add header
func prepareFile() error {
	csvFile, err := os.OpenFile(csvFileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	checkErr(err, "Cannot create csv file")

	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	checkErr(err, "Cannot read from file")

	if len(records) < 1 {
		w := csv.NewWriter(csvFile)
		if err := w.Write(csvHeader); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
		w.Flush()
	}

	return csvFile.Close()
}

func checkErr(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}
