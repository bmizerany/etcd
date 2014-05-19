// Code generated by protoc-gen-gogo.
// source: snapshot_recovery_request.proto
// DO NOT EDIT!

package protobuf

import testing35 "testing"
import math_rand35 "math/rand"
import time35 "time"
import code_google_com_p_gogoprotobuf_proto20 "github.com/coreos/etcd/third_party/code.google.com/p/gogoprotobuf/proto"
import testing36 "testing"
import math_rand36 "math/rand"
import time36 "time"
import encoding_json5 "encoding/json"
import testing37 "testing"
import math_rand37 "math/rand"
import time37 "time"
import code_google_com_p_gogoprotobuf_proto21 "github.com/coreos/etcd/third_party/code.google.com/p/gogoprotobuf/proto"
import math_rand38 "math/rand"
import time38 "time"
import testing38 "testing"
import fmt10 "fmt"
import math_rand39 "math/rand"
import time39 "time"
import testing39 "testing"
import code_google_com_p_gogoprotobuf_proto22 "github.com/coreos/etcd/third_party/code.google.com/p/gogoprotobuf/proto"
import math_rand40 "math/rand"
import time40 "time"
import testing40 "testing"
import fmt11 "fmt"
import go_parser5 "go/parser"
import math_rand41 "math/rand"
import time41 "time"
import testing41 "testing"
import code_google_com_p_gogoprotobuf_proto23 "github.com/coreos/etcd/third_party/code.google.com/p/gogoprotobuf/proto"

