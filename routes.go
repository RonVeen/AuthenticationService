package main

func (s *Server) routes() {
	s.router.POST("/login", s.handleLogin())
	s.router.POST("/", s.handleCreate())
	s.router.GET("/token/:token", s.handleToken())
}
