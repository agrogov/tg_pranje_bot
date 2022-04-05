package iterator

import "tg_pranje_bot/node"

type IIterator interface {
	HasNext() bool
	Next() *node.Node
}
