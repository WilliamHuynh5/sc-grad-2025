package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

/*
Blackbox tests for GetAllFolders Error cases.
*/
func Test_GetAllFolders_Err(t *testing.T) {
	t.Run("Error: request is `nil`", func(t *testing.T) {
		res, err := folders.GetAllFolders(nil)
		var fetchErr *folders.FetchFolderError
		// Check that error is ErrInvalidRequest
		if assert.ErrorAs(t, err, &fetchErr) {
			assert.Equal(t, folders.ErrInvalidRequest, fetchErr.ErrCode)
		}
		assert.Nil(t, res)
	})

	t.Run("Error: orgID is `nil`", func(t *testing.T) {
		res, err := folders.GetAllFolders(
			&folders.FetchFolderRequest{
				OrgID: uuid.Nil,
			})
		var fetchErr *folders.FetchFolderError
		// Check that error is ErrInvalidUUID
		if assert.ErrorAs(t, err, &fetchErr) {
			assert.Equal(t, folders.ErrInvalidUUID, fetchErr.ErrCode)
		}
		assert.Nil(t, res)
	})
}

/*
Blackbox tests for GetAllFolders Success cases.
*/
func Test_GetAllFolders_Success(t *testing.T) {
	t.Run("Success: Valid Request with existing orgID", func(t *testing.T) {
		defaultOrgID := uuid.FromStringOrNil(folders.DefaultOrgID)

		res, err := folders.GetAllFolders(
			&folders.FetchFolderRequest{
				OrgID: defaultOrgID,
			})

		assert.NoError(t, err)
		assert.NotNil(t, res)

		for _, folder := range res.Folders {
			assert.Equal(t, defaultOrgID, folder.OrgId)
		}
	})

	t.Run("Success: Valid Request with non-existing orgID", func(t *testing.T) {
		nonExistingOrgID := uuid.Must(uuid.NewV4())

		res, err := folders.GetAllFolders(
			&folders.FetchFolderRequest{
				OrgID: nonExistingOrgID,
			})

		assert.NoError(t, err)
		assert.Empty(t, res.Folders)
	})
}
