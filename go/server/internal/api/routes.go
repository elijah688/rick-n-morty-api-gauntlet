package api

func (s *Server) Routes() {

	s.router.Get("/character", s.getCharacters)

	s.router.Get("/character/{id}", s.getCharacterByID)

	s.router.Post("/character", s.upsertCharacter)

	s.router.Delete("/character/{id}", s.deleteCharacter)

	s.router.Post("/location/list", s.getLocationByIDs)

	s.router.Post("/character/list/episodes", s.getCharacterEpisodesByIDs)

	s.router.Post("/character/list/debut", s.getDebuts)

}
