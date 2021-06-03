package blockchain

// https://www.jianshu.com/p/4a06b8bcf6e7
type block02 struct {
	Index      int
	TimeStamp  string
	BPM        int
	Hash       string //本区块hash
	PrevHash   string //上一个区块hash
	Difficulty int    //难度系数
	Nonce      string
}


