package notes

import (
	"context"
	"os"
	"testing"

	"singularity.com/pz8-mongo/internal/db"
)

func TestCreateAndGet(t *testing.T) {
	ctx := context.Background()
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		uri = "mongodb://root:secret@localhost:27017/?authSource=admin"
	}

	deps, err := db.ConnectMongo(ctx, uri, "pz8_test")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { deps.Client.Disconnect(ctx) })
	r, err := NewRepo(deps.Database)
	if err != nil {
		t.Fatal(err)
	}

	created, err := r.Create(ctx, "T1", "C1")
	if err != nil {
		t.Fatal(err)
	}

	got, err := r.ByID(ctx, created.ID.Hex())
	if err != nil {
		t.Fatal(err)
	}
	if got.Title != "T1" {
		t.Fatalf("want T1 got %s", got.Title)
	}
}
