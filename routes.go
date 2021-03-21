package main

func (s *Server) routes() {
	s.Router.HandleFunc("/", s.handleHome())
	s.Router.HandleFunc("/about", s.handleAbout())
	s.Router.HandleFunc("/outline", s.handleOutline())
	s.Router.HandleFunc("/devices/", s.handleDeviceIndex())
	s.Router.HandleFunc("/devices/{id}", s.handleDeviceID())
	s.Router.HandleFunc("/device", s.handleDeviceHostname()).Methods("POST")
	s.Router.HandleFunc("/scan/devices", s.HandleScanDevices()).Methods("GET")
	s.Router.HandleFunc("/scan/subnet", s.HandleScanSubnet()).Methods("GET")
}
