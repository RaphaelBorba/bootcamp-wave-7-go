package application

import (
	"app/internal/config"
	"app/internal/handler"
	"app/internal/repository"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewApplicationDefault creates a new default application.
func NewApplicationDefault(addr, filePathStore string) (a *ApplicationDefault) {
	// default config
	defaultRouter := chi.NewRouter()
	defaultAddr := ":8080"
	if addr != "" {
		defaultAddr = addr
	}

	a = &ApplicationDefault{
		rt:            defaultRouter,
		addr:          defaultAddr,
		filePathStore: filePathStore,
	}
	return
}

// ApplicationDefault is the default application.
type ApplicationDefault struct {
	// rt is the router.
	rt *chi.Mux
	// addr is the address to listen.
	addr string
	// filePathStore is the file path to store.
	filePathStore string
}

// TearDown tears down the application.
func (a *ApplicationDefault) TearDown() (err error) {
	return
}

// SetUp sets up the application.
func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// - store
	dbCfg := config.LoadDBConfig()
	db, err := sql.Open("mysql", dbCfg.DSN())
	if err != nil {
		log.Fatalf("could not open DB: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("could not ping DB: %v", err)
	}

	// 2. Create a MySQL-backed repository
	repo := repository.NewRepositoryProductMySQL(db)
	// - handler
	hd := handler.NewHandlerProduct(repo)

	// router
	// - middlewares
	a.rt.Use(middleware.Logger)
	a.rt.Use(middleware.Recoverer)
	// - endpoints
	a.rt.Route("/products", func(r chi.Router) {
		// GET /products/{id}
		r.Get("/{id}", hd.GetById())
		// POST /products
		r.Post("/", hd.Create())
		// PUT /products/{id}
		r.Put("/{id}", hd.UpdateOrCreate())
		// PATCH /products/{id}
		r.Patch("/{id}", hd.Update())
		// DELETE /products/{id}
		r.Delete("/{id}", hd.Delete())
	})

	return
}

// Run runs the application.
func (a *ApplicationDefault) Run() (err error) {
	log.Println("👉 Listening on :8080")
	err = http.ListenAndServe(a.addr, a.rt)
	return
}
