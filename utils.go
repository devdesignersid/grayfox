package grayfox

import "os"

// requries single or no argument  
// if no argument is provided set's host to localhost and port to 8080
func resolveAddress (addr []string) string {
  switch len(addr) {
  case 0:
    if port := os.Getenv("PORT"); port != "" {
      return ":" + port
    }
    return ":8080"
  case 1:
    return addr[0]
  default:
    panic("too many arguments")
  }
}
