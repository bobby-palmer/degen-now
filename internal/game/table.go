package game

type Table struct{}

func NewTable() *Table {
	return &Table{}
}

func (t *Table) Join(name string, stack int64) error {
	return nil // TODO
}
