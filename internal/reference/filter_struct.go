package reference

type FilterStruc struct {
	Field     string `form:"field"`
	Condition string `form:"condition"`
	Value     string `form:"value"`
}
