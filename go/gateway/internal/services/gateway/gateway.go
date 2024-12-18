package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"riki_gateway/internal/config"
)

type GatewayService struct {
	cfg    *config.Config
	client *http.Client
}

func New(cfg *config.Config) *GatewayService {

	return &GatewayService{cfg, new(http.Client)}
}

func marshalToJSONBody(in map[string]any) (*bytes.Reader, error) {
	body, err := json.Marshal(in)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal ids: %w", err)
	}

	return bytes.NewReader(body), nil
}
