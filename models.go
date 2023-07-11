package main

import (
	"kimb/internal/database"
	"time"

	"github.com/google/uuid"
)

/**USER MODELS */
type User struct {
	ID           uuid.UUID `json:"id"`
	ApiKey       string    `json:"api_key"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	DisplayName  string    `json:"display_name"`
	Email        string    `json:"email"`
	Password     []byte    `json:"password"`
	Description  string    `json:"description"`
	Verified     bool      `json:"verified"`
	Organization bool      `json:"organization"`
}

func returnUser(dbUser database.User) User {
	return User{
		ID:          dbUser.ID,
		ApiKey:      dbUser.ApiKey,
		CreatedAt:   dbUser.CreatedAt,
		UpdatedAt:   dbUser.UpdatedAt,
		Name:        dbUser.Name,
		DisplayName: dbUser.DisplayName,
		Description: dbUser.Description,
	}
	// Not returning password not email
}

type LoginUser struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

/**POST MODELS */
type Post struct {
	ID        uuid.UUID `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uuid.UUID `json:"user_id"`
}

func returnPost(dbPost database.Post) Post {
	return Post{
		ID:      dbPost.ID,
		Content: dbPost.Content,
		UserID:  dbPost.UserID,
	}
}

func returnPosts(dbPosts []database.Post) []Post {
	allPosts := []Post{}
	for _, dbUserPost := range dbPosts {
		allPosts = append(allPosts, returnPost(dbUserPost))
	}
	return allPosts
}

func returnUserPosts(dbUserPosts []database.Post) []Post {
	allUserPosts := []Post{}
	for _, dbUserPost := range dbUserPosts {
		allUserPosts = append(allUserPosts, returnPost(dbUserPost))
	}
	return allUserPosts
}

/**LIKE POST MODELS */
type LikedPost struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UserID    uuid.UUID `json:"user_id"`
	PostID    uuid.UUID `json:"post_id"`
}

func returnLikedPost(dbLikedPost database.LikedPost) LikedPost {
	return LikedPost{
		ID:        dbLikedPost.ID,
		CreatedAt: dbLikedPost.CreatedAt,
		UserID:    dbLikedPost.UserID,
		PostID:    dbLikedPost.PostID,
	}
}

func returnAllLikedPosts(dbLikedPosts []database.LikedPost) []LikedPost {
	allUserLikes := []LikedPost{}
	for _, dbLikedPost := range dbLikedPosts {
		allUserLikes = append(allUserLikes, returnLikedPost(dbLikedPost))
	}
	return allUserLikes
}

func returnPostLikes(dbLikedPosts []database.LikedPost) []LikedPost {
	allPostLikes := []LikedPost{}
	for _, dbLikedPost := range dbLikedPosts {
		allPostLikes = append(allPostLikes, returnLikedPost(dbLikedPost))
	}
	return allPostLikes
}

/**FOLLOW MODELS */
type Follow struct {
	ID         uuid.UUID `json:"id"`
	UserID     uuid.UUID `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	FollowedID uuid.UUID `json:"followed_id"`
}

func returnFollow(dbFollow database.Follow) Follow {
	return Follow{
		ID:         dbFollow.ID,
		UserID:     dbFollow.UserID,
		CreatedAt:  dbFollow.CreatedAt,
		FollowedID: dbFollow.FollowedID,
	}
}

func returnFollowing(dbFollowing []database.Follow) []Follow {
	following := []Follow{}
	for _, dbFollow := range dbFollowing {
		following = append(following, returnFollow(dbFollow))
	}
	return following
}

func returnFollowers(dbFollowers []database.Follow) []Follow {
	followers := []Follow{}
	for _, dbFollower := range dbFollowers {
		followers = append(followers, returnFollow(dbFollower))
	}
	return followers
}
