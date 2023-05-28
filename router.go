package grayfox

import (
	"net/http"
	"path"
	"strings"
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
  requestPaths := strings.Split(req.URL.Path, "/")
  controllerPath := preparePath(requestPaths[1])
  httpMethod := req.Method
  controller := r.routes[controllerPath];
  handlerPath := path.Join(requestPaths[2:]...)
  handler := controller.getHandler(handlerPath, httpMethod)
  handler(w, req)
}

func (r *Router) Route(path string, controller Controller){
  routePath := preparePath(path)
  r.routes[routePath] = controller
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
         routePath := preparePath(path)
         if c.routeHandlerMap[routePath] == nil {
           c.routeHandlerMap[routePath] = make(map[string]func(http.ResponseWriter, *http.Request))
         }
         c.routeHandlerMap[routePath][method] = handlerFunc
}

func (c *Controller) getHandler(path string, method string) func(http.ResponseWriter, *http.Request) {
  handlerFunc := c.routeHandlerMap[path][method]
  return handlerFunc
}







