package common

// JSONAPIClient denotes the responsibilities of a JSON API client
type JSONAPIClient interface {
	GetJSON(url string) (map[string]interface{}, error)
}
