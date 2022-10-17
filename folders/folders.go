package folders

import (
	"github.com/gofrs/uuid"
)

// This function returns the response where filtered data matching ORG ID is put into a struct of memory addresses ready for pretty print.
/*
*** Improvements ***
1. Unused variables name should be removed or replaced with blank identifier as Go will not run. This is likely due to save resources and ensure good coding practice.

2. Descriptive variable names should be used to increase readability. e.g k and v should be named folder and content or key and value instead.

3. Line 40 should be deleted as it is updated in line 42 after
*/
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {
	//These variables are not being used, should be implemented in the codes below
	var (
		err error
		f1  Folder
		fs  []*Folder
	)
	//Slice of structs
	f := []Folder{}
	// r is the resfolder being returned
	r, _ := FetchAllFoldersByOrgID(req.OrgID)
	// k should be replaced with a blank identifier for program to run smoothly
	for k, v := range r {
		//Adding into f the actual value of r iteratively
		f = append(f, *v)
	}
	//a slice of structs
	var fp []*Folder
	// k1 should be replaced with a blank identifier for program to run smoothly
	for k1, v1 := range f {
		//Adding into fp the memory address of the values of f
		fp = append(fp, &v1)
	}
	//This can be deleted as it is updated in the next line
	var ffr *FetchFolderResponse
	//Updating value of Folders in the address of FetchFolderResponse
	ffr = &FetchFolderResponse{Folders: fp}
	//Returns a struct and error type
	return ffr, nil
}

// Finds in orgID with UUID as type from sample.json
// Function returns two types
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	//Returns a struct that includes data from sample.json
	folders := GetSampleData()

	//Empty slice of structs
	resFolder := []*Folder{}
	//Loop through folders
	for _, folder := range folders {
		//If orgIDs match add into resFolder
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
