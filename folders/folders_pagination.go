package folders

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
)

/*
For this component, I opted for cursor-based pagination because it enables us to directly access the next
set of data without needing to perform a linear search to determine the starting point.

FetchFoldersRequestPaginated: This struct represents the request for fetching paginated folders. It includes:
- `Cursor`: A string representing the starting point for the pagination. It is typically encoded and decoded to ensure it is handled correctly.
- `Limit`: The maximum number of folders to be returned in the response.
- `OrgID`: The UUID representing the organization whose folders are being fetched.

FetchFoldersResponsePaginated: This struct represents the response of the paginated fetch operation. It includes:
- `NextCursor`: A string that represents the cursor for fetching the next set of folders.
- `Folders`: A slice of `Folder` pointers, which contains the folders fetched in the current request.

### Functionality:

GetAllFoldersPaginated: This function handles the paginated retrieval of folders. It starts by validating the request, checking the `Limit`,
and decoding the `Cursor` to determine the starting point in the folder list.
- The `FetchAllFoldersByOrgID` function retrieves all folders for the given organization ID.
- The function then calculates the `end` index, ensuring it does not exceed the total number of folders available.
- If there are more folders available beyond the current `end`, the `NextCursor` is calculated using the `EncodeCursor` function.
- The function returns a paginated response with the folders and the `NextCursor` for subsequent requests.

EncodeCursor: Encodes the index of the next starting point into a cursor string using base64 encoding.
This cursor is used in the response to indicate where the next batch of data should start.

DecodeCursor: Decodes the cursor string from the request to retrieve the starting index for the current pagination.
If the cursor is invalid or empty, it defaults to the beginning.
*/

type FetchFoldersRequestPaginated struct {
	Cursor string
	Limit  int
	OrgID  uuid.UUID
}

type FetchFoldersResponsePaginated struct {
	NextCursor string
	Folders    []*Folder
}

// Copy over the `GetFolders` and `FetchAllFoldersByOrgID` to get started
func GetAllFoldersPaginated(req *FetchFoldersRequestPaginated) (*FetchFoldersResponsePaginated, error) {
	if req == nil {
		return nil, NewFetchFolderError(ErrInvalidRequest)
	}

	if req.Limit <= 0 {
		return nil, NewFetchFolderError(ErrInvalidLimit)
	}

	start, err := DecodeCursor(req.Cursor)
	if err != nil {
		return nil, NewFetchFolderError(ErrInvalidCursor)
	}

	folders, err := FetchAllFoldersByOrgID(req.OrgID)
	if err != nil {
		return nil, err
	}

	end := start + req.Limit
	if end > len(folders) {
		end = len(folders)
	}

	nextCursor := ""
	if end != len(folders) {
		nextCursor = EncodeCursor(end)
	}

	return &FetchFoldersResponsePaginated{Folders: folders[start:end], NextCursor: nextCursor}, nil
}

func EncodeCursor(index int) string {
	return base64.StdEncoding.EncodeToString([]byte("next cursor:" + strconv.Itoa(index)))
}

func DecodeCursor(cursor string) (int, error) {
	if cursor == "" {
		return 0, nil
	}

	decodedCursor, err := base64.StdEncoding.DecodeString(cursor)

	if err != nil {
		return 0, err
	}
	index, err := strconv.Atoi(strings.Split(string(decodedCursor), ":")[1])

	if err != nil {
		return 0, err
	}

	return index, nil
}
