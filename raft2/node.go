// Package raft implements raft.
package raft

import "code.google.com/p/go.net/context"

type stateResp struct {
	st          State
	ents, cents []Entry
	msgs        []Message
}

func (a State) Equal(b State) bool {
	return a.Term == b.Term && a.Vote == b.Vote && a.LastIndex == b.LastIndex
}

func (sr stateResp) containsUpdates(prev stateResp) bool {
	return !prev.st.Equal(sr.st) || len(sr.ents) > 0 || len(sr.cents) > 0 || len(sr.msgs) > 0
}

type Node struct {
	ctx    context.Context
	propc  chan []byte
	recvc  chan Message
	statec chan stateResp
	tickc  chan struct{}
}

func Start(ctx context.Context, id int64) *Node {
	n := &Node{
		ctx:    ctx,
		propc:  make(chan []byte),
		recvc:  make(chan Message),
		statec: make(chan stateResp),
		tickc:  make(chan struct{}),
	}
	r := &raft{raftLog: newLog(), id: id}
	go n.run(r)
	return n
}

func (n *Node) run(r *raft) {
	propc := n.propc
	statec := n.statec

	var prev stateResp
	for {
		if r.hasLeader() {
			propc = n.propc
		} else {
			// We cannot accept proposals because we don't know who
			// to send them to, so we'll apply back-pressure and
			// block senders.
			propc = nil
		}

		sr := stateResp{
			r.State,
			r.raftLog.unstableEnts(),
			r.raftLog.nextEnts(),
			r.msgs,
		}

		if sr.containsUpdates(prev) {
			statec = n.statec
		} else {
			statec = nil
		}

		select {
		case p := <-propc:
			r.propose(p)
		case m := <-n.recvc:
			r.Step(m) // raft never returns an error
		case <-n.tickc:
			// r.tick()
		case statec <- sr:
			r.raftLog.resetNextEnts()
			r.raftLog.resetUnstable()
			r.msgs = nil
		case <-n.ctx.Done():
			return
		}
	}
}

func (n *Node) Tick() error {
	select {
	case n.tickc <- struct{}{}:
		return nil
	case <-n.ctx.Done():
		return n.ctx.Err()
	}
}

// Propose proposes data be appended to the log.
func (n *Node) Propose(ctx context.Context, data []byte) error {
	select {
	case n.propc <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-n.ctx.Done():
		return n.ctx.Err()
	}
}

// Step advances the state machine using m.
func (n *Node) Step(m Message) error {
	select {
	case n.recvc <- m:
		return nil
	case <-n.ctx.Done():
		return n.ctx.Err()
	}
}

// ReadState returns the current point-in-time state.
func (n *Node) ReadState(ctx context.Context) (st State, ents, cents []Entry, msgs []Message, err error) {
	select {
	case sr := <-n.statec:
		return sr.st, sr.ents, sr.cents, sr.msgs, nil
	case <-ctx.Done():
		return State{}, nil, nil, nil, ctx.Err()
	case <-n.ctx.Done():
		return State{}, nil, nil, nil, n.ctx.Err()
	}
}
