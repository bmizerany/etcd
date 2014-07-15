package raft

import (
	"reflect"
	"sort"
	"testing"
)

const (
	a = iota + 1
	b
	c
)

type messages []Message

func (ms messages) Len() int           { return len(ms) }
func (ms messages) Less(i, j int) bool { return ms[i].To < ms[j].To }
func (ms messages) Swap(i, j int)      { ms[i], ms[j] = ms[j], ms[i] }

func TestStateIsLeader(t *testing.T) {
	tests := []struct {
		s State
		w bool
	}{
		{State{Id: 1, Lead: 1}, true},
		{State{Id: 2, Lead: 2}, true},
		{State{Id: 1, Lead: 2}, false},
		{State{Id: 2, Lead: 1}, false},
	}

	for i, tt := range tests {
		if g := tt.s.IsLeader(); g != tt.w {
			t.Errorf("#%d: IsLeader = %v, want %v", i, g, tt.w)
		}
	}
}

func TestStateIsCandidate(t *testing.T) {
	tests := []struct {
		s State
		w bool
	}{
		{State{Id: 1, Vote: 1, Lead: 0}, true},
		{State{Id: 2, Vote: 2, Lead: 0}, true},
		{State{Id: 1, Vote: 2, Lead: 0}, false},
		{State{Id: 2, Vote: 1, Lead: 0}, false},
		{State{Id: 2, Vote: 2, Lead: 2}, false},
	}

	for i, tt := range tests {
		if g := tt.s.IsCandidate(); g != tt.w {
			t.Fatalf("#%d: IsCandidate != %v", i, g, tt.w)
		}
	}
}

func ents(n ...int64) []Entry {
	if len(n)%2 != 0 {
		panic("uneven number of args")
	}

	es := []Entry{{}}
	for len(n) > 0 {
		es = append(es, Entry{Term: n[0], Index: n[1]})
		n = n[2:]
	}
	return es
}

func votes(term int64, ns ...int64) map[int64]State {
	c := make(map[int64]State)
	for i, v := range ns {
		s := State{Id: int64(i + 1), Term: term, Vote: v}
		c[s.Id] = s
	}
	return c
}

func TestCampaign(t *testing.T) {
	type votes []struct {
		id, term, vote int64
	}

	tests := []struct {
		ps votes
		w  bool
	}{
		{
			votes{{b, 1, b}, {c, 1, c}},
			false,
		},
		{
			votes{{b, 1, a}, {c, 1, b}},
			true,
		},
		{
			votes{{b, 1, a}, c: {c, 1, a}},
			true,
		},

		// votes are all in previous term
		{
			votes{{b, 1, b}, {c, 1, c}},
			false,
		},
	}

	for i, tt := range tests {
		n := New(a, b, c)
		n.Campaign()
		if !n.s.IsCandidate() {
			t.Fatal("not candidate")
		}
		for _, v := range tt.ps {
			s := State{Id: v.id, Term: v.term, Vote: v.vote}
			n.Step(Message{State: s})
		}
		if g := n.hasMajority(); g != tt.w {
			t.Errorf("#%d: hasMajority = %v, want %v", i, g, tt.w)
		}
		if g := n.s.IsLeader(); g != tt.w {
			t.Errorf("#%d: IsLeader = %v, want %v", i, g, tt.w)
		}
	}
}

func TestCampaignMessages(t *testing.T) {
	n := New(a, b, c)
	n.log = append(n.log, Entry{1, 1})
	n.Campaign()
	g := messages(n.ReadMessages())
	w := make(messages, 0)
	for id, _ := range n.peers {
		w = append(w, Message{To: id, State: *n.s, Mark: Mark{1, 1}})
	}
	sort.Sort(g)
	sort.Sort(w)
	if !reflect.DeepEqual(g, w) {
		t.Errorf("\ng  = %+v\nwant %+v", g, w)
	}
}
