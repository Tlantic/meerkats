package benchmark

import (
	"fmt"
	"testing"

	"github.com/uber-go/zap"
)



func zap_fakeFields() []zap.Field {
	return []zap.Field{
		zap.Int("int", 1),
		zap.Int64("int64", 2),
		zap.Float64("float", 3.0),
		zap.String("string", "four!"),
		zap.Bool("bool", true),
		zap.String("another string", "done!"),
	}
}

func zap_fakeMessages(n int) []string {
	messages := make([]string, n)
	for i := range messages {
		messages[i] = fmt.Sprintf("Test logging, but use a somewhat realistic message length. (#%v)", i)
	}
	return messages
}

func BenchmarkZapDisabledLevelsWithoutFields(b *testing.B) {
	logger := zap.New(zap.NewTextEncoder(), zap.ErrorLevel,zap.DiscardOutput)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("Should be discarded.")
		}
	})
}

func BenchmarkZapDisabledLevelsAccumulatedContext(b *testing.B) {
	context := zap_fakeFields()
	logger := zap.New(
		zap.NewTextEncoder(),
		zap.ErrorLevel,
		zap.DiscardOutput,
		zap.Fields(context...),
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("Should be discarded.")
		}
	})
}

func BenchmarkZapDisabledLevelsAddingFields(b *testing.B) {
	logger := zap.New(
		zap.NewTextEncoder(),
		zap.ErrorLevel,
		zap.DiscardOutput,
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("Should be discarded.", zap_fakeFields()...)
		}
	})
}

func BenchmarkZapAddingFields(b *testing.B) {
	logger := zap.New(
		zap.NewTextEncoder(),
		zap.DebugLevel,
		zap.DiscardOutput,
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("Go fast.", zap_fakeFields()...)
		}
	})
}

func BenchmarkZapWithAccumulatedContext(b *testing.B) {
	context := zap_fakeFields()
	logger := zap.New(
		zap.NewTextEncoder(),
		zap.DebugLevel,
		zap.Fields(context...),
		zap.DiscardOutput,
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("Go really fast.")
		}
	})
}

func BenchmarkZapWithoutFields(b *testing.B) {
	logger := zap.New(
		zap.NewTextEncoder(),
		zap.DebugLevel,
		zap.DiscardOutput,
	)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("Go fast.")
		}
	})
}