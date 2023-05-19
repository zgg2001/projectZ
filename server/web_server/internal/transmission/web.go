package transmission

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/zgg2001/projectZ/server/user_server/pkg/rpc"
	"github.com/zgg2001/projectZ/server/web_server/internal/operate"
	"google.golang.org/grpc"
)

func StartHttpServer(conn *grpc.ClientConn) error {
	log.Println("Listen and serve http server ...")

	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	operate.RPCConn = conn
	operate.RPCClient = rpc.NewProjectServiceClient(conn)
	http.Handle("/register", corsHandler(http.HandlerFunc(operate.HandleRegisterRequest)))
	http.Handle("/login", corsHandler(http.HandlerFunc(operate.HandleLoginRequest)))
	http.Handle("/logout", corsHandler(http.HandlerFunc(operate.HandleLogoutRequest)))
	http.Handle("/info", corsHandler(http.HandlerFunc(operate.HandleInfoRequest)))
	http.Handle("/recharge", corsHandler(http.HandlerFunc(operate.HandleRechargeRequest)))
	http.Handle("/operator", corsHandler(http.HandlerFunc(operate.HandleOperatorRequest)))

	err := http.ListenAndServe(":2222", nil)
	if err != nil {
		return err
	}
	return nil
}
