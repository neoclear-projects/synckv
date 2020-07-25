package synckv

import (
	"log"
	"net/rpc"
	"strconv"
)

/**
 * Return the rpc Client object at a certain port
 * The return object is used for rpc calls
 **/
func connect(port int64) *rpc.Client {
	client, err := rpc.Dial("tcp", ":"+strconv.FormatInt(port, 10))
	if err != nil {
		log.Fatal("dialing:", err)
	}
	return client
}

/**
 * Get value of key to server database given port and key
 **/
func ClientGet(key string, port int64) string {
	client := connect(port)
	send := GetSend{Key: key}
	result := GetResult{}

	err := client.Call("KVCache.Get", &send, &result)
	if err != nil {
		log.Fatal("error:", err)
	}

	client.Close()
	return result.Value
}

func ClientPut(key, value string, port int64) {
	client := connect(port)
	send := PutSend{Key: key, Value: value}
	result := PutResult{}

	err := client.Call("KVCache.Put", &send, &result)
	if err != nil {
		log.Fatal("error:", err)
	}

	client.Close()
}