func TestSnapshotRecoveryRequestProto(t *testing35.T) {
	popr := math_rand35.New(math_rand35.NewSource(time35.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto20.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &SnapshotRecoveryRequest{}
	if err := code_google_com_p_gogoprotobuf_proto20.Unmarshal(data, msg); err != nil {
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

func TestSnapshotRecoveryRequestMarshalTo(t *testing35.T) {
	popr := math_rand35.New(math_rand35.NewSource(time35.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &SnapshotRecoveryRequest{}
	if err := code_google_com_p_gogoprotobuf_proto20.Unmarshal(data, msg); err != nil {
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

func BenchmarkSnapshotRecoveryRequestProtoMarshal(b *testing35.B) {
	popr := math_rand35.New(math_rand35.NewSource(616))
	total := 0
	pops := make([]*SnapshotRecoveryRequest, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedSnapshotRecoveryRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto20.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkSnapshotRecoveryRequestProtoUnmarshal(b *testing35.B) {
	popr := math_rand35.New(math_rand35.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto20.Marshal(NewPopulatedSnapshotRecoveryRequest(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &SnapshotRecoveryRequest{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := code_google_com_p_gogoprotobuf_proto20.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestSnapshotRecoveryRequest_PeerProto(t *testing35.T) {
	popr := math_rand35.New(math_rand35.NewSource(time35.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest_Peer(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto20.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &SnapshotRecoveryRequest_Peer{}
	if err := code_google_com_p_gogoprotobuf_proto20.Unmarshal(data, msg); err != nil {
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

func TestSnapshotRecoveryRequest_PeerMarshalTo(t *testing35.T) {
	popr := math_rand35.New(math_rand35.NewSource(time35.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest_Peer(popr, false)
	size := p.Size()
	data := make([]byte, size)
	for i := range data {
		data[i] = byte(popr.Intn(256))
	}
	_, err := p.MarshalTo(data)
	if err != nil {
		panic(err)
	}
	msg := &SnapshotRecoveryRequest_Peer{}
	if err := code_google_com_p_gogoprotobuf_proto20.Unmarshal(data, msg); err != nil {
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

func BenchmarkSnapshotRecoveryRequest_PeerProtoMarshal(b *testing35.B) {
	popr := math_rand35.New(math_rand35.NewSource(616))
	total := 0
	pops := make([]*SnapshotRecoveryRequest_Peer, 10000)
	for i := 0; i < 10000; i++ {
		pops[i] = NewPopulatedSnapshotRecoveryRequest_Peer(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto20.Marshal(pops[i%10000])
		if err != nil {
			panic(err)
		}
		total += len(data)
	}
	b.SetBytes(int64(total / b.N))
}

func BenchmarkSnapshotRecoveryRequest_PeerProtoUnmarshal(b *testing35.B) {
	popr := math_rand35.New(math_rand35.NewSource(616))
	total := 0
	datas := make([][]byte, 10000)
	for i := 0; i < 10000; i++ {
		data, err := code_google_com_p_gogoprotobuf_proto20.Marshal(NewPopulatedSnapshotRecoveryRequest_Peer(popr, false))
		if err != nil {
			panic(err)
		}
		datas[i] = data
	}
	msg := &SnapshotRecoveryRequest_Peer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += len(datas[i%10000])
		if err := code_google_com_p_gogoprotobuf_proto20.Unmarshal(datas[i%10000], msg); err != nil {
			panic(err)
		}
	}
	b.SetBytes(int64(total / b.N))
}

func TestSnapshotRecoveryRequestJSON(t *testing36.T) {
	popr := math_rand36.New(math_rand36.NewSource(time36.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest(popr, true)
	jsondata, err := encoding_json5.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &SnapshotRecoveryRequest{}
	err = encoding_json5.Unmarshal(jsondata, msg)
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
func TestSnapshotRecoveryRequest_PeerJSON(t *testing36.T) {
	popr := math_rand36.New(math_rand36.NewSource(time36.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest_Peer(popr, true)
	jsondata, err := encoding_json5.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &SnapshotRecoveryRequest_Peer{}
	err = encoding_json5.Unmarshal(jsondata, msg)
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
func TestSnapshotRecoveryRequestProtoText(t *testing37.T) {
	popr := math_rand37.New(math_rand37.NewSource(time37.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest(popr, true)
	data := code_google_com_p_gogoprotobuf_proto21.MarshalTextString(p)
	msg := &SnapshotRecoveryRequest{}
	if err := code_google_com_p_gogoprotobuf_proto21.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestSnapshotRecoveryRequestProtoCompactText(t *testing37.T) {
	popr := math_rand37.New(math_rand37.NewSource(time37.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest(popr, true)
	data := code_google_com_p_gogoprotobuf_proto21.CompactTextString(p)
	msg := &SnapshotRecoveryRequest{}
	if err := code_google_com_p_gogoprotobuf_proto21.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestSnapshotRecoveryRequest_PeerProtoText(t *testing37.T) {
	popr := math_rand37.New(math_rand37.NewSource(time37.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest_Peer(popr, true)
	data := code_google_com_p_gogoprotobuf_proto21.MarshalTextString(p)
	msg := &SnapshotRecoveryRequest_Peer{}
	if err := code_google_com_p_gogoprotobuf_proto21.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestSnapshotRecoveryRequest_PeerProtoCompactText(t *testing37.T) {
	popr := math_rand37.New(math_rand37.NewSource(time37.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest_Peer(popr, true)
	data := code_google_com_p_gogoprotobuf_proto21.CompactTextString(p)
	msg := &SnapshotRecoveryRequest_Peer{}
	if err := code_google_com_p_gogoprotobuf_proto21.UnmarshalText(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
	}
	if !p.Equal(msg) {
		t.Fatalf("%#v !Proto %#v", msg, p)
	}
}

func TestSnapshotRecoveryRequestStringer(t *testing38.T) {
	popr := math_rand38.New(math_rand38.NewSource(time38.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest(popr, false)
	s1 := p.String()
	s2 := fmt10.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestSnapshotRecoveryRequest_PeerStringer(t *testing38.T) {
	popr := math_rand38.New(math_rand38.NewSource(time38.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest_Peer(popr, false)
	s1 := p.String()
	s2 := fmt10.Sprintf("%v", p)
	if s1 != s2 {
		t.Fatalf("String want %v got %v", s1, s2)
	}
}
func TestSnapshotRecoveryRequestSize(t *testing39.T) {
	popr := math_rand39.New(math_rand39.NewSource(time39.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto22.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto22.Marshal(p)
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
	size3 := code_google_com_p_gogoprotobuf_proto22.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkSnapshotRecoveryRequestSize(b *testing39.B) {
	popr := math_rand39.New(math_rand39.NewSource(616))
	total := 0
	pops := make([]*SnapshotRecoveryRequest, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedSnapshotRecoveryRequest(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestSnapshotRecoveryRequest_PeerSize(t *testing39.T) {
	popr := math_rand39.New(math_rand39.NewSource(time39.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest_Peer(popr, true)
	size2 := code_google_com_p_gogoprotobuf_proto22.Size(p)
	data, err := code_google_com_p_gogoprotobuf_proto22.Marshal(p)
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
	size3 := code_google_com_p_gogoprotobuf_proto22.Size(p)
	if size3 != size {
		t.Fatalf("size %v != after marshal proto.Size %v", size, size3)
	}
}

func BenchmarkSnapshotRecoveryRequest_PeerSize(b *testing39.B) {
	popr := math_rand39.New(math_rand39.NewSource(616))
	total := 0
	pops := make([]*SnapshotRecoveryRequest_Peer, 1000)
	for i := 0; i < 1000; i++ {
		pops[i] = NewPopulatedSnapshotRecoveryRequest_Peer(popr, false)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		total += pops[i%1000].Size()
	}
	b.SetBytes(int64(total / b.N))
}

func TestSnapshotRecoveryRequestGoString(t *testing40.T) {
	popr := math_rand40.New(math_rand40.NewSource(time40.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest(popr, false)
	s1 := p.GoString()
	s2 := fmt11.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser5.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestSnapshotRecoveryRequest_PeerGoString(t *testing40.T) {
	popr := math_rand40.New(math_rand40.NewSource(time40.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest_Peer(popr, false)
	s1 := p.GoString()
	s2 := fmt11.Sprintf("%#v", p)
	if s1 != s2 {
		t.Fatalf("GoString want %v got %v", s1, s2)
	}
	_, err := go_parser5.ParseExpr(s1)
	if err != nil {
		panic(err)
	}
}
func TestSnapshotRecoveryRequestVerboseEqual(t *testing41.T) {
	popr := math_rand41.New(math_rand41.NewSource(time41.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto23.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &SnapshotRecoveryRequest{}
	if err := code_google_com_p_gogoprotobuf_proto23.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}
func TestSnapshotRecoveryRequest_PeerVerboseEqual(t *testing41.T) {
	popr := math_rand41.New(math_rand41.NewSource(time41.Now().UnixNano()))
	p := NewPopulatedSnapshotRecoveryRequest_Peer(popr, false)
	data, err := code_google_com_p_gogoprotobuf_proto23.Marshal(p)
	if err != nil {
		panic(err)
	}
	msg := &SnapshotRecoveryRequest_Peer{}
	if err := code_google_com_p_gogoprotobuf_proto23.Unmarshal(data, msg); err != nil {
		panic(err)
	}
	if err := p.VerboseEqual(msg); err != nil {
		t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
	}
}

//These tests are generated by code.google.com/p/gogoprotobuf/plugin/testgen
