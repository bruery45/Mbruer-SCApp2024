package folder_test

import (
	"fmt"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	// generating organisation IDs for testing
	orgID_A := uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")
	orgID_B := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
	orgID_C := uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385")

	// Contains singular organisation
	sampleFolders_A := []folder.Folder{
		{Name: "a", Paths: "a", OrgId: orgID_A},

		{Name: "b", Paths: "a.b", OrgId: orgID_A},
		{Name: "b", Paths: "a.b.c", OrgId: orgID_A},

		{Name: "d", Paths: "a.d", OrgId: orgID_A},
		{Name: "e", Paths: "a.d.e", OrgId: orgID_A},
	}
	// contains two organisations, contains sampleFolders_A
	sampleFolders_B := []folder.Folder{
		// org_A
		{Name: "a", Paths: "a", OrgId: orgID_A},
		{Name: "b", Paths: "a.b", OrgId: orgID_A},
		{Name: "b", Paths: "a.b.c", OrgId: orgID_A},

		{Name: "d", Paths: "a.d", OrgId: orgID_A},
		{Name: "e", Paths: "a.d.e", OrgId: orgID_A},

		// org_B
		{Name: "f", Paths: "f", OrgId: orgID_B},

		{Name: "g", Paths: "f.g", OrgId: orgID_B},
		{Name: "h", Paths: "f.g.h", OrgId: orgID_B},
		{Name: "i", Paths: "f.g.i", OrgId: orgID_B},

		{Name: "j", Paths: "f.j", OrgId: orgID_B},

		{Name: "k", Paths: "f.k", OrgId: orgID_B},
		{Name: "b", Paths: "f.k.b", OrgId: orgID_B},
	}

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
		{
			name:    "sole organisation",
			orgID:   orgID_A,
			folders: sampleFolders_A,
			want:    sampleFolders_A,
		},

		{
			name:    "multiple organisations",
			orgID:   orgID_A,
			folders: sampleFolders_B,
			want:    sampleFolders_A,
		},

		{
			name:    "other organisation",
			orgID:   orgID_B,
			folders: sampleFolders_B,
			want: []folder.Folder{

				{Name: "f", Paths: "f", OrgId: orgID_B},

				{Name: "g", Paths: "f.g", OrgId: orgID_B},
				{Name: "h", Paths: "f.g.h", OrgId: orgID_B},
				{Name: "i", Paths: "f.g.i", OrgId: orgID_B},

				{Name: "j", Paths: "f.j", OrgId: orgID_B},

				{Name: "k", Paths: "f.k", OrgId: orgID_B},
				{Name: "b", Paths: "f.k.b", OrgId: orgID_B},
			},
		},

		{
			name:    "organisation not found",
			orgID:   orgID_C,
			folders: sampleFolders_B,
			want:    []folder.Folder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)

			// matching actual output vs expected output
			assert.ElementsMatch(t, tt.want, get, "Folders do not match")
		})
	}
}

// Tests for GetAllChildFolders

