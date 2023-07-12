package models

type ConfModel struct {
	Links []string `json:"rss"`
	Timer int      `json:"request_period"`
}
