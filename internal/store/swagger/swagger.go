package swagger

import (
	"encoding/json"
	"hitss/pkg/helper/logger"
	"hitss/pkg/helper/openapi"
	"os"

	"github.com/getkin/kin-openapi/openapi3"
	"gopkg.in/yaml.v2"
)

type store struct {
	storage *openapi3.T
}

func New() *store {
	return &store{
		storage: openapi.New(),
	}
}

func (s store) GetJSON() ([]byte, error) {
	data, err := json.Marshal(s.storage)
	if err != nil {
		logger.Write(err)
		return nil, err
	}

	if err := os.WriteFile(os.Getenv("ROOT")+"/asset/swagger/openapi3.json", data, 0644); err != nil {
		logger.Write(err)
		return nil, err
	}

	return data, nil
}

func (s store) GetYAML() ([]byte, error) {
	data, err := yaml.Marshal(s.storage)
	if err != nil {
		logger.Write(err)
		return nil, err
	}

	if err := os.WriteFile(os.Getenv("ROOT")+"/asset/swagger/openapi3.yaml", data, 0644); err != nil {
		logger.Write(err)
		return nil, err
	}

	return data, nil
}
