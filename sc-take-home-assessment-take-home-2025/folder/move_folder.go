package folder

import (
	"errors"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	/*
		As the method signature does not contain orgIDs, we assume that only one
		instance of name and dst occur within the file structure.

		Otherwise, we would implement explicit checking when getting their folders
	*/

	// checking if moving folder into itself
	if name == dst {

		return nil, errors.New("Error: Cannot move a folder to itself")
	}

	// getting desired folders
	destination, destExists := f.getFolder(dst)
	source, sourceExists := f.getFolder(name)

	// checking if destination folder exists

	if !destExists {
		return nil, errors.New("Error: Destination folder does not exist")
	}

	// checking if source folder exists

	if !sourceExists {
		return nil, errors.New("Error: Source folder does not exist")
	}

	// checking if orgIDs match

	if source.OrgId != destination.OrgId {

		return nil, errors.New("Error: Cannot move a folder to a different organization")
	}

	// folders := f.folders

	// res := []Folder{}

}

func (f *driver) getFolder(name string) (Folder, bool) {

	found := false
	folders := f.folders

	var folder Folder

	for _, item := range folders {

		//
		if item.Name == name {
			found = true
			folder = item
		}
	}

	return folder, found
}
