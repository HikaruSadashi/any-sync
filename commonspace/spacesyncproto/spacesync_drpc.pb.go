// Code generated by protoc-gen-go-drpc. DO NOT EDIT.
// protoc-gen-go-drpc version: v0.0.32
// source: commonspace/spacesyncproto/protos/spacesync.proto

package spacesyncproto

import (
	bytes "bytes"
	context "context"
	errors "errors"
	jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	drpc "storj.io/drpc"
	drpcerr "storj.io/drpc/drpcerr"
)

type drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto struct{}

func (drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto) Marshal(msg drpc.Message) ([]byte, error) {
	return proto.Marshal(msg.(proto.Message))
}

func (drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto) Unmarshal(buf []byte, msg drpc.Message) error {
	return proto.Unmarshal(buf, msg.(proto.Message))
}

func (drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto) JSONMarshal(msg drpc.Message) ([]byte, error) {
	var buf bytes.Buffer
	err := new(jsonpb.Marshaler).Marshal(&buf, msg.(proto.Message))
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto) JSONUnmarshal(buf []byte, msg drpc.Message) error {
	return jsonpb.Unmarshal(bytes.NewReader(buf), msg.(proto.Message))
}

type DRPCSpaceSyncClient interface {
	DRPCConn() drpc.Conn

	HeadSync(ctx context.Context, in *HeadSyncRequest) (*HeadSyncResponse, error)
	SpacePush(ctx context.Context, in *SpacePushRequest) (*SpacePushResponse, error)
	SpacePull(ctx context.Context, in *SpacePullRequest) (*SpacePullResponse, error)
	ObjectSyncStream(ctx context.Context) (DRPCSpaceSync_ObjectSyncStreamClient, error)
}

type drpcSpaceSyncClient struct {
	cc drpc.Conn
}

func NewDRPCSpaceSyncClient(cc drpc.Conn) DRPCSpaceSyncClient {
	return &drpcSpaceSyncClient{cc}
}

func (c *drpcSpaceSyncClient) DRPCConn() drpc.Conn { return c.cc }

func (c *drpcSpaceSyncClient) HeadSync(ctx context.Context, in *HeadSyncRequest) (*HeadSyncResponse, error) {
	out := new(HeadSyncResponse)
	err := c.cc.Invoke(ctx, "/spacesync.SpaceSync/HeadSync", drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcSpaceSyncClient) SpacePush(ctx context.Context, in *SpacePushRequest) (*SpacePushResponse, error) {
	out := new(SpacePushResponse)
	err := c.cc.Invoke(ctx, "/spacesync.SpaceSync/SpacePush", drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcSpaceSyncClient) SpacePull(ctx context.Context, in *SpacePullRequest) (*SpacePullResponse, error) {
	out := new(SpacePullResponse)
	err := c.cc.Invoke(ctx, "/spacesync.SpaceSync/SpacePull", drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{}, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *drpcSpaceSyncClient) ObjectSyncStream(ctx context.Context) (DRPCSpaceSync_ObjectSyncStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, "/spacesync.SpaceSync/ObjectSyncStream", drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{})
	if err != nil {
		return nil, err
	}
	x := &drpcSpaceSync_ObjectSyncStreamClient{stream}
	return x, nil
}

type DRPCSpaceSync_ObjectSyncStreamClient interface {
	drpc.Stream
	Send(*ObjectSyncMessage) error
	Recv() (*ObjectSyncMessage, error)
}

type drpcSpaceSync_ObjectSyncStreamClient struct {
	drpc.Stream
}

func (x *drpcSpaceSync_ObjectSyncStreamClient) Send(m *ObjectSyncMessage) error {
	return x.MsgSend(m, drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{})
}

func (x *drpcSpaceSync_ObjectSyncStreamClient) Recv() (*ObjectSyncMessage, error) {
	m := new(ObjectSyncMessage)
	if err := x.MsgRecv(m, drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcSpaceSync_ObjectSyncStreamClient) RecvMsg(m *ObjectSyncMessage) error {
	return x.MsgRecv(m, drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{})
}

type DRPCSpaceSyncServer interface {
	HeadSync(context.Context, *HeadSyncRequest) (*HeadSyncResponse, error)
	SpacePush(context.Context, *SpacePushRequest) (*SpacePushResponse, error)
	SpacePull(context.Context, *SpacePullRequest) (*SpacePullResponse, error)
	ObjectSyncStream(DRPCSpaceSync_ObjectSyncStreamStream) error
}

type DRPCSpaceSyncUnimplementedServer struct{}

func (s *DRPCSpaceSyncUnimplementedServer) HeadSync(context.Context, *HeadSyncRequest) (*HeadSyncResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCSpaceSyncUnimplementedServer) SpacePush(context.Context, *SpacePushRequest) (*SpacePushResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCSpaceSyncUnimplementedServer) SpacePull(context.Context, *SpacePullRequest) (*SpacePullResponse, error) {
	return nil, drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

func (s *DRPCSpaceSyncUnimplementedServer) ObjectSyncStream(DRPCSpaceSync_ObjectSyncStreamStream) error {
	return drpcerr.WithCode(errors.New("Unimplemented"), drpcerr.Unimplemented)
}

type DRPCSpaceSyncDescription struct{}

func (DRPCSpaceSyncDescription) NumMethods() int { return 4 }

func (DRPCSpaceSyncDescription) Method(n int) (string, drpc.Encoding, drpc.Receiver, interface{}, bool) {
	switch n {
	case 0:
		return "/spacesync.SpaceSync/HeadSync", drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCSpaceSyncServer).
					HeadSync(
						ctx,
						in1.(*HeadSyncRequest),
					)
			}, DRPCSpaceSyncServer.HeadSync, true
	case 1:
		return "/spacesync.SpaceSync/SpacePush", drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCSpaceSyncServer).
					SpacePush(
						ctx,
						in1.(*SpacePushRequest),
					)
			}, DRPCSpaceSyncServer.SpacePush, true
	case 2:
		return "/spacesync.SpaceSync/SpacePull", drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return srv.(DRPCSpaceSyncServer).
					SpacePull(
						ctx,
						in1.(*SpacePullRequest),
					)
			}, DRPCSpaceSyncServer.SpacePull, true
	case 3:
		return "/spacesync.SpaceSync/ObjectSyncStream", drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{},
			func(srv interface{}, ctx context.Context, in1, in2 interface{}) (drpc.Message, error) {
				return nil, srv.(DRPCSpaceSyncServer).
					ObjectSyncStream(
						&drpcSpaceSync_ObjectSyncStreamStream{in1.(drpc.Stream)},
					)
			}, DRPCSpaceSyncServer.ObjectSyncStream, true
	default:
		return "", nil, nil, nil, false
	}
}

