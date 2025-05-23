package util

var lastGeneratedUserId = 0
var lastGeneratedAccountId = 10000
var lastGeneratedTransactionId = 20000

func GetUniqueUserID() int {
	lastGeneratedUserId++
	return lastGeneratedUserId
}

func GetUniqueAccountID() int {
	lastGeneratedAccountId++
	return lastGeneratedAccountId
}

func GetUniqueTransactionId() int {
	lastGeneratedTransactionId++
	return lastGeneratedTransactionId
}
