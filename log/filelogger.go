package logger

import "os"

type FileLogger struct {
    IOLogger
    Filepath string
}

func (fl *FileLogger) SetOutput(path string) {
    fl.filepath = path
    file, err := os.Create(path)
    if err != nil {
        panic(err)
    }
    fl.out = file
}
