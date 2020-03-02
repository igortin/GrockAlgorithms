package main

import (
	"fmt"
)

func main() {
	// graph for scan
	graph := make(map[string]map[string]int)
	m1 := map[string]int{"a": 7, "b": 2} // m1["c"] = 5
	graph["start"] = m1
	m2 := map[string]int{"fin": 1}
	graph["a"] = m2
	m3 := map[string]int{"a": 3, "fin": 5}
	graph["b"] = m3
	graph["fin"] = map[string]int{}
	// fmt.Println("graph:", graph)
	// map costs evaulate
	costs := make(map[string]int)
	costs["a"] = 7
	costs["b"] = 2
	costs["fin"] = 999999
	// fmt.Println("costs:", costs)
	// map parent evaulate
	parents := make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	parents["fin"] = ""
	// fmt.Println("parents:", parents)
	visited := make(map[string]bool, 0)
	node := findLowestCostNode(costs, visited)

	fmt.Println(node)

}

func findLowestCostNode(costs map[string]int, visited map[string]bool) string {

	lowest_cost := 99999

	lowest_cost_node := ""

	for node, val := range costs {

		if _, ok := visited[node]; !ok {
			if val < lowest_cost {
				lowest_cost = val
				lowest_cost_node = node
			}
		}
	}
	return lowest_cost_node
}
