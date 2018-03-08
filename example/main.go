package main

import (
	"fmt"

	"github.com/WindomZ/quizzee-db"
	_ "github.com/WindomZ/quizzee-db/bolt" // 加载数据库驱动
)

func main() {
	// 加载数据库文件
	err := quizzee_db.Open([]byte("example"), "./data/data.db")
	if err != nil {
		panic(err)
	}
	// 退出时，释放数据库
	defer quizzee_db.Close()

	err = BaseUsage()
	if err != nil {
		panic(err)
	}

	err = QuizUsage()
	if err != nil {
		panic(err)
	}
}

// BaseUsage 最基础用法
func BaseUsage() error {
	key := []byte("hello")
	value := []byte("world")

	// 存储
	err := quizzee_db.Put(key, value)
	if err != nil {
		return err
	}

	// 取数
	value = quizzee_db.Get([]byte("hello"))
	fmt.Println("基础.value:", string(value))

	// 总数
	count := quizzee_db.Count()
	fmt.Println("基础.count:", count)

	return nil
}

// QuizUsage 问答题用法(推荐)
func QuizUsage() error {
	// 创建问题
	quiz := quizzee_db.NewQuiz("《浮生六记》的作者是？")
	quiz.Options = []string{"苏轼", "李白", "沈复", "杜甫"}
	quiz.Answer = "沈复"

	// 存储
	err := quiz.Store()
	if err != nil {
		return err
	}

	// 取数
	quiz = quizzee_db.GetQuiz("《浮生六记》的作者是？")
	for _, option := range quiz.Options {
		fmt.Println("问答.option.1:", option)
	}
	fmt.Println("问答.answer.1:", quiz.Answer)

	// 遍历
	quizzee_db.Iterator(func(q *quizzee_db.Quiz) bool {
		fmt.Println("问答.question.2:", q.Question)
		for _, option := range quiz.Options {
			fmt.Println("问答.option.2:", option)
		}
		fmt.Println("问答.answer.2:", q.Answer)
		return true
	})

	return nil
}
