package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const filename = "test.txt"

func input(ctx context.Context) <-chan []byte {
	ch := make(chan []byte, 1)

	go func() {
		sc := bufio.NewScanner(os.Stdin)
		sc.Split(bufio.ScanLines)
		defer close(ch)
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

func writeToFile(ctx context.Context, ch <-chan []byte) (<-chan struct{}, error) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	done := make(chan struct{})
	go func() {
		defer func() {
			if err := f.Close(); err != nil {
				log.Println(err)
			}
			done <- struct{}{}
		}()

		writeToFileFn := func(buf []byte, ok bool) bool {
			if !ok {
				return false
			}
			_, err := f.Write(buf)
			if err != nil {
				log.Println(err)
				return false
			}
			_, err = f.WriteString("\n")
			if err != nil {
				log.Println(err)
				return false
			}
			return true
		}

		for {
			select {
			case <-ctx.Done():
				// context canceled, last chance to write data to file
				select {
				// any data in channel buffer?
				case buf, ok := <-ch:
					if !writeToFileFn(buf, ok) {
						return
					}
				default:
				}
				return
			case buf, ok := <-ch:
				if !writeToFileFn(buf, ok) {
					return
				}
			}
		}
	}()

	return done, nil
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM)
	defer stop()

	ch := input(ctx)
	doneCh, err := writeToFile(ctx, ch)
	if err != nil {
		log.Fatal(err)
	}

	<-doneCh
	fmt.Println()
}
