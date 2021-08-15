package data

import (
	"fmt"
)

type Query string
type QueryHistory struct {
	Queries []Query
	Index   int
}

func (queryHistory *QueryHistory) AddQuery(query Query) {
	queryHistory.Queries = append(queryHistory.Queries, query)
	queryHistory.Index = len(queryHistory.Queries) - 1
}
func (queryHistory *QueryHistory) HasQuery() bool {
	if len(queryHistory.Queries) > 0 {
		return true
	}
	return false
}
func (queryHistory *QueryHistory) GetQuerySize() int {
	return len(queryHistory.Queries)
}
func (queryHistory *QueryHistory) GetQuery() (Query, error) {
	if !queryHistory.HasQuery() {
		return "", fmt.Errorf("Please run a query before this action")
	}
	return queryHistory.Queries[queryHistory.Index], nil
}
func (queryHistory *QueryHistory) GoToNextQuery() error {
	return queryHistory.goToIndex(queryHistory.Index + 1)
}
func (queryHistory *QueryHistory) GoToPreviousQuery() error {
	return queryHistory.goToIndex(queryHistory.Index - 1)
}
func (queryHistory *QueryHistory) goToIndex(index int) error {
	if queryHistory.GetQuerySize() == 0 {
		return fmt.Errorf("No query found in the history")
	}
	if index >= len(queryHistory.Queries) {
		queryHistory.Index = len(queryHistory.Queries) - 1
		return fmt.Errorf("You are on the last query")
	} else if index < 0 {
		queryHistory.Index = 0
		return fmt.Errorf("You are on the first query")
	} else {
		queryHistory.Index = index
		return nil
	}
}
