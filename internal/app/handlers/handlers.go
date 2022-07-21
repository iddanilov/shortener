package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/shortener/internal/app/middleware"
	"github.com/shortener/internal/app/models"
	"github.com/shortener/internal/app/storage"
)

type routerGroup struct {
	rg *gin.RouterGroup
	s  *storage.Storage
}

func NewRouterGroup(rg *gin.RouterGroup, s *storage.Storage) *routerGroup {
	return &routerGroup{
		rg: rg,
		s:  s,
	}
}

func (h *routerGroup) Routes() {
	h.rg.GET("/:id", middleware.Middleware(h.GetURLByID))
	h.rg.POST("/", middleware.Middleware(h.SaveURL))
}

func (h *routerGroup) GetURLByID(c *gin.Context) error {
	result, ok := h.s.URL[c.Param("id")]
	if !ok {
		return middleware.ErrNotFound
	}
	c.Redirect(http.StatusTemporaryRedirect, result.URL)

	return nil
}

func (h *routerGroup) SaveURL(c *gin.Context) error {
	var url string

	id := RandStringRunes(10)
	value, err := ioutil.ReadAll(c.Request.Body)

	//_, err := c.Request.Body.Read(value)
	if err != nil {
		return err
	}
	url = string(value)
	if !strings.Contains(url, "http") {
		url = strings.Join([]string{"https://", url}, "")
	}

	h.s.URL[id] = models.URL{URL: url}
	log.Printf("Save url by id %s", id)
	c.Writer.WriteHeader(http.StatusCreated)
	c.Writer.Write([]byte(fmt.Sprintf("http://localhost:8080/%s", id)))

	return nil
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
