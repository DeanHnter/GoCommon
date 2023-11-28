package logger

import (
    "fmt"
    "io"
)

type IOLogger struct {
    level Level
    out io.Writer
}

func (bl *IOLogger) Log(level Level, message string) {
    if level >= bl.level {
        fmt.Fprintln(bl.out, message)
    }
}

func (bl *IOLogger) SetOutput(out io.Writer) {
    bl.out = out
}

func (bl *IOLogger) SetLevel(level Level) {
    bl.level = level
}
