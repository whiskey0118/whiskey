package main

import (
	"fmt"
	"time"
)

//type Twitter struct {
//	UserId int
//	TweetId int
//	//Followed string
//	TweetTime time.Time
//	user User
//}

type Tweet struct {
	TweetId int
	Twitter *Twitter
	Text    string
	Time    time.Time
}

type Twitter struct {
	UserId       int
	Tweet        []*Tweet
	TweetNums    int
	Followed     []string
	FollowedNums int
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{}
}

func (this *Twitter) FindTwitter() *Twitter {
	var tw Twitter
	return &tw
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.UserId = userId
	user := this.FindTwitter()
	tweet := Tweet{
		TweetId: tweetId,
		Twitter: user,
		Text:    "",
		Time:    time.Time{},
	}
	fmt.Println(tweet)
}

func (this *Twitter) FindUserFollwed() []string {
	return []string{}
}

func (this *Twitter) GetTenTweet() []Tweet {
	return []Tweet{}
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
//func (this *Twitter) GetNewsFeed(userId int) []int {
//	var (
//		user Twitter
//		followed []string
//		followedNums int
//		res Tweet
//	)
//
//	user.UserId = userId
//	followed = user.FindUserFollwed()
//	followedNums  = user.FollowedNums
//	for i:=range followed{
//
//	}
//
//}

///** Follower follows a followee. If the operation is invalid, it should be a no-op. */
//func (this *Twitter) Follow(followerId int, followeeId int)  {
//	this.user.FollowedNums = this.user.FollowedNums+1
//	followed := this.user.Followed
//	for i :=range followed{
//		if followeerId
//	}
//}
//
//
///** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
//func (this *Twitter) Unfollow(followerId int, followeeId int)  {
//	this.user.FollowedNums = this.user.FollowedNums-1
//}
