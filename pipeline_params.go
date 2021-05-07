package gom

// PipeSortParams = params model for pipe sort multiple
type PipeSortParams struct {
	Field     string
	Ascending bool
}

// PipeSwitchCaseParams = case model for pipe switch cases
type PipeSwitchCaseParams struct {
	Case *Filter
	Then interface{}
}

// PipeSwitchParams = params model for pipe switch
type PipeSwitchParams struct {
	Cases   []PipeSwitchCaseParams
	Default interface{}
}
