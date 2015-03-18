package clip

import (
	"encoding/json"
	"log"
	"net/http"
)

// for writing standard errors
func SendError(status int, data string, writer http.ResponseWriter) {
	what := WebResponse{}
	what.Status = status
	what.Message = data
	body, _ := json.Marshal(what)
	out := string(body)
	log.Print(out)
	http.Error(writer, out, status)
	return
}

// for when you just want a simple success message
func SendSuccess(writer http.ResponseWriter) {
	what := WebResponse{}
	what.Status = 200
	what.Message = "success"
	body, _ := json.Marshal(what)
	writer.Write(body)
}
