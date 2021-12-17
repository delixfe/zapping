package internal

import (
	"time"

	"go.uber.org/zap/zapcore"
)

// Code generated by gowrap. DO NOT EDIT.
// template: templates/decorator
// gowrap: http://github.com/hexdigest/gowrap

//go:generate gowrap gen -p go.uber.org/zap/zapcore -i Encoder -t templates/decorator -o enveloping_encoder_gen.go -v DecoratorName=envelopingEncoder -v IgnoreMethod1=EncodeEntry -v IgnoreMethod2=Clone -v noStruct -l ""

var _ zapcore.Encoder = &envelopingEncoder{}

// AddArray redirects to inner
func (r *envelopingEncoder) AddArray(key string, marshaler zapcore.ArrayMarshaler) (err error) {
	return r.inner.AddArray(key, marshaler)
}

// AddBinary redirects to inner
func (r *envelopingEncoder) AddBinary(key string, value []byte) {
	r.inner.AddBinary(key, value)
	return
}

// AddBool redirects to inner
func (r *envelopingEncoder) AddBool(key string, value bool) {
	r.inner.AddBool(key, value)
	return
}

// AddByteString redirects to inner
func (r *envelopingEncoder) AddByteString(key string, value []byte) {
	r.inner.AddByteString(key, value)
	return
}

// AddComplex128 redirects to inner
func (r *envelopingEncoder) AddComplex128(key string, value complex128) {
	r.inner.AddComplex128(key, value)
	return
}

// AddComplex64 redirects to inner
func (r *envelopingEncoder) AddComplex64(key string, value complex64) {
	r.inner.AddComplex64(key, value)
	return
}

// AddDuration redirects to inner
func (r *envelopingEncoder) AddDuration(key string, value time.Duration) {
	r.inner.AddDuration(key, value)
	return
}

// AddFloat32 redirects to inner
func (r *envelopingEncoder) AddFloat32(key string, value float32) {
	r.inner.AddFloat32(key, value)
	return
}

// AddFloat64 redirects to inner
func (r *envelopingEncoder) AddFloat64(key string, value float64) {
	r.inner.AddFloat64(key, value)
	return
}

// AddInt redirects to inner
func (r *envelopingEncoder) AddInt(key string, value int) {
	r.inner.AddInt(key, value)
	return
}

// AddInt16 redirects to inner
func (r *envelopingEncoder) AddInt16(key string, value int16) {
	r.inner.AddInt16(key, value)
	return
}

// AddInt32 redirects to inner
func (r *envelopingEncoder) AddInt32(key string, value int32) {
	r.inner.AddInt32(key, value)
	return
}

// AddInt64 redirects to inner
func (r *envelopingEncoder) AddInt64(key string, value int64) {
	r.inner.AddInt64(key, value)
	return
}

// AddInt8 redirects to inner
func (r *envelopingEncoder) AddInt8(key string, value int8) {
	r.inner.AddInt8(key, value)
	return
}

// AddObject redirects to inner
func (r *envelopingEncoder) AddObject(key string, marshaler zapcore.ObjectMarshaler) (err error) {
	return r.inner.AddObject(key, marshaler)
}

// AddReflected redirects to inner
func (r *envelopingEncoder) AddReflected(key string, value interface{}) (err error) {
	return r.inner.AddReflected(key, value)
}

// AddString redirects to inner
func (r *envelopingEncoder) AddString(key string, value string) {
	r.inner.AddString(key, value)
	return
}

// AddTime redirects to inner
func (r *envelopingEncoder) AddTime(key string, value time.Time) {
	r.inner.AddTime(key, value)
	return
}

// AddUint redirects to inner
func (r *envelopingEncoder) AddUint(key string, value uint) {
	r.inner.AddUint(key, value)
	return
}

// AddUint16 redirects to inner
func (r *envelopingEncoder) AddUint16(key string, value uint16) {
	r.inner.AddUint16(key, value)
	return
}

// AddUint32 redirects to inner
func (r *envelopingEncoder) AddUint32(key string, value uint32) {
	r.inner.AddUint32(key, value)
	return
}

// AddUint64 redirects to inner
func (r *envelopingEncoder) AddUint64(key string, value uint64) {
	r.inner.AddUint64(key, value)
	return
}

// AddUint8 redirects to inner
func (r *envelopingEncoder) AddUint8(key string, value uint8) {
	r.inner.AddUint8(key, value)
	return
}

// AddUintptr redirects to inner
func (r *envelopingEncoder) AddUintptr(key string, value uintptr) {
	r.inner.AddUintptr(key, value)
	return
}

// OpenNamespace redirects to inner
func (r *envelopingEncoder) OpenNamespace(key string) {
	r.inner.OpenNamespace(key)
	return
}
