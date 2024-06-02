package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	defaultFilename         = "dz3.json"
	descriptionFilenameFlag = "file path to db json file"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", defaultFilename, descriptionFilenameFlag)
	flag.Parse()

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var db db
	if err := json.NewDecoder(f).Decode(&db); err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Для выхода введите 0")
	fmt.Println()

	fmt.Print("Номер задания для вывода:  ")
	var number string
	for scanner.Scan() {
		number = strings.TrimSpace(scanner.Text())

		switch number {
		case "0":
			return
		case "3":
			printTableStudentsResults(db)
		case "4":
			printTableMeanScoreByObjectAndGrade(db)
		case "5":
			printTableStudentsResultsGeneric(db)
		case "6":
			printTableMeanScoreByObjectAndGradeFunctional(db)
		case "7":
			printTableAStudentsResults(db)
		default:
			fmt.Println("Номер задания должен быть от 3 до 7)")
		}
		fmt.Print("Номер задания для вывода:  ")
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
