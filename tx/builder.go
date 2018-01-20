package tx

import (
	"encoding/binary"
	"math/big"

	"github.com/vechain/thor/thor"
)

// Builder to make it easy to build transaction.
type Builder struct {
	body body
}

// Clause add a clause.
func (b *Builder) Clause(c *Clause) *Builder {
	b.body.Clauses = append(b.body.Clauses, c)
	return b
}

// GasPrice set gas price.
func (b *Builder) GasPrice(price *big.Int) *Builder {
	b.body.GasPrice = new(big.Int).Set(price)
	return b
}

// Gas set gas provision for tx.
func (b *Builder) Gas(gas uint64) *Builder {
	b.body.Gas = gas
	return b
}

// BlockRef set block reference.
func (b *Builder) BlockRef(br BlockRef) *Builder {
	b.body.BlockRef = binary.BigEndian.Uint64(br[:])
	return b
}

// Nonce set nonce.
func (b *Builder) Nonce(nonce uint64) *Builder {
	b.body.Nonce = nonce
	return b
}

// DependsOn set depended tx.
func (b *Builder) DependsOn(txHash *thor.Hash) *Builder {
	if txHash == nil {
		b.body.DependsOn = nil
	} else {
		cpy := *txHash
		b.body.DependsOn = &cpy
	}
	return b
}

// Build build tx object.
func (b *Builder) Build() *Transaction {
	if b.body.GasPrice == nil {
		b.body.GasPrice = &big.Int{}
	}
	tx := Transaction{body: b.body}
	return &tx
}
