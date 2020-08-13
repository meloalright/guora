package rdbservice

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/meloalright/guora/database"
	"github.com/meloalright/guora/model"
)

var ctx = context.Background()

func RedisAddSupporter(AnswerID int, MemberID int) (err error) {

	redisKey := getRedisKey(MemberID)
	isMember, err := database.RDB.SIsMember(ctx, redisKey, AnswerID).Result()
	if err != nil {
		return
	}
	if isMember == true {
		err = errors.New("Has upvoted")
		return
	}
	err = database.RDB.SAdd(ctx, redisKey, AnswerID).Err()
	return

}

func RedisRemoveSupporter(AnswerID int, MemberID int) (err error) {

	redisKey := getRedisKey(MemberID)
	isMember, err := database.RDB.SIsMember(ctx, redisKey, AnswerID).Result()
	if err != nil {
		return
	}
	if isMember == false {
		err = errors.New("Hasn't upvoted")
		return
	}
	err = database.RDB.SRem(ctx, redisKey, AnswerID).Err()
	return

}

func RedisWrapSupported(answer *model.Answer, MemberID int) (err error) {

	memMap, err := getRedisSmemMap(MemberID)

	if err != nil {
		return
	}

	IDString := strconv.Itoa(answer.ID)
	answer.Supported = memMap[IDString]

	return
}

func RedisWrapListSupported(answers []model.Answer, MemberID int) (err error) {

	memMap, err := getRedisSmemMap(MemberID)

	if err != nil {
		return
	}

	for key, answer := range answers {
		IDString := strconv.Itoa(answer.ID)
		answers[key].Supported = memMap[IDString]
	}

	return
}

func getRedisKey(MemberID int) (redisKey string) {
	redisKey = fmt.Sprintf("supporter:profile_id:%v", MemberID)
	return
}

func getRedisSmemMap(MemberID int) (memMap map[string]bool, err error) {

	redisKey := getRedisKey(MemberID)
	memList, err := database.RDB.SMembers(ctx, redisKey).Result()

	if err != nil {
		return
	}

	memMap = make(map[string]bool)
	for _, _AnswerId := range memList {
		memMap[_AnswerId] = true
	}

	return
}
