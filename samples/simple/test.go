package main

import (
	"github.com/kurtinlane/submarine/server"
	"github.com/kurtinlane/submarine/apps"
	"github.com/kurtinlane/submarine/keys"
	"github.com/kurtinlane/submarine/samples/simple/storage"
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"
)

func main() {
	fmt.Println("Starting submarine webservice")
	go server.StartServer()

	fmt.Println("Creating new app")
	app := createApp()
	fmt.Println("Created app:", app.Id)
	
	fmt.Println("Signing up new customer")
	key := createKey(app)
	fmt.Println("Signed up new customer", key.Email)
	
	fmt.Println("Encrypt and store the user's age")
	decodedKey, _ := base64.StdEncoding.DecodeString(key.DO_NOT_STORE_DO_NOT_LOG)
	encryptedAge, err := storage.Encrypt(decodedKey, []byte("4"))
	
	if(err != nil) {
		fmt.Println("ERROR encrypting:", err.Error())
	}
	
	storage.StoreData(key.Email+"encrypted_age", string(encryptedAge[:]))
	
	fmt.Println("Age is now hidden in the data", storage.RetrieveData(key.Email+"encrypted_age"))
	
	fmt.Println("Later, user logs back in and wants to see their age")
	fmt.Println("We ask submarine for the user's key and decrypt the data")
	storedEncryptedAge := storage.RetrieveData(key.Email+"encrypted_age")
	
	decryptedAge, _ := storage.Decrypt(decodedKey, []byte(storedEncryptedAge))
	fmt.Println("Now we have the decrypted age", string(decryptedAge[:]))
}   

func createApp() apps.App{
	var app apps.App
	appJson := doPost("http://localhost:3001/apps", `{"Name":"simple"}`, nil)
	parseJson(&app, appJson)
	
	return app
}

func createKey(app apps.App) keys.Key{
	apiKeyHeader := &Header{"API-KEY", app.SECRET_API_KEY}
	
	var key keys.Key
	keyJson := doPost("http://localhost:3000/keys", `{ "Email":"troy@bettrnet.com", "App": 0 }`, apiKeyHeader) //Passing in app.Id doesn't work... not sure why
	parseJson(&key, keyJson)
	
	return key
}

func parseJson(o interface{}, jsonStr []byte){
	err := json.Unmarshal(jsonStr, o)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	}
}

type Header struct{
	Name string
	Value string
}

func doPost(url, jsonStr string, header *Header) []byte {
	var json = []byte(jsonStr)
	req, _ := http.NewRequest("POST", url,  bytes.NewBuffer(json))
	client := &http.Client{}
	
	if header != nil {
		req.Header.Set(header.Name, header.Value)
	}
	
	resp, err := client.Do(req)
	defer resp.Body.Close()
	
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	}
	
	//fmt.Println("response Status:", resp.Status)
    //fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
	
	return body
	
}