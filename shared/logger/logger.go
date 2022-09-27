package logger

import "io"

type Level int8

type Logger struct {
	out io.Writer
}

func New(out io.Writer) *Logger {
	return &Logger{
		out: out,
	}
}

func (l *Logger) Print(message string) (int, error) {
	return l.out.Write([]byte(message))
}
