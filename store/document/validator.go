package document

import (
	"github.com/irisnet/irishub-sync/logger"
	"github.com/irisnet/irishub-sync/store"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CollectionNmStakeRoleCandidate = "stake_role_candidate"

	Candidate_Field_Address = "address"
	Candidate_Field_Tokens  = "tokens"
)

type (
	Candidate struct {
		Address         string         `bson:"address"` // owner, identity key
		PubKey          string         `bson:"pub_key"`
		PubKeyAddr      string         `bson:"pub_key_addr"`
		Jailed          bool           `bson:"jailed"` // has the validator been revoked from bonded status
		Tokens          float64        `bson:"tokens"`
		OriginalTokens  string         `bson:"original_tokens"`
		DelegatorShares float64        `bson:"delegator_shares"`
		VotingPower     float64        `bson:"voting_power"` // Voting power if pubKey is a considered a validator
		Description     ValDescription `bson:"description"`  // Description terms for the candidate
		BondHeight      int64          `bson:"bond_height"`
		Status          string         `bson:"status"`
	}
)

func (d Candidate) Name() string {
	return CollectionNmStakeRoleCandidate
}

func (d Candidate) PkKvPair() map[string]interface{} {
	return bson.M{Candidate_Field_Address: d.Address}
}

func (d Candidate) Query(query bson.M, selector interface{}, sorts ...string) (
	results []Candidate, err error) {
	exop := func(c *mgo.Collection) error {
		if sorts[0] == "" {
			return c.Find(query).Select(selector).All(&results)
		} else {
			return c.Find(query).Select(selector).Sort(sorts...).All(&results)
		}
	}
	return results, store.ExecCollection(d.Name(), exop)
}

func (d Candidate) Remove(query bson.M) error {
	remove := func(c *mgo.Collection) error {
		changeInfo, err := c.RemoveAll(query)
		logger.Info("Remove candidates", logger.Any("changeInfo", changeInfo))
		return err
	}
	return store.ExecCollection(d.Name(), remove)
}

func (d Candidate) QueryAll() (candidates []Candidate) {
	selector := bson.M{
		"address":      1,
		"pub_key_addr": 1,
		"tokens":       1,
		"jailed":       1,
		"status":       1,
	}
	candidates, err := d.Query(nil, selector, "")

	if err != nil {
		logger.Error("candidate collection is empty")
	}
	return candidates
}

func (d Candidate) RemoveCandidates() error {
	query := bson.M{}
	return d.Remove(query)
}

func (d Candidate) SaveAll(candidates []Candidate) error {
	var docs []interface{}

	if len(candidates) == 0 {
		return nil
	}

	for _, v := range candidates {
		docs = append(docs, v)
	}

	err := store.SaveAll(d.Name(), docs)

	return err
}
