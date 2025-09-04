package logger

var logChan = make(chan logEntry, 100) // buffer de logs