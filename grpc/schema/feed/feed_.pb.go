// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feed/feed_.proto

/*
Package feed is a generated protocol buffer package.

It is generated from these files:
	feed/feed_.proto

It has these top-level messages:
	Post
	Feed
	SocialAccount
*/
package feed

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import text "github.com/21stio/go-playground/grpc/schema/text"
import liketermkind "github.com/21stio/go-playground/grpc/schema/liketermkind"
import commenttermkind "github.com/21stio/go-playground/grpc/schema/commenttermkind"
import sharetermkind "github.com/21stio/go-playground/grpc/schema/sharetermkind"
import media "github.com/21stio/go-playground/grpc/schema/media"
import id "github.com/21stio/go-playground/grpc/schema/id"
import service "github.com/21stio/go-playground/grpc/schema/service"
import feedkind "github.com/21stio/go-playground/grpc/schema/feedkind"
import info "github.com/21stio/go-playground/grpc/schema/info"
import socialaccountkind "github.com/21stio/go-playground/grpc/schema/socialaccountkind"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Post struct {
	Author           *SocialAccount                   `protobuf:"bytes,1,opt,name=author" json:"author,omitempty"`
	Content          *text.Text                       `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	LikeTerm         *liketermkind.LikeTermKind       `protobuf:"varint,3,opt,name=likeTerm,enum=liketermkind.LikeTermKind" json:"likeTerm,omitempty"`
	LikesCount       *int32                           `protobuf:"varint,4,opt,name=likesCount" json:"likesCount,omitempty"`
	LikedBy          []*SocialAccount                 `protobuf:"bytes,5,rep,name=likedBy" json:"likedBy,omitempty"`
	CommentTerm      *commenttermkind.CommentTermKind `protobuf:"varint,6,opt,name=commentTerm,enum=commenttermkind.CommentTermKind" json:"commentTerm,omitempty"`
	CommentsCount    *int32                           `protobuf:"varint,7,opt,name=commentsCount" json:"commentsCount,omitempty"`
	CommentedBy      []*SocialAccount                 `protobuf:"bytes,8,rep,name=commentedBy" json:"commentedBy,omitempty"`
	Comments         []*Post                          `protobuf:"bytes,9,rep,name=comments" json:"comments,omitempty"`
	Commenting       *Post                            `protobuf:"bytes,10,opt,name=commenting" json:"commenting,omitempty"`
	ShareTerm        *sharetermkind.ShareTermKind     `protobuf:"varint,11,opt,name=shareTerm,enum=sharetermkind.ShareTermKind" json:"shareTerm,omitempty"`
	SharesCount      *int32                           `protobuf:"varint,12,opt,name=sharesCount" json:"sharesCount,omitempty"`
	Shared           *Post                            `protobuf:"bytes,13,opt,name=shared" json:"shared,omitempty"`
	SharedBy         []*SocialAccount                 `protobuf:"bytes,14,rep,name=sharedBy" json:"sharedBy,omitempty"`
	Attachments      []*media.Media                   `protobuf:"bytes,15,rep,name=attachments" json:"attachments,omitempty"`
	Hash             *string                          `protobuf:"bytes,16,opt,name=hash" json:"hash,omitempty"`
	Ids              []*id.Id                         `protobuf:"bytes,17,rep,name=ids" json:"ids,omitempty"`
	Meta             *service.TypeMeta                `protobuf:"bytes,18,opt,name=meta" json:"meta,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *Post) Reset()                    { *m = Post{} }
func (m *Post) String() string            { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()               {}
func (*Post) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Post) GetAuthor() *SocialAccount {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *Post) GetContent() *text.Text {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Post) GetLikeTerm() liketermkind.LikeTermKind {
	if m != nil && m.LikeTerm != nil {
		return *m.LikeTerm
	}
	return liketermkind.LikeTermKind_LIKE
}

func (m *Post) GetLikesCount() int32 {
	if m != nil && m.LikesCount != nil {
		return *m.LikesCount
	}
	return 0
}

func (m *Post) GetLikedBy() []*SocialAccount {
	if m != nil {
		return m.LikedBy
	}
	return nil
}

func (m *Post) GetCommentTerm() commenttermkind.CommentTermKind {
	if m != nil && m.CommentTerm != nil {
		return *m.CommentTerm
	}
	return commenttermkind.CommentTermKind_COMMENT
}

func (m *Post) GetCommentsCount() int32 {
	if m != nil && m.CommentsCount != nil {
		return *m.CommentsCount
	}
	return 0
}

func (m *Post) GetCommentedBy() []*SocialAccount {
	if m != nil {
		return m.CommentedBy
	}
	return nil
}

func (m *Post) GetComments() []*Post {
	if m != nil {
		return m.Comments
	}
	return nil
}

func (m *Post) GetCommenting() *Post {
	if m != nil {
		return m.Commenting
	}
	return nil
}

func (m *Post) GetShareTerm() sharetermkind.ShareTermKind {
	if m != nil && m.ShareTerm != nil {
		return *m.ShareTerm
	}
	return sharetermkind.ShareTermKind_SHARE
}

func (m *Post) GetSharesCount() int32 {
	if m != nil && m.SharesCount != nil {
		return *m.SharesCount
	}
	return 0
}

func (m *Post) GetShared() *Post {
	if m != nil {
		return m.Shared
	}
	return nil
}

func (m *Post) GetSharedBy() []*SocialAccount {
	if m != nil {
		return m.SharedBy
	}
	return nil
}

func (m *Post) GetAttachments() []*media.Media {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Post) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func (m *Post) GetIds() []*id.Id {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *Post) GetMeta() *service.TypeMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type Feed struct {
	Kind             *feedkind.FeedKind `protobuf:"varint,1,opt,name=kind,enum=feedkind.FeedKind" json:"kind,omitempty"`
	Info             *info.Info         `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	CreatedBy        *SocialAccount     `protobuf:"bytes,3,opt,name=createdBy" json:"createdBy,omitempty"`
	Count            *int32             `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
	Posts            []*Post            `protobuf:"bytes,5,rep,name=posts" json:"posts,omitempty"`
	Hash             *string            `protobuf:"bytes,6,opt,name=hash" json:"hash,omitempty"`
	Ids              []*id.Id           `protobuf:"bytes,7,rep,name=ids" json:"ids,omitempty"`
	Meta             *service.TypeMeta  `protobuf:"bytes,8,opt,name=meta" json:"meta,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *Feed) Reset()                    { *m = Feed{} }
func (m *Feed) String() string            { return proto.CompactTextString(m) }
func (*Feed) ProtoMessage()               {}
func (*Feed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Feed) GetKind() feedkind.FeedKind {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return feedkind.FeedKind_NEWS
}

func (m *Feed) GetInfo() *info.Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Feed) GetCreatedBy() *SocialAccount {
	if m != nil {
		return m.CreatedBy
	}
	return nil
}

func (m *Feed) GetCount() int32 {
	if m != nil && m.Count != nil {
		return *m.Count
	}
	return 0
}

func (m *Feed) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *Feed) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func (m *Feed) GetIds() []*id.Id {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *Feed) GetMeta() *service.TypeMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type SocialAccount struct {
	Info             *info.Info                           `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Kind             *socialaccountkind.SocialAccountKind `protobuf:"varint,2,opt,name=kind,enum=socialaccountkind.SocialAccountKind" json:"kind,omitempty"`
	FollowersCount   *int32                               `protobuf:"varint,3,opt,name=followersCount" json:"followersCount,omitempty"`
	Followers        []*SocialAccount                     `protobuf:"bytes,4,rep,name=followers" json:"followers,omitempty"`
	FollowingCount   *int32                               `protobuf:"varint,5,opt,name=followingCount" json:"followingCount,omitempty"`
	Following        []*SocialAccount                     `protobuf:"bytes,6,rep,name=following" json:"following,omitempty"`
	FriendsCount     *int32                               `protobuf:"varint,7,opt,name=friendsCount" json:"friendsCount,omitempty"`
	Friends          []*SocialAccount                     `protobuf:"bytes,8,rep,name=friends" json:"friends,omitempty"`
	Feed             *Feed                                `protobuf:"bytes,9,opt,name=feed" json:"feed,omitempty"`
	Posts            []*Post                              `protobuf:"bytes,10,rep,name=posts" json:"posts,omitempty"`
	Hash             *string                              `protobuf:"bytes,11,opt,name=hash" json:"hash,omitempty"`
	Ids              []*id.Id                             `protobuf:"bytes,12,rep,name=ids" json:"ids,omitempty"`
	Meta             *service.TypeMeta                    `protobuf:"bytes,13,opt,name=meta" json:"meta,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *SocialAccount) Reset()                    { *m = SocialAccount{} }
func (m *SocialAccount) String() string            { return proto.CompactTextString(m) }
func (*SocialAccount) ProtoMessage()               {}
func (*SocialAccount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SocialAccount) GetInfo() *info.Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *SocialAccount) GetKind() socialaccountkind.SocialAccountKind {
	if m != nil && m.Kind != nil {
		return *m.Kind
	}
	return socialaccountkind.SocialAccountKind_MASTODON
}

func (m *SocialAccount) GetFollowersCount() int32 {
	if m != nil && m.FollowersCount != nil {
		return *m.FollowersCount
	}
	return 0
}

func (m *SocialAccount) GetFollowers() []*SocialAccount {
	if m != nil {
		return m.Followers
	}
	return nil
}

func (m *SocialAccount) GetFollowingCount() int32 {
	if m != nil && m.FollowingCount != nil {
		return *m.FollowingCount
	}
	return 0
}

func (m *SocialAccount) GetFollowing() []*SocialAccount {
	if m != nil {
		return m.Following
	}
	return nil
}

func (m *SocialAccount) GetFriendsCount() int32 {
	if m != nil && m.FriendsCount != nil {
		return *m.FriendsCount
	}
	return 0
}

func (m *SocialAccount) GetFriends() []*SocialAccount {
	if m != nil {
		return m.Friends
	}
	return nil
}

func (m *SocialAccount) GetFeed() *Feed {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *SocialAccount) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

func (m *SocialAccount) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func (m *SocialAccount) GetIds() []*id.Id {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *SocialAccount) GetMeta() *service.TypeMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*Post)(nil), "feed.Post")
	proto.RegisterType((*Feed)(nil), "feed.Feed")
	proto.RegisterType((*SocialAccount)(nil), "feed.SocialAccount")
}

func init() { proto.RegisterFile("feed/feed_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 745 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdf, 0x6b, 0xdb, 0x48,
	0x10, 0xc6, 0xb1, 0xfc, 0x6b, 0xfc, 0xe3, 0x9c, 0xbd, 0xe3, 0x58, 0xc2, 0x11, 0x84, 0xc9, 0x19,
	0x93, 0x23, 0xab, 0x8b, 0xa1, 0xa5, 0xf4, 0xad, 0x0e, 0x14, 0x42, 0x1b, 0x28, 0x4a, 0x9e, 0xfa,
	0x52, 0x54, 0x69, 0x6d, 0x2f, 0xb1, 0xb5, 0x46, 0x5a, 0xb7, 0xf1, 0x9f, 0xd0, 0x87, 0xf6, 0x6f,
	0x2e, 0x3b, 0xda, 0x95, 0x64, 0x1b, 0x25, 0x2f, 0xeb, 0x9d, 0x6f, 0xbe, 0x1d, 0xcd, 0x7c, 0x3b,
	0xb3, 0x86, 0xe1, 0x9c, 0xf3, 0xc8, 0xd3, 0xcb, 0x17, 0xb6, 0x49, 0xa4, 0x92, 0xc4, 0xd1, 0xc6,
	0xd9, 0x50, 0xf1, 0x27, 0xe5, 0xe9, 0xc5, 0xe0, 0x67, 0xee, 0x4a, 0x3c, 0x72, 0xc5, 0x93, 0xf5,
	0xa3, 0x88, 0x23, 0xaf, 0x6c, 0x58, 0xc6, 0x38, 0x94, 0xeb, 0x35, 0x8f, 0x55, 0x4e, 0x3a, 0xb0,
	0x2d, 0x6f, 0x94, 0x2e, 0x83, 0xa4, 0x08, 0xb5, 0x67, 0x59, 0x0e, 0x59, 0xf3, 0x48, 0x04, 0x1e,
	0xae, 0x16, 0xeb, 0x89, 0xc8, 0x13, 0x39, 0xe3, 0xef, 0x94, 0x27, 0xdf, 0x44, 0xc8, 0x3d, 0xf3,
	0x6b, 0x71, 0xaa, 0xf3, 0xc7, 0xc0, 0x76, 0x63, 0x3d, 0x43, 0x11, 0xcf, 0xa5, 0xa7, 0x17, 0x8b,
	0x5c, 0xa6, 0x32, 0x14, 0xc1, 0x2a, 0x08, 0x43, 0xb9, 0x8d, 0x55, 0x96, 0xcd, 0x21, 0x62, 0xb8,
	0xa3, 0x9f, 0x4d, 0x70, 0x3e, 0xc9, 0x54, 0x91, 0xff, 0xa0, 0x19, 0x6c, 0xd5, 0x52, 0x26, 0xb4,
	0xe6, 0xd6, 0x26, 0xdd, 0xe9, 0x9f, 0x4c, 0x7f, 0x88, 0xdd, 0xe3, 0xc1, 0x77, 0xd9, 0x41, 0xdf,
	0x50, 0xc8, 0x05, 0xb4, 0x42, 0x19, 0x2b, 0x1e, 0x2b, 0x7a, 0x82, 0x6c, 0x60, 0x5a, 0x54, 0xf6,
	0xc0, 0x9f, 0x94, 0x6f, 0x5d, 0xe4, 0x35, 0xb4, 0xb5, 0xa0, 0x0f, 0x3c, 0x59, 0xd3, 0xba, 0x5b,
	0x9b, 0x0c, 0xa6, 0x67, 0xac, 0xac, 0x30, 0xfb, 0x68, 0xbc, 0x1f, 0x44, 0x1c, 0xf9, 0x39, 0x97,
	0x9c, 0x03, 0xe8, 0x7d, 0x7a, 0xa3, 0xbf, 0x49, 0x1d, 0xb7, 0x36, 0x69, 0xf8, 0x25, 0x84, 0x5c,
	0x41, 0x4b, 0x5b, 0xd1, 0x6c, 0x47, 0x1b, 0x6e, 0xbd, 0x2a, 0x57, 0xcb, 0x21, 0x33, 0xe8, 0x9a,
	0x2b, 0xc3, 0x4c, 0x9a, 0x98, 0x89, 0xcb, 0x0e, 0xae, 0x91, 0xdd, 0x14, 0x1c, 0xcc, 0xa7, 0x7c,
	0x88, 0x5c, 0x40, 0xdf, 0x98, 0x26, 0xab, 0x16, 0x66, 0xb5, 0x0f, 0x92, 0x57, 0xf9, 0x97, 0x30,
	0xb9, 0x76, 0x75, 0x72, 0x65, 0x1e, 0x19, 0x43, 0xdb, 0xc6, 0xa1, 0x1d, 0x3c, 0x03, 0xd9, 0x19,
	0x7d, 0x31, 0x7e, 0xee, 0x23, 0x97, 0x00, 0x66, 0x2f, 0xe2, 0x05, 0x05, 0x23, 0x7c, 0xc1, 0x2c,
	0x79, 0xc9, 0x5b, 0xe8, 0x60, 0x07, 0x62, 0xc9, 0x5d, 0x2c, 0xf9, 0x1f, 0xb6, 0xd7, 0x93, 0xec,
	0xde, 0xfa, 0xb1, 0xdc, 0x82, 0x4e, 0x5c, 0xe8, 0xa2, 0x61, 0x4a, 0xed, 0x61, 0xa9, 0x65, 0x88,
	0x8c, 0xa0, 0x89, 0x66, 0x44, 0xfb, 0x47, 0x59, 0x18, 0x0f, 0xf1, 0xa0, 0x9d, 0xed, 0x66, 0x3b,
	0x3a, 0xa8, 0x56, 0x22, 0x27, 0x11, 0x06, 0xdd, 0x40, 0xa9, 0x20, 0x5c, 0x66, 0x4a, 0xfc, 0x81,
	0x67, 0x7a, 0x0c, 0x87, 0x85, 0xdd, 0xe9, 0xd5, 0x2f, 0x13, 0x08, 0x01, 0x67, 0x19, 0xa4, 0x4b,
	0x3a, 0x74, 0x6b, 0x93, 0x8e, 0x8f, 0x7b, 0x42, 0xa1, 0x2e, 0xa2, 0x94, 0x9e, 0xe2, 0xd9, 0x26,
	0x13, 0x11, 0xbb, 0x8d, 0x7c, 0x0d, 0x91, 0x7f, 0xc1, 0x59, 0x73, 0x15, 0x50, 0x82, 0x09, 0x9f,
	0x32, 0x33, 0x5f, 0xec, 0x61, 0xb7, 0xe1, 0x77, 0x5c, 0x05, 0x3e, 0xba, 0x47, 0xbf, 0x4e, 0xc0,
	0x79, 0xcf, 0x79, 0x44, 0xc6, 0xe0, 0x68, 0x95, 0x70, 0x1a, 0x06, 0x53, 0xc2, 0xec, 0xd8, 0x31,
	0xed, 0x45, 0xc5, 0xd0, 0x4f, 0xce, 0xc1, 0xd1, 0xb3, 0x97, 0xcf, 0x81, 0x36, 0xd8, 0x6d, 0x3c,
	0x97, 0x3e, 0xe2, 0xe4, 0x1a, 0x3a, 0x61, 0xc2, 0x83, 0xac, 0x23, 0xea, 0xd5, 0xa3, 0x55, 0xb0,
	0xc8, 0x5f, 0xd0, 0x08, 0x4b, 0xad, 0x9f, 0x19, 0xc4, 0x85, 0xc6, 0x46, 0xa6, 0x2a, 0x35, 0x3d,
	0x5f, 0x96, 0x3c, 0x73, 0xe4, 0x82, 0x34, 0x8f, 0x05, 0x69, 0x55, 0x0b, 0xd2, 0x7e, 0x5e, 0x90,
	0x1f, 0x0e, 0xf4, 0xf7, 0x32, 0xcd, 0x2b, 0xae, 0x55, 0x54, 0xfc, 0xc6, 0x28, 0x77, 0x82, 0xca,
	0x5d, 0xb0, 0xa3, 0xb7, 0x67, 0xbf, 0xf2, 0x92, 0x96, 0x63, 0x18, 0xcc, 0xe5, 0x6a, 0x25, 0xbf,
	0xf3, 0xc4, 0xf4, 0x5e, 0x1d, 0x15, 0x38, 0x40, 0xb5, 0xa6, 0x39, 0x42, 0x9d, 0xea, 0xde, 0x2a,
	0x58, 0x45, 0x68, 0x11, 0x2f, 0xb2, 0xd0, 0x8d, 0x72, 0x68, 0x8b, 0x16, 0xa1, 0xf5, 0x88, 0x35,
	0x5f, 0x0c, 0xad, 0x47, 0x6d, 0x04, 0xbd, 0x79, 0x22, 0x78, 0x1c, 0xed, 0x3d, 0x0d, 0x7b, 0x98,
	0x7e, 0xb2, 0x8c, 0xfd, 0xdc, 0xab, 0x60, 0x39, 0x5a, 0x62, 0xed, 0xa6, 0x9d, 0xf2, 0x74, 0xe9,
	0xc6, 0xf3, 0x11, 0x2f, 0x7a, 0x01, 0x5e, 0xea, 0x85, 0xee, 0x71, 0x2f, 0xf4, 0xaa, 0x7b, 0xa1,
	0xff, 0x6c, 0x2f, 0xcc, 0xa6, 0x9f, 0xff, 0x5f, 0x08, 0xb5, 0xdc, 0x7e, 0xd5, 0x8f, 0xa7, 0x37,
	0xbd, 0x4e, 0x95, 0x90, 0xde, 0x42, 0x5e, 0x6d, 0x56, 0xc1, 0x6e, 0x91, 0xc8, 0x6d, 0x1c, 0x79,
	0x8b, 0x64, 0x13, 0x7a, 0x69, 0xb8, 0xe4, 0xeb, 0x00, 0xff, 0xa9, 0x7e, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x11, 0x29, 0xa7, 0x6a, 0x8b, 0x07, 0x00, 0x00,
}
