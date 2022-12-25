package routes

import (
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
	"picket/src/config"
	answersheet_repository "picket/src/features/answersheet/repository"
	answersheet_transport "picket/src/features/answersheet/transport"
	answersheet_usecase "picket/src/features/answersheet/usecase"
	answersheetpb "picket/src/pb/answer_sheet"
)

type handler = func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

func RouteGrpc(ctx context.Context, s *grpc.Server, config config.IConfig) {
	answersheetRepository := answersheet_repository.New(config.GetMongo())
	answersheetUsecase := answersheet_usecase.New(answersheetRepository)
	answerSheetController := answersheet_transport.New(ctx, answersheetUsecase)

	answersheetpb.RegisterAnswerSheetServiceServer(s, answerSheetController)
}

func RouteGw(ctx context.Context, gw *runtime.ServeMux, conn *grpc.ClientConn) {
	lists := []handler{
		//authpb.RegisterAuthServiceHandler,
		//userpb.RegisterUserServiceHandler,
		//classpb.RegisterClassServiceHandler,
		//testpb.RegisterTestServiceHandler,
		//answersheetpb.RegisterAnswerSheetServiceHandler,
	}

	for _, item := range lists {
		err := item(ctx, gw, conn)
		if err != nil {
			zap.S().Fatalln(err)
		}
	}
	gw.HandlePath(http.MethodGet, "/health", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		res := map[string]string{
			"message": "server is running 123",
		}
		json.NewEncoder(w).Encode(res)
	})
}
