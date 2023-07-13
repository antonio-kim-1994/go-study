package main

import (
	"aijinet/ncp-with-go/api"
	"aijinet/ncp-with-go/signatures"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	err := godotenv.Load("./ncp.env")
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	// Global Variable
	accKey := os.Getenv("NCP_ACC_KEY")
	secKey := os.Getenv("NCP_SEC_KEY")
	now := time.Now().UnixMilli()
	timestamp := strconv.FormatInt(now, 10)

	// User Input
	var input int
	fmt.Println("===== Aijinet NCP API Selector =====")
	fmt.Println("1 : Export DB Server Log to Object Storage")
	fmt.Println("2 : Get DB Server Log List")
	fmt.Print("Enter a Number : ")
	_, _ = fmt.Scan(&input)
	if input == 1 {
		fmt.Println(" * Export DB Server Log to Object Storage will proceed. *")
	} else if input == 2 {
		fmt.Println("* Get DB Server Log List will proceed. *")
	} else {
		fmt.Println(" !! You entered wrong number !! ")
	}

	// Goroutine standby setting
	var wg sync.WaitGroup
	wg.Add(1)

	go func(input int) {
		defer wg.Done()

		var apiSelector api.ReturnValue

		switch input {
		case 1:
			apiSelector = api.ExportDbServerLogToObjectStorage()
		case 2:
			apiSelector = api.GetDbServerLogList()
		default:
			fmt.Println("Input type isn't correct. Check your command.")
			return
		}

		method := apiSelector.Method
		uri := apiSelector.Uri
		baseUrl := apiSelector.BaseUrl
		fullUrl := baseUrl + uri

		// Generate Signature key
		signature, err := signatures.GenerateHmac(method, uri, timestamp, accKey, secKey)
		if err != nil {
			log.Err(err)
		}

		req, err := http.NewRequest(method, fullUrl, nil)
		if err != nil {
			log.Err(err)
			return
		}

		req.Header.Add("Content-Type", "application/json; charset=utf-8")
		req.Header.Add("x-ncp-apigw-timestamp", timestamp)
		req.Header.Add("x-ncp-iam-access-key", accKey)
		req.Header.Add("x-ncp-apigw-signature-v2", signature)

		client := &http.Client{
			Timeout: 60 * time.Second,
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Err(err)
			return
		}

		defer resp.Body.Close()

		log.Printf("Response status : %s", resp.Status)

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Err(err)
			return
		}

		var jsonResult bytes.Buffer

		err = json.Indent(&jsonResult, []byte(string(body)), "", " ")
		if err != nil {
			log.Err(err)
			return
		}

		log.Printf("\n%v", jsonResult.String())
	}(input)

	wg.Wait()
}
