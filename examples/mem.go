// Copyright 2019 Google LLC
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

package main

import (
	"context"
	"flag"
	"log"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"errors"
	"github.com/google/pprof/profile"
	pb "github.com/datacommonsorg/mixer/internal/proto"
	"google.golang.org/grpc"
)


var (
	addr = flag.String("addr", "127.0.0.1:12345", "Address of grpc server.")
	temp_path = flag.String("temp_path", "http_memprof_out", "Folder to store temporary output of memory profiles over HTTP")
)

func readProfile (filename string) (*profile.Profile, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("file could not be opened: " + filename)
	}

	prof, err := profile.Parse(reader)
	if err != nil {
		return nil, fmt.Errorf("could not parse proto file: %v", filename)
	}
	return prof, err
}

func GetTotalSpaceAllocFromProfile(filename string) (int64) {
	// Modified from https://github.com/google/pprof/blob/c488b8fa1db3fa467bf30beb5a1d6f4f10bb1b87/internal/report/report.go
	// =======================
	// computeTotal computes the sum of the absolute value of all sample values.
	// If any samples have label indicating they belong to the diff base, then the
	// total will only include samples with that label.
	computeTotal := func (prof *profile.Profile, value func(v []int64) int64) int64 {
		var total, diffTotal int64
		for _, sample := range prof.Sample {
			var v int64
			v = value(sample.Value)
			fmt.Println(v)
			if v < 0 {
				v = -v
			}
			total += v
			if sample.DiffBaseSample() {
				diffTotal += v
			}
		}
		if diffTotal > 0 {
			total = diffTotal
		}
		return total
	}
	// sample index of alloc_space in neap profiles from http/net/pprof
	const alloc_space_sample_index = 1;

	prof, err := readProfile(filename)
	if err != nil {
		log.Fatal(err)
	}
	valuef := func(v []int64) int64 {
		return v[alloc_space_sample_index]
	}
	total := computeTotal(prof, valuef)
	return total
}

func BytesToMegabytes(bytes int64) int64 {
	return bytes / (1024 * 1024);
}

func save_profile(outPath string) {
	output_location_flag := fmt.Sprintf("-output=%v", outPath)
	cmd := exec.Command("go", "tool", "pprof", output_location_flag, "-proto", "http://localhost:6060/debug/pprof/heap?gc=1")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
}

func compute_memory_profDiff (profBefore, profAfter, outPath string) {
	// TODO(snny): outPath is currently not used
	cmd := exec.Command("go", "tool", "pprof", "-top", "-nodecount=10", "-sample_index=alloc_space", "-base", profBefore, profAfter)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		return
	}
	fmt.Println(stdout.String())
}

func run_with_profile (f_identifier string, f func()) {
	profile_location_template := *temp_path + "/go_http_memprof.%v.%v.pb"

	profBefore := fmt.Sprintf(profile_location_template, "before", f_identifier)
	profAfter := fmt.Sprintf(profile_location_template, "after", f_identifier)
	//profDiff := fmt.Sprintf(profile_location_template, "diff", f_identifier)

	save_profile(profBefore)
	f()
	save_profile(profAfter)
	allocBefore := GetTotalSpaceAllocFromProfile(profBefore)
	allocAfter := GetTotalSpaceAllocFromProfile(profAfter)
	allocDiff := allocAfter - allocBefore
	allocDiffMb := BytesToMegabytes(allocDiff)
	fmt.Printf("That call took %d MB [%d bytes]", allocDiffMb, allocDiff)

	// TODO(snny): use this so that a detailed proto file of the diff is
	// available for further inspection
	// compute_memory_profDiff(profBefore, profAfter, profDiff)
	//totalMb := BytesToMegabytes(GetTotalSpaceAllocFromProfile(profDiff))
	//fmt.Printf("That call took %d", totalMb)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.Parse()

	// Ensure the temp output path exists
	err := os.MkdirAll(*temp_path, 0o0700)
	if err != nil {
		log.Fatalf("could not create temp directory at %v: %v", temp_path, err)
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100000000 /* 100M */)),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMixerClient(conn)
	ctx := context.Background()

	funcs_to_profile := make(map[string]func())

	funcs_to_profile["BulkObservationsSeriesLinked_USA"] = func(){
			req := &pb.BulkObservationsSeriesLinkedRequest{
				Variables:      []string{"Median_Age_Person_AmericanIndianOrAlaskaNativeAlone"},
				EntityType:     "City",
				LinkedEntity:   "country/USA",
				LinkedProperty: "containedInPlace",
				AllFacets:      false,
			}
			_, err := c.BulkObservationsSeriesLinked(ctx,req)
			if err != nil {
				log.Fatalf("could not run BulkObservationsSeriesLinked: %s", err)
			}
			fmt.Printf("BulkObservationsSeriesLinked returned succesfully\n")
			// Commenting for now because output floods screen
			// fmt.Printf("%d\n", proto.MarshalTextString(r))
		}

		for key, f := range funcs_to_profile {
			run_with_profile(key, f)
		}


	return;
}
