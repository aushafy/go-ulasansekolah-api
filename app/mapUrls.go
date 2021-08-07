package app

import (
	"github.com/aushafy/go-ulasansekolah-api/controllers/ping"
	"github.com/aushafy/go-ulasansekolah-api/controllers/schools"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	v1 := router.Group("/api/v1")
	{
		school := v1.Group("/schools")
		{
			school.GET("search/:npsn", schools.GetSchool)
			school.GET("search", schools.SearchSchool)
			school.POST("", schools.CreateSchool)
			school.DELETE(":npsn", schools.DeleteSchool)
		}
	}
}
