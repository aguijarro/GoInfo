package main



import (
	"goinfo/data/mongo"
	"net/http"
)

func analisis(writer http.ResponseWriter, request *http.Request) {

		results, err := mongo.ShowContentMongo()

	if err != nil {
		//error_message(writer, request, "Cannot read thread")
	} else {
		generateHTML(writer, results, "layout", "navbar", "nosql/analisis", "footer")
	}

}