func DRPCRegisterSpaceSync(mux drpc.Mux, impl DRPCSpaceSyncServer) error {
	return mux.Register(impl, DRPCSpaceSyncDescription{})
}

type DRPCSpaceSync_HeadSyncStream interface {
	drpc.Stream
	SendAndClose(*HeadSyncResponse) error
}

type drpcSpaceSync_HeadSyncStream struct {
	drpc.Stream
}

func (x *drpcSpaceSync_HeadSyncStream) SendAndClose(m *HeadSyncResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCSpaceSync_SpacePushStream interface {
	drpc.Stream
	SendAndClose(*SpacePushResponse) error
}

type drpcSpaceSync_SpacePushStream struct {
	drpc.Stream
}

func (x *drpcSpaceSync_SpacePushStream) SendAndClose(m *SpacePushResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCSpaceSync_SpacePullStream interface {
	drpc.Stream
	SendAndClose(*SpacePullResponse) error
}

type drpcSpaceSync_SpacePullStream struct {
	drpc.Stream
}

func (x *drpcSpaceSync_SpacePullStream) SendAndClose(m *SpacePullResponse) error {
	if err := x.MsgSend(m, drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{}); err != nil {
		return err
	}
	return x.CloseSend()
}

type DRPCSpaceSync_ObjectSyncStreamStream interface {
	drpc.Stream
	Send(*ObjectSyncMessage) error
	Recv() (*ObjectSyncMessage, error)
}

type drpcSpaceSync_ObjectSyncStreamStream struct {
	drpc.Stream
}

func (x *drpcSpaceSync_ObjectSyncStreamStream) Send(m *ObjectSyncMessage) error {
	return x.MsgSend(m, drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{})
}

func (x *drpcSpaceSync_ObjectSyncStreamStream) Recv() (*ObjectSyncMessage, error) {
	m := new(ObjectSyncMessage)
	if err := x.MsgRecv(m, drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{}); err != nil {
		return nil, err
	}
	return m, nil
}

func (x *drpcSpaceSync_ObjectSyncStreamStream) RecvMsg(m *ObjectSyncMessage) error {
	return x.MsgRecv(m, drpcEncoding_File_commonspace_spacesyncproto_protos_spacesync_proto{})
}