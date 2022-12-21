package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type RegistrationRequest struct {
	SmartphoneID int    `json:"smartphoneID"`
	OSVersion    string `json:"osVersion"`
	SellingPrice int    `json:"sellerPrice"`
}

type Smartphones struct {
	smartphoneID int    `gorm:"column:smartphoneID"`
	osVersion    string `gorm:"column:osVersion"`
	sellingPrice int    `gorm:"column:sellingPrice"`
}

func main() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		db, err := GetGormConn()
		if err != nil {
			err = errors.New("database connection error:" + err.Error())
			PostResponse(w, http.StatusInternalServerError, err.Error())
		}
		var regReq RegistrationRequest

		err = regReq.SetParameter(r)
		if err != nil {
			PostResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		regReq.RegistrateSmartPhone(db)
		PostResponse(w, http.StatusOK, "")
	default:
		err := errors.New("this method is not allowed")
		PostResponse(w, http.StatusMethodNotAllowed, err.Error())
	}
}

func GetGormConn() (*gorm.DB, error) {
	return gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
}

func PostResponse(w http.ResponseWriter, status int, message string) {
	response := Response{Status: status, Message: message}
	json.NewEncoder(w).Encode(&response)
}

func (regReq *RegistrationRequest) SetParameter(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(&regReq); err != nil {
		err = errors.New("incorrect parameters")
		return err
	}
	return nil
}

func (regReq *RegistrationRequest) RegistrateSmartPhone(db *gorm.DB) {
	phone := Smartphones{
		smartphoneID: regReq.SmartphoneID,
		osVersion:    regReq.OSVersion,
		sellingPrice: regReq.SellingPrice,
	}
	db.NewRecord(phone)
	db.Create(&phone)
}
