package helpers

import (
	"final_project/internal/reference"

	"gorm.io/gorm"
)

func Filter(query *gorm.DB, filters []reference.FilterStruc, tableName string) {
	for _, value := range filters {
		value.Field = tableName + "." + value.Field

		switch value.Condition {
		case "contains":
			{
				query.Where(value.Field+" LIKE ?", "%"+value.Value+"%")
				break
			}
		case "notcontains":
			{
				query.Where(value.Field+" NOT LIKE ?", "%"+value.Value+"%")
				break
			}
		case "startswith":
			{
				query.Where(value.Field+" LIKE ?", value.Value+"%")
				break
			}
		case "endswith":
			{
				query.Where(value.Field+" LIKE ?", "%"+value.Value)
				break
			}
		case "=":
			{
				query.Where(value.Field+" = ?", value.Value)
				break
			}
		case "<>":
			{
				query.Where(value.Field+" != ?", value.Value)
				break
			}
		case ">":
			{
				query.Where(value.Field+" > ?", value.Value)
				break
			}
		case "<":
			{
				query.Where(value.Field+" < ?", value.Value)
				break
			}
		case ">=":
			{
				query.Where(value.Field+" >= ?", value.Value)
				break
			}
		case "<=":
			{
				query.Where(value.Field+" <= ?", value.Value)
				break
			}
		}
	}
}
