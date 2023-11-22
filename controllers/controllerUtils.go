package controllers

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kaleabbyh/golang-santim-echo/config"
	"gorm.io/gorm"
)


var db *gorm.DB
func init(){
	var err error
	db, err = config.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
}


type User struct {
	ID       		uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name    		string   	   `json:"name"`
	Email    		string   	   `json:"email"`
	Password		string    	   `json:"password"`
	Role     		string         `json:"role"`
	CreatedAt		time.Time	   `json:"created_at"`
	UpdatedAt 		time.Time	   `json:"updated_at"`
}

type UserResponse struct {
	Status 		int    		`json:"status"`
	User   		*User  		`json:"user"`
	Token  		string 		`json:"token"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


type Account struct {
	ID       		uuid.UUID      	`gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt 	`gorm:"index"`
	UserID     		uuid.UUID  		`gorm:"type:uuid;not null"`
	AccountNumber	string  	 	`gorm:"not null"`
	Balance      	float64 		`gorm:"not null"`
	CreatedBy    	uuid.UUID  		`gorm:"type:uuid;not null"`
}

type AccountResponse struct {
	Account  		Account 		`json:"account"`
	User     		User    		`json:"user"`
	Message  		string  		`json:"message"`
}