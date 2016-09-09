package models

import (
    "time"
    "fmt"
    )

type Movie struct {
  Id int
  Name string
  Path string
  Thumb string
  Views int
  Description string
  Date time.Time
  Tag1 string
  Tag2 string
  Tag3 string
  Tag4 string
}

func (movie Movie) DateString() string {
	return fmt.Sprintf("更新日時：%s",movie.Date.Format("Mon, 02 Jan 2006 15:04:05 MST"))
}
