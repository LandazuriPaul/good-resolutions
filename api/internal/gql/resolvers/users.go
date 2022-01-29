package resolvers

import (
	"context"

	log "github.com/LandazuriPaul/good-resolutions/api/pkg/logger"

	"github.com/LandazuriPaul/good-resolutions/api/internal/gql/models"
	tf "github.com/LandazuriPaul/good-resolutions/api/internal/gql/resolvers/transformations"
	dbm "github.com/LandazuriPaul/good-resolutions/api/internal/orm/models"
)

// CreateUser creates a record
func (r *mutationResolver) CreateUser(ctx context.Context, input models.UserInput) (*models.User, error) {
	return userCreateUpdate(r, input, false)
}

// UpdateUser updates a record
func (r *mutationResolver) UpdateUser(ctx context.Context, input models.UpdateUserInput) (*models.User, error) {
	return userCreateUpdate(r, *input.Patch, true, input.ID)
}

// DeleteUser deletes a record
func (r *mutationResolver) DeleteUser(ctx context.Context, input models.DeleteUserInput) (bool, error) {
	return userDelete(r, input.ID)
}

// Users lists records
func (r *queryResolver) Users(ctx context.Context, filter *models.FilterInput) (*models.Users, error) {
	if filter != nil {
		return userList(r, filter.ID)
	}
	return userList(r, nil)
}

/**
 * Helper functions
 */

func userCreateUpdate(r *mutationResolver, input models.UserInput, update bool, ids ...string) (*models.User, error) {
	dbo, err := tf.GQLInputUserToDBUser(&input, update, ids...)
	if err != nil {
		return nil, err
	}
	// Create scoped clean db interface
	db := r.ORM.DB.New().Begin()
	if !update {
		db = db.Create(dbo).First(dbo) // Create the user
	} else {
		db = db.Model(&dbo).Update(dbo).First(dbo) // Or update it
	}
	gql, err := tf.DBUserToGQLUser(dbo)
	if err != nil {
		db.RollbackUnlessCommitted()
		return nil, err
	}
	db = db.Commit()
	return gql, db.Error
}

func userDelete(r *mutationResolver, id string) (bool, error) {
	return false, nil
}

func userList(r *queryResolver, id *string) (*models.Users, error) {
	entity := "users"
	whereID := "id = ?"
	record := &models.Users{}
	dbRecords := []*dbm.User{}
	db := r.ORM.DB.New()
	if id != nil {
		db = db.Where(whereID, *id)
	}
	db = db.Find(&dbRecords).Count(&record.Count)
	for _, dbRec := range dbRecords {
		if rec, err := tf.DBUserToGQLUser(dbRec); err != nil {
			log.Errorfn(entity, err)
		} else {
			record.List = append(record.List, rec)
		}
	}
	return record, db.Error
}
