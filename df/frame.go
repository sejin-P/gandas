package df

type FileFormat interface {
	ToCSV(path string, sep string) error
	FromDict()
	FillNa(value interface{}, axis int)
}

type Dataframe struct {
	Columns []string
	IdxNum  int
}
