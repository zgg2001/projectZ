package transmission

import (
	"log"
	"net/http"

	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
	"github.com/zgg2001/projectZ/server/web_server/internal/operate"
	"google.golang.org/grpc"
)

func StartHttpServer(conn *grpc.ClientConn) error {
	log.Println("Listen and serve http server ...")

	operate.RPCConn = conn
	operate.RPCClient = rpc.NewProjectServiceClient(conn)
	http.HandleFunc("/", operate.HandleRootRequest)
	http.HandleFunc("/login", operate.HandleLoginRequest)

	err := http.ListenAndServe(":2222", nil)
	if err != nil {
		return err
	}
	return nil
}
