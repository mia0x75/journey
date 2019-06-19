package methods

import (
	"log"

	"github.com/mia0x75/pages/database"
	"github.com/mia0x75/pages/date"
	"github.com/mia0x75/pages/structure"
)

// SavePost TODO
func SavePost(p *structure.Post) error {
	tagIds := make([]int64, 0)
	// Insert tags
	for _, tag := range p.Tags {
		// Tag slug might already be in database
		tagID, err := database.RetrieveTagIDBySlug(tag.Slug)
		if err != nil {
			// Tag is probably not in database yet
			tagID, err = database.InsertTag(tag.Name, tag.Slug, date.GetCurrentTime(), p.Author.ID)
			if err != nil {
				return err
			}
		}
		if tagID != 0 {
			tagIds = append(tagIds, tagID)
		}
	}
	// Insert post
	postID, err := database.InsertPost(p.Title, p.Slug, p.Markdown, p.HTML, p.IsFeatured, p.IsPage, p.IsPublished, p.MetaDescription, p.Image, *p.Date, p.Author.ID)
	if err != nil {
		return err
	}
	// Insert postTags
	for _, tagID := range tagIds {
		err = database.InsertPostTag(postID, tagID)
		if err != nil {
			return err
		}
	}
	// Generate new global blog
	err = GenerateBlog()
	if err != nil {
		log.Panic("Error: couldn't generate blog data:", err)
	}
	return nil
}

// UpdatePost TODO
func UpdatePost(p *structure.Post) error {
	tagIds := make([]int64, 0)
	// Insert tags
	for _, tag := range p.Tags {
		// Tag slug might already be in database
		tagID, err := database.RetrieveTagIDBySlug(tag.Slug)
		if err != nil {
			// Tag is probably not in database yet
			tagID, err = database.InsertTag(tag.Name, tag.Slug, date.GetCurrentTime(), p.Author.ID)
			if err != nil {
				return err
			}
		}
		if tagID != 0 {
			tagIds = append(tagIds, tagID)
		}
	}
	// Update post
	err := database.UpdatePost(p.ID, p.Title, p.Slug, p.Markdown, p.HTML, p.IsFeatured, p.IsPage, p.IsPublished, p.MetaDescription, p.Image, *p.Date, p.Author.ID)
	if err != nil {
		return err
	}
	// Delete old postTags
	err = database.DeletePostTagsForPostID(p.ID)
	// Insert postTags
	if err != nil {
		return err
	}
	for _, tagID := range tagIds {
		err = database.InsertPostTag(p.ID, tagID)
		if err != nil {
			return err
		}
	}
	// Generate new global blog
	err = GenerateBlog()
	if err != nil {
		log.Panic("Error: couldn't generate blog data:", err)
	}
	return nil
}

// DeletePost TODO
func DeletePost(postID int64) error {
	err := database.DeletePostByID(postID)
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
