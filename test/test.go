package main

import "touch-secure/src/entity"

func main() {
	var as entity.Answer
	as.Answer = "人面桃花相映红"
	as.QuestionID = 1
	as.AddAnswer(as)
}
