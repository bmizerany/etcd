// Code generated by protoc-gen-gogo.
// source: request_vote_responses.proto
// DO NOT EDIT!

package protobuf

import testing28 "testing"
import math_rand28 "math/rand"
import time28 "time"
import code_google_com_p_gogoprotobuf_proto16 "github.com/coreos/etcd/third_party/code.google.com/p/gogoprotobuf/proto"
import testing29 "testing"
import math_rand29 "math/rand"
import time29 "time"
import encoding_json4 "encoding/json"
import testing30 "testing"
import math_rand30 "math/rand"
import time30 "time"
import code_google_com_p_gogoprotobuf_proto17 "github.com/coreos/etcd/third_party/code.google.com/p/gogoprotobuf/proto"
import math_rand31 "math/rand"
import time31 "time"
import testing31 "testing"
import fmt8 "fmt"
import math_rand32 "math/rand"
import time32 "time"
import testing32 "testing"
import code_google_com_p_gogoprotobuf_proto18 "github.com/coreos/etcd/third_party/code.google.com/p/gogoprotobuf/proto"
import math_rand33 "math/rand"
import time33 "time"
import testing33 "testing"
import fmt9 "fmt"
import go_parser4 "go/parser"
import math_rand34 "math/rand"
import time34 "time"
import testing34 "testing"
import code_google_com_p_gogoprotobuf_proto19 "github.com/coreos/etcd/third_party/code.google.com/p/gogoprotobuf/proto"

func TestRequestVoteResponseProto(t *testing28.T) {
	popr := math_rand28.New(math_rand28.NewSource(time28.Now().UnixNano()))
	p := NewPopulatedRequestVoteResponse(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto16.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &RequestVoteResponse{}
	if err := code_google_com_p_gogoprotobuf_proto16.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestRequestVoteResponseMarshalTo(t *testing28.T) {
	popr := math_rand28.New(math_rand28.NewSource(time28.Now().UnixNano()))
	p := NewPopulatedRequestVoteResponse(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &RequestVoteResponse{}
	if err := code_google_com_p_gogoprotobuf_proto16.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func BenchmarkRequestVoteResponseProtoMarshal(b *testing28.B) {
	popr := math_rand28.New(math_rand28.NewSource(616))
	total := 0
	pops := make([]*RequestVoteResponse, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedRequestVoteResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto16.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkRequestVoteResponseProtoUnmarshal(b *testing28.B) {
	popr := math_rand28.New(math_rand28.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto16.Marshal(NewPopulatedRequestVoteResponse(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &RequestVoteResponse{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := code_google_com_p_gogoprotobuf_proto16.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestRequestVoteResponseJSON(t *testing29.T) {
	popr := math_rand29.New(math_rand29.NewSource(time29.Now().UnixNano()))
	p := NewPopulatedRequestVoteResponse(popr, true)
	jsondata, err := encoding_json4.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &RequestVoteResponse{}
	err = encoding_json4.Unmarshal(jsondata, msg)
	if err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Json Equal %#v", msg, p)
	}
}
func TestRequestVoteResponseProtoText(t *testing30.T) {
	popr := math_rand30.New(math_rand30.NewSource(time30.Now().UnixNano()))
	p := NewPopulatedRequestVoteResponse(popr, true)
	data := code_google_com_p_gogoprotobuf_proto17.MarshalTextString(p)
	msg := &RequestVoteResponse{}
	if err := code_google_com_p_gogoprotobuf_proto17.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestRequestVoteResponseProtoCompactText(t *testing30.T) {
	popr := math_rand30.New(math_rand30.NewSource(time30.Now().UnixNano()))
	p := NewPopulatedRequestVoteResponse(popr, true)
	data := code_google_com_p_gogoprotobuf_proto17.CompactTextString(p)
	msg := &RequestVoteResponse{}
	if err := code_google_com_p_gogoprotobuf_proto17.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestRequestVoteResponseStringer(t *testing31.T) {
	popr := math_rand31.New(math_rand31.NewSource(time31.Now().UnixNano()))
	p := NewPopulatedRequestVoteResponse(popr, false)
	s1 := p.String()
	s2 := fmt8.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestRequestVoteResponseSize(t *testing32.T) {
	popr := math_rand32.New(math_rand32.NewSource(time32.Now().UnixNano()))
	p := NewPopulatedRequestVoteResponse(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto18.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto18.Marshal(p)
	if err != nil {
		panic(err)
	}
	size := p.Size()
	if len(data) != size {
		t.Fatalf("size %v != marshalled size %v", size, len(data))
	}
	if size2 != size {
		t.Fatalf("size %v != before marshal proto.Size %v", size, size2)
	}
	size3 := code_google_com_p_gogoprotobuf_proto18.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkRequestVoteResponseSize(b *testing32.B) {
	popr := math_rand32.New(math_rand32.NewSource(616))
	total := 0
	pops := make([]*RequestVoteResponse, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedRequestVoteResponse(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestRequestVoteResponseGoString(t *testing33.T) {
	popr := math_rand33.New(math_rand33.NewSource(time33.Now().UnixNano()))
	p := NewPopulatedRequestVoteResponse(popr, false)
	s1 := p.GoString()
	s2 := fmt9.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser4.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestRequestVoteResponseVerboseEqual(t *testing34.T) {
	popr := math_rand34.New(math_rand34.NewSource(time34.Now().UnixNano()))
	p := NewPopulatedRequestVoteResponse(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto19.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &RequestVoteResponse{}
	if err := code_google_com_p_gogoprotobuf_proto19.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}

//These tests are generated by code.google.com/p/gogoprotobuf/plugin/testgen
