package main

import (
	"fmt"
)

func main() {
	graph := make(map[string][]string)
	graph["me"] = append(graph["me"], "Claire")
	graph["me"] = append(graph["me"], "Alena")
	graph["me"] = append(graph["me"], "Gaal")
	graph["Claire"] = append(graph["Claire"], "Alex")
	graph["Claire"] = append(graph["Claire"], "Aleksandr")
	graph["Claire"] = append(graph["Claire"], "Cameron")
	graph["Alena"] = append(graph["Alena"], "Lena")
	graph["Alena"] = append(graph["Alena"], "Nikolai")
	graph["Alena"] = append(graph["Alena"], "Getta")
	graph["Gaal"] = append(graph["Gaal"], "Getta")
	graph["Gaal"] = append(graph["Gaal"], "Pat")
	graph["Gaal"] = append(graph["Gaal"], "Gergia")
	// map check is visited before
	visited := make(map[string]bool)
	// recurse queue fifo for process
	search_queue := make([]string, 0)
	// start func from vertex me
	val := recurseFind("me", "Lena", graph, visited, search_queue)
	// print is exist interested item
	fmt.Println(val)
}

func recurseFind(start string, interested_item string, g map[string][]string, visited map[string]bool, search_queue []string) bool {
	// queue append friends of friend
	search_queue = append(search_queue, g[start]...)

	// if queue len 0 exit

	if len(search_queue) == 0 {
		return false
	}

	// get vertex[0] for check interested
	vertex := search_queue[0]

	// remove vertex from scan queue
	search_queue = search_queue[1:]

	// check is vertex visited before
	if _, ok := visited[vertex]; ok {
		// if visited get next vertex
		vertex = search_queue[0]
		// start recurse for next vertex
		return recurseFind(vertex, interested_item, g, visited, search_queue)
	}

	// if vertex was not visisted
	visited[vertex] = true

	// check vertex is interested item
	if testVertex(vertex, interested_item) {
		fmt.Println("Found", vertex)
		return true
	}

	// re-start recurse process to scan
	return recurseFind(vertex, interested_item, g, visited, search_queue)
}

// func to check interested item
func testVertex(name string, interested_item string) bool {
	// fmt.Println("name:", name, "interested_item", interested_item)

	if name == interested_item {
		return true
	}

	return false
}

/* The main idea:
create queue
append to queue friends of me
get my fisrt friend - queue element with index 0
check fisrt friend
start recurse func
	append friends of first friend to queue (please notice fifo, it means that order of queue in the begin wll be my friends and after first friend's friends and so ...)
	get my second friend from queue
	check second friend
	start recurse func
		append friends of second friend to queue
		get my third friend
		check third friend
*/
