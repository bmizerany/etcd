package raft

type State struct {
	Id   int64
	Vote int64
	Lead int64

	Term int64
}

func (s *State) IsLeader() bool {
	return s.Lead == s.Id
}

func (s *State) IsCandidate() bool {
	return !s.IsLeader() && s.Vote == s.Id
}

type Message struct {
	State

	To      int64
	Mark    Mark
	Entries []Entry

	// Quit signals the reciever to quit participating if they are know by
	// sender to have been removed.
	Quit bool
}

type Mark struct {
	Index int64
	Term  int64
}

type Entry struct {
	Index int64
	Term  int64
}

type peers map[int64]State

type Node struct {
	s *State

	// the state of all peers, not including self
	peers peers

	log []Entry

	msgs []Message
}

func New(id int64, peerids ...int64) *Node {
	if id == 0 {
		panic("raft: id cannot be 0")
	}
	ps := make(peers)
	for _, pid := range peerids {
		ps[pid] = State{}
	}
	n := &Node{
		s:     &State{Id: id},
		peers: ps,
		log:   []Entry{{}},
	}
	return n
}

func (n *Node) Campaign() {
	n.s.Term++
	n.s.Vote = n.s.Id

	if len(n.peers) == 0 {
		// we have no peers so we automaticly win the election
		n.s.Lead = n.s.Id
	} else {
		l := n.log[len(n.log)-1]
		for id := range n.peers {
			n.send(Message{To: id, Mark: Mark{l.Index, l.Term}})
		}
	}
}

func (n *Node) isRemoved(id int64) bool {
	// TODO: implement me.
	return false
}

func (n *Node) isPeer(id int64) bool {
	_, ok := n.peers[id]
	return ok
}

// Step advances nodes state based on m.
func (n *Node) Step(m Message) {
	switch {
	case n.isRemoved(m.Id):
		n.send(Message{To: m.Id, Quit: true})
		return
	case !n.isPeer(m.Id):
		// The sender isn't know to have been removed, but also isn't
		// known to have been added. We can only ignore it for now; We
		// must assume it will work things out on its own.
		return
	}

	n.peers[m.Id] = m.State

	if n.hasMajority() {
		n.s.Lead = n.s.Id
	}
}

func (n *Node) hasMajority() bool {
	g := 0
	for _, s := range n.peers {
		if s.Term == n.s.Term && s.Vote == n.s.Id {
			g++
		}
	}

	k := len(n.peers)
	q := k / 2 // no need to +1 since we don't include ourself in n.peers
	return g >= q
}

// send queues m in the outbox of messages. Messages sent to self are ignored.
func (n *Node) send(m Message) {
	if m.To == n.s.Id {
		return
	}
	m.State = *n.s
	n.msgs = append(n.msgs, m)
}

func (n *Node) ReadMessages() []Message {
	msgs := n.msgs
	n.msgs = nil
	return msgs
}