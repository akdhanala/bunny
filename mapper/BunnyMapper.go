package mapper

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/akdhanala/bunny/entity"
)

func HttpRequestToResolveDestinationRequest(r *http.Request) entity.ResolveDestinationRequest {
	response := entity.ResolveDestinationRequest{
		Command: "g",
		Query: "",
	}
	
	rawString := r.URL.Query().Get("q")
	if (len(rawString) == 0) {
		return response
	}

	tokens := strings.Split(rawString, " ")

	response.Command = tokens[0]
	if (len(tokens) >= 2) {
		remainingQuery := strings.Join(tokens[1:], " ")
		encodedQuery := url.QueryEscape(remainingQuery)
		response.Query = encodedQuery
	}
	
	return response
}
