package cmd

import (
	"fmt"

	"github.com/isaacp/rete/core/entities"
	"github.com/isaacp/rete/core/interfaces"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Do Stuff",
	Long:  `The doing of stuff`,
	Run: func(cmd *cobra.Command, args []string) {
		names := []string{"Isaac", "Jermaine", "Reggie"}
		nodes := make([]interfaces.INode, 0)

		for _, name := range names {
			nodes = append(nodes, entities.NewNode(name))
		}

		nodes[0].AddNeighbor(nodes[2])
		nodes[1].AddNeighbor(nodes[0])
		nodes[2].AddNeighbor(nodes[1])

		for _, node := range nodes {
			for _, neighbor := range node.Neighbors() {
				fmt.Println(node.Name() + " <------> " + neighbor.Name())
				fmt.Println(node.Id() + " <------> " + neighbor.Id())
				fmt.Print("\n")
			}
			fmt.Print("\n")
		}
		// find all connected
		fmt.Println("Finding all connected for: " + nodes[0].Name())
		visited := make(map[string]bool)
		var printTree func(interfaces.INode, map[string]bool)
		printTree = func(node interfaces.INode, visited map[string]bool) {
			if visited[node.Name()] {
				return
			}
			visited[node.Name()] = true
			fmt.Println(node.Name())
			for _, n := range node.Neighbors() {
				printTree(n, visited)
			}
		}

		printConnections := func(root interfaces.INode, visited map[string]bool) {
			visited[root.Name()] = true
			for _, node := range root.Neighbors() {
				printTree(node, visited)
			}
		}

		printConnections(nodes[2], visited)
	},
}
