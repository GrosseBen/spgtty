package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

const symbolsInChunk = 1024

type putCodeRequest struct {
	ID     string `json:"id"`
	Code   string `json:"code"`
	Append bool   `json:"append"`
}

func putChunk(host, id string, data string, append bool) error {
	url := fmt.Sprintf("http://%s/rpc/Script.PutCode", host)
	req := putCodeRequest{
		ID:     id,
		Code:   data,
		Append: append,
	}
	reqData, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqData))
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	fmt.Println(result)
	return nil
}

func upload(host, id, file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("failed to read file: %v\n", err)
		os.Exit(1)
	}

	code := string(content)
	fmt.Printf("total %d bytes\n", len(code))

	pos := 0
	append := false
	for pos < len(code) {
		end := pos + symbolsInChunk
		if end > len(code) {
			end = len(code)
		}
		chunk := code[pos:end]
		if err := putChunk(host, id, chunk, append); err != nil {
			fmt.Printf("failed to put chunk: %v\n", err)
			os.Exit(1)
		}
		pos = end
		append = true
	}
}

func main() {
	host := flag.String("host", "", "IP address or hostname of the Shelly device")
	id := flag.String("id", "", "ID of the script being uploaded")
	file := flag.String("file", "", "Local file containing the script code to upload")
	flag.Parse()

	if *host == "" || *id == "" || *file == "" {
		fmt.Println("host, id, and file are required")
		os.Exit(1)
	}

	upload(*host, *id, *file)
}
