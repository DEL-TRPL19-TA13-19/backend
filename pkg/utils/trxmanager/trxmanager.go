package trxmanager

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"ta13-svc/internal/factory"

	"github.com/sirupsen/logrus"
)

type trxManager struct {
	db *gorm.DB
}

//type trxFn func(ctx *abstraction.Context) error

type trxFnV2 func(ctx context.Context, factory *factory.Factory) error

func New(db *gorm.DB) *trxManager {
	return &trxManager{db}
}

//func (g *trxManager) WithTrx(pCtx *abstraction.Context, fn trxFn) (err error) {
//	tx := g.db.Begin()
//	pCtx.Trx = &abstraction.TrxContext{
//		Db: tx,
//	}
//
//	defer func() {
//		if p := recover(); p != nil {
//			// a panic occurred, rollback and repanic
//			tx.Rollback()
//			logrus.Error(p)
//			err = errors.New("panic happened because: " + fmt.Sprintf("%v", p))
//		} else if err != nil {
//			// error occurred, rollback
//			tx.Rollback()
//		} else {
//			// all good, commit
//			err = tx.Commit().Error
//		}
//	}()
//
//	err = fn(pCtx)
//	return err
//}

func (g *trxManager) WithTrxV2(ctx context.Context, fn trxFnV2) (err error) {
	tx := g.db.Begin()
	f := factory.NewFactoryV2(tx)

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			tx.Rollback()
			logrus.Error(p)
			err = errors.New("panic happened because: " + fmt.Sprintf("%v", p))
		} else if err != nil {
			// error occurred, rollback
			tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit().Error
		}
	}()

	err = fn(ctx, f)
	return err
}
