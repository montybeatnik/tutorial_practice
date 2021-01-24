package main

func (s *server) routes() {
	s.router.HandleFunc("/", s.handleHome())
	s.router.HandleFunc("/about", s.handleAbout())
	s.router.HandleFunc("/outline", s.handleOutline())
	s.router.HandleFunc("/device/{id}", s.handleDeviceID())
	s.router.HandleFunc("/device", s.handleDeviceHostname()).Methods("POST")
	s.router.HandleFunc("/scan", s.HandleScan()).Methods("GET")
}
