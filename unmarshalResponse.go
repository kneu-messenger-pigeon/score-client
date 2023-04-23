package score

import (
	"encoding/json"
	"io"
	"net/http"
)

func unmarshalResponse(response *http.Response, v any) error {
	body, err := io.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &v)
	}
	return err
}
