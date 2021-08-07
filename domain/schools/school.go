package schools

type School struct {
	Id              string `json:"id"`
	Sekolah         string `json:"sekolah"`
	Npsn            string `json:"npsn"`
	Bentuk          string `json:"bentuk"`
	Status          string `json:"status"`
	Province        string `json:"propinsi"`
	ProvinceCode    string `json:"kode_prop"`
	City            string `json:"kabupaten_kota"`
	CityCode        string `json:"kode_kab_kota"`
	SubDistrictCode string `json:"kode_kec"`
	SubDistrict     string `json:"kecamatan"`
	Address         string `json:"alamat_jalan"`
	Latitude        string `json:"lintang"`
	Longitude       string `json:"bujur"`
}
