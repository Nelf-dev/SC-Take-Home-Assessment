package folders

import (
	"fmt"

	"github.com/gofrs/uuid"
)

/*The methods from folders.go are not working as supposed as the unused variables stop the programming from running correctly. I replaced the for loop variables with a blank identifier and removed the variables at the start as they were not being used.

When I run the program, it only outputs the last folder with the name "correct-scalphunter" as there may be a bug in the for loops as it is not updating the correct memory address and only using the last.
*/

// *Returns the FFR with a slice of structs which acts as chunks allowing to be more readable *//
func Pagination(req *FetchFolderRequest) [][]Folder {
	f := pGetAllFolders(req)

	//created new type that takes in nested slices
	var s = [][]Folder{}

	//looping through f and allocated then in nested slices of 50 folders each
	for i := 0; i <= len(f)-1; i += 50 {
		s = append(s, f[i:i+49])
	}
	fmt.Println(s)
	return s
	//how to make res countable so it can be used in a for loop
}

// I have tweaked this function to return each folder with the same org ID, It seems to have stopped the bug where only the last folder was being printed out.
func pGetAllFolders(req *FetchFolderRequest) []Folder {

	f := []Folder{}
	r, _ := pFetchAllFoldersByOrgID(req.OrgID)

	for _, v := range r {
		f = append(f, *v)
	}

	return f
}

// copied from folders.go
func pFetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
