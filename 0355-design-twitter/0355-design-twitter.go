type tweet struct {
    id   int
    user int
    prev *tweet
    next *tweet
}

type Twitter struct {
    // For each user, who do they follow?
    following map[int]map[int]struct{}
    // Head of linked list containing tweets
    head      *tweet
}

func Constructor() Twitter {
    return Twitter{
        following: make(map[int]map[int]struct{}),
        head: nil,
    }
}

// O(1) insertion
func (this *Twitter) PostTweet(userId int, tweetId int)  {
    node := &tweet{
        id: tweetId,
        user: userId,
        prev: nil,
        next: nil,
    }

    if this.head == nil {
        this.head = node
    } else {
        this.head.prev = node
        node.next = this.head
        this.head = node
    }
}

// Ten most recent tweet IDs for user's following (or users')
func (this *Twitter) GetNewsFeed(userId int) []int {
    tweets := make([]int, 0, 10)
    count := 0
    for curr := this.head; curr != nil && count < 10; curr = curr.next {

        if _, found := this.following[userId][curr.user]; found || curr.user == userId {
            tweets = append(tweets, curr.id)
            count++
        }
    }
    return tweets
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    if this.following[followerId] == nil {
        this.following[followerId] = make(map[int]struct{})
    }
    this.following[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    delete(this.following[followerId], followeeId)
}
