package libs

import (
	"bufio"
	"os"
	"testing"
)

func TestDijkstra(t *testing.T) {
	t.Parallel()

	input, err := os.Open("../../testdata/dijkstra_input.txt")
	if err != nil {
		t.Fatal("failed to open input file")
	}
	defer input.Close()
	sc := bufio.NewScanner(input)
	sc.Split(bufio.ScanWords)

	n, m := NextInt(sc), NextInt(sc)
	graph := make(Graph, n)
	for i := 0; i < m; i++ {
		a, b, c := NextInt(sc)-1, NextInt(sc)-1, NextInt(sc)
		graph[a] = append(graph[a], Edge{To: b, Weight: c})
		graph[b] = append(graph[b], Edge{To: a, Weight: c})
	}

	// Dijkstra
	from1 := Dijkstra(graph, 0)
	fromN := Dijkstra(graph, n-1)

	output, err := os.Open("../../testdata/dijkstra_ouput.txt")
	if err != nil {
		t.Fatal("failed to open output file")
	}
	defer output.Close()
	sc = bufio.NewScanner(output)
	for k := 0; k < n; k++ {
		exp := NextInt(sc)
		if got := from1[k] + fromN[k]; got != exp {
			t.Errorf("expected value is %d, but got %d", exp, got)
		}
	}
}
