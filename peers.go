package main

import "sync"

type Peers struct {
	m map[string]chan<- Message
	mu sync.RWMutex
}

func (peers *Peers) Add(address string) <-chan Message {
	peers.mu.Lock()
	defer peers.mu.Unlock()

	if _, ok := peers.m[address]; ok {
		return nil
	}

	messages := make(chan Message)
	peers.m[address] = messages
	return messages
}

func (peers *Peers) List() []chan<- Message {
	peers.mu.RLock()
	defer peers.mu.RUnlock()

	list := make([]chan<- Message, 0, len(peers.m))
	for _, messages := range peers.m {
		list = append(list, messages)
	}
	return list
}

func (peers *Peers) Remove(address string) {
	peers.mu.Lock()
	defer peers.mu.Unlock()

	delete(peers.m, address)
}