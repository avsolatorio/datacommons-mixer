// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package recon

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	cbt "cloud.google.com/go/bigtable"
	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/datacommonsorg/mixer/internal/store"
	"github.com/datacommonsorg/mixer/internal/store/bigtable"
	"github.com/golang/geo/s2"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const (
	gridSize         float64 = 0.2
	geoJSONPredicate string  = "geoJsonCoordinates"
)

// TODO(Spaceenter): Also add place types to the results.

// ResolveCoordinates implements API for ReconServer.ResolveCoordinates.
func ResolveCoordinates(
	ctx context.Context, in *pb.ResolveCoordinatesRequest, store *store.Store) (
	*pb.ResolveCoordinatesResponse, error,
) {
	// Map: lat^lng => normalized lat^lng.
	normCoordinateMap := map[string]string{}
	coordinateLookupKeys := map[string]struct{}{}

	// Read request.
	for _, coordinate := range in.GetCoordinates() {
		nKey := normalizedCoordinateKey(coordinate)
		normCoordinateMap[coordinateKey(coordinate)] = nKey
		coordinateLookupKeys[nKey] = struct{}{}
	}

	// Read coordinate recon cache.
	reconRowList := cbt.RowList{}
	for key := range coordinateLookupKeys {
		reconRowList = append(reconRowList,
			fmt.Sprintf("%s%s", bigtable.BtCoordinateReconPrefix, key))
	}
	reconDataList, err := bigtable.Read(
		ctx,
		store.BtGroup,
		reconRowList,
		func(dcid string, jsonRaw []byte, isProto bool) (interface{}, error) {
			var recon pb.CoordinateRecon
			if isProto {
				if err := proto.Unmarshal(jsonRaw, &recon); err != nil {
					return nil, err
				}
			} else {
				if err := protojson.Unmarshal(jsonRaw, &recon); err != nil {
					return nil, err
				}
			}
			return &recon, nil
		},

		func(rowKey string) (string, error) {
			return strings.TrimPrefix(rowKey, bigtable.BtCoordinateReconPrefix), nil
		},
	)
	if err != nil {
		return nil, err
	}

	// Collect places that don't fully cover the tiles that the coordinates are in.
	questionablePlaces := map[string]struct{}{}
	for _, recon := range reconDataList[0] {
		for _, place := range recon.(*pb.CoordinateRecon).GetPlaces() {
			if !place.GetFull() {
				questionablePlaces[place.GetDcid()] = struct{}{}
			}
		}
	}

	// Read place GeoJson cache.
	geoJSONRowList := cbt.RowList{}
	for place := range questionablePlaces {
		geoJSONRowList = append(geoJSONRowList,
			fmt.Sprintf("%s%s^%s", bigtable.BtOutPropValPrefix, place, geoJSONPredicate))
	}
	geoJSONDataList, err := bigtable.Read(
		ctx,
		store.BtGroup,
		geoJSONRowList,
		func(dcid string, jsonRaw []byte, isProto bool) (interface{}, error) {
			var info pb.EntityInfoCollection
			if isProto {
				if err := proto.Unmarshal(jsonRaw, &info); err != nil {
					return nil, err
				}
			} else {
				if err := protojson.Unmarshal(jsonRaw, &info); err != nil {
					return nil, err
				}
			}
			return &info, nil
		},

		func(rowKey string) (string, error) {
			l := strings.TrimPrefix(rowKey, bigtable.BtOutPropValPrefix)
			return strings.TrimSuffix(l, fmt.Sprintf("^%s", geoJSONPredicate)), nil
		},
	)
	if err != nil {
		return nil, err
	}
	geoJSONMap := map[string]string{}
	for place, info := range geoJSONDataList[0] {
		// A place should only have a single geoJsonCooridnates out arc.
		typedInfo := info.(*pb.EntityInfoCollection)
		if typedInfo.GetTotalCount() != 1 {
			continue
		}
		geoJSONMap[place] = typedInfo.GetEntities()[0].GetValue()
	}

	// Assemble response.
	res := &pb.ResolveCoordinatesResponse{}
	for _, co := range in.GetCoordinates() {
		nKey := normCoordinateMap[coordinateKey(co)]

		recon, ok := reconDataList[0][nKey]
		if !ok {
			continue
		}

		placeCoordinates := &pb.ResolveCoordinatesResponse_PlaceCoordinate{
			Latitude:  co.GetLatitude(),
			Longitude: co.GetLongitude(),
		}
		for _, place := range recon.(*pb.CoordinateRecon).GetPlaces() {
			if place.GetFull() {
				placeCoordinates.PlaceDcids = append(placeCoordinates.PlaceDcids,
					place.GetDcid())
			} else { // Not fully cover the tile.
				geoJSON, ok := geoJSONMap[place.GetDcid()]
				if !ok {
					continue
				}
				contained, err := isContainedIn(geoJSON,
					co.GetLatitude(), co.GetLongitude())
				if err != nil {
					return res, err
				}
				if contained {
					placeCoordinates.PlaceDcids = append(placeCoordinates.PlaceDcids,
						place.GetDcid())
				}
			}
		}

		res.PlaceCoordinates = append(res.PlaceCoordinates, placeCoordinates)
	}

	return res, nil
}

