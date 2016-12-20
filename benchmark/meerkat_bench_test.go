package benchmark

import (
	"fmt"
	"testing"


	"github.com/Tlantic/meerkats/handlers/writer"
	"github.com/Tlantic/meerkats"
)

func meerkat_fakeFields() []meerkats.KeyValue {
	return []meerkats.KeyValue{
		meerkats.Int("int", 1),
		meerkats.Int64("int64", 2),
		meerkats.Float64("float", 3.0),
		meerkats.String("string", "four!"),
		meerkats.Bool("bool", true),
		meerkats.String("another string", "done!"),
	}
}

func meerkat_fakeMessages(n int) []string {
	messages := make([]string, n)
	for i := range messages {
		messages[i] = fmt.Sprintf("Test logging, but use a somewhat realistic message length. (#%v)", i)
	}
	return messages
}

func BenchmarkMeerkatDisabledLevelsWithoutFields(b *testing.B) {
	logger := meerkats.New(
		meerkats.ERROR,
		writer.Register(writer.DiscardOutput),
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("Should be discarded.")
		}
	})
}

func BenchmarkMeerkatDisabledLevelsAccumulatedContext(b *testing.B) {
	logger := meerkats.New(
		meerkats.ERROR,
		writer.Register(writer.DiscardOutput),
		meerkats.WithFields(meerkat_fakeFields()...))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("Should be discarded.")
		}
	})
}

func BenchmarkMeerkatDisabledLevelsAddingFields(b *testing.B) {
	logger := meerkats.New(
		meerkats.ERROR,
		writer.Register(writer.DiscardOutput),
	)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("Should be discarded.", meerkat_fakeFields()...)
		}
	})
}

func BenchmarkMeerkatAddingFields(b *testing.B) {
	logger := meerkats.New(
		meerkats.DEBUG,
		writer.Register(writer.DiscardOutput),
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("Go fast.", meerkat_fakeFields()...)
		}
	})
}

func BenchmarkMeerkatWithAccumulatedContext(b *testing.B) {
	logger := meerkats.New(
		meerkats.INFO,
		writer.Register(writer.DiscardOutput),
		meerkats.WithFields(meerkat_fakeFields()...))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("Go really fast.")
		}
	})
}

func BenchmarkMeerkatWithoutFields(b *testing.B) {
	logger := meerkats.New(
		meerkats.TRACE,
		writer.Register(writer.DiscardOutput),
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("Go fast.")
		}
	})
}