package transmission

import (
	"log"
	"net/http"

	"github.com/zgg2001/projectZ/server/web_server/internal/operate"
)

func StartHttpServer() error {
	log.Println("Listen and serve http server ...")

	http.HandleFunc("/", operate.HandleRootRequest)
	http.HandleFunc("/login", operate.HandleLoginRequest)

	err := http.ListenAndServe(":2222", nil)
	if err != nil {
		return err
	}
	return nil
}
