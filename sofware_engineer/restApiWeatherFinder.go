package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func restApi(name string) int32 {
	degrees := int32(0)
	url := "https://jsonmock.hackerrank.com/api/weather?name=" + name

	method := "GET"

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	res, _ := client.Do(req)
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	data := string(body)

	var weatherData map[string]interface{}
	if err := json.Unmarshal([]byte(data), &weatherData); err != nil {
		fmt.Println(err)
		return int32(0)
	}

	if dataArray, ok := weatherData["data"].([]interface{}); ok {
		for _, arr := range dataArray {
			numStr := strings.Split(arr.(map[string]interface{})["weather"].(string), " ")[0]
			num, _ := strconv.Atoi(numStr)
			degrees += int32(num)
		}
	}
	return degrees
}

func main() {
	fmt.Println(restApi("Dallas"))
}
