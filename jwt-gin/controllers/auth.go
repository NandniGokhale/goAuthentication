package controllers

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"jwt/models"
	"jwt/utils/token"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	gomail "gopkg.in/mail.v2"
)

var otpnumber string
var token1 string

func EmailSend(c *gin.Context) error {
	number, err := getRandNum()
	otpnumber = number
	if err != nil {
		log.Fatal(err)
	}
	m := gomail.NewMessage()
	m.SetHeader("From", "nandanigeitpl@gmail.com")
	m.SetHeader("To", "nandni.g@goldeneagle.ai")
	m.SetHeader("subject", "Otp-Verification")
	s := fmt.Sprintf("OTP  %s", number)
	m.SetBody("text/plain", s)
	d := gomail.NewDialer("smtp.gmail.com", 587, "nandanigeitpl@gmail.com", "lmkiflsfevrhpdgb")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
		return err
	}
	return nil
	//c.JSON(http.StatusOK, gin.H{"message": "Successfully send Otp"})
	//number, _ := getRandNum()
	//var otp string
	//fmt.Println("Successfully send message")

}
func getRandNum() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		return "", e
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}

func CurrentUser(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	u, err := models.GetUserById(user_id)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type OtpVerfication struct {
	Otp string `json:"otp"`
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username, u.Password)
	token1 = token
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect!"})
		return
	}
	//c.JSON(http.StatusOK, gin.H{"token": token})
	//fmt.Println(token1)
	err1 := EmailSend(c)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Successfully send otp"})
	}
}
func VerifyOtp(c *gin.Context) {
	var input OtpVerfication
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if otpnumber == input.Otp {
		c.JSON(http.StatusOK, gin.H{"token": token1})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Otp"})
	}
	//fmt.Println("the number is", otpnumber)
	//fmt.Println("the input otp", input.Otp)
}

type RegisterInput struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Contact   string `json:"contact"`
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.User{}
	u.Firstname = input.Firstname
	u.Lastname = input.Lastname
	u.Contact = input.Contact
	u.Username = input.Username
	u.Password = input.Password

	err1 := u.BeforeSave()
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	_, err := u.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Register successfully"})
}
