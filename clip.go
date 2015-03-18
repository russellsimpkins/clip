package clip

import (
	"encoding/json"
	"log"
	"net/http"
)

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
