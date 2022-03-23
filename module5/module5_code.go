package module5

import (
	"fmt"
	"net/http"
)

func GetExampleDotCom() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println("something went wrong")
	}

	defer resp.Body.Close()
}
