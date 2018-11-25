// Code generated by protoc-gen-go. DO NOT EDIT.
// source: book.proto

package book

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 请求书详情的参数结构  book_id 32位整形
type BookInfoParams struct {
	BookId               int32    `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookInfoParams) Reset()         { *m = BookInfoParams{} }
func (m *BookInfoParams) String() string { return proto.CompactTextString(m) }
func (*BookInfoParams) ProtoMessage()    {}
func (*BookInfoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{0}
}

func (m *BookInfoParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookInfoParams.Unmarshal(m, b)
}
func (m *BookInfoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookInfoParams.Marshal(b, m, deterministic)
}
func (m *BookInfoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookInfoParams.Merge(m, src)
}
func (m *BookInfoParams) XXX_Size() int {
	return xxx_messageInfo_BookInfoParams.Size(m)
}
func (m *BookInfoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_BookInfoParams.DiscardUnknown(m)
}

var xxx_messageInfo_BookInfoParams proto.InternalMessageInfo

func (m *BookInfoParams) GetBookId() int32 {
	if m != nil {
		return m.BookId
	}
	return 0
}

// 书详情信息的结构   book_name字符串类型
type BookInfo struct {
	BookId               int32    `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	BookName             string   `protobuf:"bytes,2,opt,name=book_name,json=bookName,proto3" json:"book_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookInfo) Reset()         { *m = BookInfo{} }
func (m *BookInfo) String() string { return proto.CompactTextString(m) }
func (*BookInfo) ProtoMessage()    {}
func (*BookInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{1}
}

func (m *BookInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookInfo.Unmarshal(m, b)
}
func (m *BookInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookInfo.Marshal(b, m, deterministic)
}
func (m *BookInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookInfo.Merge(m, src)
}
func (m *BookInfo) XXX_Size() int {
	return xxx_messageInfo_BookInfo.Size(m)
}
func (m *BookInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BookInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BookInfo proto.InternalMessageInfo

func (m *BookInfo) GetBookId() int32 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *BookInfo) GetBookName() string {
	if m != nil {
		return m.BookName
	}
	return ""
}

// 请求书列表的参数结构  page、limit   32位整形
type BookListParams struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookListParams) Reset()         { *m = BookListParams{} }
func (m *BookListParams) String() string { return proto.CompactTextString(m) }
func (*BookListParams) ProtoMessage()    {}
func (*BookListParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{2}
}

func (m *BookListParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookListParams.Unmarshal(m, b)
}
func (m *BookListParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookListParams.Marshal(b, m, deterministic)
}
func (m *BookListParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookListParams.Merge(m, src)
}
func (m *BookListParams) XXX_Size() int {
	return xxx_messageInfo_BookListParams.Size(m)
}
func (m *BookListParams) XXX_DiscardUnknown() {
	xxx_messageInfo_BookListParams.DiscardUnknown(m)
}

var xxx_messageInfo_BookListParams proto.InternalMessageInfo

func (m *BookListParams) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *BookListParams) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// 书列表的结构    BookInfo结构数组
type BookList struct {
	BookList             []*BookInfo `protobuf:"bytes,1,rep,name=book_list,json=bookList,proto3" json:"book_list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *BookList) Reset()         { *m = BookList{} }
func (m *BookList) String() string { return proto.CompactTextString(m) }
func (*BookList) ProtoMessage()    {}
func (*BookList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{3}
}

func (m *BookList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookList.Unmarshal(m, b)
}
func (m *BookList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookList.Marshal(b, m, deterministic)
}
func (m *BookList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookList.Merge(m, src)
}
func (m *BookList) XXX_Size() int {
	return xxx_messageInfo_BookList.Size(m)
}
func (m *BookList) XXX_DiscardUnknown() {
	xxx_messageInfo_BookList.DiscardUnknown(m)
}

var xxx_messageInfo_BookList proto.InternalMessageInfo

func (m *BookList) GetBookList() []*BookInfo {
	if m != nil {
		return m.BookList
	}
	return nil
}

func init() {
	proto.RegisterType((*BookInfoParams)(nil), "book.BookInfoParams")
	proto.RegisterType((*BookInfo)(nil), "book.BookInfo")
	proto.RegisterType((*BookListParams)(nil), "book.BookListParams")
	proto.RegisterType((*BookList)(nil), "book.BookList")
}

func init() { proto.RegisterFile("book.proto", fileDescriptor_1e89d0eaa98dc5d8) }

var fileDescriptor_1e89d0eaa98dc5d8 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xca, 0xcf, 0xcf,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x34, 0xb9, 0xf8, 0x9c, 0xf2,
	0xf3, 0xb3, 0x3d, 0xf3, 0xd2, 0xf2, 0x03, 0x12, 0x8b, 0x12, 0x73, 0x8b, 0x85, 0xc4, 0xb9, 0xd8,
	0x41, 0x32, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x6c, 0x20, 0xae, 0x67,
	0x8a, 0x92, 0x03, 0x17, 0x07, 0x4c, 0x29, 0x4e, 0x45, 0x42, 0xd2, 0x5c, 0x9c, 0x60, 0x89, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x0e, 0x90, 0x80, 0x5f, 0x62, 0x6e,
	0xaa, 0x92, 0x15, 0xc4, 0x32, 0x9f, 0xcc, 0xe2, 0x12, 0xa8, 0x65, 0x42, 0x5c, 0x2c, 0x05, 0x89,
	0xe9, 0xa9, 0x50, 0x43, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0x9c, 0xcc, 0xdc, 0xcc, 0x12, 0xb0,
	0x76, 0xd6, 0x20, 0x08, 0x47, 0xc9, 0x1c, 0x62, 0x3b, 0x48, 0xaf, 0x90, 0x36, 0xd4, 0x92, 0x9c,
	0xcc, 0xe2, 0x12, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x3e, 0x3d, 0xb0, 0xd7, 0x60, 0x0e,
	0x84, 0x58, 0x0a, 0x52, 0x6c, 0x54, 0xcd, 0xc5, 0x0d, 0x12, 0x0d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0x15, 0x32, 0xe5, 0xe2, 0x76, 0x4f, 0x2d, 0x81, 0x7b, 0x44, 0x04, 0x55, 0x1f, 0xc4, 0x59,
	0x52, 0x68, 0xa6, 0x29, 0x31, 0x20, 0x69, 0x03, 0xbb, 0x00, 0x49, 0x1b, 0xc2, 0x37, 0xc8, 0xda,
	0x40, 0xa2, 0x4a, 0x0c, 0x49, 0x6c, 0xe0, 0xb0, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x92,
	0xf4, 0xdd, 0x02, 0x79, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookServiceClient is the client API for BookService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookServiceClient interface {
	GetBookInfo(ctx context.Context, in *BookInfoParams, opts ...grpc.CallOption) (*BookInfo, error)
	GetBookList(ctx context.Context, in *BookListParams, opts ...grpc.CallOption) (*BookList, error)
}

type bookServiceClient struct {
	cc *grpc.ClientConn
}

func NewBookServiceClient(cc *grpc.ClientConn) BookServiceClient {
	return &bookServiceClient{cc}
}

func (c *bookServiceClient) GetBookInfo(ctx context.Context, in *BookInfoParams, opts ...grpc.CallOption) (*BookInfo, error) {
	out := new(BookInfo)
	err := c.cc.Invoke(ctx, "/book.BookService/GetBookInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookServiceClient) GetBookList(ctx context.Context, in *BookListParams, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := c.cc.Invoke(ctx, "/book.BookService/GetBookList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServiceServer is the server API for BookService service.
type BookServiceServer interface {
	GetBookInfo(context.Context, *BookInfoParams) (*BookInfo, error)
	GetBookList(context.Context, *BookListParams) (*BookList, error)
}

func RegisterBookServiceServer(s *grpc.Server, srv BookServiceServer) {
	s.RegisterService(&_BookService_serviceDesc, srv)
}

func _BookService_GetBookInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookInfoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBookInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/GetBookInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBookInfo(ctx, req.(*BookInfoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookService_GetBookList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookListParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServiceServer).GetBookList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.BookService/GetBookList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServiceServer).GetBookList(ctx, req.(*BookListParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "book.BookService",
	HandlerType: (*BookServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookInfo",
			Handler:    _BookService_GetBookInfo_Handler,
		},
		{
			MethodName: "GetBookList",
			Handler:    _BookService_GetBookList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}
