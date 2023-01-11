package migrations

type startEmptyMigration struct{}

func (s *startEmptyMigration) up() string {
	return "select 1"
}

func (s *startEmptyMigration) down() string {
	return "select 2"
}
