package service

import "log"

var (
	ann Announcement
)

func AdList(venture, groupType, orderByPrice string, offSet, limitPerPage int) (Announcement, error) {

	props, err := GetProperties(venture, groupType, orderByPrice, offSet, limitPerPage)

	if err != nil {
		log.Println("Error found.")
	} else {
		ann.Properties = props
		ann.Attr.Total = len(props)
		ann.Attr.PerPage = 2
		ann.Attr.PageSize = 0
	}
	return ann, nil
}

type (
	Announcement struct {
		Properties []Property `json:"properties"`
		Attr       Attr       `json: "@attr"`
	}

	Attr struct {
		GroupStore string `json:"groupStore"`
		PageSize   int    `json:"pageSize"`
		PerPage    int    `json:"perPage"`
		Total      int    `json:"total"`
	}
)
