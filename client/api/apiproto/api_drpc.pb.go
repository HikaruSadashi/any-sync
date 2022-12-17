// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.32
// source: api/apiproto/protos/api.proto

package apiproto

import (
	bytes "bytes"
	context "context"
	errors "errors"
	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_api_apiproto_protos_api_proto struct{}

func (drpcEncoding_File_api_apiproto_protos_api_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_api_apiproto_protos_api_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_api_apiproto_protos_api_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_api_apiproto_protos_api_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCClientApiClient interface {
	DRPCConn() drpc.Conn

	CreateSpace(ctx context.Context, in *CreateSpaceRequest) (*CreateSpaceResponse, error)
	DeriveSpace(ctx context.Context, in *DeriveSpaceRequest) (*DeriveSpaceResponse, error)
	CreateDocument(ctx context.Context, in *CreateDocumentRequest) (*CreateDocumentResponse, error)
	DeleteDocument(ctx context.Context, in *DeleteDocumentRequest) (*DeleteDocumentResponse, error)
	AddText(ctx context.Context, in *AddTextRequest) (*AddTextResponse, error)
	DumpTree(ctx context.Context, in *DumpTreeRequest) (*DumpTreeResponse, error)
	TreeParams(ctx context.Context, in *TreeParamsRequest) (*TreeParamsResponse, error)
	AllTrees(ctx context.Context, in *AllTreesRequest) (*AllTreesResponse, error)
	AllSpaces(ctx context.Context, in *AllSpacesRequest) (*AllSpacesResponse, error)
	LoadSpace(ctx context.Context, in *LoadSpaceRequest) (*LoadSpaceResponse, error)
	PutFile(ctx context.Context, in *PutFileRequest) (*PutFileResponse, error)
	GetFile(ctx context.Context, in *GetFileRequest) (*GetFileResponse, error)
	DeleteFile(ctx context.Context, in *DeleteFileRequest) (*DeleteFileResponse, error)
}

type drpcClientApiClient struct {
	cc drpc.Conn
}

func NewDRPCClientApiClient(cc drpc.Conn) DRPCClientApiClient {
	return &drpcClientApiClient{cc}
}

func (c *drpcClientApiClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcClientApiClient) CreateSpace(ctx context.Context, in *CreateSpaceRequest) (*CreateSpaceResponse, error) {
	out := new(CreateSpaceResponse)
	err := c.cc.Invoke(ctx, "/clientapi.ClientApi/CreateSpace", drpcEncoding_File_api_apiproto_protos_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcClientApiClient) DeriveSpace(ctx context.Context, in *DeriveSpaceRequest) (*DeriveSpaceResponse, error) {
	out := new(DeriveSpaceResponse)
	err := c.cc.Invoke(ctx, "/clientapi.ClientApi/DeriveSpace", drpcEncoding_File_api_apiproto_protos_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcClientApiClient) CreateDocument(ctx context.Context, in *CreateDocumentRequest) (*CreateDocumentResponse, error) {
	out := new(CreateDocumentResponse)
	err := c.cc.Invoke(ctx, "/clientapi.ClientApi/CreateDocument", drpcEncoding_File_api_apiproto_protos_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcClientApiClient) DeleteDocument(ctx context.Context, in *DeleteDocumentRequest) (*DeleteDocumentResponse, error) {
	out := new(DeleteDocumentResponse)
	err := c.cc.Invoke(ctx, "/clientapi.ClientApi/DeleteDocument", drpcEncoding_File_api_apiproto_protos_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcClientApiClient) AddText(ctx context.Context, in *AddTextRequest) (*AddTextResponse, error) {
	out := new(AddTextResponse)
	err := c.cc.Invoke(ctx, "/clientapi.ClientApi/AddText", drpcEncoding_File_api_apiproto_protos_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcClientApiClient) DumpTree(ctx context.Context, in *DumpTreeRequest) (*DumpTreeResponse, error) {
	out := new(DumpTreeResponse)
	err := c.cc.Invoke(ctx, "/clientapi.ClientApi/DumpTree", drpcEncoding_File_api_apiproto_protos_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcClientApiClient) TreeParams(ctx context.Context, in *TreeParamsRequest) (*TreeParamsResponse, error) {
	out := new(TreeParamsResponse)
	err := c.cc.Invoke(ctx, "/clientapi.ClientApi/TreeParams", drpcEncoding_File_api_apiproto_protos_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcClientApiClient) AllTrees(ctx context.Context, in *AllTreesRequest) (*AllTreesResponse, error) {
	out := new(AllTreesResponse)
	err := c.cc.Invoke(ctx, "/clientapi.ClientApi/AllTrees", drpcEncoding_File_api_apiproto_protos_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcClientApiClient) AllSpaces(ctx context.Context, in *AllSpacesRequest) (*AllSpacesResponse, error) {
	out := new(AllSpacesResponse)
	err := c.cc.Invoke(ctx, "/clientapi.ClientApi/AllSpaces", drpcEncoding_File_api_apiproto_protos_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcClientApiClient) LoadSpace(ctx context.Context, in *LoadSpaceRequest) (*LoadSpaceResponse, error) {
	out := new(LoadSpaceResponse)
	err := c.cc.Invoke(ctx, "/clientapi.ClientApi/LoadSpace", drpcEncoding_File_api_apiproto_protos_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcClientApiClient) PutFile(ctx context.Context, in *PutFileRequest) (*PutFileResponse, error) {
	out := new(PutFileResponse)
	err := c.cc.Invoke(ctx, "/clientapi.ClientApi/PutFile", drpcEncoding_File_api_apiproto_protos_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcClientApiClient) GetFile(ctx context.Context, in *GetFileRequest) (*GetFileResponse, error) {
	out := new(GetFileResponse)
	err := c.cc.Invoke(ctx, "/clientapi.ClientApi/GetFile", drpcEncoding_File_api_apiproto_protos_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcClientApiClient) DeleteFile(ctx context.Context, in *DeleteFileRequest) (*DeleteFileResponse, error) {
	out := new(DeleteFileResponse)
	err := c.cc.Invoke(ctx, "/clientapi.ClientApi/DeleteFile", drpcEncoding_File_api_apiproto_protos_api_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DRPCClientApiServer interface {
	CreateSpace(context.Context, *CreateSpaceRequest) (*CreateSpaceResponse, error)
	DeriveSpace(context.Context, *DeriveSpaceRequest) (*DeriveSpaceResponse, error)
	CreateDocument(context.Context, *CreateDocumentRequest) (*CreateDocumentResponse, error)
	DeleteDocument(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error)
	AddText(context.Context, *AddTextRequest) (*AddTextResponse, error)
	DumpTree(context.Context, *DumpTreeRequest) (*DumpTreeResponse, error)
	TreeParams(context.Context, *TreeParamsRequest) (*TreeParamsResponse, error)
	AllTrees(context.Context, *AllTreesRequest) (*AllTreesResponse, error)
	AllSpaces(context.Context, *AllSpacesRequest) (*AllSpacesResponse, error)
	LoadSpace(context.Context, *LoadSpaceRequest) (*LoadSpaceResponse, error)
	PutFile(context.Context, *PutFileRequest) (*PutFileResponse, error)
	GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error)
	DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error)
}

type DRPCClientApiUnimplementedServer struct{}

func (s *DRPCClientApiUnimplementedServer) CreateSpace(context.Context, *CreateSpaceRequest) (*CreateSpaceResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCClientApiUnimplementedServer) DeriveSpace(context.Context, *DeriveSpaceRequest) (*DeriveSpaceResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCClientApiUnimplementedServer) CreateDocument(context.Context, *CreateDocumentRequest) (*CreateDocumentResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCClientApiUnimplementedServer) DeleteDocument(context.Context, *DeleteDocumentRequest) (*DeleteDocumentResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCClientApiUnimplementedServer) AddText(context.Context, *AddTextRequest) (*AddTextResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCClientApiUnimplementedServer) DumpTree(context.Context, *DumpTreeRequest) (*DumpTreeResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCClientApiUnimplementedServer) TreeParams(context.Context, *TreeParamsRequest) (*TreeParamsResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCClientApiUnimplementedServer) AllTrees(context.Context, *AllTreesRequest) (*AllTreesResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCClientApiUnimplementedServer) AllSpaces(context.Context, *AllSpacesRequest) (*AllSpacesResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCClientApiUnimplementedServer) LoadSpace(context.Context, *LoadSpaceRequest) (*LoadSpaceResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCClientApiUnimplementedServer) PutFile(context.Context, *PutFileRequest) (*PutFileResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCClientApiUnimplementedServer) GetFile(context.Context, *GetFileRequest) (*GetFileResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCClientApiUnimplementedServer) DeleteFile(context.Context, *DeleteFileRequest) (*DeleteFileResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCClientApiDescription struct{}

func (DRPCClientApiDescription) NumMethods() int { return 13 }

func (DRPCClientApiDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/clientapi.ClientApi/CreateSpace", drpcEncoding_File_api_apiproto_protos_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCClientApiServer).
					CreateSpace(
						ctx,
						in1.(*CreateSpaceRequest),
					)
			}, DRPCClientApiServer.CreateSpace, true
	case 1:
		return "/clientapi.ClientApi/DeriveSpace", drpcEncoding_File_api_apiproto_protos_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCClientApiServer).
					DeriveSpace(
						ctx,
						in1.(*DeriveSpaceRequest),
					)
			}, DRPCClientApiServer.DeriveSpace, true
	case 2:
		return "/clientapi.ClientApi/CreateDocument", drpcEncoding_File_api_apiproto_protos_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCClientApiServer).
					CreateDocument(
						ctx,
						in1.(*CreateDocumentRequest),
					)
			}, DRPCClientApiServer.CreateDocument, true
	case 3:
		return "/clientapi.ClientApi/DeleteDocument", drpcEncoding_File_api_apiproto_protos_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCClientApiServer).
					DeleteDocument(
						ctx,
						in1.(*DeleteDocumentRequest),
					)
			}, DRPCClientApiServer.DeleteDocument, true
	case 4:
		return "/clientapi.ClientApi/AddText", drpcEncoding_File_api_apiproto_protos_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCClientApiServer).
					AddText(
						ctx,
						in1.(*AddTextRequest),
					)
			}, DRPCClientApiServer.AddText, true
	case 5:
		return "/clientapi.ClientApi/DumpTree", drpcEncoding_File_api_apiproto_protos_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCClientApiServer).
					DumpTree(
						ctx,
						in1.(*DumpTreeRequest),
					)
			}, DRPCClientApiServer.DumpTree, true
	case 6:
		return "/clientapi.ClientApi/TreeParams", drpcEncoding_File_api_apiproto_protos_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCClientApiServer).
					TreeParams(
						ctx,
						in1.(*TreeParamsRequest),
					)
			}, DRPCClientApiServer.TreeParams, true
	case 7:
		return "/clientapi.ClientApi/AllTrees", drpcEncoding_File_api_apiproto_protos_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCClientApiServer).
					AllTrees(
						ctx,
						in1.(*AllTreesRequest),
					)
			}, DRPCClientApiServer.AllTrees, true
	case 8:
		return "/clientapi.ClientApi/AllSpaces", drpcEncoding_File_api_apiproto_protos_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCClientApiServer).
					AllSpaces(
						ctx,
						in1.(*AllSpacesRequest),
					)
			}, DRPCClientApiServer.AllSpaces, true
	case 9:
		return "/clientapi.ClientApi/LoadSpace", drpcEncoding_File_api_apiproto_protos_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCClientApiServer).
					LoadSpace(
						ctx,
						in1.(*LoadSpaceRequest),
					)
			}, DRPCClientApiServer.LoadSpace, true
	case 10:
		return "/clientapi.ClientApi/PutFile", drpcEncoding_File_api_apiproto_protos_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCClientApiServer).
					PutFile(
						ctx,
						in1.(*PutFileRequest),
					)
			}, DRPCClientApiServer.PutFile, true
	case 11:
		return "/clientapi.ClientApi/GetFile", drpcEncoding_File_api_apiproto_protos_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCClientApiServer).
					GetFile(
						ctx,
						in1.(*GetFileRequest),
					)
			}, DRPCClientApiServer.GetFile, true
	case 12:
		return "/clientapi.ClientApi/DeleteFile", drpcEncoding_File_api_apiproto_protos_api_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCClientApiServer).
					DeleteFile(
						ctx,
						in1.(*DeleteFileRequest),
					)
			}, DRPCClientApiServer.DeleteFile, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterClientApi(mux drpc.Mux, impl DRPCClientApiServer) error {
	return mux.Register(impl, DRPCClientApiDescription{})
}

