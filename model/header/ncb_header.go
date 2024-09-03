package header

import (
	"github.com/KyawBo/common-library/json"
	"github.com/KyawBo/common-library/validator"

	"github.com/google/uuid"
)

type NCBHeader struct {
	RequestID     string `header:"X-Request-ID" json:"X-Request-ID" binding:"required" validate:"required"`
	CorrelationID string `header:"X-Correlation-ID" json:"X-Correlation-ID" binding:"required" validate:"required"`
	XAuthToken    string `header:"X-Auth-Token" json:"X-Auth-Token"`
	XDevopsSrc    string `header:"x-devops-src" json:"x-devops-src,omitempty"`
	XDevopsDest   string `header:"x-devops-dest" json:"x-devops-dest,omitempty"`
	XDevopsKey    string `header:"x-devops-key" json:"x-devops-key,omitempty"`
}

func BuildNewNCBHeader(header NCBHeader) (NCBHeader, error) {
	newHeader := NCBHeader{
		RequestID:     uuid.New().String(),
		CorrelationID: uuid.New().String(),
		XAuthToken:    header.XAuthToken,
		XDevopsSrc:    header.XDevopsSrc,
		XDevopsDest:   header.XDevopsDest,
		XDevopsKey:    header.XDevopsKey,
	}

	validate := validator.GetValidator()

	err := validate.Struct(newHeader)
	if err != nil {
		return NCBHeader{}, err
	}

	return newHeader, nil
}

func (h *NCBHeader) ToMap() map[string]string {
	tempHeader, _ := json.Json.Marshal(h)

	var mapHeader map[string]string
	_ = json.Json.Unmarshal(tempHeader, &mapHeader)

	return mapHeader
}
