// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.32
// source: commonfile/fileproto/protos/file.proto

package fileproto

import (
	bytes "bytes"
	context "context"
	errors "errors"
	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_commonfile_fileproto_protos_file_proto struct{}

func (drpcEncoding_File_commonfile_fileproto_protos_file_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_commonfile_fileproto_protos_file_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_commonfile_fileproto_protos_file_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_commonfile_fileproto_protos_file_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCFileClient interface {
	DRPCConn() drpc.Conn

	BlockGet(ctx context.Context, in *BlockGetRequest) (*BlockGetResponse, error)
	BlockPush(ctx context.Context, in *BlockPushRequest) (*BlockPushResponse, error)
	BlocksCheck(ctx context.Context, in *BlocksCheckRequest) (*BlocksCheckResponse, error)
	BlocksBind(ctx context.Context, in *BlocksBindRequest) (*BlocksBindResponse, error)
	FilesDelete(ctx context.Context, in *FilesDeleteRequest) (*FilesDeleteResponse, error)
	FilesCheck(ctx context.Context, in *FilesCheckRequest) (*FilesCheckResponse, error)
	Check(ctx context.Context, in *CheckRequest) (*CheckResponse, error)
	CheckUsage(ctx context.Context, in *CheckUsageRequest) (*CheckUsageResponse, error)
}

type drpcFileClient struct {
	cc drpc.Conn
}

func NewDRPCFileClient(cc drpc.Conn) DRPCFileClient {
	return &drpcFileClient{cc}
}

func (c *drpcFileClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcFileClient) BlockGet(ctx context.Context, in *BlockGetRequest) (*BlockGetResponse, error) {
	out := new(BlockGetResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/BlockGet", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) BlockPush(ctx context.Context, in *BlockPushRequest) (*BlockPushResponse, error) {
	out := new(BlockPushResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/BlockPush", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) BlocksCheck(ctx context.Context, in *BlocksCheckRequest) (*BlocksCheckResponse, error) {
	out := new(BlocksCheckResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/BlocksCheck", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) BlocksBind(ctx context.Context, in *BlocksBindRequest) (*BlocksBindResponse, error) {
	out := new(BlocksBindResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/BlocksBind", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) FilesDelete(ctx context.Context, in *FilesDeleteRequest) (*FilesDeleteResponse, error) {
	out := new(FilesDeleteResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/FilesDelete", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) FilesCheck(ctx context.Context, in *FilesCheckRequest) (*FilesCheckResponse, error) {
	out := new(FilesCheckResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/FilesCheck", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) Check(ctx context.Context, in *CheckRequest) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/Check", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcFileClient) CheckUsage(ctx context.Context, in *CheckUsageRequest) (*CheckUsageResponse, error) {
	out := new(CheckUsageResponse)
	err := c.cc.Invoke(ctx, "/filesync.File/CheckUsage", drpcEncoding_File_commonfile_fileproto_protos_file_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCFileServer interface {
	BlockGet(context.Context, *BlockGetRequest) (*BlockGetResponse, error)
	BlockPush(context.Context, *BlockPushRequest) (*BlockPushResponse, error)
	BlocksCheck(context.Context, *BlocksCheckRequest) (*BlocksCheckResponse, error)
	BlocksBind(context.Context, *BlocksBindRequest) (*BlocksBindResponse, error)
	FilesDelete(context.Context, *FilesDeleteRequest) (*FilesDeleteResponse, error)
	FilesCheck(context.Context, *FilesCheckRequest) (*FilesCheckResponse, error)
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	CheckUsage(context.Context, *CheckUsageRequest) (*CheckUsageResponse, error)
}

type DRPCFileUnimplementedServer struct{}

func (s *DRPCFileUnimplementedServer) BlockGet(context.Context, *BlockGetRequest) (*BlockGetResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) BlockPush(context.Context, *BlockPushRequest) (*BlockPushResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) BlocksCheck(context.Context, *BlocksCheckRequest) (*BlocksCheckResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) BlocksBind(context.Context, *BlocksBindRequest) (*BlocksBindResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) FilesDelete(context.Context, *FilesDeleteRequest) (*FilesDeleteResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) FilesCheck(context.Context, *FilesCheckRequest) (*FilesCheckResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCFileUnimplementedServer) CheckUsage(context.Context, *CheckUsageRequest) (*CheckUsageResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCFileDescription struct{}

func (DRPCFileDescription) NumMethods() int { return 8 }

func (DRPCFileDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/filesync.File/BlockGet", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					BlockGet(
						ctx,
						in1.(*BlockGetRequest),
					)
			}, DRPCFileServer.BlockGet, true
	case 1:
		return "/filesync.File/BlockPush", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					BlockPush(
						ctx,
						in1.(*BlockPushRequest),
					)
			}, DRPCFileServer.BlockPush, true
	case 2:
		return "/filesync.File/BlocksCheck", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					BlocksCheck(
						ctx,
						in1.(*BlocksCheckRequest),
					)
			}, DRPCFileServer.BlocksCheck, true
	case 3:
		return "/filesync.File/BlocksBind", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					BlocksBind(
						ctx,
						in1.(*BlocksBindRequest),
					)
			}, DRPCFileServer.BlocksBind, true
	case 4:
		return "/filesync.File/FilesDelete", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					FilesDelete(
						ctx,
						in1.(*FilesDeleteRequest),
					)
			}, DRPCFileServer.FilesDelete, true
	case 5:
		return "/filesync.File/FilesCheck", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					FilesCheck(
						ctx,
						in1.(*FilesCheckRequest),
					)
			}, DRPCFileServer.FilesCheck, true
	case 6:
		return "/filesync.File/Check", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					Check(
						ctx,
						in1.(*CheckRequest),
					)
			}, DRPCFileServer.Check, true
	case 7:
		return "/filesync.File/CheckUsage", drpcEncoding_File_commonfile_fileproto_protos_file_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCFileServer).
					CheckUsage(
						ctx,
						in1.(*CheckUsageRequest),
					)
			}, DRPCFileServer.CheckUsage, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterFile(mux drpc.Mux, impl DRPCFileServer) error {
	return mux.Register(impl, DRPCFileDescription{})
}

type DRPCFile_BlockGetStream interface {
	drpc.Stream
	SendAndClose(*BlockGetResponse) error
}

type drpcFile_BlockGetStream struct {
	drpc.Stream
}

func (x *drpcFile_BlockGetStream) SendAndClose(m *BlockGetResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_BlockPushStream interface {
	drpc.Stream
	SendAndClose(*BlockPushResponse) error
}

type drpcFile_BlockPushStream struct {
	drpc.Stream
}

func (x *drpcFile_BlockPushStream) SendAndClose(m *BlockPushResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_BlocksCheckStream interface {
	drpc.Stream
	SendAndClose(*BlocksCheckResponse) error
}

type drpcFile_BlocksCheckStream struct {
	drpc.Stream
}

func (x *drpcFile_BlocksCheckStream) SendAndClose(m *BlocksCheckResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_BlocksBindStream interface {
	drpc.Stream
	SendAndClose(*BlocksBindResponse) error
}

type drpcFile_BlocksBindStream struct {
	drpc.Stream
}

func (x *drpcFile_BlocksBindStream) SendAndClose(m *BlocksBindResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_FilesDeleteStream interface {
	drpc.Stream
	SendAndClose(*FilesDeleteResponse) error
}

type drpcFile_FilesDeleteStream struct {
	drpc.Stream
}

func (x *drpcFile_FilesDeleteStream) SendAndClose(m *FilesDeleteResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_FilesCheckStream interface {
	drpc.Stream
	SendAndClose(*FilesCheckResponse) error
}

type drpcFile_FilesCheckStream struct {
	drpc.Stream
}

func (x *drpcFile_FilesCheckStream) SendAndClose(m *FilesCheckResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_CheckStream interface {
	drpc.Stream
	SendAndClose(*CheckResponse) error
}

type drpcFile_CheckStream struct {
	drpc.Stream
}

func (x *drpcFile_CheckStream) SendAndClose(m *CheckResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCFile_CheckUsageStream interface {
	drpc.Stream
	SendAndClose(*CheckUsageResponse) error
}

type drpcFile_CheckUsageStream struct {
	drpc.Stream
}

func (x *drpcFile_CheckUsageStream) SendAndClose(m *CheckUsageResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonfile_fileproto_protos_file_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}
