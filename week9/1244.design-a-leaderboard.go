type Leaderboard struct {
  scoreBoard map[int]*Score
  data maxHeap
}

func (this Leaderboard) Print() {
  for id, s := range this.scoreBoard {
    fmt.Printf("%v, %v\n", id, s.score)
  }
  fmt.Println()
}


func Constructor() Leaderboard {
  lb := Leaderboard {
    scoreBoard: map[int]*Score{},
    data: make(maxHeap, 0),
  }
  return lb
}


func (this *Leaderboard) AddScore(playerId int, score int)  {
  if s, exist := this.scoreBoard[playerId]; exist {
    s.score += score
    heap.Fix(&this.data, s.index)
    return
  }
  this.scoreBoard[playerId] = &Score{
    score: score,
    index: 0,
  }
  heap.Push(&this.data, this.scoreBoard[playerId])
}


func (this *Leaderboard) Top(k int) int {
  topK := make([]*Score, k)
  sum := 0
  for i := range topK {
    topK[i] = heap.Pop(&this.data).(*Score)
    sum += topK[i].score
  }
  
  for _, s := range topK {
    heap.Push(&this.data, s)
  }
  return sum
}


func (this *Leaderboard) Reset(playerId int) {
  score := this.scoreBoard[playerId]
  heap.Remove(&this.data, score.index)
  delete(this.scoreBoard, playerId)
}

type Score struct {
  score int
  index int
}

type maxHeap []*Score

func (s maxHeap) Len() int {
  return len(s)
}

func (s maxHeap) Less(i, j int) bool {
  return s[j].score < s[i].score
}

func (s maxHeap) Swap(i, j int) {
  s[i], s[j] = s[j], s[i]
  s[i].index = i
  s[j].index = j
}

func (s *maxHeap) Push(v interface{}) {
  score := v.(*Score)
  score.index = len(*s)
  *s = append(*s, score)
}

func (s *maxHeap) Pop() interface{} {
  old := *s
  n := len(old)
  v := old[n-1]
  old[n-1] = nil
  v.index = -1
  *s = old[:n-1]
  return v
}


/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */