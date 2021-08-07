package schools

import (
	"net/http"

	"github.com/aushafy/go-ulasansekolah-api/domain/schools"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateSchool(c *gin.Context) {
	var input schools.School

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database, err := gorm.Open(sqlite.Open("ulasansekolah.db"))

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&schools.School{})

	data := schools.School{
		Id:              input.Id,
		Sekolah:         input.Sekolah,
		Npsn:            input.Npsn,
		Bentuk:          input.Bentuk,
		Status:          input.Status,
		Province:        input.Province,
		ProvinceCode:    input.ProvinceCode,
		City:            input.City,
		CityCode:        input.CityCode,
		SubDistrictCode: input.SubDistrictCode,
		SubDistrict:     input.SubDistrict,
		Address:         input.Address,
		Latitude:        input.Latitude,
		Longitude:       input.Longitude,
	}
	database.Create(&data)

	c.JSON(http.StatusOK, gin.H{"message": "inserrt data success"})
}

func GetSchool(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

func SearchSchool(c *gin.Context) {
	var data []schools.School

	database, err := gorm.Open(sqlite.Open("ulasansekolah.db"))

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.Find(&data)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}
