package main

import(
	"fmt"
	"crypto/sha256"
	"sort"
)

// consistent hashing Methodology
// For now no replicas is added to the ring

type ConsistentHash struct{
	sorted_ring []int
	hash_ring map[int]string

}

func (chash * ConsistentHash) getEncode(value string){
	return sha256.Sum256([]byte(value))
}

func (chash * ConsistentHash) insertServer(server string){

	server_key := getEncode(server)

	chash.hash_ring[server_key] = server
	chash.sorted_ring = append(sorted_ring,server_key)

	// hash the sorted ring for keeping the node on correct index on ring
	sort.Ints(chash.sorted_ring)
}

func (chash * ConsistentHash) getServer(request string){
	request_key = getEncode(request)
	
	// search based on binary search on the hash ring to get the server on clockwise
	server_key = sort.Search(len(chash.sorted_ring),func (i int) bool{
		return nums[i] >= request_key
	})

	if server_key < len(nums){
		return chash.hash_ring[server_key]
	}else{
		// return the start node server as it exceeds the len of the ring 
		start_node = chash.sorted_ring[0]
		return chash.hash_ring[start_node]
	}
}	

func main(){

	consistentHash := ConsistentHash{}
	servers := []string{"server1","server2","server3"}
	
	for value := range servers{
		
	}
}