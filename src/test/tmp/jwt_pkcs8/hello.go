package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"time"
)

func main() {
	privateKeyPath := "tulong.PKCS8"
	publicKeyPath := "tulong.PKCS8.pub"
	privateKeypemData, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		err = fmt.Errorf("read key file: %s", err)
		return
	}
	publicKeypemData, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		err = fmt.Errorf("read key file: %s", err)
		return
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeypemData)
	if err != nil {
		log.Fatal(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, UserClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add((10 * time.Minute)).Unix(),
		},
		IsAdmin:  true,
		Username: "liu",
		Verify:   true,
	})
	jwtToken, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("token:\n %v \n", jwtToken)

	token1 := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add((10 * time.Minute)).Unix(),
		},
		IsAdmin:  true,
		Username: "liu",
		Verify:   true,
	})
	key := []byte("123jkliopnmb123123")
	jwtToken1, err := token1.SignedString(key)
	fmt.Printf("token1:\n %v \n", jwtToken1)

	// RSA	validate
	f1 := func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeypemData)
		return publicKey, err
	}
	t, err := jwt.ParseWithClaims(jwtToken, &UserClaim{}, f1)
	claims, ok := t.Claims.(*UserClaim)
	fmt.Println("RS256:")
	fmt.Println(claims)
	if ok && t.Valid {
		fmt.Println("RSA validated")
	} else {
		fmt.Println(err)
	}

	//	hs256
	t1, err := jwt.ParseWithClaims(jwtToken1, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return key, err
	})
	claims, ok = t1.Claims.(*UserClaim)
	fmt.Println("HS256:")
	fmt.Println(claims)
	if ok && t1.Valid {
		fmt.Println("t1 HS256 validated")
	} else {
		fmt.Println(err)
	}

}

type UserClaim struct {
	jwt.StandardClaims
	IsAdmin  bool   `json:"is_admin"`
	Username string `json:"username"`
	Verify   bool   `json:"verify"`
}

func (c UserClaim) Valid() (err error) {
	if err = c.StandardClaims.Valid(); err != nil {
		return err
	}
	return
}
