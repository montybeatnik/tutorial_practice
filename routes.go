package main

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleHome())
	s.router.HandleFunc("/about", s.handleAbout())
}
