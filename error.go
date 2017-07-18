package main

// jsonErr is Error structure for use with JSON.
type jsonErr struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
