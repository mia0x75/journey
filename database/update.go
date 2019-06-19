package database

import (
	"time"
)

const stmtUpdatePost = "UPDATE posts SET title = ?, slug = ?, markdown = ?, html = ?, featured = ?, page = ?, status = ?, meta_description = ?, image = ?, updated_at = ?, updated_by = ? WHERE id = ?"
const stmtUpdatePostPublished = "UPDATE posts SET title = ?, slug = ?, markdown = ?, html = ?, featured = ?, page = ?, status = ?, meta_description = ?, image = ?, updated_at = ?, updated_by = ?, published_at = ?, published_by = ? WHERE id = ?"
const stmtUpdateSettings = "UPDATE settings SET value = ?, updated_at = ?, updated_by = ? WHERE key = ?"
const stmtUpdateUser = "UPDATE users SET name = ?, slug = ?, email = ?, image = ?, cover = ?, bio = ?, website = ?, location = ?, updated_at = ?, updated_by = ? WHERE id = ?"
const stmtUpdateLastLogin = "UPDATE users SET last_login = ? WHERE id = ?"
const stmtUpdateUserPassword = "UPDATE users SET password = ?, updated_at = ?, updated_by = ? WHERE id = ?"

// UpdatePost TODO
func UpdatePost(postID int64, title []byte, slug string, markdown []byte, html []byte, featured bool, isPage bool, published bool, metaDescription []byte, image []byte, updatedAt time.Time, updatedBy int64) error {
	currentPost, err := RetrievePostByID(postID)
	if err != nil {
		return err
	}
	status := "draft"
	if published {
		status = "published"
	}
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}
	// If the updated post is published for the first time, add publication date and user
	if published && !currentPost.IsPublished {
		_, err = writeDB.Exec(stmtUpdatePostPublished, title, slug, markdown, html, featured, isPage, status, metaDescription, image, updatedAt, updatedBy, updatedAt, updatedBy, postID)
	} else {
		_, err = writeDB.Exec(stmtUpdatePost, title, slug, markdown, html, featured, isPage, status, metaDescription, image, updatedAt, updatedBy, postID)
	}
	if err != nil {
		writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}

// UpdateSettings TODO
func UpdateSettings(title []byte, description []byte, logo []byte, cover []byte, postsPerPage int64, activeTheme string, navigation []byte, updatedAt time.Time, updatedBy int64) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}
	// Title
	_, err = writeDB.Exec(stmtUpdateSettings, title, updatedAt, updatedBy, "title")
	if err != nil {
		writeDB.Rollback()
		return err
	}
	// Description
	_, err = writeDB.Exec(stmtUpdateSettings, description, updatedAt, updatedBy, "description")
	if err != nil {
		writeDB.Rollback()
		return err
	}
	// Logo
	_, err = writeDB.Exec(stmtUpdateSettings, logo, updatedAt, updatedBy, "logo")
	if err != nil {
		writeDB.Rollback()
		return err
	}
	// Cover
	_, err = writeDB.Exec(stmtUpdateSettings, cover, updatedAt, updatedBy, "cover")
	if err != nil {
		writeDB.Rollback()
		return err
	}
	// PostsPerPage
	_, err = writeDB.Exec(stmtUpdateSettings, postsPerPage, updatedAt, updatedBy, "postsPerPage")
	if err != nil {
		writeDB.Rollback()
		return err
	}
	// ActiveTheme
	_, err = writeDB.Exec(stmtUpdateSettings, activeTheme, updatedAt, updatedBy, "activeTheme")
	if err != nil {
		writeDB.Rollback()
		return err
	}
	// Navigation
	_, err = writeDB.Exec(stmtUpdateSettings, navigation, updatedAt, updatedBy, "navigation")
	if err != nil {
		writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}

// UpdateActiveTheme TODO
func UpdateActiveTheme(activeTheme string, updatedAt time.Time, updatedBy int64) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}
	_, err = writeDB.Exec(stmtUpdateSettings, activeTheme, updatedAt, updatedBy, "activeTheme")
	if err != nil {
		writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}

// UpdateUser TODO
func UpdateUser(userID int64, name []byte, slug string, email []byte, image []byte, cover []byte, bio []byte, website []byte, location []byte, updatedAt time.Time, updatedBy int64) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}
	_, err = writeDB.Exec(stmtUpdateUser, name, slug, email, image, cover, bio, website, location, updatedAt, updatedBy, userID)
	if err != nil {
		writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}

// UpdateLastLogin TODO
func UpdateLastLogin(logInDate time.Time, userID int64) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}
	_, err = writeDB.Exec(stmtUpdateLastLogin, logInDate, userID)
	if err != nil {
		writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}

// UpdateUserPassword TODO
func UpdateUserPassword(userID int64, password string, updatedAt time.Time, updatedBy int64) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}
	_, err = writeDB.Exec(stmtUpdateUserPassword, password, updatedAt, updatedBy, userID)
	if err != nil {
		writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}
