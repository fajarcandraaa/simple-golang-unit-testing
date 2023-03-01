package routers

import (
	repositories "github.com/fajarcandraaa/simple-golang-unit-testing/model"
	"github.com/fajarcandraaa/simple-golang-unit-testing/route"
	"github.com/fajarcandraaa/simple-golang-unit-testing/src/user"
)

func (se *Serve) initializeRoutes() {
	//======================== REPOSITORIES ========================
	//initiate repository
	r := repositories.NewRepository(se.DB)

	//======================== ROUTER ========================
	//Setting Services
	//Setting User Service

	//=== USER ===
	s := user.NewService(r)
	h := route.NewUserHandler(s)
	//=========================================================

	//======================== ENDPOINT ========================
	//Initialize endpoint route

	//=== USER ===
	se.Router.HandleFunc("/user/create", h.RegisterNewUser).Methods("POST")
	se.Router.HandleFunc("/user/{id}/find", h.FindUserByUserID).Methods("GET")
	se.Router.HandleFunc("/user", h.GetUsers).Methods("GET")
	se.Router.HandleFunc("/user", h.UpdateDataUsers).Methods("PUT")
	se.Router.HandleFunc("/user/{id}", h.UserDelete).Methods("DELETE")
	//==========================================================

}
