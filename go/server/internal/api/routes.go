package api

func (s *Server) Routes() {

	s.router.Get("/character", s.getCharacters)

	s.router.Get("/character/{id}", s.getCharacterByID)

	s.router.Post("/character", s.upsertCharacter)

	s.router.Delete("/character/{id}", s.deleteCharacter)

	s.router.Get("/location/{id}", s.getLocationByID)

	s.router.Get("/character/{id}/episodes", s.getCharacterEpisodes)

	s.router.Get("/character/{id}/debut", s.getDebut)

}
