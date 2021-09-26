func findCircleNum(isConnected [][]int) int {
  visited := make([]bool, len(isConnected))
  count := 0
  for i := range isConnected {
    if !visited[i] {
      dfs(isConnected, i, visited)
      count++
    }
  }
  return count
}

func dfs(graph [][]int, i int, visited []bool) {
  visited[i] = true
  for j, c := range graph[i] {
    if c == 1 && !visited[j] {
      dfs(graph, j, visited)
    }
  }
}