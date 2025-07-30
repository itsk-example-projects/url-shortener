package web

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewEngine(h *Handlers) *gin.Engine {
	r := gin.Default()

	r.Static("/static", "static")

	r.GET("/", gin.WrapH(http.HandlerFunc(h.Index)))
	r.GET("/:slug", gin.WrapH(http.HandlerFunc(h.Index)))
	r.POST("/shorten", gin.WrapH(http.HandlerFunc(h.Shorten)))
	r.POST("/delete/:slug", gin.WrapH(http.HandlerFunc(h.Delete)))

	r.GET("/export-id", gin.WrapH(http.HandlerFunc(h.ExportID)))
	r.POST("/import-id", gin.WrapH(http.HandlerFunc(h.ImportID)))

	r.GET("/admin", gin.WrapH(http.HandlerFunc(h.Admin)))
	r.POST("/admin/delete/:slug", gin.WrapH(http.HandlerFunc(h.AdminDelete)))

	return r
}

func RunServer(r *gin.Engine) {
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start web server: %v", err)
	}
}
