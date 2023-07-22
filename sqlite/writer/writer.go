package writer

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	pbv2 "github.com/datacommonsorg/mixer/internal/proto/v2"
	"github.com/datacommonsorg/mixer/internal/server/resource"
	"github.com/datacommonsorg/mixer/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type observation struct {
	entity   string
	variable string
	date     string
	value    string
}

type triple struct {
	subjectID   string
	predicate   string
	objectID    string
	objectValue string
}

// Writer is an object for SQLite writer.
type Writer struct {
	inputDir           string
	outputDir          string
	resourceMetadata   *resource.Metadata
	httpClient         *http.Client
	observationsHeader []string
	triplesHeader      []string
}

// New creates a new Writer instance.
func New(inputDir, outputDir string) *Writer {
	return &Writer{
		inputDir:  inputDir,
		outputDir: outputDir,
		resourceMetadata: &resource.Metadata{
			RemoteMixerDomain: "https://api.datacommons.org",
			RemoteMixerAPIKey: "AIzaSyCTI4Xz-UW_G2Q2RfknhcfdAnTHq5X5XuI",
		},
		httpClient:         &http.Client{},
		observationsHeader: []string{"entity", "variable", "date", "value"},
		triplesHeader:      []string{"subject_id", "predicate", "object_id", "object_value"},
	}
}

// Write performs the writing.
func (w *Writer) Write() error {
	csvFiles, err := listCSVFiles(w.inputDir)
	if err != nil {
		return err
	}
	if len(csvFiles) == 0 {
		return status.Errorf(codes.FailedPrecondition, "No CSV files found in %s", w.inputDir)
	}

	observationList := []*observation{}
	variableSet := map[string]struct{}{}
	for _, csvFile := range csvFiles {
		observations, variables, err := w.processCSVFile(csvFile)
		if err != nil {
			return err
		}
		observationList = append(observationList, observations...)
		for _, v := range variables {
			variableSet[v] = struct{}{}
		}
	}

	tripleList := []*triple{}
	for variable := range variableSet {
		tripleList = append(tripleList, &triple{
			subjectID: variable,
			predicate: "typeOf",
			objectID:  "StatisticalVariable",
		})
	}

	return w.writeOutput(observationList, tripleList)
}

func listCSVFiles(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var res []string
	for _, file := range files {
		if fName := file.Name(); strings.HasSuffix(fName, ".csv") {
			res = append(res, file.Name())
		}
	}

	return res, nil
}

func (w *Writer) processCSVFile(csvFile string) ([]*observation,
	[]string, // A list of variables.
	error) {
	// Read the CSV file.
	f, err := os.Open(filepath.Join(w.inputDir, csvFile))
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	records, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, nil, err
	}
	numRecords := len(records)
	if numRecords < 2 {
		return nil, nil, status.Errorf(codes.FailedPrecondition,
			"Empty CSV file %s", csvFile)
	}

	// Load header.
	header := records[0]
	if len(header) < 3 {
		return nil, nil, status.Errorf(codes.FailedPrecondition,
			"Less than 3 columns in CSV file %s", csvFile)
	}
	numColumns := len(header)

	// Resolve places.
	places := []string{}
	for i := 1; i < numRecords; i++ {
		places = append(places, records[i][0])
	}
	resolvedPlaceMap, err := w.resolvePlaces(places, header[0])
	if err != nil {
		return nil, nil, err
	}

	// Generate observations.
	observations := []*observation{}
	for i := 1; i < numRecords; i++ {
		record := records[i]

		resolvedPlace, ok := resolvedPlaceMap[record[0]]
		if !ok {
			// If a place cannot be resolved, simply ignore it.
			continue
		}

		for j := 2; j < numColumns; j++ {
			observations = append(observations, &observation{
				entity:   resolvedPlace,
				variable: header[j],
				date:     record[1],
				value:    record[j],
			})
		}
	}

	return observations, header[2:], nil
}

func (w *Writer) resolvePlaces(places []string, placeHeader string) (map[string]string, error) {
	placeToDCID := map[string]string{}

	if placeHeader == "lat#lng" {
		// TODO(ws): lat#lng recon.
	} else if placeHeader == "name" {
		// TODO(ws): name recon.
	} else {
		resp := &pbv2.ResolveResponse{}
		if err := util.FetchRemote(w.resourceMetadata, w.httpClient, "/v2/resolve",
			&pbv2.ResolveRequest{
				Nodes:    places,
				Property: fmt.Sprintf("<-%s->dcid", placeHeader),
			}, resp); err != nil {
			return nil, err
		}
		for _, entity := range resp.GetEntities() {
			if _, ok := placeToDCID[entity.GetNode()]; ok {
				continue
			}
			// TODO(ws): Handle the case with multiple ResolvedIds.
			placeToDCID[entity.GetNode()] = entity.GetResolvedIds()[0]
		}
	}

	return placeToDCID, nil
}

func (w *Writer) writeOutput(observations []*observation,
	triples []*triple) error {
	// Observations.
	fObservations, err := os.Create(filepath.Join(w.outputDir, "observations.csv"))
	if err != nil {
		return err
	}
	defer fObservations.Close()
	wObservations := csv.NewWriter(fObservations)
	wObservations.Write(w.observationsHeader)
	for _, o := range observations {
		wObservations.Write([]string{o.entity, o.variable, o.date, o.value})
	}
	wObservations.Flush()

	// Triples.
	fTriples, err := os.Create(filepath.Join(w.outputDir, "triples.csv"))
	if err != nil {
		return err
	}
	defer fTriples.Close()
	wTriples := csv.NewWriter(fTriples)
	wTriples.Write(w.triplesHeader)
	for _, t := range triples {
		wTriples.Write([]string{t.subjectID, t.predicate, t.objectID, t.objectValue})
	}
	wTriples.Flush()

	return nil
}
