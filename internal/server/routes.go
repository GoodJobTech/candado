package server

func (s *Server) routes() {
	s.router.HandleFunc("/acquire/{id}", s.AcquireHandler)
	s.router.HandleFunc("/heartbeat/{id}", s.HeartbeatHandler)
	s.router.HandleFunc("/release/{id}", s.ReleaseHandler)
}
