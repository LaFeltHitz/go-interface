package api

import (
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type TestDatasource struct{}

func (t *TestDatasource) Query(in string) (Record, error) {
	for _, change := range changes {
		if change.ChangeNumber == in {
			return change, nil
		}
	}
	return nil, ErrRecordNotFound
}

func NewTD() *TestDatasource {
	return &TestDatasource{}
}

var changes = []*Change{
	&Change{ChangeNumber: "CHG000000", Assignee: "DH9514", ApplicationOrService: "edw-infra-sandbox19090000"},
	&Change{ChangeNumber: "CHG000001", Assignee: "DH9514", ApplicationOrService: "project-number-3-1378"},
	&Change{ChangeNumber: "CHG000002", Assignee: "DH9514", ApplicationOrService: "forseti-demo-thd"},
}
