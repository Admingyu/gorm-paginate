package paginate

import "github.com/jinzhu/gorm"

type PageParams struct {
	Serializer interface{}
	Order      []string
	PageIndex  int
	PageSize   int
}

func Pagenate(query *gorm.DB, p PageParams) interface{} {
	if len(p.Order) >= 1 {
		for _, i := range p.Order {
			query = query.Order(i)
		}
	}
	rows := query.Offset((p.PageIndex - 1) * p.PageSize).Limit(p.PageSize)
	rows.Scan(&p.Serializer)
	return p.Serializer
}
