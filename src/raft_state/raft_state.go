package raft_state

type NodeRole int

const (
	Follower NodeRole = iota
	Leader
	Candidate
)

const NotVoted = -1

// VolatileNodeState struct for volatile state kept on all nodes
type VolatileNodeState struct {
	// Current role of a node
	Role NodeRole
	// Index of highest log entry known to be committed
	CommitIndex uint
	// Index of highest log entry applied to state machine
	LastApplied uint
	// Id of current leader
	LeaderId uint
}

// VolatileLeaderState struct for volatile state kept only on leader (reinitialized after election)
type VolatileLeaderState struct {
	// Index of the next log entry to send for given node (array index is equal to node id)
	NextIndex []uint
	// Index of highest log entry known to be replicated on given node (array index is equal to node id)
	MatchIndex []uint
}

type VolatileState struct {
	// Current role of a node
	Role NodeRole
	// Index of highest log entry known to be committed
	CommitIndex uint
	// Index of highest log entry applied to state machine
	LastApplied uint
	// Id of current leader
	LeaderId uint
}

type LogEntry struct {
	// Index of given log entry
	Index uint
	// command for a state machine
	Command string
	// term in which entry was received by leader
	Term uint
}

// PersistentState struct for persistent state kept on all nodes
type PersistentState struct {
	// Id of this node
	NodeId uint
	// latest term server has seen
	CurrentTerm uint
	// Id of candidate that received vote in current term
	VotedFor int
	// Log entries for state machine
	Log []LogEntry
}
