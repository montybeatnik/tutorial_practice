package main

func (s *Server) routes() {
	s.Router.HandleFunc("/", s.handleHome())
	s.Router.HandleFunc("/about", s.handleAbout())
	s.Router.HandleFunc("/outline", s.handleOutline())
	s.Router.HandleFunc("/device/{id}", s.handleDeviceID())
	s.Router.HandleFunc("/device", s.handleDeviceHostname()).Methods("POST")
	s.Router.HandleFunc("/scan", s.HandleScan()).Methods("GET")
}
