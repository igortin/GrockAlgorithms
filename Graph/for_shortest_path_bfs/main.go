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
	graph["Gaal"] = append(graph["Gaal"], "Gergia777")

	// start func from vertex me
	val := recurseFindPath("me", "Gergia777", graph)
	fmt.Println(val)
	//fmt.Println("Good")
}

func recurseFindPath(start string, interested_item string, g map[string][]string) []string {
	// create new slice for start point to proceed
	startSlice := []string{start}
	// create the map visited for check when proceed vertex
	visited := make(map[string]bool)
	// create new queue of slices for scan interested item, it will be used for iterate main loop
	search_queue := make([][]string, 0)
	// append slice startSlice ["me"] as first element to procced

	search_queue = append(search_queue, startSlice)
	// check a special condition

	if start == interested_item {
		return startSlice
	}

	// main loop to scan interested item in search_queue, where search_queue is queue of slices
	for len(search_queue) > 0 {
		// create new slice and copy to it the search_queue first element (slice), for temprorary saving
		path := search_queue[0][:]

		// remove first element (slice) from search_queue, because it is going to procced in this iteration
		search_queue = search_queue[1:]

		// get last item of slice "path" for proceeding
		vertex := path[len(path)-1]
		// check vertex was visited before or not?
		if _, ok := visited[vertex]; ok {
			continue
		}

		// iterate throgh vertex's neighbours
		for _, neighbour := range g[vertex] {
			// create new slice in each iteration and save temprorary slice to it for saving possible path from start vertex to intersted vertex
			checkPath := path[:]
			// append vertex's neighbour
			checkPath = append(checkPath, neighbour)
			fmt.Println("slice path", checkPath)

			// append slice to main serach queue for futher proceeding
			search_queue = append(search_queue, checkPath)
			// check is neighbor is interested item
			if neighbour == interested_item {
				return checkPath
			}
		}
	}
	return []string{}
}

/* main idea
work with not primitive  string, except work with slice and get last item in slice as vertex
*/
