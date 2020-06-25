package models

type Tag struct {
	Model

	Name       string `json:"name,omitempty"`
	CreatedBy  string `json:"created_by,omitempty"`
	ModifiedBy string `json:"modified_by,omitempty"`
	State      int    `json:"state,omitempty"`
}

func GetTags(PageNum int, pagesize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(PageNum).Limit(pagesize).Find(&tags)
	return tags
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return count
}
