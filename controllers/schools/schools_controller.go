package schools

import (
	"net/http"

	"github.com/aushafy/go-ulasansekolah-api/configs"
	"github.com/aushafy/go-ulasansekolah-api/domain/schools"
	"github.com/gin-gonic/gin"
)

func CreateSchool(c *gin.Context) {
	var input schools.School

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	configs.ConnectionDB()

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
	configs.DB.Create(&data)

	c.JSON(http.StatusOK, gin.H{"message": "inserrt data success"})
}

func GetSchool(c *gin.Context) {
	var data schools.School

	param := c.Param("npsn")

	configs.ConnectionDB()

	configs.DB.Raw("SELECT * FROM schools WHERE npsn = ?", param).Scan(&data)

	if data.Npsn == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not Found!"})

	} else {
		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}

func SearchSchool(c *gin.Context) {
	var data []schools.School

	configs.ConnectionDB()

	configs.DB.Find(&data)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})
}

func DeleteSchool(c *gin.Context) {
	param := c.Param("npsn")

	configs.ConnectionDB()

	query := "DELETE FROM schools WHERE npsn='" + param + "'"

	configs.DB.Exec(query)

	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
