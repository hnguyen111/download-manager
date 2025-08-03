package grpc

import (
	"context"

	"github.com/hnguyen111/download-manager/internal/generated/grpc/idm"
	"google.golang.org/grpc"
)

// Handler implements the gRPC server for the Download Service.
type Handler struct {
	idm.UnimplementedDownloadServiceServer
}

// NewHandler creates a new gRPC handler for the Download Service.
func NewHandler() *Handler {
	return &Handler{}
}

func (Handler) CreateAccount(_ context.Context, _ *idm.CreateAccountRequest) (_ *idm.CreateAccountResponse, _ error) {
	panic("not implemented") // TODO: Implement
}

func (Handler) CreateSession(_ context.Context, _ *idm.CreateSessionRequest) (_ *idm.CreateSessionResponse, _ error) {
	panic("not implemented") // TODO: Implement
}

func (Handler) CreateDownloadTask(_ context.Context, _ *idm.CreateDownloadTaskRequest) (_ *idm.CreateDownloadTaskResponse, _ error) {
	panic("not implemented") // TODO: Implement
}

func (Handler) GetDownloadTaskList(_ context.Context, _ *idm.GetDownloadTaskListRequest) (_ *idm.GetDownloadTaskListResponse, _ error) {
	panic("not implemented") // TODO: Implement
}

func (Handler) UpdateDownloadTask(_ context.Context, _ *idm.UpdateDownloadTaskRequest) (_ *idm.UpdateDownloadTaskResponse, _ error) {
	panic("not implemented") // TODO: Implement
}

func (Handler) DeleteDownloadTask(_ context.Context, _ *idm.DeleteDownloadTaskRequest) (_ *idm.DeleteDownloadTaskResponse, _ error) {
	panic("not implemented") // TODO: Implement
}

func (Handler) GetDownloadTaskFile(_ *idm.GetDownloadTaskFileRequest, _ grpc.ServerStreamingServer[idm.GetDownloadTaskFileResponse]) (_ error) {
	panic("not implemented") // TODO: Implement
}

func (Handler) mustEmbedUnimplementedDownloadServiceServer() {
	panic("not implemented") // TODO: Implement
}
