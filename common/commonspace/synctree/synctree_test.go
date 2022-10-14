package synctree

import (
	"context"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/spacesyncproto"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/syncservice"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/syncservice/mock_syncservice"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/synctree/updatelistener"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/synctree/updatelistener/mock_updatelistener"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/pkg/acl/list"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/pkg/acl/list/mock_list"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/pkg/acl/storage"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/pkg/acl/storage/mock_storage"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/pkg/acl/tree"
	mock_tree "github.com/anytypeio/go-anytype-infrastructure-experiments/pkg/acl/tree/mock_objecttree"
	"github.com/anytypeio/go-anytype-infrastructure-experiments/pkg/acl/treechangeproto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

type syncTreeMatcher struct {
	objTree  tree.ObjectTree
	client   syncservice.SyncClient
	listener updatelistener.UpdateListener
}

func (s syncTreeMatcher) Matches(x interface{}) bool {
	t, ok := x.(*SyncTree)
	if !ok {
		return false
	}
	return s.objTree == t.ObjectTree && t.syncClient == s.client && t.listener == s.listener
}

func (s syncTreeMatcher) String() string {
	return ""
}

func Test_DeriveSyncTree(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	updateListenerMock := mock_updatelistener.NewMockUpdateListener(ctrl)
	syncClientMock := mock_syncservice.NewMockSyncClient(ctrl)
	aclListMock := mock_list.NewMockACLList(ctrl)
	createStorage := storage.TreeStorageCreatorFunc(func(payload storage.TreeStorageCreatePayload) (storage.TreeStorage, error) {
		return nil, nil
	})
	objTreeMock := mock_tree.NewMockObjectTree(ctrl)
	createDerivedObjectTree = func(payload tree.ObjectTreeCreatePayload, l list.ACLList, create storage.TreeStorageCreatorFunc) (objTree tree.ObjectTree, err error) {
		require.Equal(t, l, aclListMock)
		return objTreeMock, nil
	}
	headUpdate := &spacesyncproto.ObjectSyncMessage{}
	syncClientMock.EXPECT().CreateHeadUpdate(syncTreeMatcher{objTreeMock, syncClientMock, updateListenerMock}, gomock.Nil()).Return(headUpdate)
	syncClientMock.EXPECT().BroadcastAsync(gomock.Eq(headUpdate)).Return(nil)

	_, err := DeriveSyncTree(ctx, tree.ObjectTreeCreatePayload{}, syncClientMock, updateListenerMock, aclListMock, createStorage)
	require.NoError(t, err)
}

func Test_CreateSyncTree(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	updateListenerMock := mock_updatelistener.NewMockUpdateListener(ctrl)
	syncClientMock := mock_syncservice.NewMockSyncClient(ctrl)
	aclListMock := mock_list.NewMockACLList(ctrl)
	createStorage := storage.TreeStorageCreatorFunc(func(payload storage.TreeStorageCreatePayload) (storage.TreeStorage, error) {
		return nil, nil
	})
	objTreeMock := mock_tree.NewMockObjectTree(ctrl)
	createObjectTree = func(payload tree.ObjectTreeCreatePayload, l list.ACLList, create storage.TreeStorageCreatorFunc) (objTree tree.ObjectTree, err error) {
		require.Equal(t, l, aclListMock)
		return objTreeMock, nil
	}
	headUpdate := &spacesyncproto.ObjectSyncMessage{}
	syncClientMock.EXPECT().CreateHeadUpdate(syncTreeMatcher{objTreeMock, syncClientMock, updateListenerMock}, gomock.Nil()).Return(headUpdate)
	syncClientMock.EXPECT().BroadcastAsync(gomock.Eq(headUpdate)).Return(nil)

	_, err := CreateSyncTree(ctx, tree.ObjectTreeCreatePayload{}, syncClientMock, updateListenerMock, aclListMock, createStorage)
	require.NoError(t, err)
}

func Test_BuildSyncTree(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	updateListenerMock := mock_updatelistener.NewMockUpdateListener(ctrl)
	syncClientMock := mock_syncservice.NewMockSyncClient(ctrl)
	aclListMock := mock_list.NewMockACLList(ctrl)
	storageMock := mock_storage.NewMockTreeStorage(ctrl)
	objTreeMock := mock_tree.NewMockObjectTree(ctrl)
	buildObjectTree = func(store storage.TreeStorage, l list.ACLList) (objTree tree.ObjectTree, err error) {
		require.Equal(t, aclListMock, l)
		require.Equal(t, store, storageMock)
		return objTreeMock, nil
	}
	headUpdate := &spacesyncproto.ObjectSyncMessage{}
	syncClientMock.EXPECT().CreateHeadUpdate(syncTreeMatcher{objTreeMock, syncClientMock, updateListenerMock}, gomock.Nil()).Return(headUpdate)
	syncClientMock.EXPECT().BroadcastAsyncOrSendResponsible(gomock.Eq(headUpdate)).Return(nil)

	tr, err := BuildSyncTree(ctx, syncClientMock, storageMock, updateListenerMock, aclListMock)
	require.NoError(t, err)

	t.Run("AddRawChanges update", func(t *testing.T) {
		changes := []*treechangeproto.RawTreeChangeWithId{{Id: "some"}}
		expectedRes := tree.AddResult{
			Added: changes,
			Mode:  tree.Append,
		}
		objTreeMock.EXPECT().AddRawChanges(gomock.Any(), gomock.Eq(changes)).
			Return(expectedRes, nil)
		updateListenerMock.EXPECT().Update(tr)

		syncClientMock.EXPECT().CreateHeadUpdate(gomock.Eq(tr), gomock.Eq(changes)).Return(headUpdate)
		syncClientMock.EXPECT().BroadcastAsync(gomock.Eq(headUpdate)).Return(nil)
		res, err := tr.AddRawChanges(ctx, changes...)
		require.NoError(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("AddRawChanges rebuild", func(t *testing.T) {
		changes := []*treechangeproto.RawTreeChangeWithId{{Id: "some"}}
		expectedRes := tree.AddResult{
			Added: changes,
			Mode:  tree.Rebuild,
		}
		objTreeMock.EXPECT().AddRawChanges(gomock.Any(), gomock.Eq(changes)).
			Return(expectedRes, nil)
		updateListenerMock.EXPECT().Rebuild(tr)

		syncClientMock.EXPECT().CreateHeadUpdate(gomock.Eq(tr), gomock.Eq(changes)).Return(headUpdate)
		syncClientMock.EXPECT().BroadcastAsync(gomock.Eq(headUpdate)).Return(nil)
		res, err := tr.AddRawChanges(ctx, changes...)
		require.NoError(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("AddRawChanges nothing", func(t *testing.T) {
		changes := []*treechangeproto.RawTreeChangeWithId{{Id: "some"}}
		expectedRes := tree.AddResult{
			Added: changes,
			Mode:  tree.Nothing,
		}
		objTreeMock.EXPECT().AddRawChanges(gomock.Any(), gomock.Eq(changes)).
			Return(expectedRes, nil)

		res, err := tr.AddRawChanges(ctx, changes...)
		require.NoError(t, err)
		require.Equal(t, expectedRes, res)
	})

	t.Run("AddContent", func(t *testing.T) {
		changes := []*treechangeproto.RawTreeChangeWithId{{Id: "some"}}
		content := tree.SignableChangeContent{
			Data: []byte("abcde"),
		}
		expectedRes := tree.AddResult{
			Mode:  tree.Append,
			Added: changes,
		}
		objTreeMock.EXPECT().AddContent(gomock.Any(), gomock.Eq(content)).
			Return(expectedRes, nil)

		syncClientMock.EXPECT().CreateHeadUpdate(gomock.Eq(tr), gomock.Eq(changes)).Return(headUpdate)
		syncClientMock.EXPECT().BroadcastAsync(gomock.Eq(headUpdate)).Return(nil)
		res, err := tr.AddContent(ctx, content)
		require.NoError(t, err)
		require.Equal(t, expectedRes, res)
	})
}