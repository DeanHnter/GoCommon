package logger

import "os"

type FileLogger struct {
    IOLogger
    Filepath string
}

func (fl *FileLogger) SetOutput(path string) {
    fl.Filepath = path
    file, err := os.Create(path)
    if err != nil {
        panic(err)
    }
    fl.Out = file
}
