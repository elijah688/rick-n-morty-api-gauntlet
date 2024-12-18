package gateway

import (
	"context"
	"errors"
	"riki_gateway/internal/model"
)

func (cs *GatewayService) CompileCharacter(ctx context.Context, id int) (*model.Character, error) {
	character, err := cs.getCharacterByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return cs.compileCharacter(ctx, character)
}

func (cs *GatewayService) compileCharacter(ctx context.Context, character *model.Character) (*model.Character, error) {
	if character == nil {
		return nil, nil
	}
	pID := character.ID
	if pID == nil {
		return nil, errors.New("id cannot be nil")
	}

	id := *character.ID

	if err := cs.fetchAndSetLocations(ctx, character); err != nil {
		return nil, err
	}

	es, err := cs.getEpisodes(ctx, id)
	if err != nil {
		return nil, err
	}

	if es != nil {
		character.Episodes = &es
	}

	dbt, err := cs.getDebut(ctx, id)
	if err != nil {
		return nil, err
	}

	character.Debut = dbt

	return character, nil
}
