package models

type Bible struct {
	Id        int
	Book      int
	Chapter   int
	Verse     int
	Scripture string
}

func (b *Bible) TableName() string {
	return "bible"
}

type Titles struct {
	Id        int
	Book      int
	Chapter   int
	Verse     int
	Scripture string
}

func (t *Titles) TableName() string {
	return "titles"
}

type Chapname struct {
	Id       int
	Book     int
	Jianchen string
	Bookname string
	Chnum    int
}

func (c *Chapname) TableName() string {
	return "chapname"
}
