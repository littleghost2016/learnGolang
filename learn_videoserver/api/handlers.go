package main

import (
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"learn_videoserver/api/defs"
	"learn_videoserver/api/dbops"
	"learn_videoserver/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// 演示部分
	// io.WriteString(w, "Create User Handler")
	res, _ := ioutil.ReadAll(r.Body)

	ubody := &defs.UserCredential{}
	// ubody := defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
