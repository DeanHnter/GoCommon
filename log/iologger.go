package logger

import (
    "fmt"
    "io"
)

type IOLogger struct {
    Level Level
    Out io.Writer
}

func (bl IOLogger) Log(level Level, message string) {
    if level >= bl.Level {
        fmt.Fprintln(bl.Out, message)
    }
}

func (bl IOLogger) SetOutput(out io.Writer) {
    bl.Out = out
}

func (bl IOLogger) SetLevel(level Level) {
    bl.Level = level
}
