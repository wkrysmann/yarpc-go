// Code generated by thriftrw v1.16.0. DO NOT EDIT.
// @generated

package atomic

import (
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Store_CompareAndSwap_Args represents the arguments for the Store.compareAndSwap function.
//
// The arguments for compareAndSwap are sent and received over the wire as this struct.
type Store_CompareAndSwap_Args struct {
	Request *CompareAndSwap `json:"request,omitempty"`
}

// ToWire translates a Store_CompareAndSwap_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Store_CompareAndSwap_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Request != nil {
		w, err = v.Request.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _CompareAndSwap_Read(w wire.Value) (*CompareAndSwap, error) {
	var v CompareAndSwap
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Store_CompareAndSwap_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Store_CompareAndSwap_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Store_CompareAndSwap_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Store_CompareAndSwap_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Request, err = _CompareAndSwap_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a Store_CompareAndSwap_Args
// struct.
func (v *Store_CompareAndSwap_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Request != nil {
		fields[i] = fmt.Sprintf("Request: %v", v.Request)
		i++
	}

	return fmt.Sprintf("Store_CompareAndSwap_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Store_CompareAndSwap_Args match the
// provided Store_CompareAndSwap_Args.
//
// This function performs a deep comparison.
func (v *Store_CompareAndSwap_Args) Equals(rhs *Store_CompareAndSwap_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Request == nil && rhs.Request == nil) || (v.Request != nil && rhs.Request != nil && v.Request.Equals(rhs.Request))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Store_CompareAndSwap_Args.
func (v *Store_CompareAndSwap_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Request != nil {
		err = multierr.Append(err, enc.AddObject("request", v.Request))
	}
	return err
}

// GetRequest returns the value of Request if it is set or its
// zero value if it is unset.
func (v *Store_CompareAndSwap_Args) GetRequest() (o *CompareAndSwap) {
	if v != nil && v.Request != nil {
		return v.Request
	}

	return
}

// IsSetRequest returns true if Request is not nil.
func (v *Store_CompareAndSwap_Args) IsSetRequest() bool {
	return v != nil && v.Request != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "compareAndSwap" for this struct.
func (v *Store_CompareAndSwap_Args) MethodName() string {
	return "compareAndSwap"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Store_CompareAndSwap_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Store_CompareAndSwap_Helper provides functions that aid in handling the
// parameters and return values of the Store.compareAndSwap
// function.
var Store_CompareAndSwap_Helper = struct {
	// Args accepts the parameters of compareAndSwap in-order and returns
	// the arguments struct for the function.
	Args func(
		request *CompareAndSwap,
	) *Store_CompareAndSwap_Args

	// IsException returns true if the given error can be thrown
	// by compareAndSwap.
	//
	// An error can be thrown by compareAndSwap only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for compareAndSwap
	// given the error returned by it. The provided error may
	// be nil if compareAndSwap did not fail.
	//
	// This allows mapping errors returned by compareAndSwap into a
	// serializable result struct. WrapResponse returns a
	// non-nil error if the provided error cannot be thrown by
	// compareAndSwap
	//
	//   err := compareAndSwap(args)
	//   result, err := Store_CompareAndSwap_Helper.WrapResponse(err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from compareAndSwap: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(error) (*Store_CompareAndSwap_Result, error)

	// UnwrapResponse takes the result struct for compareAndSwap
	// and returns the erorr returned by it (if any).
	//
	// The error is non-nil only if compareAndSwap threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   err := Store_CompareAndSwap_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Store_CompareAndSwap_Result) error
}{}

func init() {
	Store_CompareAndSwap_Helper.Args = func(
		request *CompareAndSwap,
	) *Store_CompareAndSwap_Args {
		return &Store_CompareAndSwap_Args{
			Request: request,
		}
	}

	Store_CompareAndSwap_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *IntegerMismatchError:
			return true
		default:
			return false
		}
	}

	Store_CompareAndSwap_Helper.WrapResponse = func(err error) (*Store_CompareAndSwap_Result, error) {
		if err == nil {
			return &Store_CompareAndSwap_Result{}, nil
		}

		switch e := err.(type) {
		case *IntegerMismatchError:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for Store_CompareAndSwap_Result.Mismatch")
			}
			return &Store_CompareAndSwap_Result{Mismatch: e}, nil
		}

		return nil, err
	}
	Store_CompareAndSwap_Helper.UnwrapResponse = func(result *Store_CompareAndSwap_Result) (err error) {
		if result.Mismatch != nil {
			err = result.Mismatch
			return
		}
		return
	}

}

// Store_CompareAndSwap_Result represents the result of a Store.compareAndSwap function call.
//
// The result of a compareAndSwap execution is sent and received over the wire as this struct.
type Store_CompareAndSwap_Result struct {
	Mismatch *IntegerMismatchError `json:"mismatch,omitempty"`
}

// ToWire translates a Store_CompareAndSwap_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Store_CompareAndSwap_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Mismatch != nil {
		w, err = v.Mismatch.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	if i > 1 {
		return wire.Value{}, fmt.Errorf("Store_CompareAndSwap_Result should have at most one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _IntegerMismatchError_Read(w wire.Value) (*IntegerMismatchError, error) {
	var v IntegerMismatchError
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Store_CompareAndSwap_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Store_CompareAndSwap_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Store_CompareAndSwap_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Store_CompareAndSwap_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Mismatch, err = _IntegerMismatchError_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Mismatch != nil {
		count++
	}
	if count > 1 {
		return fmt.Errorf("Store_CompareAndSwap_Result should have at most one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Store_CompareAndSwap_Result
// struct.
func (v *Store_CompareAndSwap_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Mismatch != nil {
		fields[i] = fmt.Sprintf("Mismatch: %v", v.Mismatch)
		i++
	}

	return fmt.Sprintf("Store_CompareAndSwap_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Store_CompareAndSwap_Result match the
// provided Store_CompareAndSwap_Result.
//
// This function performs a deep comparison.
func (v *Store_CompareAndSwap_Result) Equals(rhs *Store_CompareAndSwap_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Mismatch == nil && rhs.Mismatch == nil) || (v.Mismatch != nil && rhs.Mismatch != nil && v.Mismatch.Equals(rhs.Mismatch))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Store_CompareAndSwap_Result.
func (v *Store_CompareAndSwap_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Mismatch != nil {
		err = multierr.Append(err, enc.AddObject("mismatch", v.Mismatch))
	}
	return err
}

// GetMismatch returns the value of Mismatch if it is set or its
// zero value if it is unset.
func (v *Store_CompareAndSwap_Result) GetMismatch() (o *IntegerMismatchError) {
	if v != nil && v.Mismatch != nil {
		return v.Mismatch
	}

	return
}

// IsSetMismatch returns true if Mismatch is not nil.
func (v *Store_CompareAndSwap_Result) IsSetMismatch() bool {
	return v != nil && v.Mismatch != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "compareAndSwap" for this struct.
func (v *Store_CompareAndSwap_Result) MethodName() string {
	return "compareAndSwap"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Store_CompareAndSwap_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
