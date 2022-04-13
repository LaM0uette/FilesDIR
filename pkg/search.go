package pkg

type Search struct {

	// Flags
	Cls      bool
	Compiler bool

	//..
	Mode      string
	Word      string
	Ext       string
	PoolSize  int
	Maj       bool
	Devil     bool
	Super     bool
	BlackList bool

	// Search
	SrcPath string
	DstPath string
}

//...
// Functions
func (s *Search) RunSearch() {

	s.initSearch()
}

func (s *Search) initSearch() {
	DrawInitSearch()
}
