/*
 * Forum application
 *
 * Forum application api
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package code_help_forum_api

import (
	"context"
	"net/http"
)



// CategoryAPIRouter defines the required methods for binding the api requests to a responses for the CategoryAPI
// The CategoryAPIRouter implementation should parse necessary information from the http request,
// pass the data to a CategoryAPIServicer to perform the required actions, then write the service results to the http response.
type CategoryAPIRouter interface { 
	CreateCategory(http.ResponseWriter, *http.Request)
	GetAllCategories(http.ResponseWriter, *http.Request)
}
// CommentAPIRouter defines the required methods for binding the api requests to a responses for the CommentAPI
// The CommentAPIRouter implementation should parse necessary information from the http request,
// pass the data to a CommentAPIServicer to perform the required actions, then write the service results to the http response.
type CommentAPIRouter interface { 
	CommentOnPost(http.ResponseWriter, *http.Request)
	DeleteComment(http.ResponseWriter, *http.Request)
	GetCommentReplies(http.ResponseWriter, *http.Request)
	GetCommentsForPost(http.ResponseWriter, *http.Request)
	ReplyToComment(http.ResponseWriter, *http.Request)
	UpdateComment(http.ResponseWriter, *http.Request)
}
// CommunityAPIRouter defines the required methods for binding the api requests to a responses for the CommunityAPI
// The CommunityAPIRouter implementation should parse necessary information from the http request,
// pass the data to a CommunityAPIServicer to perform the required actions, then write the service results to the http response.
type CommunityAPIRouter interface { 
	AddANewModeratorToTheCommunity(http.ResponseWriter, *http.Request)
	CreateCommunity(http.ResponseWriter, *http.Request)
	DeleteCommunity(http.ResponseWriter, *http.Request)
	GetAllCommunities(http.ResponseWriter, *http.Request)
	GetCommunityByUid(http.ResponseWriter, *http.Request)
	GetCommunityModerators(http.ResponseWriter, *http.Request)
	JoinCommunity(http.ResponseWriter, *http.Request)
	LeaveCommunity(http.ResponseWriter, *http.Request)
	RemoveModerator(http.ResponseWriter, *http.Request)
	UpdateCommunity(http.ResponseWriter, *http.Request)
}
// PostAPIRouter defines the required methods for binding the api requests to a responses for the PostAPI
// The PostAPIRouter implementation should parse necessary information from the http request,
// pass the data to a PostAPIServicer to perform the required actions, then write the service results to the http response.
type PostAPIRouter interface { 
	CreateCommunityPost(http.ResponseWriter, *http.Request)
	DeletePost(http.ResponseWriter, *http.Request)
	GetPost(http.ResponseWriter, *http.Request)
	GetPosts(http.ResponseWriter, *http.Request)
	UpdatePost(http.ResponseWriter, *http.Request)
}


// CategoryAPIServicer defines the api actions for the CategoryAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type CategoryAPIServicer interface { 
	CreateCategory(context.Context, CategoryCreate) (ImplResponse, error)
	GetAllCategories(context.Context) (ImplResponse, error)
}


// CommentAPIServicer defines the api actions for the CommentAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type CommentAPIServicer interface { 
	CommentOnPost(context.Context, string, CommentRequest) (ImplResponse, error)
	DeleteComment(context.Context, string) (ImplResponse, error)
	GetCommentReplies(context.Context, string) (ImplResponse, error)
	GetCommentsForPost(context.Context, string) (ImplResponse, error)
	ReplyToComment(context.Context, string, CommentRequest) (ImplResponse, error)
	UpdateComment(context.Context, string, CommentRequest) (ImplResponse, error)
}


// CommunityAPIServicer defines the api actions for the CommunityAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type CommunityAPIServicer interface { 
	AddANewModeratorToTheCommunity(context.Context, string, ModeratorRequest) (ImplResponse, error)
	CreateCommunity(context.Context, CommunityRequest) (ImplResponse, error)
	DeleteCommunity(context.Context, string) (ImplResponse, error)
	GetAllCommunities(context.Context) (ImplResponse, error)
	GetCommunityByUid(context.Context, string) (ImplResponse, error)
	GetCommunityModerators(context.Context, string) (ImplResponse, error)
	JoinCommunity(context.Context, string) (ImplResponse, error)
	LeaveCommunity(context.Context, string) (ImplResponse, error)
	RemoveModerator(context.Context, string, string) (ImplResponse, error)
	UpdateCommunity(context.Context, string, CommunityRequest) (ImplResponse, error)
}


// PostAPIServicer defines the api actions for the PostAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type PostAPIServicer interface { 
	CreateCommunityPost(context.Context, string, PostRequest) (ImplResponse, error)
	DeletePost(context.Context, string) (ImplResponse, error)
	GetPost(context.Context, string) (ImplResponse, error)
	GetPosts(context.Context, string) (ImplResponse, error)
	UpdatePost(context.Context, string, PostRequest) (ImplResponse, error)
}