func coordinateKey(c *pb.ResolveCoordinatesRequest_Coordinate) string {
	return fmt.Sprintf("%f^%f", c.GetLatitude(), c.GetLongitude())
}

func normalizedCoordinateKey(c *pb.ResolveCoordinatesRequest_Coordinate) string {
	// Normalize to South-West of the grid points.
	lat := float64(int((c.GetLatitude()+90.0)/gridSize))*gridSize - 90
	lng := float64(int((c.GetLongitude()+180.0)/gridSize))*gridSize - 180
	return fmt.Sprintf("%.1f^%.1f", lat, lng)
}

// Polygon represents a polygon shape.
type Polygon struct {
	Loops [][][]float64
}

// MultiPolygon represents a list of polygons.
type MultiPolygon struct {
	Polygons [][][][]float64
}

// GeoJSON represents the geoJson data for a place.
type GeoJSON struct {
	Type         string          `json:"type"`
	Coordinates  json.RawMessage `json:"coordinates"`
	Polygon      Polygon         `json:"-"`
	MultiPolygon MultiPolygon    `json:"-"`
}

func buildS2Loops(loops [][][]float64) ([]*s2.Loop, error) {
	res := []*s2.Loop{}
	for i, loop := range loops {
		if l := len(loop); l < 4 {
			return nil, fmt.Errorf("geoJson requires >= 4 points for a loop, got %d", l)
		}

		s2Points := []s2.Point{}
		// NOTE: We have to skip the last point when constructing the s2Loop.
		// In GeoJson, the last point is the same as the first point for a loop.
		// If not skipping, it sometimes leads to wrong result for containment calculation.
		for _, point := range loop[:len(loop)-1] {
			if len(point) != 2 {
				return nil, fmt.Errorf("wrong point format: %+v", point)
			}
			// NOTE: GeoJson has longitude comes before latitude.
			s2Points = append(s2Points,
				s2.PointFromLatLng(s2.LatLngFromDegrees(point[1], point[0])))
		}
		s2Loop := s2.LoopFromPoints(s2Points)
		if i == 0 {
			// The first ring of a polygon is a shell, it should be normalized to counter-clockwise.
			//
			// This step ensures that the planar polygon loop follows the "right-hand rule"
			// and reverses the orientation when that is not the case. This is specified by
			// RFC 7946 GeoJSON spec (https://tools.ietf.org/html/rfc7946), but is commonly
			// disregarded. Since orientation is easy to deduce on the plane, we assume the
			// obvious orientation is intended. We reverse orientation to ensure that all
			// loops follow the right-hand rule. This corresponds to S2's "interior-on-the-
			// left rule", and allows us to create these polygon as oriented S2 polygons.
			//
			// Also see https://en.wikipedia.org/wiki/Curve_orientation.
			s2Loop.Normalize()
		}
		res = append(res, s2Loop)
	}
	return res, nil
}

func parseGeoJSON(geoJSON string) (*s2.Polygon, error) {
	g := &GeoJSON{}
	if err := json.Unmarshal([]byte(geoJSON), g); err != nil {
		return nil, err
	}

	switch g.Type {
	case "Polygon":
		if err := json.Unmarshal(g.Coordinates, &g.Polygon.Loops); err != nil {
			return nil, err
		}
		s2Loops, err := buildS2Loops(g.Polygon.Loops)
		if err != nil {
			return nil, err
		}
		return s2.PolygonFromOrientedLoops(s2Loops), nil
	case "MultiPolygon":
		if err := json.Unmarshal(g.Coordinates, &g.MultiPolygon.Polygons); err != nil {
			return nil, err
		}
		s2Loops := []*s2.Loop{}
		for _, polygon := range g.MultiPolygon.Polygons {
			lps, err := buildS2Loops(polygon)
			if err != nil {
				return nil, err
			}
			s2Loops = append(s2Loops, lps...)
		}
		return s2.PolygonFromOrientedLoops(s2Loops), nil
	default:
		return nil, fmt.Errorf("unrecognized GeoJson object: %+v", g.Type)
	}
}

func isContainedIn(geoJSON string, lat float64, lng float64) (bool, error) {
	s2Polygon, err := parseGeoJSON(geoJSON)
	if err != nil {
		return false, err
	}
	s2Point := s2.PointFromLatLng(s2.LatLngFromDegrees(lat, lng))
	return s2Polygon.ContainsPoint(s2Point), nil
}
