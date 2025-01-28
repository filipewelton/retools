package internal

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/render"
)

func ReadRequestBody[T any](w http.ResponseWriter, r *http.Request) (T, bool) {
	var parsedBody T

	body, err := io.ReadAll(r.Body)

	if err != nil {
		render.Status(r, 500)
		render.PlainText(w, r, "Internal server error")

		return parsedBody, false
	}

	err = json.Unmarshal(body, &parsedBody)

	if err != nil {
		render.Status(r, 500)
		render.PlainText(w, r, "Internal server error")

		return parsedBody, false
	}

	return parsedBody, true
}
