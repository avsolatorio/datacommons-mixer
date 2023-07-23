package main

import (
	"flag"
	"log"

	"github.com/datacommonsorg/mixer/internal/server/resource"
	"github.com/datacommonsorg/mixer/sqlite/writer"
)

var (
	sqlite_dir = flag.String("sqlite_dir", "/Users/wsws/Software/mixer/sqlite/writer/test_data", "SQLite directory.")
)

func main() {
	if err := writer.WriteCSV(
		&resource.Metadata{
			SQLitePath:        *sqlite_dir,
			RemoteMixerDomain: "https://api.datacommons.org",
			RemoteMixerAPIKey: "AIzaSyCTI4Xz-UW_G2Q2RfknhcfdAnTHq5X5XuI",
		}); err != nil {
		log.Fatalf("writer.Write() = %v", err)
	}
}
