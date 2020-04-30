package util

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, o interface{}) {
	b, err := json.Marshal(o)
	if err != nil {
		logrus.Errorln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	w.Write(b)

}
