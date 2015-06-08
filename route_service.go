package main

import (
	"goinfo/data/services"
	"net/http"
)

func continente(writer http.ResponseWriter, request *http.Request) {

		results, err := services.ShowContent("http://quito.cloudapi.junar.com/datastreams/invoke/ESTAD-PROME-DIAS-DE-TURIS?auth_key=780dab5f961cc66d490610e5be5c0364ffd9b43a&output=json_array",82,87)

	if err != nil {
		//error_message(writer, request, "Cannot read thread")
	} else {
		generateHTML(writer, results, "layout", "navbar", "services/continente", "footer")
	}

}
