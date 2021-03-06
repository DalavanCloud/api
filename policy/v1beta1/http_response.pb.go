// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: policy/v1beta1/http_response.proto

package v1beta1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strconv "strconv"

import strings "strings"
import reflect "reflect"
import sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// HTTP response codes.
// For more details: http://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
type HttpStatusCode int32

const (
	// Empty - This code not part of the HTTP status code specification, but it is needed for proto
	// `enum` type.
	Empty                         HttpStatusCode = 0
	Continue                      HttpStatusCode = 100
	OK                            HttpStatusCode = 200
	Created                       HttpStatusCode = 201
	Accepted                      HttpStatusCode = 202
	NonAuthoritativeInformation   HttpStatusCode = 203
	NoContent                     HttpStatusCode = 204
	ResetContent                  HttpStatusCode = 205
	PartialContent                HttpStatusCode = 206
	MultiStatus                   HttpStatusCode = 207
	AlreadyReported               HttpStatusCode = 208
	IMUsed                        HttpStatusCode = 226
	MultipleChoices               HttpStatusCode = 300
	MovedPermanently              HttpStatusCode = 301
	Found                         HttpStatusCode = 302
	SeeOther                      HttpStatusCode = 303
	NotModified                   HttpStatusCode = 304
	UseProxy                      HttpStatusCode = 305
	TemporaryRedirect             HttpStatusCode = 307
	PermanentRedirect             HttpStatusCode = 308
	BadRequest                    HttpStatusCode = 400
	Unauthorized                  HttpStatusCode = 401
	PaymentRequired               HttpStatusCode = 402
	Forbidden                     HttpStatusCode = 403
	NotFound                      HttpStatusCode = 404
	MethodNotAllowed              HttpStatusCode = 405
	NotAcceptable                 HttpStatusCode = 406
	ProxyAuthenticationRequired   HttpStatusCode = 407
	RequestTimeout                HttpStatusCode = 408
	Conflict                      HttpStatusCode = 409
	Gone                          HttpStatusCode = 410
	LengthRequired                HttpStatusCode = 411
	PreconditionFailed            HttpStatusCode = 412
	PayloadTooLarge               HttpStatusCode = 413
	URITooLong                    HttpStatusCode = 414
	UnsupportedMediaType          HttpStatusCode = 415
	RangeNotSatisfiable           HttpStatusCode = 416
	ExpectationFailed             HttpStatusCode = 417
	MisdirectedRequest            HttpStatusCode = 421
	UnprocessableEntity           HttpStatusCode = 422
	Locked                        HttpStatusCode = 423
	FailedDependency              HttpStatusCode = 424
	UpgradeRequired               HttpStatusCode = 426
	PreconditionRequired          HttpStatusCode = 428
	TooManyRequests               HttpStatusCode = 429
	RequestHeaderFieldsTooLarge   HttpStatusCode = 431
	InternalServerError           HttpStatusCode = 500
	NotImplemented                HttpStatusCode = 501
	BadGateway                    HttpStatusCode = 502
	ServiceUnavailable            HttpStatusCode = 503
	GatewayTimeout                HttpStatusCode = 504
	HTTPVersionNotSupported       HttpStatusCode = 505
	VariantAlsoNegotiates         HttpStatusCode = 506
	InsufficientStorage           HttpStatusCode = 507
	LoopDetected                  HttpStatusCode = 508
	NotExtended                   HttpStatusCode = 510
	NetworkAuthenticationRequired HttpStatusCode = 511
)

