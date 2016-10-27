package base

//Paginator paginator
type Paginator struct {
	Href       string
	PageNo     int64
	PageSize   int64
	PrevPageNo int64
	NextPageNo int64
	Pages      []int64
	List       interface{}
}

//NewPaginator new paginator
func NewPaginator(href string, count int64, pageNo int64, pageSize int64, list interface{}) *Paginator {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	prev := pageNo - 1
	if prev <= 0 {
		prev = 1
	}
	next := pageNo + 1
	if next >= tp {
		next = tp
	}
	var pages []int64
	for i := int64(1); i <= tp; i++ {
		pages = append(pages, i)
	}
	return &Paginator{
		Href:     href,
		PageNo:   pageNo,
		PageSize: pageSize,
		// TotalPage:  tp,
		// TotalCount: count,
		Pages:      pages,
		PrevPageNo: prev,
		NextPageNo: next,
		List:       list,
	}
}
