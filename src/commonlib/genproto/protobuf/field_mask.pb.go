// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/protobuf/field_mask.proto
// DO NOT EDIT!

package descriptor // import "commonlib/genproto/protobuf"

import proto "commonlib/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `FieldMask` represents a set of symbolic field paths, for example:
//
//     paths: "f.a"
//     paths: "f.b.d"
//
// Here `f` represents a field in some root message, `a` and `b`
// fields in the message found in `f`, and `d` a field found in the
// message in `f.b`.
//
// Field masks are used to specify a subset of fields that should be
// returned by a get operation or modified by an update operation.
// Field masks also have a custom JSON encoding (see below).
//
// # Field Masks in Projections
//
// When used in the context of a projection, a response message or
// sub-message is filtered by the API to only contain those fields as
// specified in the mask. For example, if the mask in the previous
// example is applied to a response message as follows:
//
//     f {
//       a : 22
//       b {
//         d : 1
//         x : 2
//       }
//       y : 13
//     }
//     z: 8
//
// The result will not contain specific values for fields x,y and z
// (their value will be set to the default, and omitted in proto text
// output):
//
//
//     f {
//       a : 22
//       b {
//         d : 1
//       }
//     }
//
// A repeated field is not allowed except at the last position of a
// field mask.
//
// If a FieldMask object is not present in a get operation, the
// operation applies to all fields (as if a FieldMask of all fields
// had been specified).
//
// Note that a field mask does not necessarily apply to the
// top-level response message. In case of a REST get operation, the
// field mask applies directly to the response, but in case of a REST
// list operation, the mask instead applies to each individual message
// in the returned resource list. In case of a REST custom method,
// other definitions may be used. Where the mask applies will be
// clearly documented together with its declaration in the API.  In
// any case, the effect on the returned resource/resources is required
// behavior for APIs.
//
// # Field Masks in Update Operations
//
// A field mask in update operations specifies which fields of the
// targeted resource are going to be updated. The API is required
// to only change the values of the fields as specified in the mask
// and leave the others untouched. If a resource is passed in to
// describe the updated values, the API ignores the values of all
// fields not covered by the mask.
//
// If a repeated field is specified for an update operation, the existing
// repeated values in the target resource will be overwritten by the new values.
// Note that a repeated field is only allowed in the last position of a field
// mask.
//
// If a sub-message is specified in the last position of the field mask for an
// update operation, then the existing sub-message in the target resource is
// overwritten. Given the target message:
//
//     f {
//       b {
//         d : 1
//         x : 2
//       }
//       c : 1
//     }
//
// And an update message:
//
//     f {
//       b {
//         d : 10
//       }
//     }
//
// then if the field mask is:
//
//  paths: "f.b"
//
// then the result will be:
//
//     f {
//       b {
//         d : 10
//       }
//       c : 1
//     }
//
// However, if the update mask was:
//
//  paths: "f.b.d"
//
// then the result would be:
//
//     f {
//       b {
//         d : 10
//         x : 2
//       }
//       c : 1
//     }
//
// In order to reset a field's value to the default, the field must
// be in the mask and set to the default value in the provided resource.
// Hence, in order to reset all fields of a resource, provide a default
// instance of the resource and set all fields in the mask, or do
// not provide a mask as described below.
//
// If a field mask is not present on update, the operation applies to
// all fields (as if a field mask of all fields has been specified).
// Note that in the presence of schema evolution, this may mean that
// fields the client does not know and has therefore not filled into
// the request will be reset to their default. If this is unwanted
// behavior, a specific service may require a client to always specify
// a field mask, producing an error if not.
//
// As with get operations, the location of the resource which
// describes the updated values in the request message depends on the
// operation kind. In any case, the effect of the field mask is
// required to be honored by the API.
//
// ## Considerations for HTTP REST
//
// The HTTP kind of an update operation which uses a field mask must
// be set to PATCH instead of PUT in order to satisfy HTTP semantics
// (PUT must only be used for full updates).
//
// # JSON Encoding of Field Masks
//
// In JSON, a field mask is encoded as a single string where paths are
// separated by a comma. Fields name in each path are converted
// to/from lower-camel naming conventions.
//
// As an example, consider the following message declarations:
//
//     message Profile {
//       User user = 1;
//       Photo photo = 2;
//     }
//     message User {
//       string display_name = 1;
//       string address = 2;
//     }
//
// In proto a field mask for `Profile` may look as such:
//
//     mask {
//       paths: "user.display_name"
//       paths: "photo"
//     }
//
// In JSON, the same mask is represented as below:
//
//     {
//       mask: "user.displayName,photo"
//     }
//
// # Field Masks and Oneof Fields
//
// Field masks treat fields in oneofs just as regular fields. Consider the
// following message:
//
//     message SampleMessage {
//       oneof test_oneof {
//         string name = 4;
//         SubMessage sub_message = 9;
//       }
//     }
//
// The field mask can be:
//
//     mask {
//       paths: "name"
//     }
//
// Or:
//
//     mask {
//       paths: "sub_message"
//     }
//
// Note that oneof type names ("test_oneof" in this case) cannot be used in
// paths.
type FieldMask struct {
	// The set of field mask paths.
	Paths []string `protobuf:"bytes,1,rep,name=paths" json:"paths,omitempty"`
}

func (m *FieldMask) Reset()                    { *m = FieldMask{} }
func (m *FieldMask) String() string            { return proto.CompactTextString(m) }
func (*FieldMask) ProtoMessage()               {}
func (*FieldMask) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func init() {
	proto.RegisterType((*FieldMask)(nil), "google.protobuf.FieldMask")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/protobuf/field_mask.proto", fileDescriptor2)
}

var fileDescriptor2 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4b, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f, 0x4f,
	0xcd, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x07, 0x93, 0x49, 0xa5, 0x69, 0xfa, 0x69, 0x99, 0xa9,
	0x39, 0x29, 0xf1, 0xb9, 0x89, 0xc5, 0xd9, 0x7a, 0x60, 0x31, 0x21, 0x7e, 0xa8, 0x2e, 0x98, 0x0a,
	0x25, 0x45, 0x2e, 0x4e, 0x37, 0x90, 0x22, 0xdf, 0xc4, 0xe2, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0x82,
	0xc4, 0x92, 0x8c, 0x62, 0x09, 0x46, 0x05, 0x66, 0x0d, 0xce, 0x20, 0x08, 0xc7, 0x29, 0x90, 0x4b,
	0x38, 0x39, 0x3f, 0x57, 0x0f, 0x4d, 0xa7, 0x13, 0x1f, 0x5c, 0x5f, 0x00, 0x48, 0x28, 0x80, 0x71,
	0x01, 0x23, 0xe3, 0x22, 0x26, 0x66, 0xf7, 0x00, 0xa7, 0x55, 0x4c, 0x72, 0xee, 0x10, 0xc5, 0x01,
	0x50, 0xc5, 0x7a, 0xe1, 0xa9, 0x39, 0x39, 0xde, 0x79, 0xf9, 0xe5, 0x79, 0x21, 0x95, 0x05, 0xa9,
	0xc5, 0x49, 0x6c, 0x60, 0x53, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x4c, 0x96, 0xee,
	0xc5, 0x00, 0x00, 0x00,
}
