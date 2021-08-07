package app

import (
	"github.com/aushafy/go-ulasansekolah-api/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	v1 := router.Group("/api/v1")
	{
		// _ := v1.Group("/schools")
		// {

		// }
	}
}
