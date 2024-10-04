package graph

import (
	"bbs-gql-project/graph/model"
	"bbs-gql-project/models"
	"context"
	"strconv"
)

// 新規投稿作成のリゾルバ
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	if input.Title == "" {
		return nil, models.BadRequestError("title is required", "title is required")
	}
	if input.Content == "" {
		return nil, models.BadRequestError("content is required", "content is required")
	}

	newID := len(models.Posts) + 1
	newPost := models.Post{
		ID:      newID,
		Title:   input.Title,
		Content: input.Content,
	}
	models.Posts = append(models.Posts, newPost)
	return &model.Post{
		ID:      strconv.Itoa(newPost.ID),
		Title:   newPost.Title,
		Content: newPost.Content,
	}, nil
}

// 投稿の更新のリゾルバ
func (r *mutationResolver) UpdatePost(ctx context.Context, id string, input model.UpdatePost) (*model.Post, error) {
	postID, err := strconv.Atoi(id)
	if err != nil {
		return nil, models.BadRequestError("invalid ID format", "invalid ID format")
	}
	for i, post := range models.Posts {
		if post.ID == postID {
			models.Posts[i].Title = input.Title
			models.Posts[i].Content = input.Content
			return &model.Post{
				ID:      strconv.Itoa(models.Posts[i].ID),
				Title:   models.Posts[i].Title,
				Content: models.Posts[i].Content,
			}, nil
		}
	}
	return nil, models.NotFoundError("post not found", "post not found")
}

// 投稿の削除のリゾルバ
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	postID, err := strconv.Atoi(id)
	if err != nil {
		return false, models.BadRequestError("invalid ID format", "invalid ID format")
	}
	for i, post := range models.Posts {
		if post.ID == postID {
			models.Posts = append(models.Posts[:i], models.Posts[i+1:]...)
			return true, nil
		}
	}
	return false, models.NotFoundError("post not found", "post not found")
}

// 投稿一覧取得のリゾルバ
func (r *queryResolver) GetAllPosts(ctx context.Context, page int, perPage int) ([]*model.Post, error) {
	start := (page - 1) * (perPage)
	end := start + (perPage)

	if start >= len(models.Posts) {
		return []*model.Post{}, nil
	}
	if end > len(models.Posts) {
		end = len(models.Posts)
	}

	result := models.Posts[start:end]
	var postPointers []*model.Post
	for _, post := range result {
		postPointers = append(postPointers, &model.Post{
			ID:      strconv.Itoa(post.ID),
			Title:   post.Title,
			Content: post.Content,
		})
	}

	return postPointers, nil
}

// 投稿の詳細取得のリゾルバ
func (r *queryResolver) GetPost(ctx context.Context, id string) (*model.Post, error) {
	postID, err := strconv.Atoi(id)
	if err != nil {
		return nil, models.BadRequestError("invalid ID format", "invalid ID format")
	}

	for _, post := range models.Posts {
		if post.ID == postID {
			return &model.Post{
				ID:      strconv.Itoa(post.ID),
				Title:   post.Title,
				Content: post.Content,
			}, nil
		}
	}
	return nil, models.NotFoundError("post not found", "post not found")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
