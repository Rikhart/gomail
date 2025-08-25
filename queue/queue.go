package queue

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mikestefanello/backlite"
	"github.com/mikestefanello/backlite/ui"
)

// Importar la tarea si está en otro archivo
// Si está en el mismo paquete, ya no hace falta
// type NewOrderEmailTask struct { ... }

func Connect() (*backlite.Client, *http.ServeMux, error) {
	db, err := sql.Open("sqlite3", "data.db?_journal=WAL&_timeout=5000")
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

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

	// Registrar la tarea **antes de iniciar los workers**
	// ✅ ahora funciona porque Config() devuelve puntero

	processor := func(ctx context.Context, task NewOrderEmailTask) error {
		// return email.Send(ctx, task.EmailAddress, fmt.Sprintf("Order %s received", task.OrderID))
		return nil
	}

	queue := backlite.NewQueue[NewOrderEmailTask](processor)
	bl.Register(queue)

	// Verifica que las migraciones se hayan realizado
	if err := bl.Install(); err != nil {
		return nil, nil, fmt.Errorf("error al migrar tablas: %w", err)
	}

	mux := http.NewServeMux()
	h, _ := ui.NewHandler(ui.Config{
		DB:       db,
		BasePath: "/dashboard",
	})
	h.Register(mux)

	// Redirigir /dashboard → /dashboard/
	mux.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/dashboard/", http.StatusMovedPermanently)
	})

	return bl, mux, nil
}
