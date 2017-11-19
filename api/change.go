package api

import (
	"fmt"
)

type Change struct {
	ChangeNumber         string
	Assignee             string
	ApplicationOrService string
}

func (c *Change) Number() string {
	return c.ChangeNumber
}

func (c *Change) CI() string {
	return c.ApplicationOrService
}

func (c *Change) AssignedTo() string {
	return c.Assignee
}

func (c *Change) String() string {
	return fmt.Sprintf("Number: %s\nAssignee: %s\nApplicationOrService: %s\n", c.ChangeNumber, c.Assignee, c.ApplicationOrService)
}
