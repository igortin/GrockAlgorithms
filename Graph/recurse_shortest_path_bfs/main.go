package main

import (
	"fmt"
)

func main() {
	graph := make(map[string][]string)
	graph["Heaven"] = append(graph["Heaven"], "Claire")
	graph["Heaven"] = append(graph["Heaven"], "Alena")
	graph["Heaven"] = append(graph["Heaven"], "Gaal")
	graph["Claire"] = append(graph["Claire"], "Alex")
	graph["Claire"] = append(graph["Claire"], "Aleksandr")
	graph["Claire"] = append(graph["Claire"], "Cameron")
	graph["Alena"] = append(graph["Alena"], "Lena")
	graph["Alena"] = append(graph["Alena"], "Nikolai")
	graph["Alena"] = append(graph["Alena"], "Getta")
	graph["Gaal"] = append(graph["Gaal"], "Getta")
	graph["Gaal"] = append(graph["Gaal"], "Pat")
	graph["Gaal"] = append(graph["Gaal"], "Gergia777")
	visited := make(map[string]bool)
	searchQueue := make([][]string, 0)
	goal := "Gergia777"
	start := "Heaven"
	searchQueue = append(searchQueue, []string{start})
	fmt.Println(recurseFindShortestPathVertex(start, goal, graph, searchQueue, visited))
}

func recurseFindShortestPathVertex(start string, interestedItem string, graph map[string][]string, searchQueue [][]string, visited map[string]bool) []string {
	if len(searchQueue) == 0 {
		return []string{}
	}
	tmpList := searchQueue[0]

	searchQueue = searchQueue[1:]

	node := tmpList[len(tmpList)-1]
	if _, ok := visited[node]; ok {
		tmpList := searchQueue[0]
		searchQueue = searchQueue[1:]
		node = tmpList[len(tmpList)-1]
		return recurseFindShortestPathVertex(node, interestedItem, graph, searchQueue, visited)
	}

	neighbors := graph[node]

	for _, socedNode := range neighbors {
		newPath := make([]string, 0)
		newPath = append(newPath, tmpList...)
		newPath = append(newPath, socedNode)
		if socedNode == interestedItem {
			return newPath
		}
		searchQueue = append(searchQueue, newPath)
	}
	fmt.Println("searchQueue is list of:", searchQueue)
	return recurseFindShortestPathVertex(node, interestedItem, graph, searchQueue, visited)
}
