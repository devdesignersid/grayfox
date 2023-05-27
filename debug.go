package grayfox

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// debugLog logs a debug message with the provided format and arguments.
// It includes the timestamp, file name, and line number in the log message.
// The format string and arguments are similar to the fmt.Printf function.
func debugLog(format string, v ...interface{}){
_, file, line, _ := runtime.Caller(1)
  fmt.Printf("[%s] [DEBUG] [%s:%d] ", time.Now().Format("2006-01-02 15:04:05"), getFileName(file), line)
  fmt.Printf(format, v...)
  fmt.Println()
}


// getFileName extracts the file name from the full file path.
// It splits the file path using the forward slash ("/") delimiter and returns the last segment, which represents the file name.
func getFileName(file string) string {
  segments := strings.Split(file, "/")
  return segments[len(segments)-1]
}
