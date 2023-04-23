package score

import (
	"errors"
	"fmt"
	scoreApi "github.com/kneu-messenger-pigeon/score-api"
	"net/http"
	"strconv"
)

const Version = "/v1"

type ClientInterface interface {
	GetStudentDisciplines(studentId uint32) (response scoreApi.DisciplineScoreResults, err error)
	GetStudentDiscipline(studentId uint32, disciplineId int) (response scoreApi.DisciplineScoreResult, err error)
}

type Client struct {
	Host string
}

func (client *Client) GetStudentDisciplines(studentId uint32) (response scoreApi.DisciplineScoreResults, err error) {
	err = client.doRequest(fmt.Sprintf("/students/%d/disciplines", studentId), &response)
	return
}

func (client *Client) GetStudentDiscipline(studentId uint32, disciplineId int) (response scoreApi.DisciplineScoreResult, err error) {
	err = client.doRequest(fmt.Sprintf("/students/%d/disciplines/%d", studentId, disciplineId), &response)
	return
}

func (client *Client) doRequest(requestUri string, responseInterface any) error {
	var response *http.Response

	request, err := http.NewRequest(http.MethodGet, client.Host+Version+requestUri, nil)

	if err == nil {
		response, err = http.DefaultClient.Do(request)

		if err == nil && response.StatusCode != 200 {
			errorResponse := scoreApi.ErrorResponse{}
			err = unmarshalResponse(response, &errorResponse)
			if err == nil && errorResponse.Error != "" {
				err = errors.New("API error: " + errorResponse.Error)
			} else {
				err = errors.New("Receive http code: " + strconv.Itoa(response.StatusCode))
			}
		}
	}

	if err == nil {
		err = unmarshalResponse(response, &responseInterface)
	}

	return err
}
