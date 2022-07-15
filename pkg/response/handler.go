package response

import (
	"log"
	"net/http"
)

func HandleMessage(w http.ResponseWriter, code int, customError error) {
	// w.WriteHeader(code)
	// err := json.NewEncoder(w).Encode(response{
	// 	Data: customError.Error(),
	// })
	// if err != nil {
	// 	panic(err)
	// }
	log.Print("error: %w", customError)
}
