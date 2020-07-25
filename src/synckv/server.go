package synckv

import (
	"log"
	"net"
	"net/rpc"
	"strconv"
	"sync"
)

/**
 * The constant string to indicate error type
 **/
const (
	OK       = "OK"
	ErrNoKey = "ErrNoKey"
)

/**
 * The struct used to provide lock feature and remote mapping
 **/
type KVCache struct {
	mtx   sync.Mutex
	cache map[string]string
}

type Err string

type GetSend struct {
	Key string
}

type GetResult struct {
	Value string
	Err   Err
}

type PutSend struct {
	Key   string
	Value string
}

type PutResult struct {
	Err Err
}

func (kv *KVCache) Get(send *GetSend, result *GetResult) error {
	kv.mtx.Lock()
	defer kv.mtx.Unlock()

	val, ok := kv.cache[send.Key]
	if ok {
		result.Err = OK
		result.Value = val
	} else {
		result.Err = ErrNoKey
		result.Value = ""
	}

	return nil
}

func (kv *KVCache) Put(send *PutSend, result *PutResult) error {
	kv.mtx.Lock()
	defer kv.mtx.Unlock()

	kv.cache[send.Key] = send.Value
	result.Err = OK

	return nil
}

func StartServer(port int64) {
	kv := new(KVCache)

	// Init cache
	kv.cache = map[string]string{}

	// Start new server and register object
	remotePC := rpc.NewServer()
	remotePC.Register(kv)

	l, e := net.Listen("tcp", ":"+strconv.FormatInt(port, 10))
	if e != nil {
		log.Fatal("Listen error:", e)
	}

	// Run server concurrently
	go func() {
		for {
			conn, err := l.Accept()
			if err == nil {
				go remotePC.ServeConn(conn)
			} else {
				break
			}
		}
		l.Close()
	}()
}
