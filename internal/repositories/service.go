package repositories

//func DBConnection() (*sqlx.DB, error) {
//	if err := godotenv.Load(); err != nil {
//		return nil, fmt.Errorf("unable to read env")
//	}
//	db, err := NewPostgresDB(Config{
//		Host:     "localhost",
//		Port:     os.Getenv("POSTGRES_PORT"),
//		Username: os.Getenv("POSTGRES_USER"),
//		Password: os.Getenv("POSTGRES_PASSWORD"),
//		DBName:   os.Getenv("POSTGRES_DB"),
//		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
//	})
//	if err != nil {
//		return nil, fmt.Errorf("failed to initialize db")
//	}
//	return db, nil
//}
