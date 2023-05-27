package grayfox

import (
	"net/http"
)

type Router struct {
  routes map[string] Controller
  handler http.ServeMux
}

func NewRouter() *Router {
  sm := http.NewServeMux()
  return &Router{
    routes: make(map[string]Controller),
    handler: *sm,
  }
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  path := req.URL.Path
  method := req.Method
  debugLog("Request path: %s:%s", path, method)
  controller := r.routes[path];
  handler := controller.getHandler(path, method)
  handler(w, req)
}
  



type Controller struct {
  router Router
  routeHandlerMap map[string]map[string]func(http.ResponseWriter, *http.Request)
}

func NewController(router Router) *Controller {
  return &Controller{
    router: router,
    routeHandlerMap: make(map[string]map[string]func(http.ResponseWriter, *http.Request)),
  }
}

func (c *Controller) Route(path string, method string, handlerFunc func(http.ResponseWriter, *http.Request)){
         if c.routeHandlerMap[path] == nil {
           c.routeHandlerMap[path] = make(map[string]func(http.ResponseWriter, *http.Request))
         }
         c.routeHandlerMap[path][method] = handlerFunc
}

func (c *Controller) getHandler(path string, method string) func(http.ResponseWriter, *http.Request) {
  handlerFunc := c.routeHandlerMap[path][method]
  return handlerFunc
}

func (r *Router) Route(path string, controller Controller){
  r.routes[path] = controller
}





