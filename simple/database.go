package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	db := &Database{
		Name: "PostgreSQL",
	}
	return (*DatabasePostgreSQL)(db)
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	db := &Database{
		Name: "MongoDB",
	}
	return (*DatabaseMongoDB)(db)
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreSQL: postgreSQL,
		DatabaseMongoDB:    mongoDB,
	}
}
