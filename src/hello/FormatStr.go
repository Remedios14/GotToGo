package hello

import (
	"fmt"
)

func main() {
	var stockcode = 123
	var enddate = "2022-09-02"
	var url = "Code=%d & endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
}