package api

type Query struct {
	r Retriever
}

func NewQuery(r Retriever) *Query {
	return &Query{r: r}
}

func (q *Query) Query(input string) (Record, error) {
	return q.r.Query(input)
}

type Retriever interface {
	Query(input string) (Record, error)
}
