package types

import (

	"go.mongodb.org/mongo-driver/mongo"
)

type ParsedMessage struct {
	_type			string	`bson:"type"`
	data			[]byte
}

type Fee struct {
	denom			string
	amount		int64
}

type ParsedUpdates struct {
	ChangeBalances map[string]map[string]int64	`bson:"ChangeBalances"`
}

type ParsedBlock struct {
	hash			string
	height		int64
	chainId		string
	timestamp	int64
	elapsed		int64
	proposer	string
	txs				[]string
	gasUsed		int64
	gasWanted int64
	updates		ParsedUpdates
}

type ParsedTxResult struct {
	height		int64
	_error		string | nil	`bson:"error"`
	gasUsed		int64
	updates		ParsedUpdates
	receivers	[]string
}

type ParsedTx struct {
	hash			string
	chainId		string
	sender		string
	timestamp	string
	memo			string
	fee				Fee
	gasWanted	int64
	result		ParsedTxResult
}