package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
)

const (
	defaultFilename         = "problems.csv"
	defaultNeedShuffle      = false
	descriptionFilenameFlag = "file path to CSV file with problems"
	descriptionShuffleFlag  = "shuffle the order of problems"
	problemBatchSize        = 50 // max number of problems read at once from CSV file
)

func main() {
	var (
		filename    string
		needShuffle bool
	)

	flag.StringVar(&filename, "file", defaultFilename, descriptionFilenameFlag)
	flag.StringVar(&filename, "f", defaultFilename, descriptionFilenameFlag+" (shorthand)")
	flag.BoolVar(&needShuffle, "shuffle", defaultNeedShuffle, descriptionShuffleFlag)
	flag.Parse()

	if err := run(filename, needShuffle); err != nil {
		log.Fatal(err)
	}
}

func run(filename string, needShuffle bool) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	var (
		problemsIndex int
		rightAnswers  int
		problems      [][2]string
	)
	for problemsIndex = 0; ; problemsIndex++ {
		record, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}
		if len(record) != 2 {
			return fmt.Errorf("expected 2 fields, not %d", len(record))
		}

		// batch of problems readed
		if problemsIndex%problemBatchSize == 0 {
			ra, err := qa(problems, needShuffle)
			if err != nil {
				return errors.New("bad input: " + err.Error())
			}
			rightAnswers += ra
			problems = make([][2]string, 0, problemBatchSize)
		}
		problems = append(problems, [2]string{strings.TrimSpace(record[0]), strings.TrimSpace(record[1])})
	}

	ra, err := qa(problems, needShuffle)
	if err != nil {
		return errors.New("bad input: " + err.Error())
	}
	rightAnswers += ra

	fmt.Printf("___________\n\nВаш счет\nверно:    %d\nне верно: %d\n", rightAnswers, problemsIndex-rightAnswers)

	return nil
}

func qa(problems [][2]string, needShuffle bool) (int, error) {
	if len(problems) == 0 {
		return 0, nil
	}

	var rightAnswers int

	// not to modify the original slice
	indexes := make([]int, len(problems))
	for i := range problems {
		indexes[i] = i
	}
	if needShuffle {
		rand.Shuffle(len(indexes), func(i, j int) {
			indexes[i], indexes[j] = indexes[j], indexes[i]
		})
	}

	scanner := bufio.NewScanner(os.Stdin)
	for _, i := range indexes {
		fmt.Println("вопрос:", problems[i][0])
		fmt.Print("ответ:  ")

		var answer string
		if !scanner.Scan() {
			return 0, scanner.Err()
		}
		answer = strings.TrimSpace(scanner.Text())

		if strings.EqualFold(answer, problems[i][1]) {
			rightAnswers++
		}
	}

	return rightAnswers, nil
}
