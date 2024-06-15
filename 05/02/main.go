package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const filename = "test.txt"

func input(ctx context.Context, wg *sync.WaitGroup) <-chan []byte {
	ch := make(chan []byte, 1)

	go func() {
		sc := bufio.NewScanner(os.Stdin)
		sc.Split(bufio.ScanBytes)

		defer close(ch)
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !sc.Scan() {
					return
				}
				ch <- bytes.Clone(sc.Bytes())
			}
		}
	}()

	return ch
}

func writeToFile(ctx context.Context, wg *sync.WaitGroup, ch <-chan []byte) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	go func() {
		defer wg.Done()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				// context canceled, last chance to write data to file
				select {
				// any data in channel buffer?
				case buf, ok := <-ch:
					if !ok {
						break LOOP
					}
					_, err := f.Write(buf)
					if err != nil {
						log.Println(err)
					}
				default:
				}
				break LOOP
			case buf, ok := <-ch:
				if !ok {
					break LOOP
				}
				_, err := f.Write(buf)
				if err != nil {
					log.Println(err)
					break LOOP
				}
			}
		}

		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}()

	return nil
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM)
	defer stop()

	var wg = &sync.WaitGroup{}

	wg.Add(1)
	ch := input(ctx, wg)

	wg.Add(1)
	err := writeToFile(ctx, wg, ch)
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
	fmt.Println()
}
