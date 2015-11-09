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
	"time"
	"bufio"
	"os"
	"strings"
)

func main() {
	fmt.Println("\nStarting submarine webservice\n")
	go server.StartServer()
	time.Sleep(2 * time.Second)

	fmt.Println("\nStarting webservice \"StoreMyData\"\n")
	time.Sleep(2 * time.Second)
	app := createApp()
	
	fmt.Print("\nSign up for StoreMyData (enter your email address): \n")
	scanner := bufio.NewScanner(os.Stdin)
	
	scanner.Scan()
	email := scanner.Text()

	key := createKey(app, strings.Replace(email, "\n", "", -1))

	fmt.Print("\nTell StoreMyData something:\n")
	scanner.Scan()
	data := scanner.Text()
	
	encryptedData, _ := storage.Encrypt(decodeKey(key.DO_NOT_STORE_DO_NOT_LOG), []byte(data))
	
	storage.StoreData(key.Email+"encrypted_data", string(encryptedData[:]))
	
	fmt.Println("\nStored the data\n")
	fmt.Println()
	time.Sleep(2 * time.Second)
	
	fmt.Println("\nChoose one of the following:")
	fmt.Println("1 - Show me my secret.")
	fmt.Println("2 - Show my hidden secret.")
	fmt.Println("3 - Forget my secret forever.\n")
	for scanner.Scan() {
		input := scanner.Text()
		time.Sleep(1 * time.Second)
		
		switch input {
		case "1":
			storedEncryptedData := storage.RetrieveData(key.Email+"encrypted_data")
			decryptedAge, err := storage.Decrypt(decodeKey(key.DO_NOT_STORE_DO_NOT_LOG), []byte(storedEncryptedData))
			
			if err == nil {
				fmt.Println("\nYour secret:\n")
				fmt.Println(string(decryptedAge[:]))
			} else {
				fmt.Println("\nCould not decrypt:\n")
				fmt.Println(err.Error())
			}
		case "2":
			storedEncryptedData := storage.RetrieveData(key.Email+"encrypted_data")
			fmt.Println("\nYour encrypted secret:\n")
			fmt.Println(string(storedEncryptedData))
		case "3":
			fmt.Println("\nRemoving your key from submarine\n")
			apiKeyHeader := &Header{"API-KEY", app.SECRET_API_KEY}
			doDelete("http://localhost:3000/keys/"+string(key.Id), apiKeyHeader)
			key.DO_NOT_STORE_DO_NOT_LOG = ""
			
		}
		time.Sleep(2 * time.Second)

		fmt.Println("\n\n\nChoose one of the following:")
		fmt.Println("1 Show me my secret.")
		fmt.Println("2 Show my hidden secret.")
		fmt.Println("3 Forget my secret forever.\n")
	}
	
	
}   

func decodeKey(encodedKey string) []byte {
	decodedKey, _ := base64.StdEncoding.DecodeString(encodedKey)
	
	return decodedKey
}

func createApp() apps.App{
	var app apps.App
	appJson := doPost("http://localhost:3001/apps", `{"Name":"simple"}`, nil)
	parseJson(&app, appJson)
	
	return app
}

func createKey(app apps.App, email string) keys.Key{
	apiKeyHeader := &Header{"API-KEY", app.SECRET_API_KEY}
	
	var key keys.Key
	reqJson := "{ \"Email\":\"" + email + "\", \"App\": 0 }"
	
	fmt.Println(reqJson)
	keyJson := doPost("http://localhost:3000/keys", reqJson, apiKeyHeader)

	parseJson(&key, keyJson)
	
	return key
}

func getKey(app apps.App, email string) keys.Key{
	apiKeyHeader := &Header{"API-KEY", app.SECRET_API_KEY}
	
	var key keys.Key
	reqJson := "{ \"Email\":\"" + email + "\", \"App\": 0 }"
	
	fmt.Println(reqJson)
	keyJson := doPost("http://localhost:3000/keys", reqJson, apiKeyHeader)

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
    //fmt.Println("response Body:", string(body))
	
	return body
	
}

func doDelete(url string, header *Header) {
	req, _ := http.NewRequest("DELETE", url, nil)
	client := &http.Client{}
	if header != nil {
		req.Header.Set(header.Name, header.Value)
	}
	
	resp, _ := client.Do(req)
	defer resp.Body.Close()
}