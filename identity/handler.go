package identity

import (
	"net/http"
	"encoding/json"
	"github.com/ONSdigital/go-ns/log"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

var (
	authHeaderKey     = "Authorization"
	userAuthHeaderKey = "X-Florence-Token"
)

type Stub struct {
	profiles identityProfiles
}

type response struct {
	Identifier string `json:"identifier,omitempty"`
	Message    string `json:"message,omitempty"`
}

type identityProfile struct {
	AuthorizationToken string `yaml:"authorization-token,omitempty"`
	Identifier         string `yaml:"identifier,omitempty"`
	Status             int    `yaml:"response-status,omitempty"`
	Message            string `yaml:"message,omitempty"`
	Scenario           string `yaml:"scenario,omitempty"`
}

type identityProfiles struct {
	List []identityProfile `yaml:"identities"`
}

func NewStub() (*Stub, error) {
	source, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.ErrorC("error attempting to load config.yml", err, nil)
		return nil, err
	}

	var profiles identityProfiles
	if err := yaml.Unmarshal(source, &profiles); err != nil {
		log.ErrorC("error attempting to marshal config.yml", err, nil)
		return nil, err
	}

	b, _ := json.MarshalIndent(profiles, "", " ")
	log.Debug("loaded identities config", log.Data{"config": string(b)})

	return &Stub{profiles: profiles}, nil
}

func (s *Stub) Handle(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get(authHeaderKey)

	for _, identity := range s.profiles.List {
		if identity.AuthorizationToken == authToken {
			log.Info(identity.Scenario, log.Data{"status": identity.Status})

			b, _ := json.Marshal(response{Identifier: identity.Identifier, Message: identity.Message})
			w.WriteHeader(identity.Status)
			w.Write(b)
			return
		}
	}

	log.Info("no matching identity profile found, returning default not authenticated response", nil)
	b, _ := json.Marshal(response{Message: "not authenticated"})
	w.WriteHeader(401)
	w.Write(b)
}
