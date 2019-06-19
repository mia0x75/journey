package database

const stmtDeletePostTagsByPostID = "DELETE FROM posts_tags WHERE post_id = ?"
const stmtDeletePostByID = "DELETE FROM posts WHERE id = ?"

// DeletePostTagsForPostID TODO
func DeletePostTagsForPostID(postID int64) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}
	_, err = writeDB.Exec(stmtDeletePostTagsByPostID, postID)
	if err != nil {
		writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}

// DeletePostByID TODO
func DeletePostByID(postID int64) error {
	writeDB, err := readDB.Begin()
	if err != nil {
		writeDB.Rollback()
		return err
	}
	_, err = writeDB.Exec(stmtDeletePostByID, postID)
	if err != nil {
		writeDB.Rollback()
		return err
	}
	return writeDB.Commit()
}
