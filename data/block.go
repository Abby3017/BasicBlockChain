package data

// Block Structure for blockchain
type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

// Message structure for blockchain
type Message struct {
	BPM int
}
