package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFoldersPaginated_Error(t *testing.T) {
	t.Run("Error: Invalid Limit", func(t *testing.T) {
		defaultOrgID := uuid.FromStringOrNil(folders.DefaultOrgID)

		res, err := folders.GetAllFoldersPaginated(
			&folders.FetchFoldersRequestPaginated{
				Cursor: "",
				Limit:  -1,
				OrgID:  defaultOrgID,
			})

		var fetchErr *folders.FetchFolderError
		// Check that error is ErrInvalidLimit
		if assert.ErrorAs(t, err, &fetchErr) {
			assert.Equal(t, folders.ErrInvalidLimit, fetchErr.ErrCode)
		}
		assert.Nil(t, res)
	})

	t.Run("Error: Invalid Cursor", func(t *testing.T) {
		defaultOrgID := uuid.FromStringOrNil(folders.DefaultOrgID)

		res, err := folders.GetAllFoldersPaginated(
			&folders.FetchFoldersRequestPaginated{
				Cursor: "abc123",
				Limit:  5,
				OrgID:  defaultOrgID,
			})

		var fetchErr *folders.FetchFolderError
		// Check that error is ErrInvalidCursor
		if assert.ErrorAs(t, err, &fetchErr) {
			assert.Equal(t, folders.ErrInvalidCursor, fetchErr.ErrCode)
		}
		assert.Nil(t, res)
	})

	t.Run("Error: request is `nil`", func(t *testing.T) {
		res, err := folders.GetAllFoldersPaginated(nil)
		var fetchErr *folders.FetchFolderError
		// Check that error is ErrInvalidRequest
		if assert.ErrorAs(t, err, &fetchErr) {
			assert.Equal(t, folders.ErrInvalidRequest, fetchErr.ErrCode)
		}
		assert.Nil(t, res)
	})

	t.Run("Error: orgID is `nil`", func(t *testing.T) {
		res, err := folders.GetAllFoldersPaginated(
			&folders.FetchFoldersRequestPaginated{
				Cursor: "",
				Limit:  5,
				OrgID:  uuid.Nil,
			})
		var fetchErr *folders.FetchFolderError
		// Check that error is ErrInvalidUUID
		if assert.ErrorAs(t, err, &fetchErr) {
			assert.Equal(t, folders.ErrInvalidUUID, fetchErr.ErrCode)
		}
		assert.Nil(t, res)
	})
}

func Test_GetAllFoldersPaginated_Success(t *testing.T) {
	t.Run("Success: Valid Request for first 5 folders", func(t *testing.T) {
		defaultOrgID := uuid.FromStringOrNil(folders.DefaultOrgID)

		res, err := folders.GetAllFoldersPaginated(
			&folders.FetchFoldersRequestPaginated{
				Cursor: "",
				Limit:  5,
				OrgID:  defaultOrgID,
			})

		assert.NoError(t, err)
		expected, _ := folders.FetchAllFoldersByOrgID(defaultOrgID)
		assert.Equal(t, expected[0:5], res.Folders)
	})

	t.Run("Success: Two Valid Request back to back", func(t *testing.T) {
		defaultOrgID := uuid.FromStringOrNil(folders.DefaultOrgID)

		// Fetch first 5 folders
		res, err := folders.GetAllFoldersPaginated(
			&folders.FetchFoldersRequestPaginated{
				Cursor: "",
				Limit:  5,
				OrgID:  defaultOrgID,
			})

		assert.NoError(t, err)
		expected, _ := folders.FetchAllFoldersByOrgID(defaultOrgID)
		assert.Equal(t, expected[0:5], res.Folders)

		// Fetch next 4 folders
		res2, err := folders.GetAllFoldersPaginated(
			&folders.FetchFoldersRequestPaginated{
				Cursor: res.NextCursor,
				Limit:  4,
				OrgID:  defaultOrgID,
			})

		assert.NoError(t, err)
		assert.Equal(t, expected[5:9], res2.Folders)
	})

	t.Run("Success: Valid request for last 3 folders", func(t *testing.T) {
		defaultOrgID := uuid.FromStringOrNil(folders.DefaultOrgID)

		expected, _ := folders.FetchAllFoldersByOrgID(defaultOrgID)
		cursor := folders.EncodeCursor(len(expected) - 3)

		res, err := folders.GetAllFoldersPaginated(&folders.FetchFoldersRequestPaginated{
			OrgID:  defaultOrgID,
			Limit:  5,
			Cursor: cursor,
		})

		assert.NoError(t, err)
		assert.Equal(t, expected[len(expected)-3:], res.Folders)
	})

}
