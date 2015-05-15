package main

import (
	"encoding/json"
	"github.com/gorilla/context"
	"net/http"
)

const (
	REQT_DEC = iota
	RESP_ENC
)

// generic decoder type
type Decoder interface {
	Decode(v interface{}) error
}

// generic encoder type
type Encoder interface {
	Encode(v interface{}) error
}

// CodecProvdr is the middle to dynamically provide
// codec according to header
type CodecProvdr struct {
}

func (p *CodecProvdr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: set decoder dynamically according to header / extension
	context.Set(r, REQT_DEC, json.NewDecoder(r.Body))

	// TODO: set encoder dynamically according to header
	context.Set(r, RESP_ENC, json.NewEncoder(w))
}

// getReqtDec gets request decoder
func GetReqtDec(r *http.Request) Decoder {
	if val, ok := context.GetOk(r, REQT_DEC); ok {
		if dec, ok := val.(Decoder); ok {
			return dec
		}
	}
	return nil
}

// getRespEnc gets response encoder
func GetRespEnc(w http.ResponseWriter, r *http.Request) Encoder {
	if val, ok := context.GetOk(r, RESP_ENC); ok {
		if enc, ok := val.(Encoder); ok {
			return enc
		}
	}
	return nil
}