var HttpStatusCode_name = map[int32]string{
	0:   "Empty",
	100: "Continue",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "NonAuthoritativeInformation",
	204: "NoContent",
	205: "ResetContent",
	206: "PartialContent",
	207: "MultiStatus",
	208: "AlreadyReported",
	226: "IMUsed",
	300: "MultipleChoices",
	301: "MovedPermanently",
	302: "Found",
	303: "SeeOther",
	304: "NotModified",
	305: "UseProxy",
	307: "TemporaryRedirect",
	308: "PermanentRedirect",
	400: "BadRequest",
	401: "Unauthorized",
	402: "PaymentRequired",
	403: "Forbidden",
	404: "NotFound",
	405: "MethodNotAllowed",
	406: "NotAcceptable",
	407: "ProxyAuthenticationRequired",
	408: "RequestTimeout",
	409: "Conflict",
	410: "Gone",
	411: "LengthRequired",
	412: "PreconditionFailed",
	413: "PayloadTooLarge",
	414: "URITooLong",
	415: "UnsupportedMediaType",
	416: "RangeNotSatisfiable",
	417: "ExpectationFailed",
	421: "MisdirectedRequest",
	422: "UnprocessableEntity",
	423: "Locked",
	424: "FailedDependency",
	426: "UpgradeRequired",
	428: "PreconditionRequired",
	429: "TooManyRequests",
	431: "RequestHeaderFieldsTooLarge",
	500: "InternalServerError",
	501: "NotImplemented",
	502: "BadGateway",
	503: "ServiceUnavailable",
	504: "GatewayTimeout",
	505: "HTTPVersionNotSupported",
	506: "VariantAlsoNegotiates",
	507: "InsufficientStorage",
	508: "LoopDetected",
	510: "NotExtended",
	511: "NetworkAuthenticationRequired",
}
var HttpStatusCode_value = map[string]int32{
	"Empty":                         0,
	"Continue":                      100,
	"OK":                            200,
	"Created":                       201,
	"Accepted":                      202,
	"NonAuthoritativeInformation":   203,
	"NoContent":                     204,
	"ResetContent":                  205,
	"PartialContent":                206,
	"MultiStatus":                   207,
	"AlreadyReported":               208,
	"IMUsed":                        226,
	"MultipleChoices":               300,
	"MovedPermanently":              301,
	"Found":                         302,
	"SeeOther":                      303,
	"NotModified":                   304,
	"UseProxy":                      305,
	"TemporaryRedirect":             307,
	"PermanentRedirect":             308,
	"BadRequest":                    400,
	"Unauthorized":                  401,
	"PaymentRequired":               402,
	"Forbidden":                     403,
	"NotFound":                      404,
	"MethodNotAllowed":              405,
	"NotAcceptable":                 406,
	"ProxyAuthenticationRequired":   407,
	"RequestTimeout":                408,
	"Conflict":                      409,
	"Gone":                          410,
	"LengthRequired":                411,
	"PreconditionFailed":            412,
	"PayloadTooLarge":               413,
	"URITooLong":                    414,
	"UnsupportedMediaType":          415,
	"RangeNotSatisfiable":           416,
	"ExpectationFailed":             417,
	"MisdirectedRequest":            421,
	"UnprocessableEntity":           422,
	"Locked":                        423,
	"FailedDependency":              424,
	"UpgradeRequired":               426,
	"PreconditionRequired":          428,
	"TooManyRequests":               429,
	"RequestHeaderFieldsTooLarge":   431,
	"InternalServerError":           500,
	"NotImplemented":                501,
	"BadGateway":                    502,
	"ServiceUnavailable":            503,
	"GatewayTimeout":                504,
	"HTTPVersionNotSupported":       505,
	"VariantAlsoNegotiates":         506,
	"InsufficientStorage":           507,
	"LoopDetected":                  508,
	"NotExtended":                   510,
	"NetworkAuthenticationRequired": 511,
}

func (HttpStatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptorHttpResponse, []int{0} }

// Direct HTTP response for a client-facing error message which can be attached
// to an RPC error.
type DirectHttpResponse struct {
	// Optional HTTP status code. If not set, RPC error code is used.
	Code HttpStatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=istio.policy.v1beta1.HttpStatusCode" json:"code,omitempty"`
	// HTTP response body.
	Body string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	// Optional HTTP response headers.
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *DirectHttpResponse) Reset()                    { *m = DirectHttpResponse{} }
func (*DirectHttpResponse) ProtoMessage()               {}
func (*DirectHttpResponse) Descriptor() ([]byte, []int) { return fileDescriptorHttpResponse, []int{0} }

func (m *DirectHttpResponse) GetCode() HttpStatusCode {
	if m != nil {
		return m.Code
	}
	return Empty
}

