package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// func main() {
// 	http.HandleFunc("/", hello)
// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

// 前一个括号内是入参(指定变量名和类型)，后一个是出参(且仅需指定类型)
func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?bc17dc9a7a036937af8f566cca3b54a6&q=" + city)
	// 若有访问错误则返回空信息和错误
	if err != nil {
		return weatherData{}, err
	}
	// 成功响应则关闭响应体
	defer resp.Body.Close()

	var d weatherData
	// 解码 Get 请求的响应
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
	// 返回信息和空错误
	return d, nil
}

// type 关键词定义一个新类型为一个 struct
type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}
