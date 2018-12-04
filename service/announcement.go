package service

import "log"

var (
	ann Announcement
)

func AdList(venture, groupType, orderByPrice string, offSet, limitPerPage int) (Announcement, error) {

	props, tot, err := GetProperties(venture, groupType, orderByPrice, offSet, limitPerPage)

	if err != nil {
		log.Println("Error found.")
	} else {
		ann.Properties = props
		ann.Attr.Total = tot
		ann.Attr.PageSize = limitPerPage
	}
	return ann, nil
}

type (
	Announcement struct {
		Properties []Property `json:"properties"`
		Attr       Attr       `json: "@attr"`
	}

	Attr struct {
		PageSize   int    `json:"pageSize"`
		Total      int    `json:"total"`
	}
)
