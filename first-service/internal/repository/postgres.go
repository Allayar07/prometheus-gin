package repository

//func NewPostgresDB(ctx context.Context) (*pgxpool.Pool, error) {
//	dbURL := "postgres://" + "postgres" + ":" + url.QueryEscape("password0701") + "@" + "localhost" + ":" + "5432" + "/" + "practice" + "?sslmode=" + "disable"
//	dbPool, err := pgxpool.Connect(ctx, dbURL)
//
//	if err != nil {
//		return nil, err
//	}
//	err = dbPool.Ping(ctx)
//	if err != nil {
//		return nil, err
//	}
//	return dbPool, nil
//}
