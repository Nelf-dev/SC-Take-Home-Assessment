package folders_test

//"github.com/georgechieng-sc/interns-2022/folders"
//"github.com/stretchr/testify/assert"

var req = &folders.FetchFolderRequest{
	OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
}

func Test_GetAllFolders(t *testing.T) {
	//Checks if struct outputs every value correctly
	t.Run("test", func(t *testing.T) {
		output, _ := folders.GetAllFolders(req)
		fmt.Print(output)
		//I would like to check if each struct in FFR contains information, however I am not sure what to use to check if it is empty
		for index, content in range output {
			// output[index].id, name, orgID,deleted check if empty
		}

	})
}

func Test_FetchAllFoldersByOrgID(t *testing.T) {
	t.Run("test", func(t *testing.T) {

		//testing FetchAllfolders by OrgID if files are being appended into resfolder
		const DefaultOrgID = "c1556e17-b7c0-45a3-a6ae-9546248fb17a"

		orgId := uuid.FromStringOrNil(DefaultOrgID)
		resfolder, _ := folders.FetchAllFoldersByOrgID(orgId)

		if len(resfolder) <= 0 {
			t.Error("Folders unable to append into resfolder")
		} else {
			t.Log("Folders appended into resfolder")
		}
	})
}

