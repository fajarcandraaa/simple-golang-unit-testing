package routers

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Serve struct {
	DB     *gorm.DB
	Router *mux.Router
}
