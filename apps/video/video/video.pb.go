// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.20.1
// source: video.proto

package video

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VideoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            // 视频唯一标识
	AuthorId      int64  `protobuf:"varint,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`                // 视频作者信息
	Title         string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`                                       // 视频标题
	PlayUrl       string `protobuf:"bytes,4,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`                    // 视频播放地址
	CoverUrl      string `protobuf:"bytes,5,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`                 // 视频封面地址
	FavoriteCount int64  `protobuf:"varint,6,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int64  `protobuf:"varint,7,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`    // 视频的评论总数
	PublishTime   int64  `protobuf:"varint,8,opt,name=publish_time,json=publishTime,proto3" json:"publish_time,omitempty"`       // 视频发布时间
}

func (x *VideoItem) Reset() {
	*x = VideoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoItem) ProtoMessage() {}

func (x *VideoItem) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoItem.ProtoReflect.Descriptor instead.
func (*VideoItem) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{0}
}

func (x *VideoItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VideoItem) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *VideoItem) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *VideoItem) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *VideoItem) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *VideoItem) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *VideoItem) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *VideoItem) GetPublishTime() int64 {
	if x != nil {
		return x.PublishTime
	}
	return 0
}

type PublishVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId int64  `protobuf:"varint,1,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"` // 视频作者信息
	PlayUrl  string `protobuf:"bytes,2,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`     // 视频播放地址
	CoverUrl string `protobuf:"bytes,3,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`  // 视频封面地址​
	Title    string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`                        //视频标题
}

func (x *PublishVideoRequest) Reset() {
	*x = PublishVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishVideoRequest) ProtoMessage() {}

func (x *PublishVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishVideoRequest.ProtoReflect.Descriptor instead.
func (*PublishVideoRequest) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{1}
}

func (x *PublishVideoRequest) GetAuthorId() int64 {
	if x != nil {
		return x.AuthorId
	}
	return 0
}

func (x *PublishVideoRequest) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *PublishVideoRequest) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *PublishVideoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type PublishVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishVideoResponse) Reset() {
	*x = PublishVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishVideoResponse) ProtoMessage() {}

func (x *PublishVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishVideoResponse.ProtoReflect.Descriptor instead.
func (*PublishVideoResponse) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{2}
}

type VideoListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`                     // 用户Id
	Cursor   int64 `protobuf:"varint,2,opt,name=cursor,proto3" json:"cursor,omitempty"`                     // 游标
	PageSize int64 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"` // 分页大小
	SortType int32 `protobuf:"varint,4,opt,name=sort_type,json=sortType,proto3" json:"sort_type,omitempty"` // 排序方式
	VideoId  int64 `protobuf:"varint,5,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`    // 视频Id
}

func (x *VideoListRequest) Reset() {
	*x = VideoListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoListRequest) ProtoMessage() {}

func (x *VideoListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoListRequest.ProtoReflect.Descriptor instead.
func (*VideoListRequest) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{3}
}

func (x *VideoListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *VideoListRequest) GetCursor() int64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *VideoListRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *VideoListRequest) GetSortType() int32 {
	if x != nil {
		return x.SortType
	}
	return 0
}

func (x *VideoListRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type VideoListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoList []*VideoItem `protobuf:"bytes,1,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
	IsEnd     bool         `protobuf:"varint,2,opt,name=is_end,json=isEnd,proto3" json:"is_end,omitempty"`
	Cursor    int64        `protobuf:"varint,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	VideoId   int64        `protobuf:"varint,4,opt,name=videoId,proto3" json:"videoId,omitempty"`
}

func (x *VideoListResponse) Reset() {
	*x = VideoListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoListResponse) ProtoMessage() {}

func (x *VideoListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoListResponse.ProtoReflect.Descriptor instead.
func (*VideoListResponse) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{4}
}

func (x *VideoListResponse) GetVideoList() []*VideoItem {
	if x != nil {
		return x.VideoList
	}
	return nil
}

func (x *VideoListResponse) GetIsEnd() bool {
	if x != nil {
		return x.IsEnd
	}
	return false
}

func (x *VideoListResponse) GetCursor() int64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *VideoListResponse) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

var File_video_proto protoreflect.FileDescriptor

var file_video_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x22, 0xf5, 0x01, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a,
	0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x80, 0x01, 0x0a,
	0x13, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22,
	0x16, 0x0a, 0x14, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x97, 0x01, 0x0a, 0x10, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x72,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x6f,
	0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x64, 0x22, 0x8d, 0x01, 0x0a, 0x11, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x65,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x45, 0x6e, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x64, 0x32, 0x93, 0x01, 0x0a, 0x08, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x50, 0x43, 0x12, 0x3e,
	0x0a, 0x09, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47,
	0x0a, 0x0c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x1a,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_proto_rawDescOnce sync.Once
	file_video_proto_rawDescData = file_video_proto_rawDesc
)

func file_video_proto_rawDescGZIP() []byte {
	file_video_proto_rawDescOnce.Do(func() {
		file_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_proto_rawDescData)
	})
	return file_video_proto_rawDescData
}

var file_video_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_video_proto_goTypes = []interface{}{
	(*VideoItem)(nil),            // 0: video.VideoItem
	(*PublishVideoRequest)(nil),  // 1: video.PublishVideoRequest
	(*PublishVideoResponse)(nil), // 2: video.PublishVideoResponse
	(*VideoListRequest)(nil),     // 3: video.VideoListRequest
	(*VideoListResponse)(nil),    // 4: video.VideoListResponse
}
var file_video_proto_depIdxs = []int32{
	0, // 0: video.VideoListResponse.video_list:type_name -> video.VideoItem
	3, // 1: video.VideoRPC.VideoList:input_type -> video.VideoListRequest
	1, // 2: video.VideoRPC.PublishVideo:input_type -> video.PublishVideoRequest
	4, // 3: video.VideoRPC.VideoList:output_type -> video.VideoListResponse
	2, // 4: video.VideoRPC.PublishVideo:output_type -> video.PublishVideoResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_video_proto_init() }
func file_video_proto_init() {
	if File_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoItem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishVideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_proto_goTypes,
		DependencyIndexes: file_video_proto_depIdxs,
		MessageInfos:      file_video_proto_msgTypes,
	}.Build()
	File_video_proto = out.File
	file_video_proto_rawDesc = nil
	file_video_proto_goTypes = nil
	file_video_proto_depIdxs = nil
}
