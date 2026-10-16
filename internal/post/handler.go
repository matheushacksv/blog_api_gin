package post

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

// PostHandler represents the post repository
type PostHandler struct {
	pool *pgxpool.Pool
}

// NewPostHandler creates a new post handler
func NewPostHandler(pool *pgxpool.Pool) PostHandler {
	return PostHandler{pool}
}

// Create a new post in database
func (p PostHandler) Create(c *gin.Context) {
	var data Post
	erro := c.ShouldBindJSON(&data)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	repository := NewRepositoryPost(p.pool)

	if erro := data.Prepare(); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	id, erro := repository.CreatePost(c.Request.Context(), data)
	if erro != nil {
		log.Println(erro)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create post"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})

}

// Search for a post in database and returns it
func (p PostHandler) Search(c *gin.Context) {
	title := c.Query("title")

	repository := NewRepositoryPost(p.pool)

	posts, erro := repository.SearchPosts(c.Request.Context(), title)
	if erro != nil {
		log.Println(erro)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get posts"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// Get a single post from database
func (p PostHandler) Get(c *gin.Context) {
	param := c.Param("ID")
	postID, erro := strconv.ParseUint(param, 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	repository := NewRepositoryPost(p.pool)
	post, erro := repository.GetAPost(c.Request.Context(), postID)
	if erro != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// Update a post from database
func (p PostHandler) Update(c *gin.Context) {
	var data Post
	param := c.Param("ID")
	erro := c.ShouldBindJSON(&data)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	postID, erro := strconv.ParseUint(param, 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}
	if erro := data.Prepare(); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	repository := NewRepositoryPost(p.pool)
	erro = repository.UpdatePost(c.Request.Context(), postID, data.Title, data.Content)
	if erro != nil {
		log.Println(erro)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "could not reach the service"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

// Delete a post from database
func (p PostHandler) Delete(c *gin.Context) {
	param := c.Param("ID")
	postID, erro := strconv.ParseUint(param, 10, 64)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	repository := NewRepositoryPost(p.pool)
	if erro := repository.DeletePost(c.Request.Context(), postID); erro != nil {
		log.Println(erro)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete item"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
