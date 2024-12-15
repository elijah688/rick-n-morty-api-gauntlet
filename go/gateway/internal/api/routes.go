package api

func (s *Server) Routes() {
	s.router.Get("/character/{id}", s.getCharacterByID)
	s.router.Get("/character", s.getCharacters)

	s.router.Post("/character", s.upsertCharacter)
	s.router.Delete("/character/{id}", s.deleteCharacter)

}
