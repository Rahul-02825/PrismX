package main

import(
	"fmt"
	"crypto/sha256"
	"sort"
	"encoding/binary"
)

// consistent hashing Methodology
// For now no replicas is added to the ring

type ConsistentHash struct{
	sorted_ring []uint64
	hash_ring map[uint64]string
}

func (chash * ConsistentHash) getEncode(value string) uint64{
	hash := sha256.Sum256([]byte(value))
	return binary.BigEndian.Uint64(hash[:8])
}

func (chash * ConsistentHash) insertServer(server string){
	// initialize hash ring if not
	if chash.hash_ring == nil {
		chash.hash_ring = make(map[uint64]string)
	}
	server_key := chash.getEncode(server)

	chash.hash_ring[server_key] = server
	chash.sorted_ring = append(chash.sorted_ring,server_key)

	// hash the sorted ring for keeping the node on correct index on ring
	sort.Slice(chash.sorted_ring, func(i, j int) bool {
		return chash.sorted_ring[i] < chash.sorted_ring[j]
	})
}

// removeServer removes a server from the hash ring.
func (chash *ConsistentHash) removeServer(server string) {
	if chash.hash_ring == nil {
		return
	}
	
	// get the server key
	server_key := chash.getEncode(server)

	// check for server existence
	_,ok := chash.hash_ring[server_key]

	if ok{
		// Remove from hash and ring as well
		delete(chash.hash_ring, server_key)

		// search in log time
		idx := sort.Search(len(chash.sorted_ring), func(i int) bool {
			return chash.sorted_ring[i] >= server_key
		})

		if idx < len(chash.sorted_ring) && chash.sorted_ring[idx] == server_key {
			chash.sorted_ring = append(chash.sorted_ring[:idx], chash.sorted_ring[idx+1:]...)
		}
	}else{
		fmt.Printf("No such servers exist")
	}
}


func (chash * ConsistentHash) getServer(request string) string{
	request_key := chash.getEncode(request)
	
	// search based on binary search on the hash ring to get the server on clockwise
	server_key_index := sort.Search(len(chash.sorted_ring),func (i int) bool{
		return chash.sorted_ring[i] >= request_key
	})

	if server_key_index < len(chash.sorted_ring){
		server_key := chash.sorted_ring[server_key_index]
		return chash.hash_ring[server_key]
	}else{
		// return the start node server as it exceeds the len of the ring 
		start_node := chash.sorted_ring[0]
		return chash.hash_ring[start_node]
	}
}	

func main(){

	consistentHash := ConsistentHash{}
	servers := []string{"server1","server2","server3"}
	
	for _,value := range servers{
		consistentHash.insertServer(value)	
	}
	s1:=consistentHash.getServer("request1")
	s2:=consistentHash.getServer("request2")
	s3:=consistentHash.getServer("request5")

	fmt.Printf("%s,%s,%s\n",s1,s2,s3)

	fmt.Printf("Done\n")
}