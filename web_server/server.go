package main

import "net/http"

type Server struct {
	port   string //puerto en el que escuchara las conexiones
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path] //el map retorna true o false si existe el path en el router. El _ corresponde al valor del map
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc) //si no existe inicializo el map de map
	}
	s.router.rules[path][method] = handler //"handler" seria lo que envie desde main, HandleRoot o HandleHome
}

func (s *Server) AddMiddleware(f http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	for _, m := range middleware {
		f = m(f)
	}
	return f
}

func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
