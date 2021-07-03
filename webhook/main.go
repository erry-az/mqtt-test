package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/auth_on_register", authOnRegister)
	http.HandleFunc("/auth_on_subscribe", authOnSubscribe)
	http.HandleFunc("/auth_on_publish", authOnPublish)

	log.Println("connect on port : 8090")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type authOnRegisterReq struct {
	PeerAddr     string `json:"peer_addr"`
	PeerPort     int    `json:"peer_port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Mountpoint   string `json:"mountpoint"`
	ClientID     string `json:"client_id"`
	CleanSession bool   `json:"clean_session"`
}

type authOnSubscribeReq struct {
	ClientID   string   `json:"client_id"`
	Mountpoint string   `json:"mountpoint"`
	Username   string   `json:"username"`
	Topics     *[]topic `json:"topics"`
}

type authOnPublishReq struct {
	Username   string `json:"username"`
	ClientID   string `json:"client_id"`
	Mountpoint string `json:"mountpoint"`
	Qos        int    `json:"qos"`
	Topic      string `json:"topic"`
	Payload    string `json:"payload"`
	Retain     bool   `json:"retain"`
}

type response struct {
	Result    interface{} `json:"result"`
	Modifiers interface{} `json:"modifiers,omitempty"`
	Topics    *[]topic    `json:"topics,omitempty"`
}

type authOnRegisterModifier struct {
	MaxMessageSize      int `json:"max_message_size"`
	MaxInflightMessages int `json:"max_inflight_messages"`
	RetryInterval       int `json:"retry_interval"`
}

type errorReason struct {
	Error string `json:"error"`
}

type topic struct {
	Topic string `json:"topic"`
	Qos   int    `json:"qos"`
}

type authOnPublishModifier struct {
	Topic   string `json:"topic"`
	Qos     int    `json:"qos"`
	Payload string `json:"payload"`
	Retain  bool   `json:"retain"`
}

const (
	OK = "ok"
)

func authOnRegister(w http.ResponseWriter, r *http.Request) {
	var req authOnRegisterReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errorWithReason(w, "invalid message")
		return
	}

	log.Printf("auth_on_register: %+v", req)
	success(w)
}

func authOnSubscribe(w http.ResponseWriter, r *http.Request) {
	var req authOnSubscribeReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errorWithReason(w, "invalid message")
		return
	}

	log.Printf("auth_on_subscribe: %+v", req)
	if req.Topics != nil {
		log.Printf("auth_on_subscribe: topic %+v", *req.Topics)
	}
	success(w)
}

func authOnPublish(w http.ResponseWriter, r *http.Request) {
	var req authOnPublishReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errorWithReason(w, "invalid message")
		return
	}

	log.Printf("auth_on_publish: %+v", req)

	originalPayload, _ := base64.StdEncoding.DecodeString(req.Payload)
	newPayload := strings.TrimSpace(string(originalPayload))
	log.Print("auth_on_publish: real payload : ", newPayload)

	if newPayload == "test" {
		newPayload = newPayload + " " + req.ClientID
	}

	writeResponse(w, response{
		Result: OK,
		Modifiers: authOnPublishModifier{
			Topic:   req.Topic,
			Qos:     req.Qos,
			Payload: base64.StdEncoding.EncodeToString([]byte(newPayload)),
			Retain:  req.Retain,
		},
	})
}

func success(w http.ResponseWriter) {
	writeResponse(w, response{Result: OK})
}

func errorWithReason(w http.ResponseWriter, reason string) {
	writeResponse(w, response{
		Result: errorReason{Error: reason},
	})
}

func writeResponse(w http.ResponseWriter, res response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Print(err)
	}
}
