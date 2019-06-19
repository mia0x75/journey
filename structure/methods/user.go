package methods

import (
	"github.com/mia0x75/pages/database"
	"github.com/mia0x75/pages/date"
	"github.com/mia0x75/pages/structure"
)

// SaveUser TODO
func SaveUser(u *structure.User, hashedPassword string, createdBy int64) error {
	userID, err := database.InsertUser(u.Name, u.Slug, hashedPassword, u.Email, u.Image, u.Cover, date.GetCurrentTime(), createdBy)
	if err != nil {
		return err
	}
	err = database.InsertRoleUser(u.Role, userID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser TODO
func UpdateUser(u *structure.User, userID int64) error {
	err := database.UpdateUser(u.ID, u.Name, u.Slug, u.Email, u.Image, u.Cover, u.Bio, u.Website, u.Location, date.GetCurrentTime(), userID)
	if err != nil {
		return err
	}
	return nil
}
