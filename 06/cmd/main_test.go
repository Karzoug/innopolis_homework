package main

import (
	"06/internal/logger/best"
	"06/internal/logger/optim"
	"06/internal/logger/simple"
	"testing"

	"github.com/rs/zerolog"
)

func BenchmarkSimpleLoggerPrint(b *testing.B) {
	stream := &blackholeStream{}
	l := simple.NewLogger(stream)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Print("The quick brown fox jumps over the lazy dog", int(32), float64(64))
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkOptimLoggerPrint(b *testing.B) {
	stream := &blackholeStream{}
	l := optim.NewLogger(stream)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Print("The quick brown fox jumps over the lazy dog", int(32), float64(64))
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkBestLoggerPrint(b *testing.B) {
	stream := &blackholeStream{}
	l := best.NewLogger(stream)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Print("The quick brown fox jumps over the lazy dog", int(32), float64(64))
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkZerologPrint(b *testing.B) {
	stream := &blackholeStream{}
	zerolog.TimeFieldFormat = TimeFieldFormat
	l := zerolog.New(stream).With().Timestamp().Logger()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Print("The quick brown fox jumps over the lazy dog", int(32), float64(64))
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
