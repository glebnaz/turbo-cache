package lru

import "context"

type Interface interface {
	//Put value to lru
	Put(ctx context.Context, key string, val []byte)
	//Get by key
	//put to list again
	Get(ctx context.Context, key string) []byte
	//View get by key
	//only get, don't put to list again
	View(ctx context.Context, key string) []byte
	//Purge clean cache
	Purge(ctx context.Context) error
}
