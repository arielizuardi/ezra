package db

import (
	"errors"
	"path"
	"runtime"

	"time"

	_ "github.com/mattes/migrate/driver/mysql"
	"github.com/mattes/migrate/migrate"
)

// RunAllMigrations run the db migration
func RunAllMigrations(dbURI string) ([]error, bool) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return []error{errors.New("Failed to get filename")}, false
	}
	migrationPath := path.Join(path.Dir(filename), "migrations")

	for {
		errs, ok := migrate.ResetSync("mysql://"+dbURI, migrationPath)
		if !ok {
			for _, e := range errs {
				if e != nil {
					if e.Error() == `driver: bad connection` {
						continue
					} else {
						version, err := migrate.Version("mysql://"+dbURI, migrationPath)
						if err != nil {
							return []error{err}, false
						}

						if version > 0 {
							time.Sleep(3 * time.Second)
							return []error{}, true
						} else {
							return errs, false
						}
					}
				}
			}
		} else if ok {
			break
		}
	}

	return []error{}, true
}
