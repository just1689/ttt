package io

import (
	"github.com/just1689/ttt/domain"
	"github.com/just1689/ttt/util"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HandleList(w http.ResponseWriter, r *http.Request) {
	l := domain.LocalInstance.List()
	util.WriteResponse(w, l)
}

func HandleNew(w http.ResponseWriter, r *http.Request) {
	//Input
	req, err := ReadNewGameRequest(w, r)
	if err != nil {
		return
	}

	//Processing
	player := domain.NewPlayer(req.PlayerID, req.Secret, req.PlayerName)
	domain.LocalInstance.AddPlayer(player)
	game := domain.LocalInstance.New(req.GameName)
	game.AddPlayer(player)

	//Output
	util.WriteResponse(w, GameIDResponse{GameID: game.GameID})
}

func HandleJoin(w http.ResponseWriter, r *http.Request) {
	//Input
	req, err := ReadJoinGameRequest(w, r)
	if err != nil {
		return
	}
	player := domain.NewPlayer(req.PlayerID, req.Secret, req.PlayerName)

	//Processing
	_, gameID := domain.LocalInstance.JoinGame(req.GameID, player)
	domain.LocalInstance.AddPlayer(player)

	//Output
	util.WriteResponse(w, GameIDResponse{GameID: gameID})
}

func HandleMove(w http.ResponseWriter, r *http.Request) {
	//Input
	req, err := ReadMoveRequest(w, r)
	if err != nil {
		return
	}
	player, found := domain.LocalInstance.GetPlayer(req.PlayerID, req.Secret)
	if !found {
		w.WriteHeader(http.StatusBadRequest)
		logrus.Print("player not found")
		return
	}

	//Processing
	moved := domain.LocalInstance.Move(player, req.Location)

	//Output
	util.WriteResponse(w, BoolResponse{Result: moved})

}

func HandlePoll(w http.ResponseWriter, r *http.Request) {
	//Input
	req, err := ReadPollRequest(w, r)
	if err != nil {
		return
	}

	//Processing
	rendered, err := domain.LocalInstance.RenderGame(req.GameID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorln("could not render game")
		return
	}

	//Output
	util.WriteResponse(w, rendered)

}
