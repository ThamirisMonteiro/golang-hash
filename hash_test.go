package hash_test

import (
	"github.com/brianvoe/gofakeit"
	"github.com/ebanx/akkad/pkg/core/hash"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestHash(t *testing.T) {
	suite.Run(t, new(HashTest))
}

type HashTest struct {
	suite.Suite
}

func (t *HashTest) Test_hash_New_IsSame() {
	// Structs and variables
	type emptyStruct struct{}
	type structWithoutExportedFields struct {
		number int
	}
	type structWithExportedFields struct {
		ID       string
		SomeText string
	}
	someID := "someID"
	someText := gofakeit.Word()

	type testCase struct {
		name    string
		struct1 interface{}
		struct2 interface{}
		isSame  bool
	}
	testCases := []testCase{
		{
			name:    "hashes of two nil structs are the same",
			struct1: nil,
			struct2: nil,
			isSame:  true,
		},
		{
			name:    "hashes of two empty structs are the same",
			struct1: emptyStruct{},
			struct2: emptyStruct{},
			isSame:  true,
		},
		{
			name:    "hashes of two empty structs that do not have exported fields are the same",
			struct1: structWithoutExportedFields{},
			struct2: structWithoutExportedFields{},
			isSame:  true,
		},
		{
			name:    "hash of an empty struct is the same as a hash of an struct without exported fields",
			struct1: emptyStruct{},
			struct2: structWithoutExportedFields{},
			isSame:  true,
		},
		{
			name: "hashes of two structs without exported fields are the same regardless of the fields's values",
			struct1: structWithoutExportedFields{
				number: 0,
			},
			struct2: structWithoutExportedFields{
				number: 1,
			},
			isSame: true,
		},
		{
			name:    "hashes of two empty structs with exported fields are the same",
			struct1: structWithExportedFields{},
			struct2: structWithExportedFields{},
			isSame:  true,
		},
		{
			name: "hashes of two structs with exported fields are the same when their values are the same",
			struct1: structWithExportedFields{
				ID:       someID,
				SomeText: someText,
			},
			struct2: structWithExportedFields{
				ID:       someID,
				SomeText: someText,
			},
			isSame: true,
		},
		{
			name:    "hashes of a nil struct and an empty struct are not the same",
			struct1: nil,
			struct2: emptyStruct{},
			isSame:  false,
		},
		{
			name:    "hashes of a nil struct and an empty struct without exported fields are not the same",
			struct1: nil,
			struct2: structWithoutExportedFields{},
			isSame:  false,
		},
		{
			name:    "hashes of a nil struct and an empty struct with exported fields are not the same",
			struct1: nil,
			struct2: structWithExportedFields{},
			isSame:  false,
		},
		{
			name:    "hashes of an empty struct without exported fields and an empty struct with exported fields are not the same",
			struct1: structWithoutExportedFields{},
			struct2: structWithExportedFields{},
			isSame:  false,
		},
		{
			name: "hashes of a struct with exported fields and another struct with exported fields whose fields are different are not the same",
			struct1: structWithExportedFields{
				ID:       "",
				SomeText: gofakeit.Word(),
			},
			struct2: structWithExportedFields{
				ID:       someID,
				SomeText: someText,
			},
			isSame: false,
		},
	}
	for _, tc := range testCases {
		hashStruct1, err := hash.New(tc.struct1)
		t.Require().NoError(err)
		t.NotNil(hashStruct1)

		hashStruct2, err := hash.New(tc.struct2)
		t.Require().NoError(err)
		t.NotNil(hashStruct2)

		t.Equal(tc.isSame, hash.IsSame(hashStruct1, hashStruct2))
	}
}