type DRPCClientApi_CreateSpaceStream interface {
	drpc.Stream
	SendAndClose(*CreateSpaceResponse) error
}

type drpcClientApi_CreateSpaceStream struct {
	drpc.Stream
}

func (x *drpcClientApi_CreateSpaceStream) SendAndClose(m *CreateSpaceResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_api_apiproto_protos_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCClientApi_DeriveSpaceStream interface {
	drpc.Stream
	SendAndClose(*DeriveSpaceResponse) error
}

type drpcClientApi_DeriveSpaceStream struct {
	drpc.Stream
}

func (x *drpcClientApi_DeriveSpaceStream) SendAndClose(m *DeriveSpaceResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_api_apiproto_protos_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCClientApi_CreateDocumentStream interface {
	drpc.Stream
	SendAndClose(*CreateDocumentResponse) error
}

type drpcClientApi_CreateDocumentStream struct {
	drpc.Stream
}

func (x *drpcClientApi_CreateDocumentStream) SendAndClose(m *CreateDocumentResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_api_apiproto_protos_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCClientApi_DeleteDocumentStream interface {
	drpc.Stream
	SendAndClose(*DeleteDocumentResponse) error
}

type drpcClientApi_DeleteDocumentStream struct {
	drpc.Stream
}

func (x *drpcClientApi_DeleteDocumentStream) SendAndClose(m *DeleteDocumentResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_api_apiproto_protos_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCClientApi_AddTextStream interface {
	drpc.Stream
	SendAndClose(*AddTextResponse) error
}

type drpcClientApi_AddTextStream struct {
	drpc.Stream
}

func (x *drpcClientApi_AddTextStream) SendAndClose(m *AddTextResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_api_apiproto_protos_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCClientApi_DumpTreeStream interface {
	drpc.Stream
	SendAndClose(*DumpTreeResponse) error
}

type drpcClientApi_DumpTreeStream struct {
	drpc.Stream
}

func (x *drpcClientApi_DumpTreeStream) SendAndClose(m *DumpTreeResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_api_apiproto_protos_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCClientApi_TreeParamsStream interface {
	drpc.Stream
	SendAndClose(*TreeParamsResponse) error
}

type drpcClientApi_TreeParamsStream struct {
	drpc.Stream
}

func (x *drpcClientApi_TreeParamsStream) SendAndClose(m *TreeParamsResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_api_apiproto_protos_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCClientApi_AllTreesStream interface {
	drpc.Stream
	SendAndClose(*AllTreesResponse) error
}

type drpcClientApi_AllTreesStream struct {
	drpc.Stream
}

func (x *drpcClientApi_AllTreesStream) SendAndClose(m *AllTreesResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_api_apiproto_protos_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCClientApi_AllSpacesStream interface {
	drpc.Stream
	SendAndClose(*AllSpacesResponse) error
}

type drpcClientApi_AllSpacesStream struct {
	drpc.Stream
}

func (x *drpcClientApi_AllSpacesStream) SendAndClose(m *AllSpacesResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_api_apiproto_protos_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCClientApi_LoadSpaceStream interface {
	drpc.Stream
	SendAndClose(*LoadSpaceResponse) error
}

type drpcClientApi_LoadSpaceStream struct {
	drpc.Stream
}

func (x *drpcClientApi_LoadSpaceStream) SendAndClose(m *LoadSpaceResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_api_apiproto_protos_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCClientApi_PutFileStream interface {
	drpc.Stream
	SendAndClose(*PutFileResponse) error
}

type drpcClientApi_PutFileStream struct {
	drpc.Stream
}

func (x *drpcClientApi_PutFileStream) SendAndClose(m *PutFileResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_api_apiproto_protos_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCClientApi_GetFileStream interface {
	drpc.Stream
	SendAndClose(*GetFileResponse) error
}

type drpcClientApi_GetFileStream struct {
	drpc.Stream
}

func (x *drpcClientApi_GetFileStream) SendAndClose(m *GetFileResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_api_apiproto_protos_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCClientApi_DeleteFileStream interface {
	drpc.Stream
	SendAndClose(*DeleteFileResponse) error
}

type drpcClientApi_DeleteFileStream struct {
	drpc.Stream
}

func (x *drpcClientApi_DeleteFileStream) SendAndClose(m *DeleteFileResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_api_apiproto_protos_api_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}