package job

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"os"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes ランダムな文字列生成
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// CreateJobs jobの作成
func CreateJobs(amount int) []string {
	var jobs []string

	for i := 0; i < amount; i++ {
		jobs = append(jobs, RandStringRunes(8))
	}
	return jobs
}

// Exec 実実行
func Exec(word string, id int) {
	h := fnv.New32a()
	h.Write([]byte(word))
	time.Sleep(time.Second)
	if os.Getenv("DEBUG") == "true" {
		fmt.Printf("worker [%d] - created hash [%d] from word [%s]\n", id, h.Sum32(), word)
	}
}