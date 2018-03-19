package handler

import (
	"net/http"
	"encoding/json"
	"github.com/ONSdigital/go-ns/log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"errors"
)

var (
	authHeaderKey     = "Authorization"
	userAuthHeaderKey = "X-Florence-Token"

	internalSeverError = Response{Message: "internal server error", Status: 500}
	badRequestError    = Response{Message: "bad request", Status: 400}
	unauthorizedError  = Response{Message: "not authenticated", Status: 401}
)

type APIStub struct {
	scenarios Scenarios
}

func NewAPIStub() (*APIStub, error) {
	source, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.ErrorC("error attempting to load config.yml", err, nil)
		return nil, err
	}

	var scenarios Scenarios
	if err := yaml.Unmarshal(source, &scenarios); err != nil {
		log.ErrorC("error attempting to marshal config.yml", err, nil)
		return nil, err
	}

	b, err := json.MarshalIndent(scenarios, "", " ")
	if err != nil {
		return nil, errors.New("failed to marshal config to json")
	}

	log.Debug("stub Response configuration", log.Data{"config": string(b)})

	return &APIStub{scenarios: scenarios}, nil
}

func (api *APIStub) Identify(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get(authHeaderKey)
	xFlorenceToken := r.Header.Get(userAuthHeaderKey)

	for _, identity := range api.scenarios.Identities {
		if identity.AuthorizationToken == authToken && identity.XFlorenceToken == xFlorenceToken {
			log.Info("identity profile match", log.Data{
				"scenario": identity.Scenario,
				"status":   identity.Status,
			})
			writeResponse(
				Response{
					Identifier: identity.Identifier,
					Message:    identity.Message,
					Status:     identity.Status,
				}, w)
			return
		}
	}

	log.Info("no matching identity profile found, returning default not authenticated response", nil)
	writeResponse(unauthorizedError, w)
}

func (api *APIStub) CreateServiceAccount(w http.ResponseWriter, r *http.Request) {
	log.Info("create service account", nil)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.ErrorC("create service account error: failed to read request body", err, nil)
		writeResponse(internalSeverError, w)
		return
	}
	defer r.Body.Close()

	var service serviceRequest
	if err := json.Unmarshal(body, &service); err != nil {
		log.ErrorC("create service account error: error unmarshalling request body", err, nil)
		writeResponse(internalSeverError, w)
		return
	}

	if service.Name == "" {
		log.Debug("create service account error: request invalid, service name empty", log.Data{"responseStatus": 400})
		writeResponse(badRequestError, w)
		return
	}

	for _, name := range api.scenarios.Services {
		if name == service.Name {
			log.Info("creating new service account", log.Data{
				"name": service.Name,
			})
			writeResponse(Response{Name: name, Token: api.scenarios.NewServiceToken, Status: 201}, w)
			return
		}
	}

	log.Info("no service match found returning error response", nil)
	writeResponse(internalSeverError, w)
}

func writeResponse(r Response, w http.ResponseWriter) {
	b, _ := json.Marshal(r)
	w.WriteHeader(r.Status)
	w.Write(b)
}
