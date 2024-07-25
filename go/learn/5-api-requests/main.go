package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type InputModel struct {
	Input float32 `json:"input"`
	Name  string  `json:"name"`
}

type ResponseModel struct {
	Output float32 `json:"output"`
	Status bool    `json:"status"`
}

const url = "http://localhost:8000/predict"

func sendGetRequest() {
	resp, err := http.Get("http://localhost:8000")
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)
	fmt.Println(string(content))

}

func sendPostRequest(payload io.Reader) {
	resp, err := http.Post(url, "application/json", payload)
	if err != nil {
		fmt.Print(err)
	}
	defer resp.Body.Close()
	content, _ := io.ReadAll(resp.Body)
	fmt.Println("Response content: ", string(content))

	checkValid := json.Valid(content)
	if checkValid {
		fmt.Println("Response content is valid JSON")
		var responseModel ResponseModel
		err := json.Unmarshal(content, &responseModel)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ResponseModel: %#v\n", responseModel)
	}
}

func main() {
	sendGetRequest()

	// Method 1
	payload := strings.NewReader(`{"input": 2, "name": "Integer input"}`)
	sendPostRequest(payload)

	// Method 2
	inputModel := InputModel{Input: 5.5, Name: "Float input"}
	// jsonPayload, err := json.Marshal(inputModel)
	jsonPayload, err := json.MarshalIndent(inputModel, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("inputModel struct: %#v\n", inputModel)
	fmt.Println("jsonPayload: ", string(jsonPayload))
	reader := bytes.NewReader(jsonPayload)
	sendPostRequest(reader)
}
