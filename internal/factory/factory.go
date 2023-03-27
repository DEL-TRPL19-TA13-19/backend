package factory

import (
	"gorm.io/gorm"
	"ta13-svc/database"
	"ta13-svc/internal/repository"
)

type Factory struct {
	Db                    *gorm.DB
	UserRepository        repository.UserRepository
	TpsRepository         repository.TpsRepository
	CollectionRepository  repository.CollectionRepository
	AlternativeRepository repository.AlternativeRepository
}

func NewFactory() *Factory {
	f := &Factory{}
	f.SetupDb()
	f.SetupRepository()

	return f
}

func NewFactoryV2(db *gorm.DB) *Factory {
	f := &Factory{
		Db: db,
	}
	f.SetupRepository()

	return f
}

func (f *Factory) SetupDb() {
	db, err := database.Connection("TA13DB")
	if err != nil {
		panic("Failed setup db, connection is undefined")
	}
	f.Db = db
}

func (f *Factory) SetupRepository() {
	if f.Db == nil {
		panic("Failed setup repository, db is undefined")
	}

	f.UserRepository = repository.NewUser(f.Db)
	f.TpsRepository = repository.NewTps(f.Db)
	f.CollectionRepository = repository.NewCollection(f.Db)
	f.AlternativeRepository = repository.NewAlternative(f.Db)
}
