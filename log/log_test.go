package log

import (
	"io"
	"log/slog"
	"os"
	"reflect"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := New(NewDefaultFileWriter(), NewDefaultOptions())
	logger.Info("hello world", "the answer", 42)

	loggerStd := New(os.Stdout, NewDefaultOptions())
	loggerStd.Info("hello world", "the answer", 43)
}

func TestSLogger(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   false,
		Level:       nil,
		ReplaceAttr: nil,
	}))

	logger.Info("hello world", "the answer", 42)
}

func TestNewDefaultFileWriterWithFileName(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want io.Writer
	}{
		{"1", args{fileName: "test.log"}, NewDefaultFileWriterWithFileName("test.log")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultFileWriterWithFileName(tt.args.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultFileWriterWithFileName() = %v, want %v", got, tt.want)
			} else {
				logger := New(got, NewDefaultOptions())
				logger.Info("hello world", "the answer", 42)
			}
		})
	}
}
