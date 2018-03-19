package benchmark

import (
	"fmt"
	"testing"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
)

func zap_fakeFields() []zapcore.Field {
	return []zapcore.Field{
		zap.Int("int", 1),
		zap.Int64("int64", 2),
		zap.Float64("float", 3.0),
		zap.String("string", "four!"),
		zap.Bool("bool", true),
		zap.String("another string", "done!"),
		zap.Any("object", _jane),
	}
}

//noinspection GoUnusedFunction
func zap_fakeMessages(n int) []string {
	messages := make([]string, n)
	for i := range messages {
		messages[i] = fmt.Sprintf("Test logging, but use a somewhat realistic message length. (#%v)", i)
	}
	return messages
}

func BenchmarkZapNew(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			zap.New(zapcore.NewCore(
				zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
				zapcore.AddSync(ioutil.Discard),
				zapcore.PanicLevel,
			))
		}
	})
}

func BenchmarkZapNewWithPredefinedFields(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for pb.Next() {
				zap.New(zapcore.NewCore(
					zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
					zapcore.AddSync(ioutil.Discard),
					zapcore.PanicLevel,
				), zap.Fields(zap_fakeFields()...))
			}
		}
	})
}

func BenchmarkZapDisabledLog(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.PanicLevel,
	))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("A sample text message.")
		}
	})
}

func BenchmarkZapDisabledLogWithPredefinedFields(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.PanicLevel,
	), zap.Fields(zap_fakeFields()...))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("A sample text message.")
		}
	})
}

func BenchmarkZapDisabledLogWithFields(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.PanicLevel,
	))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("A sample text message.", zap_fakeFields()...)
		}
	})
}

func BenchmarkZapLog(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.InfoLevel,
	))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("A sample text message.")
		}
	})
}

func BenchmarkZapLogWithPredefinedFields(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.InfoLevel,
	), zap.Fields(zap_fakeFields()...))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("A sample text message.")
		}
	})
}

func BenchmarkZapLogWithFields(b *testing.B) {
	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(ioutil.Discard),
		zapcore.InfoLevel,
	))
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("A sample text message.", zap_fakeFields()...)
		}
	})
}