func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	// generating organisation IDs for testing
	orgID_A := uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")
	orgID_B := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
	orgID_C := uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385")

	// Contains singular organisation
	sampleFolders_A := []folder.Folder{
		{Name: "a", Paths: "a", OrgId: orgID_A},

		{Name: "b", Paths: "a.b", OrgId: orgID_A},
		{Name: "c", Paths: "a.b.c", OrgId: orgID_A},

		{Name: "be", Paths: "a.be", OrgId: orgID_A},
		{Name: "ce", Paths: "a.be.ce", OrgId: orgID_A},

		{Name: "d", Paths: "a.d", OrgId: orgID_A},
		{Name: "e", Paths: "a.d.e", OrgId: orgID_A},
	}
	// contains two organisations, contains sampleFolders_A
	sampleFolders_B := []folder.Folder{
		// org_A
		{Name: "a", Paths: "a", OrgId: orgID_A},

		{Name: "b", Paths: "a.b", OrgId: orgID_A},
		{Name: "c", Paths: "a.b.c", OrgId: orgID_A},

		{Name: "be", Paths: "a.be", OrgId: orgID_A},
		{Name: "ce", Paths: "a.be.ce", OrgId: orgID_A},

		{Name: "d", Paths: "a.d", OrgId: orgID_A},
		{Name: "e", Paths: "a.d.e", OrgId: orgID_A},

		// org_B
		{Name: "f", Paths: "f", OrgId: orgID_B},

		{Name: "g", Paths: "f.g", OrgId: orgID_B},
		{Name: "h", Paths: "f.g.h", OrgId: orgID_B},
		{Name: "i", Paths: "f.g.i", OrgId: orgID_B},

		{Name: "j", Paths: "f.j", OrgId: orgID_B},

		{Name: "k", Paths: "f.k", OrgId: orgID_B},
		{Name: "b", Paths: "f.k.b", OrgId: orgID_B},
	}

	// error messages
	fileError := "error: folder does not exist"
	orgError := "error: no such folder in specified organisation"

	tests := [...]struct {
		name         string
		orgID        uuid.UUID
		folderName   string
		folders      []folder.Folder
		want         []folder.Folder
		wantError    bool
		errorMessage string
	}{
		// TODO: your tests here
		{
			name:       "Whole folder",
			orgID:      orgID_A,
			folderName: "a",
			folders:    sampleFolders_A,
			want: []folder.Folder{

				{Name: "b", Paths: "a.b", OrgId: orgID_A},
				{Name: "c", Paths: "a.b.c", OrgId: orgID_A},

				{Name: "be", Paths: "a.be", OrgId: orgID_A},
				{Name: "ce", Paths: "a.be.ce", OrgId: orgID_A},

				{Name: "d", Paths: "a.d", OrgId: orgID_A},
				{Name: "e", Paths: "a.d.e", OrgId: orgID_A},
			},
			wantError: false,
		},

		{
			name:       "Whole folder, with other organisation",
			orgID:      orgID_A,
			folderName: "a",
			folders:    sampleFolders_B,
			want: []folder.Folder{

				{Name: "b", Paths: "a.b", OrgId: orgID_A},
				{Name: "c", Paths: "a.b.c", OrgId: orgID_A},

				{Name: "be", Paths: "a.be", OrgId: orgID_A},
				{Name: "ce", Paths: "a.be.ce", OrgId: orgID_A},

				{Name: "d", Paths: "a.d", OrgId: orgID_A},
				{Name: "e", Paths: "a.d.e", OrgId: orgID_A},
			},
			wantError: false,
		},

		{
			name:       "Singular child, same prefix as another",
			orgID:      orgID_A,
			folderName: "b",
			folders:    sampleFolders_B,
			want: []folder.Folder{

				{Name: "c", Paths: "a.b.c", OrgId: orgID_A},
			},
			wantError: false,
		},

		{
			name:       "Folder in multiple orgs, orgB, no children",
			orgID:      orgID_B,
			folderName: "b",
			folders:    sampleFolders_B,
			want:       []folder.Folder{},
			wantError:  false,
		},

		{
			name:       "Additional test for other folder",
			orgID:      orgID_B,
			folderName: "f",
			folders:    sampleFolders_B,
			want: []folder.Folder{

				{Name: "g", Paths: "f.g", OrgId: orgID_B},
				{Name: "h", Paths: "f.g.h", OrgId: orgID_B},
				{Name: "i", Paths: "f.g.i", OrgId: orgID_B},

				{Name: "j", Paths: "f.j", OrgId: orgID_B},

				{Name: "k", Paths: "f.k", OrgId: orgID_B},
				{Name: "b", Paths: "f.k.b", OrgId: orgID_B},
			},
			wantError: false,
		},

		// invalid inputs

		{
			name:         "folder does not exist",
			orgID:        orgID_B,
			folderName:   "fakeFolder",
			folders:      sampleFolders_A,
			want:         []folder.Folder{},
			wantError:    true,
			errorMessage: fileError,
		},

		{
			name:         "Organisation not presentt",
			orgID:        orgID_C,
			folderName:   "a",
			folders:      sampleFolders_A,
			want:         []folder.Folder{},
			wantError:    true,
			errorMessage: orgError,
		},
		{
			name:         "folder in wrong organisation",
			orgID:        orgID_A,
			folderName:   "h",
			folders:      sampleFolders_B,
			want:         []folder.Folder{},
			wantError:    true,
			errorMessage: orgError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.GetAllChildFolders(tt.orgID, tt.folderName)

			// testing expected error for invalid input
			if tt.wantError {

				// no error
				if err == nil {
					t.Errorf("error Expected")

					// wrong error
				} else if tt.errorMessage != err.Error() {
					t.Errorf("expected error message '%s' \nGot: '%s'", tt.errorMessage, err.Error())
				}
				// testing valid inputs
			} else {

				// unexpected error
				if err != nil {
					t.Errorf("no error expected \nreceived: %s", err.Error())

					// matching actual output vs expected output
				} else {

					assert.ElementsMatch(t, tt.want, get)
				}
			}

		})
	}
}

