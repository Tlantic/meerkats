package benchmark

import (
	"fmt"
	"testing"

	"github.com/Tlantic/meerkats"
	"github.com/Tlantic/meerkats/handlers/writer"
)

func meerkat_fakeFields() []meerkats.Field {
	return []meerkats.Field{
		meerkats.Int("int", 1),
		meerkats.Int64("int64", 2),
		meerkats.Float64("float", 3.0),
		meerkats.String("string", "four!"),
		meerkats.Bool("bool", true),
		meerkats.NewField("object", _jane),
	}
}

func meerkat_fakeMessages(n int) []string {
	messages := make([]string, n)
	for i := range messages {
		messages[i] = fmt.Sprintf("Test logging, but use a somewhat realistic message length. (#%v)", i)
	}
	return messages
}

func BenchmarkMeerkatsNew(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			meerkats.New(meerkats.PANIC, writer.Register())
		}
	})
}

func BenchmarkMeerkatsNewWithPredefinedFields(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			meerkats.New(
				meerkats.PANIC,
				writer.New(),
				meerkats.WithFields(meerkat_fakeFields()...))
		}
	})
}

func BenchmarkMeerkatsDisabledLog(b *testing.B) {
	logger := meerkats.New(meerkats.PANIC, writer.Register())
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("A sample text message.")
		}
	})
}

func BenchmarkMeerkatsDisabledLogWithPredefinedFields(b *testing.B) {
	logger := meerkats.New(meerkats.PANIC, writer.New(), meerkats.WithFields(meerkat_fakeFields()...))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("A sample text message.")
		}
	})
}

func BenchmarkMeerkatsDisabledLogWithFields(b *testing.B) {
	logger := meerkats.New(meerkats.PANIC, writer.New())
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("A sample text message.", meerkat_fakeFields()...)
		}
	})
}

func BenchmarkMeerkatsLog(b *testing.B) {
	logger := meerkats.New(writer.New(writer.DiscardOutput))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("A sample text message.")
		}
	})
}

func BenchmarkMeerkatsLogWithPredefinedFields(b *testing.B) {
	logger := meerkats.New(writer.New(writer.DiscardOutput), meerkats.WithFields(meerkat_fakeFields()...))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("A sample text message.")
		}
	})
}

func BenchmarkMeerkatsLogWithFields(b *testing.B) {
	logger := meerkats.New(writer.New(writer.DiscardOutput))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("A sample text message.", meerkat_fakeFields()...)
		}
	})
}
