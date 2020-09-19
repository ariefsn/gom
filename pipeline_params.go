package gom

// PipeSortParams = params model for pipe sort multiple
type PipeSortParams struct {
	Field     string
	Ascending bool
}

// Case = case model for pipe switch cases
type Case struct {
	Case *Filter
	Then interface{}
}

// PipeSwitchParams = params model for pipe switch
type PipeSwitchParams struct {
	Cases   []Case
	Default interface{}
}
