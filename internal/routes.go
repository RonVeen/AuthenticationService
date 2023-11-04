package internal

func (s *server) Routes() {
	s.Router.POST("/login", s.handleLogin())
	s.Router.POST("/", s.handleCreate())
	s.Router.GET("/token/:token", s.handleToken())
}
