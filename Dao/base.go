package dao

type QueryOption struct {
	Conditions map[string]interface{} // 查询条件
	Preloads   map[string]interface{} // 预加载条件
	Page       int
	PageSize   int
}

func NewQueryOption() *QueryOption {
	return &QueryOption{
		Conditions: make(map[string]interface{}),
		Preloads:   make(map[string]interface{}),
		Page:       1,
		PageSize:   10,
	}
}

func (q *QueryOption) AddCondition(key string, value interface{}) {
	q.Conditions[key] = value
}

func (q *QueryOption) AddPreload(key string, condition ...interface{}) {
	q.Preloads[key] = condition
}
