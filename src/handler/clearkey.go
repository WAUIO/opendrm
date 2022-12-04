package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func strPtr(str string) *string {
	return &str
}

type ClearKeyRequest struct {
	Kids []string `json:"kids"`
	Type string   `json:"type"`
}

type ClearKeySet struct {
	Key string  `json:"k"`
	Kty string  `json:"kty"`
	Kid string  `json:"kid"`
	Alg *string `json:"alg,omitempty"`
}

type ClearKeyResponse struct {
	Keys []ClearKeySet `json:"keys"`
	Type string        `json:"type"`
}

func AcquireLicenseForClearKey(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Read request failed. err=%s", err)
		return
	}

	req := &ClearKeyRequest{}
	if err := json.Unmarshal(reqData, &req); err != nil {
		log.Printf("failed to parse request: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	keyset := &ClearKeyResponse{
		Keys: []ClearKeySet{},
		Type: req.Type,
	}

	for _, k := range req.Kids {
		ck := ClearKeySet{
			// @todo: find a way to couple a kid with a key
			// for now it's hard-coded for test
			Key: "ihawK6q5S0mzeizD0FRRig",
			Kty: "oct",
			Kid: k,
		}
		ck.Alg = strPtr("A128KW")

		keyset.Keys = append(keyset.Keys, ck)
	}

	if response, err := json.Marshal(keyset); err == nil {
		w.Header().Add("Content-Type", "application/json")
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusForbidden)
}
