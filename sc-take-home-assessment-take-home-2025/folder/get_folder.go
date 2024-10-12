package folder

import (
	"github.com/gofrs/uuid"

	"strings"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

	// return []Folder{}
}

func (f *driver) 


func isChild(parentPath string, childCandidate folder) bool {

	// checking if candidate child is the parent
	if childCandidate.path == parentPath {
		return false
	}

	else if 
}
