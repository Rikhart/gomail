package queue

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3" // Asegúrate de importar el driver
	"github.com/mikestefanello/backlite"
	"github.com/mikestefanello/backlite/ui"
)

func Connect() (*backlite.Client, *http.ServeMux, error) {
	db, err := sql.Open("sqlite3", "data.db?_journal=WAL&_timeout=5000")
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	// Opcional: Validar la conexión
	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	bl, err := backlite.NewClient(backlite.ClientConfig{
		DB:              db,
		Logger:          slog.Default(),
		ReleaseAfter:    10 * time.Minute,
		NumWorkers:      10,
		CleanupInterval: time.Hour,
	})
	if err != nil {
		log.Fatalf("Error al iniciar Backlite: %v", err)
	}

	// Verifica que las migraciones se hayan realizado
	if err := bl.Install(); err != nil {
		return nil, nil, fmt.Errorf("error al migrar tablas: %w", err)
	}

	mux := http.DefaultServeMux
	ui.NewHandler(db).Register(mux)

	return bl, mux, nil
}
