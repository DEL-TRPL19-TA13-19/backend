package abstraction

import "gorm.io/gorm"

type Repository struct {
	Connection *gorm.DB
	Db         *gorm.DB
}

func (r *Repository) CheckTrx(ctx *Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
