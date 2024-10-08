package folders

import (
	"github.com/gofrs/uuid"
)

/*
The `GetAllFolders` function fetches all the folders that are associated with this `OrgID`.
It does so by recieving a `FetchFolderRequest` which contains the OrgID.
On success, a `FetchFolderResponse` is returned.
On failture, an error is returned.

There are some improvements that can be made to this code.

 1. Declared but unused variables
    `GetAllFolders` contains multiple unused variable delcarations. Not only does this
    this make the function difficult to understand, but it also blocks compilation
    as observed when running `go build`.
    Suggestion:
    These unused variables can be removed to improve the clarity of `GetAllFolders`.
    In cases where multiple values are returned, Golang provides the blank identifier
    `_`, which can be used to ignore unneeded values.

 2. Unclear variable names
    Varible names such as `f1`, `fs`, `f`, `r` and `v`, are ambiguous and make it difficult
    to understand what `GetAllFolders` is doing. This can increase the cost associated
    with debugging, or when this function is handed over to a different engineer.
    Suggestion:
    Variable names that clearly convey their purpose should be used. This will enhance the
    overall clarity of `GetAllFolders` and make the code easier to interpret.

 3. Dereferencing and re-referencing of Folder
    `GetAllFolders` performs some unnecessary transforming of the Folders returned by
    `FetchAllFoldersByOrgID`. `FetchAllFoldersByOrgID` returns a slice of Folder pointers
    which is dereferenced in the first for-loop into `f`. It is then re-referenced again
    in the subsequent for-loop into `fp`. This action is already performed by
    `FetchAllFoldersByOrgID`.
    Suggestion:
    Remove both for-loops, and directly use the slice returned by `FetchAllFoldersByOrgID`

 4. Error Handling
    `GetAllFolders` currently does not handle errors from the `FetchAllFoldersByOrgID` function
    which can lead to undefined behavior if the function encounters errors or unexpected data types.
    Additionally, the `req` parameter may be undefined, which is not accounted for by `GetAllFolders`.
    Suggestion:
    Handle errors that could be thrown by `FetchAllFoldersByOrgID`.
    Handle the case where `req` may be `nil`
    Additionally, a custom error class can be created allowing for better maintainability and readbility
    of `GetAllFolders`.

 4. Syntax / style
    The variable `ffr` is declared on one line, and then assigned on another line.
    Solution:
    Use Golang's `:=` operatior to declare and assign the value on a single line.
*/
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	if req == nil {
		return nil, NewFetchFolderError(ErrInvalidRequest)
	}
	folders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}
	return &FetchFolderResponse{Folders: folders}, nil
}

/*
The `FetchAllFoldersByOrgID` function returns all folders that match the `orgID` parameter given.
On success, a slice of Folder pointers is returned.
On failure, an error is returned

Assumption: For the function `GetSampleData()`, there is a possibility that the filepath will not exist, resulting
in a `panic`. It will be assumed then that `FetchAllFoldersByOrgID` has no requirement to  handle this gracefully
and that all inputs for this assessment will be 'valid'. (`sample.json` will always exist)

Assumption: When there are no Folders that match the given `orgID`, an empty slice will be returned.
There is no need to throw an error.

There are some improvements that can be made to this code.

 1. Error Handling
    The `orgID` passed in may be `nil`.
    Suggestion:
    Return an error accordingly if this occurs.
*/
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	if orgID == uuid.Nil {
		return nil, NewFetchFolderError(ErrInvalidUUID)
	}
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
