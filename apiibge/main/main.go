package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	name := []string{"Sergio", "Jessica", "Maria", "João", "Amanda"}

	for i := 0; i < 5; i++ {
		resp, _ := http.Get("https://servicodados.ibge.gov.br/api/v2/censos/nomes/" + name[i])
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println((string)(body))
	}

	//defer resp.Body.Close()
}
