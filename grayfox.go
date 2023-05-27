package grayfox

import ("net/http")

// App represents the application instance.
type App struct {
  Router
}

// New creates a new instance of the App.
func New(router Router) *App {
  app := &App{
    Router: router,
  }
  return app
}
  
// Run starts the HTTP server and listens on the specified address.
// It resolves the address and logs a debug message before starting the server.
// If any error occurs during server startup, it is returned.
func (app *App) Run(addr ...string) (err error) {
  address := resolveAddress(addr)
  debugLog("Listening and serving HTTP on %s", address)
  err = http.ListenAndServe(address, &app.Router)
  return 
}

// Note: using middleware stack approach for middleware.