// feel free to change how the unit test is structured
func Test_folder_getPath(t *testing.T) {
	t.Parallel()

	// generating organisation IDs for testing
	orgID_A := uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")
	orgID_B := uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a")
	orgID_C := uuid.FromStringOrNil("9b4cdb0a-cfea-4f9d-8a68-24f038fae385")

	// Contains singular organisation
	sampleFolders_A := []folder.Folder{
		{Name: "a", Paths: "a", OrgId: orgID_A},

		{Name: "b", Paths: "a.b", OrgId: orgID_A},
		{Name: "b", Paths: "a.b.c", OrgId: orgID_A},

		{Name: "d", Paths: "a.d", OrgId: orgID_A},
		{Name: "e", Paths: "a.d.e", OrgId: orgID_A},
	}
	// contains two organisations, contains sampleFolders_A
	sampleFolders_B := []folder.Folder{
		// org_A
		{Name: "a", Paths: "a", OrgId: orgID_A},
		{Name: "b", Paths: "a.b", OrgId: orgID_A},
		{Name: "b", Paths: "a.b.c", OrgId: orgID_A},

		{Name: "d", Paths: "a.d", OrgId: orgID_A},
		{Name: "e", Paths: "a.d.e", OrgId: orgID_A},

		// org_B
		{Name: "f", Paths: "f", OrgId: orgID_B},

		{Name: "g", Paths: "f.g", OrgId: orgID_B},
		{Name: "h", Paths: "f.g.h", OrgId: orgID_B},
		{Name: "i", Paths: "f.g.i", OrgId: orgID_B},

		{Name: "j", Paths: "f.j", OrgId: orgID_B},

		{Name: "k", Paths: "f.k", OrgId: orgID_B},
		{Name: "b", Paths: "f.k.b", OrgId: orgID_B},
	}

	tests := [...]struct {
		name             string
		folderName       string
		orgID            uuid.UUID
		folders          []folder.Folder
		want             string
		wantInOrg        bool
		wantFolderExists bool
	}{
		// valid inputs
		{name: "File path of root",
			folderName:       "a",
			orgID:            orgID_A,
			folders:          sampleFolders_A,
			want:             "a",
			wantInOrg:        true,
			wantFolderExists: true,
		},

		{name: "File path of root, multiple orgs",
			folderName:       "a",
			orgID:            orgID_A,
			folders:          sampleFolders_B,
			want:             "a",
			wantInOrg:        true,
			wantFolderExists: true,
		},

		{name: "File path of subfolder",
			folderName:       "e",
			orgID:            orgID_A,
			folders:          sampleFolders_B,
			want:             "a.d.e",
			wantInOrg:        true,
			wantFolderExists: true,
		},

		{name: "Subfolder in both organisations",
			folderName:       "b",
			orgID:            orgID_B,
			folders:          sampleFolders_B,
			want:             "f.k.b",
			wantInOrg:        true,
			wantFolderExists: true,
		},

		{name: "Subfolder in both organisations",
			folderName:       "b",
			orgID:            orgID_B,
			folders:          sampleFolders_B,
			want:             "f.k.b",
			wantInOrg:        true,
			wantFolderExists: true,
		},

		// invalid input
		{name: "Folder does not exist",
			folderName:       "z",
			orgID:            orgID_B,
			folders:          sampleFolders_B,
			want:             "",
			wantInOrg:        false,
			wantFolderExists: false,
		},

		{name: "Folder exists in other org",
			folderName:       "a",
			orgID:            orgID_B,
			folders:          sampleFolders_B,
			want:             "",
			wantInOrg:        false,
			wantFolderExists: true,
		},

		{name: "Organisation does not exist",
			folderName:       "b",
			orgID:            orgID_C,
			folders:          sampleFolders_B,
			want:             "",
			wantInOrg:        false,
			wantFolderExists: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			path, getOrg, getFolderExists := f.GetPath(tt.orgID, tt.folderName)

			orgErrorMessage := fmt.Sprintf("Expected inOrg: %t, received %t", tt.wantInOrg, getOrg)
			folderErrorMessage := fmt.Sprintf("Expected folderExists: %t, received %t",
				tt.wantFolderExists, getFolderExists)

			// matching actual output vs expected output
			assert.Equal(t, tt.want, path, "Folders do not match")
			assert.Equal(t, tt.wantInOrg, getOrg, orgErrorMessage)
			assert.Equal(t, tt.wantFolderExists, getFolderExists, folderErrorMessage)
		})
	}
}
