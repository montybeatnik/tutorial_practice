package main

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleHome())
	s.router.HandleFunc("/about", s.handleAbout())
	s.router.HandleFunc("/outline", s.handleOutline())
	s.router.HandleFunc("/device/{id}", s.handleDevice())
}
