package database

import (
	"database/sql"
	"fmt"
	"log"
	"project2/infra/config"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	appConfig := config.GetAppConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DBHost, appConfig.DBPort, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName,
	)

	db, err = sql.Open(appConfig.DBDialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}

}

func handleCreateRequiredTables() {
	userTable := `
    CREATE TABLE IF NOT EXISTS "user" (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) UNIQUE NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL,
        age INT NOT NULL,
        created_at timestamptz DEFAULT now(),
        updated_at timestamptz DEFAULT now()
    );
`
	photoTable := `
    CREATE TABLE IF NOT EXISTS "photo" (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        caption TEXT NOT NULL,
        photo_url TEXT NOT NULL,
        user_id INT NOT NULL,
        created_at timestamptz DEFAULT now(),
        updated_at timestamptz DEFAULT now(),
        CONSTRAINT photo_user_id_fk
            FOREIGN KEY(user_id)
                REFERENCES "user"(id)
    );
`
	commentTable := `
    CREATE TABLE IF NOT EXISTS "comment" (
        id SERIAL PRIMARY KEY,
        user_id INT NOT NULL,
        photo_id INT NOT NULL,
        message TEXT NOT NULL,
        created_at timestamptz DEFAULT now(),
        updated_at timestamptz DEFAULT now(),
        CONSTRAINT comment_user_id_fk
            FOREIGN KEY(user_id)
                REFERENCES "user"(id),
        CONSTRAINT comment_photo_id_fk
            FOREIGN KEY(photo_id)
                REFERENCES "photo"(id)
    );
`
	socialMediaTable := `
    CREATE TABLE IF NOT EXISTS "socialmedia" (
        id SERIAL PRIMARY KEY,
        name VARCHAR(255) NOT NULL,
        social_media_url TEXT NOT NULL,
        user_id INT NOT NULL,
        created_at timestamptz DEFAULT now(),
        updated_at timestamptz DEFAULT now(),
        CONSTRAINT "socialmedia_user_id_fk"
            FOREIGN KEY(user_id)
                REFERENCES "user"(id)
    );
`
	createTableQueries := fmt.Sprintf("%s %s %s %s", userTable, photoTable, commentTable, socialMediaTable)

	_, err = db.Exec(createTableQueries)

	if err != nil {
		log.Panic("error occured while trying to create required tables:", err)
	}

}

func InitiliazeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	return db
}
