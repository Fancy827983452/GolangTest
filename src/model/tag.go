package model

type Tag struct {
	ID   uint32
	Name string
}

var t Tag

func AddTag(name string) error {
	sql := "insert into tbl_tag (name) values(?)"
	_, err := db.Exec(sql, name)
	return err
}

//查找标签，返回标签id数组
func FindTag(names []string) ([]uint32) {
	var tag []uint32
	for _, name := range names {
		query := "select tag_id from tbl_tag tag where tag.name = '" + name + "'"
		rows, err := db.Query(query)
		var tid uint32
		for rows.Next(){
			if err=rows.Scan(&tid);err==nil {
				tag = append(tag, tid)
			} else {
				panic(err)
			}
		}
	}
	return tag
}
