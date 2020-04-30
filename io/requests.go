package io

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type NewGameRequest struct {
	GameName   string `json:"gameName"`
	PlayerID   string `json:"playerID"`
	Secret     string `json:"secret"`
	PlayerName string `json:"playerName"`
}

func ReadNewGameRequest(w http.ResponseWriter, r *http.Request) (*NewGameRequest, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}
	result := &NewGameRequest{}
	err = json.Unmarshal(b, result)
	if err != nil {
		logrus.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}
	return result, nil
}

type JoinGameRequest struct {
	GameID     string `json:"gameID"`
	PlayerID   string `json:"playerID"`
	Secret     string `json:"secret"`
	PlayerName string `json:"playerName"`
}

func ReadJoinGameRequest(w http.ResponseWriter, r *http.Request) (*JoinGameRequest, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}
	result := &JoinGameRequest{}
	err = json.Unmarshal(b, result)
	if err != nil {
		logrus.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}
	return result, nil
}

type MoveRequest struct {
	GameID   string `json:"gameID"`
	PlayerID string `json:"playerID"`
	Secret   string `json:"secret"`
	Location int    `json:"location"`
}

func ReadMoveRequest(w http.ResponseWriter, r *http.Request) (*MoveRequest, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}
	result := &MoveRequest{}
	err = json.Unmarshal(b, result)
	if err != nil {
		logrus.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}
	return result, nil
}

type PollRequest struct {
	GameID string `json:"gameID"`
}

func ReadPollRequest(w http.ResponseWriter, r *http.Request) (*PollRequest, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}
	result := &PollRequest{}
	err = json.Unmarshal(b, result)
	if err != nil {
		logrus.Errorln(err)
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}
	return result, nil
}