func (m *DirectHttpResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *DirectHttpResponse) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*DirectHttpResponse)(nil), "istio.policy.v1beta1.DirectHttpResponse")
	proto.RegisterEnum("istio.policy.v1beta1.HttpStatusCode", HttpStatusCode_name, HttpStatusCode_value)
}
func (x HttpStatusCode) String() string {
	s, ok := HttpStatusCode_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (this *DirectHttpResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DirectHttpResponse)
	if !ok {
		that2, ok := that.(DirectHttpResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Code != that1.Code {
		return false
	}
	if this.Body != that1.Body {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if this.Headers[i] != that1.Headers[i] {
			return false
		}
	}
	return true
}
func (this *DirectHttpResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&v1beta1.DirectHttpResponse{")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	s = append(s, "Body: "+fmt.Sprintf("%#v", this.Body)+",\n")
	keysForHeaders := make([]string, 0, len(this.Headers))
	for k, _ := range this.Headers {
		keysForHeaders = append(keysForHeaders, k)
	}
	sortkeys.Strings(keysForHeaders)
	mapStringForHeaders := "map[string]string{"
	for _, k := range keysForHeaders {
		mapStringForHeaders += fmt.Sprintf("%#v: %#v,", k, this.Headers[k])
	}
	mapStringForHeaders += "}"
	if this.Headers != nil {
		s = append(s, "Headers: "+mapStringForHeaders+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringHttpResponse(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *DirectHttpResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DirectHttpResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHttpResponse(dAtA, i, uint64(m.Code))
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHttpResponse(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	if len(m.Headers) > 0 {
		for k, _ := range m.Headers {
			dAtA[i] = 0x1a
			i++
			v := m.Headers[k]
			mapSize := 1 + len(k) + sovHttpResponse(uint64(len(k))) + 1 + len(v) + sovHttpResponse(uint64(len(v)))
			i = encodeVarintHttpResponse(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintHttpResponse(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintHttpResponse(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeVarintHttpResponse(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DirectHttpResponse) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovHttpResponse(uint64(m.Code))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovHttpResponse(uint64(l))
	}
	if len(m.Headers) > 0 {
		for k, v := range m.Headers {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovHttpResponse(uint64(len(k))) + 1 + len(v) + sovHttpResponse(uint64(len(v)))
			n += mapEntrySize + 1 + sovHttpResponse(uint64(mapEntrySize))
		}
	}
	return n
}

func sovHttpResponse(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHttpResponse(x uint64) (n int) {
	return sovHttpResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DirectHttpResponse) String() string {
	if this == nil {
		return "nil"
	}
	keysForHeaders := make([]string, 0, len(this.Headers))
	for k, _ := range this.Headers {
		keysForHeaders = append(keysForHeaders, k)
	}
	sortkeys.Strings(keysForHeaders)
	mapStringForHeaders := "map[string]string{"
	for _, k := range keysForHeaders {
		mapStringForHeaders += fmt.Sprintf("%v: %v,", k, this.Headers[k])
	}
	mapStringForHeaders += "}"
	s := strings.Join([]string{`&DirectHttpResponse{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Body:` + fmt.Sprintf("%v", this.Body) + `,`,
		`Headers:` + mapStringForHeaders + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringHttpResponse(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DirectHttpResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpResponse
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DirectHttpResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DirectHttpResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (HttpStatusCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHttpResponse
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpResponse
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHttpResponse
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Headers == nil {
				m.Headers = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowHttpResponse
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHttpResponse
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthHttpResponse
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowHttpResponse
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthHttpResponse
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipHttpResponse(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthHttpResponse
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Headers[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttpResponse(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpResponse
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHttpResponse(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHttpResponse
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHttpResponse
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHttpResponse
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHttpResponse
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHttpResponse
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHttpResponse(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHttpResponse = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHttpResponse   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("policy/v1beta1/http_response.proto", fileDescriptorHttpResponse) }

var fileDescriptorHttpResponse = []byte{
	// 1032 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0x4b, 0x8f, 0x1b, 0x45,
	0x10, 0xde, 0x71, 0xe7, 0xb5, 0x9d, 0xcd, 0xa6, 0xd3, 0xd9, 0x90, 0x25, 0x01, 0x6b, 0x15, 0x71,
	0x88, 0x38, 0x78, 0x95, 0x20, 0xa4, 0x28, 0xb7, 0xcd, 0xc6, 0x9b, 0x58, 0xac, 0x1d, 0xcb, 0x6b,
	0xe7, 0xc0, 0x05, 0xf5, 0x4e, 0x97, 0xed, 0x56, 0xc6, 0x5d, 0x93, 0x9e, 0xb2, 0x93, 0xe1, 0xc4,
	0x4f, 0x08, 0xef, 0xf7, 0xeb, 0x00, 0x44, 0x28, 0x21, 0x20, 0xb8, 0xf0, 0x0f, 0xc2, 0x3b, 0x47,
	0x8e, 0xc4, 0x5c, 0xb8, 0x91, 0x03, 0x6f, 0x10, 0xa0, 0x6e, 0x3f, 0xb4, 0x11, 0xe1, 0x36, 0x5d,
	0x53, 0x5f, 0xd5, 0x57, 0xdf, 0x57, 0x33, 0xcd, 0x8f, 0xa4, 0x98, 0x98, 0x38, 0x5f, 0x1e, 0x1c,
	0xdb, 0x04, 0x52, 0xc7, 0x96, 0xbb, 0x44, 0xe9, 0x13, 0x0e, 0xb2, 0x14, 0x6d, 0x06, 0xa5, 0xd4,
	0x21, 0xa1, 0x5c, 0x30, 0x19, 0x19, 0x2c, 0x8d, 0x32, 0x4b, 0xe3, 0xcc, 0x23, 0x3f, 0x46, 0x5c,
	0x9e, 0x36, 0x0e, 0x62, 0x3a, 0x4b, 0x94, 0x36, 0xc6, 0x10, 0x79, 0x82, 0x6f, 0x8b, 0x51, 0xc3,
	0x62, 0xb4, 0x14, 0x1d, 0x9d, 0x3f, 0xfe, 0x50, 0xe9, 0x5e, 0xd8, 0x92, 0x47, 0x6c, 0x90, 0xa2,
	0x7e, 0xb6, 0x8a, 0x1a, 0x1a, 0x01, 0x21, 0x25, 0xdf, 0xb6, 0x89, 0x3a, 0x5f, 0x2c, 0x2c, 0x45,
	0x47, 0x67, 0x1b, 0xe1, 0x59, 0x9e, 0xe3, 0x3b, 0xbb, 0xa0, 0x34, 0xb8, 0x6c, 0x91, 0x2d, 0xb1,
	0xa3, 0xbb, 0x8f, 0x3f, 0x7a, 0xef, 0x82, 0xff, 0x25, 0x52, 0x3a, 0x3b, 0xc2, 0x95, 0x2d, 0xb9,
	0xbc, 0x31, 0xa9, 0x72, 0xe8, 0x24, 0x9f, 0xdb, 0xfa, 0x42, 0x0a, 0xce, 0x2e, 0x40, 0x1e, 0xd8,
	0xce, 0x36, 0xfc, 0xa3, 0x5c, 0xe0, 0xdb, 0x07, 0x2a, 0xe9, 0xc3, 0x98, 0xc7, 0xe8, 0x70, 0xb2,
	0x70, 0x22, 0x7a, 0xf8, 0xd3, 0x59, 0x3e, 0x7f, 0x37, 0x73, 0x39, 0xcb, 0xb7, 0x97, 0x7b, 0x29,
	0xe5, 0x62, 0x46, 0xce, 0xf1, 0x5d, 0xab, 0x68, 0xc9, 0xd8, 0x3e, 0x08, 0x2d, 0x77, 0xf2, 0xc2,
	0xb9, 0xc7, 0xc4, 0xcd, 0x48, 0xce, 0xf1, 0x9d, 0xab, 0x0e, 0x14, 0x81, 0x16, 0x9f, 0x45, 0x72,
	0x0f, 0xdf, 0xb5, 0x12, 0xc7, 0x90, 0xfa, 0xe3, 0xe7, 0x91, 0x5c, 0xe2, 0x87, 0x6b, 0x68, 0x57,
	0xfa, 0xd4, 0x45, 0x67, 0x48, 0x91, 0x19, 0x40, 0xc5, 0xb6, 0xd1, 0xf5, 0x14, 0x19, 0xb4, 0xe2,
	0x8b, 0x48, 0xce, 0xf3, 0xd9, 0x1a, 0xfa, 0xba, 0x60, 0x49, 0x7c, 0x19, 0xc9, 0x7d, 0x7c, 0xae,
	0x01, 0x19, 0xd0, 0x24, 0xf4, 0x55, 0x24, 0xf7, 0xf3, 0xf9, 0xba, 0x72, 0x64, 0x54, 0x32, 0x09,
	0x7e, 0x1d, 0x49, 0xc1, 0x77, 0x57, 0xfb, 0x09, 0x99, 0x11, 0x57, 0xf1, 0x4d, 0x24, 0x17, 0xf8,
	0xde, 0x95, 0xc4, 0x81, 0xd2, 0x79, 0x03, 0x52, 0x74, 0x9e, 0xc1, 0xad, 0x48, 0xee, 0xe6, 0x3b,
	0x2a, 0xd5, 0x56, 0x06, 0x5a, 0x0c, 0x43, 0x4a, 0x00, 0xa5, 0x09, 0xac, 0x76, 0xd1, 0xc4, 0x90,
	0x89, 0x6b, 0x05, 0x79, 0x80, 0x8b, 0x2a, 0x0e, 0x40, 0xd7, 0xc1, 0xf5, 0x94, 0x05, 0x4b, 0x49,
	0x2e, 0xae, 0x17, 0x24, 0xe7, 0xdb, 0xd7, 0xb0, 0x6f, 0xb5, 0xf8, 0xa0, 0xe0, 0xc7, 0xda, 0x00,
	0x38, 0x47, 0x5d, 0x70, 0xe2, 0x46, 0xc1, 0x37, 0xaf, 0x21, 0x55, 0x51, 0x9b, 0xb6, 0x01, 0x2d,
	0x3e, 0x0c, 0x09, 0xad, 0x0c, 0xea, 0x0e, 0x2f, 0xe7, 0xe2, 0xa3, 0x82, 0xbc, 0x8f, 0xef, 0x6b,
	0x42, 0x2f, 0x45, 0xa7, 0x5c, 0xde, 0x00, 0x1d, 0xcc, 0x13, 0x1f, 0x87, 0xf8, 0xb4, 0xcb, 0x34,
	0xfe, 0x49, 0x41, 0xee, 0xe5, 0xfc, 0x94, 0xd2, 0x0d, 0xb8, 0xd8, 0x87, 0x8c, 0xc4, 0x15, 0xe6,
	0x65, 0x68, 0x59, 0x35, 0xd2, 0xed, 0x49, 0xd0, 0xe2, 0x69, 0xe6, 0xc9, 0xd7, 0x55, 0xde, 0x0b,
	0xc8, 0x8b, 0x7d, 0xe3, 0x40, 0x8b, 0x67, 0x98, 0xd7, 0x6f, 0x0d, 0xdd, 0xa6, 0xd1, 0x1a, 0xac,
	0x78, 0x96, 0x79, 0x22, 0x35, 0xa4, 0x11, 0xf1, 0xe7, 0x58, 0x98, 0x0d, 0xa8, 0x8b, 0xba, 0x86,
	0xb4, 0x92, 0x24, 0x78, 0x09, 0xb4, 0x78, 0x9e, 0x49, 0xc9, 0xf7, 0xf8, 0x40, 0x70, 0x4a, 0x6d,
	0x26, 0x20, 0x5e, 0x60, 0xde, 0xab, 0xc0, 0xdf, 0xbb, 0x05, 0x96, 0x4c, 0x1c, 0x3c, 0x9a, 0xf6,
	0x7a, 0x91, 0x79, 0x23, 0xc6, 0x14, 0x9b, 0xa6, 0x07, 0xd8, 0x27, 0xf1, 0x52, 0x68, 0xb8, 0x8a,
	0xb6, 0x9d, 0x98, 0x98, 0xc4, 0xcb, 0x4c, 0xce, 0xf2, 0x6d, 0x67, 0xd0, 0x82, 0x78, 0x25, 0xa4,
	0xaf, 0x83, 0xed, 0x50, 0x77, 0x5a, 0xe3, 0x55, 0x26, 0x0f, 0x72, 0x59, 0x77, 0x10, 0xa3, 0xd5,
	0xc6, 0x97, 0x5f, 0x53, 0x26, 0x01, 0x2d, 0x5e, 0x9b, 0x8c, 0x97, 0xa0, 0xd2, 0x4d, 0xc4, 0x75,
	0xe5, 0x3a, 0x20, 0x5e, 0x67, 0x5e, 0x98, 0x56, 0xa3, 0xe2, 0x23, 0x68, 0x3b, 0xe2, 0x0d, 0x26,
	0xef, 0xe7, 0x0b, 0x2d, 0x9b, 0xf5, 0xd3, 0x91, 0xc3, 0x55, 0xd0, 0x46, 0x35, 0xf3, 0x14, 0xc4,
	0x9b, 0x4c, 0x2e, 0xf2, 0xfd, 0x0d, 0x65, 0x3b, 0x50, 0x43, 0xda, 0x50, 0x64, 0xb2, 0xb6, 0x09,
	0xa3, 0xbd, 0xc5, 0xbc, 0xec, 0xe5, 0xcb, 0x29, 0xc4, 0xa4, 0xb6, 0xf4, 0x7c, 0x3b, 0x90, 0xa9,
	0x9a, 0x6c, 0x64, 0x03, 0x4c, 0xe5, 0x7f, 0x27, 0x94, 0x6a, 0xd9, 0xd4, 0x61, 0x0c, 0x59, 0xe6,
	0x8b, 0x94, 0x2d, 0x19, 0xca, 0xc5, 0xbb, 0xcc, 0xef, 0xd3, 0x3a, 0xc6, 0x17, 0x40, 0x8b, 0xf7,
	0x82, 0xba, 0xa3, 0x62, 0xa7, 0x21, 0x05, 0xab, 0xc1, 0xc6, 0xb9, 0xb8, 0x1a, 0x46, 0x69, 0xa5,
	0x1d, 0xa7, 0x34, 0x4c, 0x27, 0x7f, 0x3f, 0x30, 0xdf, 0x3a, 0xf9, 0xf4, 0xd5, 0xb5, 0x00, 0x68,
	0x22, 0x56, 0x95, 0xcd, 0xc7, 0x1c, 0x32, 0x71, 0x3d, 0x18, 0x32, 0x3e, 0x8e, 0xbe, 0xe8, 0x35,
	0x03, 0x89, 0xce, 0xa6, 0xea, 0xdc, 0x08, 0x34, 0x2b, 0x96, 0xc0, 0x59, 0x95, 0x6c, 0x80, 0x1b,
	0x80, 0x2b, 0x3b, 0x87, 0x4e, 0xfc, 0x14, 0xb4, 0xaf, 0x21, 0x55, 0x7a, 0x69, 0x02, 0x7e, 0x63,
	0x40, 0x8b, 0x9f, 0xd9, 0x78, 0xcb, 0xce, 0x28, 0x82, 0x4b, 0x2a, 0x17, 0xbf, 0x84, 0xf9, 0x3d,
	0xce, 0xc4, 0xd0, 0xb2, 0x6a, 0xa0, 0x4c, 0x12, 0x04, 0xfb, 0x35, 0xc0, 0xc7, 0x69, 0x13, 0xa7,
	0x7f, 0x63, 0xf2, 0x01, 0x7e, 0xf0, 0x6c, 0xb3, 0x59, 0x3f, 0x0f, 0x2e, 0x33, 0x68, 0xbd, 0xca,
	0x13, 0x1b, 0xc4, 0xef, 0x4c, 0x1e, 0xe2, 0x07, 0xce, 0x2b, 0x67, 0x94, 0xa5, 0x95, 0x24, 0xc3,
	0x1a, 0x74, 0x90, 0x8c, 0x22, 0xc8, 0xc4, 0x1f, 0x63, 0x9e, 0x59, 0xbf, 0xdd, 0x36, 0xb1, 0x01,
	0x4b, 0x1b, 0x84, 0x4e, 0x75, 0x40, 0xfc, 0x19, 0xf6, 0x7c, 0x1d, 0x31, 0x3d, 0x0d, 0x14, 0x2c,
	0x10, 0x7f, 0xb1, 0xf1, 0xc7, 0x55, 0xbe, 0x4c, 0x5e, 0x51, 0x2d, 0xfe, 0x66, 0xf2, 0x08, 0x7f,
	0xb0, 0x06, 0x74, 0x09, 0xdd, 0x85, 0xff, 0xd9, 0xcd, 0x7f, 0xd8, 0xa9, 0xca, 0xad, 0xdb, 0xc5,
	0x99, 0x6f, 0x6f, 0x17, 0x67, 0xee, 0xdc, 0x2e, 0x46, 0x4f, 0x0d, 0x8b, 0xd1, 0xd5, 0x61, 0x31,
	0xba, 0x39, 0x2c, 0x46, 0xb7, 0x86, 0xc5, 0xe8, 0xbb, 0x61, 0x31, 0xfa, 0x61, 0x58, 0x9c, 0xb9,
	0x33, 0x2c, 0x46, 0x57, 0xbe, 0x2f, 0xce, 0x3c, 0x7e, 0x78, 0xf4, 0x93, 0x35, 0xb8, 0xac, 0x52,
	0xb3, 0x7c, 0xf7, 0x15, 0xb1, 0xb9, 0x23, 0xdc, 0x0a, 0x8f, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff,
	0xd9, 0x97, 0x1d, 0x7f, 0x3b, 0x06, 0x00, 0x00,
}
