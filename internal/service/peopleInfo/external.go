package peopleinfo

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/YEgorLu/time-tracker/internal/logger"
	"github.com/YEgorLu/time-tracker/internal/service/peopleInfo/models"
)

var _ PeopleInfoService = &externalPeopleInfoService{}

type externalPeopleInfoService struct {
	log logger.Logger
}

func newExternalService(log logger.Logger) *externalPeopleInfoService {
	return &externalPeopleInfoService{log}
}

// GetInfo implements PeopleInfoService.
func (e *externalPeopleInfoService) GetInfo(ctx context.Context, passportSerie, passportNumber string) (info models.Info, err error) {
	defer func() {
		if err != nil {
			e.log.Error(err)
		}
	}()
	queryParams := url.Values{
		"passportSerie":  []string{passportSerie},
		"passportNumber": []string{passportNumber},
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, `some_url/info?`+queryParams.Encode(), nil)
	if err != nil {
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	if err = json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return
	}
	return
}
