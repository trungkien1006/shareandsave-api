package boostraps

import (
	"final_project/internal/application/app/commentapp"
	"final_project/internal/infrastructure/grpcpb"
	persistence "final_project/internal/infrastructure/persistence/repo"
	grpcserver "final_project/internal/interface/grpc/handler"
	"log"
	"net"

	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func StartGRPCServer(db *gorm.DB) {
	lis, err := net.Listen("tcp", ":50051") // Có thể đổi port
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	repo := persistence.NewCommentRepoDB(db)
	commentUC := commentapp.NewUseCase(repo)
	handler := grpcserver.NewMessageHandlerServer(commentUC)

	grpcpb.RegisterMessageHandlerServer(grpcServer, handler)

	log.Println("gRPC server is listening on :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
