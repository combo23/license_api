package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/denisbrodbeck/machineid"
)

const License_key = "TEST"

func get_hwid() string {
	id, err := machineid.ID()
	if err != nil {
		panic(err)
	}
	return id
}

func main() {
	fmt.Println("Checking license...")

	hwid := get_hwid()

	rawPayload, err := json.Marshal(map[string]interface{}{
		"hwid":        hwid,
		"license_key": License_key,
	})
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", "http://127.0.0.1:8080/verify", bytes.NewReader(rawPayload))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	jsonBody := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println(jsonBody["message"])
		// verfied
	} else {
		fmt.Println(jsonBody["error"])
		os.Exit(1)
	}
}
