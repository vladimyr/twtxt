package internal

import (
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

const (
	blogsCacheFile = "blogscache"
)

// Cache key: BlogPost.Hash() -> BlogPost
type BlogsCache map[string]*BlogPost

// Store ...
func (cache BlogsCache) Store(path string) error {
	b := new(bytes.Buffer)
	enc := gob.NewEncoder(b)
	err := enc.Encode(cache)
	if err != nil {
		log.WithError(err).Error("error encoding cache")
		return err
	}

	f, err := os.OpenFile(filepath.Join(path, blogsCacheFile), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.WithError(err).Error("error opening cache file for writing")
		return err
	}

	defer f.Close()

	if _, err = f.Write(b.Bytes()); err != nil {
		log.WithError(err).Error("error writing cache file")
		return err
	}
	return nil
}

// LoadBlogsCache ...
func LoadBlogsCache(path string) (BlogsCache, error) {
	cache := make(BlogsCache)

	f, err := os.Open(filepath.Join(path, blogsCacheFile))
	if err != nil {
		if !os.IsNotExist(err) {
			log.WithError(err).Error("error loading cache, cache not found")
			return nil, err
		}
		return cache, nil
	}
	defer f.Close()

	dec := gob.NewDecoder(f)
	err = dec.Decode(&cache)
	if err != nil {
		log.WithError(err).Error("error decoding cache")
		return nil, err
	}
	return cache, nil
}

// UpdateBlogs ...
func (cache BlogsCache) UpdateBlogs(conf *Config) {
	blogPosts, err := GetAllBlogPosts(conf)
	if err != nil {
		log.WithError(err).Error("error getting all blog posts")
		return
	}

	for _, blogPost := range blogPosts {
		cache.Add(blogPost)
	}
}

// Add ...
func (cache BlogsCache) Add(blogPost *BlogPost) {
	cache[blogPost.Hash()] = blogPost
}

// Get ...
func (cache BlogsCache) Get(hash string) (*BlogPost, bool) {
	blogPost, ok := cache[hash]
	return blogPost, ok
}

// GetAll ...
func (cache BlogsCache) GetAll() BlogPosts {
	var blogPosts BlogPosts
	for _, blogPost := range cache {
		blogPosts = append(blogPosts, blogPost)
	}
	return blogPosts
}
