package examples

import (
	"fmt"
	"net/http"
	"time"

	"github.com/MhmdRyhn/httprw"
)

func main() {
	http.HandleFunc("/timezone", getTimeByTimezone)
}

type errorResponseBody struct {
	Error string `json:"error"`
}

type requestBody struct {
	Timezone string `json:"timezone"`
}

func getTimeByTimezone(w http.ResponseWriter, r *http.Request) {
	parser := httprw.NewRequestParser(r)
	input := requestBody{}
	if r.Method == "POST" {
		_ = parser.Body(&input)
	} else if r.Method == "GET" {
		query := parser.QueryParams()
		if timezones, found := query["timezone"]; found && len(timezones) != 0 {
			input.Timezone = timezones[0]
		}
	}
	location, err := time.LoadLocation(input.Timezone)
	if err != nil {
		httprw.Response(
			w,
			errorResponseBody{Error: fmt.Sprintf("Timezone `%s` is invalid.", input.Timezone)},
			map[string]string{},
			422,
		)
		return
	}
	currentTime := time.Now().In(location)
	httprw.Response(
		w,
		struct {
			Info string `json:"info"`
		}{
			Info: fmt.Sprintf("Current time in timezone `%s` is %s.", input.Timezone, currentTime.Format(time.RFC1123)),
		},
		map[string]string{},
		200,
	)
}
