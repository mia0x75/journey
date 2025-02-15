package methods

import (
	"encoding/json"
	"log"

	"github.com/mia0x75/pages/configuration"
	"github.com/mia0x75/pages/database"
	"github.com/mia0x75/pages/date"
	"github.com/mia0x75/pages/slug"
	"github.com/mia0x75/pages/structure"
)

// Blog Global blog - thread safe and accessible by all requests
var Blog *structure.Blog

var assetPath = []byte("/assets/")

// UpdateBlog TODO
func UpdateBlog(b *structure.Blog, userID int64) error {
	// Marshal navigation items to json string
	navigation, err := json.Marshal(b.NavigationItems)
	if err != nil {
		return err
	}
	err = database.UpdateSettings(b.Title, b.Description, b.Logo, b.Cover, b.PostsPerPage, b.ActiveTheme, navigation, date.GetCurrentTime(), userID)
	if err != nil {
		return err
	}
	// Generate new global blog
	err = GenerateBlog()
	if err != nil {
		log.Panic("Error: couldn't generate blog data:", err)
	}
	return nil
}

// UpdateActiveTheme TODO
func UpdateActiveTheme(activeTheme string, userID int64) error {
	err := database.UpdateActiveTheme(activeTheme, date.GetCurrentTime(), userID)
	if err != nil {
		return err
	}
	// Generate new global blog
	err = GenerateBlog()
	if err != nil {
		log.Panic("Error: couldn't generate blog data:", err)
	}
	return nil
}

// GenerateBlog TODO
func GenerateBlog() error {
	// Write lock the global blog
	if Blog != nil {
		Blog.Lock()
		defer Blog.Unlock()
	}
	// Generate blog from db
	blog, err := database.RetrieveBlog()
	if err != nil {
		return err
	}
	// Add parameters that are not saved in db
	blog.URL = []byte(configuration.Config.Url)
	blog.AssetPath = assetPath
	// Create navigation slugs
	for index := range blog.NavigationItems {
		blog.NavigationItems[index].Slug = slug.Generate(blog.NavigationItems[index].Label, "navigation")
	}
	Blog = blog
	return nil
}
