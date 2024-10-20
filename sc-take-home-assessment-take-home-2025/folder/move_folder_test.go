package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func Test_folder_MoveFolder(t *testing.T) {

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

		{Name: "z", Paths: "z", OrgId: orgID_A},
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

		{Name: "z", Paths: "z", OrgId: orgID_A},

		// org_B
		{Name: "f", Paths: "f", OrgId: orgID_B},

		{Name: "g", Paths: "f.g", OrgId: orgID_B},
		{Name: "h", Paths: "f.g.h", OrgId: orgID_B},
		{Name: "i", Paths: "f.g.i", OrgId: orgID_B},

		{Name: "j", Paths: "f.j", OrgId: orgID_B},

		{Name: "k", Paths: "f.k", OrgId: orgID_B},
		{Name: "b", Paths: "f.k.b", OrgId: orgID_B},
		{Name: "z", Paths: "z", OrgId: orgID_B},
	}

	tests := [...]struct {
		name         string
		orgID        uuid.UUID
		sourceName   string
		dstName      string
		folders      []folder.Folder
		want         []folder.Folder
		wantError    bool
		errorMessage string
	}{

		// VALID INPUTS
		{
			name:       "Moving folder with no children to ancestor",
			orgID:      orgID_A,
			sourceName: "c",
			dstName:    "a",
			folders:    sampleFolders_A,
			want: []folder.Folder{

				{Name: "a", Paths: "a", OrgId: orgID_A},

				{Name: "b", Paths: "a.b", OrgId: orgID_A},
				{Name: "c", Paths: "a.c", OrgId: orgID_A},

				{Name: "be", Paths: "a.be", OrgId: orgID_A},
				{Name: "ce", Paths: "a.be.ce", OrgId: orgID_A},

				{Name: "d", Paths: "a.d", OrgId: orgID_A},
				{Name: "e", Paths: "a.d.e", OrgId: orgID_A},
				{Name: "z", Paths: "z", OrgId: orgID_A},
			},
			wantError: false,
		},
		{
			name:       "Moving folder with no children to different parent",
			orgID:      orgID_A,
			sourceName: "c",
			dstName:    "z",
			folders:    sampleFolders_A,
			want: []folder.Folder{

				{Name: "a", Paths: "a", OrgId: orgID_A},

				{Name: "b", Paths: "a.b", OrgId: orgID_A},
				{Name: "c", Paths: "z.c", OrgId: orgID_A},

				{Name: "be", Paths: "a.be", OrgId: orgID_A},
				{Name: "ce", Paths: "a.be.ce", OrgId: orgID_A},

				{Name: "d", Paths: "a.d", OrgId: orgID_A},
				{Name: "e", Paths: "a.d.e", OrgId: orgID_A},
				{Name: "z", Paths: "z", OrgId: orgID_A},
			},
			wantError: false,
		},

		{
			name:       "Moving folder with children within same folder",
			orgID:      orgID_A,
			sourceName: "d",
			dstName:    "b",
			folders:    sampleFolders_A,
			want: []folder.Folder{

				{Name: "a", Paths: "a", OrgId: orgID_A},

				{Name: "b", Paths: "a.b", OrgId: orgID_A},
				{Name: "c", Paths: "a.c", OrgId: orgID_A},

				{Name: "be", Paths: "a.be", OrgId: orgID_A},
				{Name: "ce", Paths: "a.be.ce", OrgId: orgID_A},

				{Name: "d", Paths: "a.b.d", OrgId: orgID_A},
				{Name: "e", Paths: "a.b.d.e", OrgId: orgID_A},
				{Name: "z", Paths: "z", OrgId: orgID_A},
			},
			wantError: false,
		},

		{
			name:       "Destination also exists within other org",
			orgID:      orgID_A,
			sourceName: "d",
			dstName:    "z",
			folders:    sampleFolders_B,
			want: []folder.Folder{

				// org_A
				{Name: "a", Paths: "a", OrgId: orgID_A},

				{Name: "b", Paths: "a.b", OrgId: orgID_A},
				{Name: "c", Paths: "a.b.c", OrgId: orgID_A},

				{Name: "be", Paths: "a.be", OrgId: orgID_A},
				{Name: "ce", Paths: "a.be.ce", OrgId: orgID_A},

				{Name: "d", Paths: "z.d", OrgId: orgID_A},
				{Name: "e", Paths: "z.d.e", OrgId: orgID_A},

				{Name: "z", Paths: "z", OrgId: orgID_A},

				// org_B
				{Name: "f", Paths: "f", OrgId: orgID_B},

				{Name: "g", Paths: "f.g", OrgId: orgID_B},
				{Name: "h", Paths: "f.g.h", OrgId: orgID_B},
				{Name: "i", Paths: "f.g.i", OrgId: orgID_B},

				{Name: "j", Paths: "f.j", OrgId: orgID_B},

				{Name: "k", Paths: "f.k", OrgId: orgID_B},
				{Name: "b", Paths: "f.k.b", OrgId: orgID_B},
				{Name: "z", Paths: "z", OrgId: orgID_B},
			},
			wantError: false,
		},
		{
			name:       "Source exists in both orgs",
			orgID:      orgID_A,
			sourceName: "b",
			dstName:    "e",
			folders:    sampleFolders_A,
			want: []folder.Folder{

				// org_A
				{Name: "a", Paths: "a", OrgId: orgID_A},

				{Name: "b", Paths: "a.d.e.b", OrgId: orgID_A},
				{Name: "c", Paths: "a.d.e.b.c", OrgId: orgID_A},

				{Name: "be", Paths: "a.be", OrgId: orgID_A},
				{Name: "ce", Paths: "a.be.ce", OrgId: orgID_A},

				{Name: "d", Paths: "a.d", OrgId: orgID_A},
				{Name: "e", Paths: "a.d.e", OrgId: orgID_A},

				{Name: "z", Paths: "z", OrgId: orgID_A},

				// org_B
				{Name: "f", Paths: "f", OrgId: orgID_B},

				{Name: "g", Paths: "f.g", OrgId: orgID_B},
				{Name: "h", Paths: "f.g.h", OrgId: orgID_B},
				{Name: "i", Paths: "f.g.i", OrgId: orgID_B},

				{Name: "j", Paths: "f.j", OrgId: orgID_B},

				{Name: "k", Paths: "f.k", OrgId: orgID_B},
				{Name: "b", Paths: "f.k.b", OrgId: orgID_B},
				{Name: "z", Paths: "z", OrgId: orgID_B},
			},
			wantError: false,
		},

		{
			name:       "Source and Destination exist in both orgs, orgA",
			orgID:      orgID_A,
			sourceName: "b",
			dstName:    "z",
			folders:    sampleFolders_B,
			want: []folder.Folder{

				// org_A
				{Name: "a", Paths: "a", OrgId: orgID_A},

				{Name: "b", Paths: "z.b", OrgId: orgID_A},
				{Name: "c", Paths: "z.b.c", OrgId: orgID_A},

				{Name: "be", Paths: "a.be", OrgId: orgID_A},
				{Name: "ce", Paths: "a.be.ce", OrgId: orgID_A},

				{Name: "d", Paths: "a.d", OrgId: orgID_A},
				{Name: "e", Paths: "a.d.e", OrgId: orgID_A},

				{Name: "z", Paths: "z", OrgId: orgID_A},

				// org_B
				{Name: "f", Paths: "f", OrgId: orgID_B},

				{Name: "g", Paths: "f.g", OrgId: orgID_B},
				{Name: "h", Paths: "f.g.h", OrgId: orgID_B},
				{Name: "i", Paths: "f.g.i", OrgId: orgID_B},

				{Name: "j", Paths: "f.j", OrgId: orgID_B},

				{Name: "k", Paths: "f.k", OrgId: orgID_B},
				{Name: "b", Paths: "f.k.b", OrgId: orgID_B},
				{Name: "z", Paths: "z", OrgId: orgID_B},
			},
			wantError: false,
		},

		{
			name:       "Source and Destination exist in both orgs, orgB",
			orgID:      orgID_B,
			sourceName: "b",
			dstName:    "z",
			folders:    sampleFolders_B,
			want: []folder.Folder{

				// org_A
				{Name: "a", Paths: "a", OrgId: orgID_A},

				{Name: "b", Paths: "a.b", OrgId: orgID_A},
				{Name: "c", Paths: "a.b.c", OrgId: orgID_A},

				{Name: "be", Paths: "a.be", OrgId: orgID_A},
				{Name: "ce", Paths: "a.be.ce", OrgId: orgID_A},

				{Name: "d", Paths: "a.d", OrgId: orgID_A},
				{Name: "e", Paths: "a.d.e", OrgId: orgID_A},

				{Name: "z", Paths: "z", OrgId: orgID_A},

				// org_B
				{Name: "f", Paths: "f", OrgId: orgID_B},

				{Name: "g", Paths: "f.g", OrgId: orgID_B},
				{Name: "h", Paths: "f.g.h", OrgId: orgID_B},
				{Name: "i", Paths: "f.g.i", OrgId: orgID_B},

				{Name: "j", Paths: "f.j", OrgId: orgID_B},

				{Name: "k", Paths: "f.k", OrgId: orgID_B},
				{Name: "b", Paths: "z.b", OrgId: orgID_B},
				{Name: "z", Paths: "z", OrgId: orgID_B},
			},
			wantError: false,
		},

		{
			name:       "Large movement of folder to new folder",
			orgID:      orgID_B,
			sourceName: "f",
			dstName:    "z",
			folders:    sampleFolders_B,
			want: []folder.Folder{

				// org_A
				{Name: "a", Paths: "a", OrgId: orgID_A},

				{Name: "b", Paths: "a.b", OrgId: orgID_A},
				{Name: "c", Paths: "a.b.c", OrgId: orgID_A},

				{Name: "be", Paths: "a.be", OrgId: orgID_A},
				{Name: "ce", Paths: "a.be.ce", OrgId: orgID_A},

				{Name: "d", Paths: "a.d", OrgId: orgID_A},
				{Name: "e", Paths: "a.d.e", OrgId: orgID_A},

				{Name: "z", Paths: "z", OrgId: orgID_A},

				// org_B
				{Name: "f", Paths: "z.f", OrgId: orgID_B},

				{Name: "g", Paths: "z.f.g", OrgId: orgID_B},
				{Name: "h", Paths: "z.f.g.h", OrgId: orgID_B},
				{Name: "i", Paths: "z.f.g.i", OrgId: orgID_B},

				{Name: "j", Paths: "z.f.j", OrgId: orgID_B},

				{Name: "k", Paths: "z.f.k", OrgId: orgID_B},
				{Name: "b", Paths: "z.f.b", OrgId: orgID_B},
				{Name: "z", Paths: "z", OrgId: orgID_B},
			},
			wantError: false,
		},

		// INVALID TEST INPUTS

		{
			name:       "Moving folder with children within same folder",
			orgID:      orgID_A,
			sourceName: "d",
			dstName:    "b",
			folders:    sampleFolders_A,
			want:       []folder.Folder{},
			wantError:  false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder(tt.sourceName, tt.dstName, tt.orgID)

		})
	}
}
