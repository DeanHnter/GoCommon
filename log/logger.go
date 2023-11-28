package logger

import "io"

type Level int

const (
    INFO = iota
    DEBUG
    WARNING
    ERROR
)

type LogInterface interface {
    Log(level Level, message string)
    SetOutput(io.Writer)
    SetLevel(level Level)
}
