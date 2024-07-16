package servers

import (
	"database/sql"
	"oop_send_file/handlers"
	"oop_send_file/repositories"
	"oop_send_file/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutesUploadBinary(router *gin.Engine, db *sql.DB) {

	r := repositories.NewRepositoryAdapter(db)
	s := services.NewServiceAdapter(r)
	h := handlers.NewHanerhandlerAdapter(s)

	router.PATCH("/api/uploadLogoBinary", h.UploadBinaryChicCRMHandlers)
	router.GET("/api/getLogoBinary/:organizeID", h.GetBinaryChicCRMHandlders)
	router.PATCH("/api/editLogoCompany", h.EditLogoProfileChiCCRMHandlers)
	router.PATCH("/api/editPersonalProfile", h.EditPersonalProfileChicCRMHandlers)
	router.GET("/api/getPersonalProfile/:personalID", h.GetPersonalProfileChicCRMHandlers)
}
