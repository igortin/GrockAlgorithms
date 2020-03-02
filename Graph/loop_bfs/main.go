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
	graph["Gaal"] = append(graph["Gaal"], "Patrick")
	graph["Gaal"] = append(graph["Gaal"], "Gergia")
	fmt.Println("NO")
	_ = bfs("me", "Getta", graph) // func paramters start - from what vertex start find, interested item - what to find , graph -  where to find
}

func bfs(start string, interested_item string, graph map[string][]string) bool {
	// create  queue for scan friends
	search_queue := make([]string, 0)

	// append styart vertex's freinds
	search_queue = append(search_queue, graph[start]...)

	// create map for visited vertexes
	visited := make(map[string]bool)

	// 	while search_queue > 0 then scan vertex and check by func checkVertex(), if func return false then delete index 0 item from queue and append to vertex friends

	for len(search_queue) > 0 {
		person := search_queue[0]

		// check the vertex visited before
		if _, ok := visited[person]; ok {

			// friends of friends know one person
			fmt.Println("DUPLICATE VERTEX: Skip vertex because it was checked before -", person)

			// remove person from queue
			search_queue = search_queue[1:]

			// do not process because person was proccesed before
			continue
		}

		// func processes the input person
		if checkVertex(person, interested_item) {
			fmt.Println("FOUND VERTEX:", person)
			return true
		}

		// remove index 0 item from queue
		search_queue = search_queue[1:]

		// mark vertex - visited
		visited[person] = true

		// append vertex's friends to queue

		search_queue = append(search_queue, graph[person]...)
	}

	fmt.Println("NOT FOUND VERTEX: item", interested_item)
	return false
}

// func to check interested item
func checkVertex(name string, interested_item string) bool {
	if name == interested_item {
		return true
	}
	return false
}
