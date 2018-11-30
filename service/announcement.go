package service

import "log"

var (
	ann Announcement
)

func AdList(GroupType string) (Announcement, error) {

	props, err := GetProperties()
	log.Println("##########")

	if err != nil {
		log.Println("Error found.")
	} else {
		ann.Properties = props
		ann.Attr.Total = len(props)
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
