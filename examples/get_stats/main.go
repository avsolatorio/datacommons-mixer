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
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"time"

	pb "github.com/datacommonsorg/mixer/proto"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "127.0.0.1:12345", "Address of grpc server.")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100000000 /* 100M */)))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMixerClient(conn)
	ctx := context.Background()

	// Read a large number of places.
	file, err := os.Open("geos.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	place := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		place = append(place, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// Total population
	runGetStats(ctx, c, []string{"geoId/06085", "geoId/06"}, "TotalPopulation")

	// Citizenship
	runGetStats(ctx, c, []string{"geoId/0649670"}, "dc/jpbdx38nyvhdd")

	// StateUnemploymentInsurance
	runGetStats(ctx, c, []string{"geoId/06085", "geoId/06"}, "dc/kgklfjwvtwx0b")

	runGetStats(ctx, c, place, "NYTCovid19CumulativeCases")
}

func runGetStats(ctx context.Context, c pb.MixerClient, place []string, statsVar string) {
	t1 := time.Now()
	resp, err := c.GetStats(ctx, &pb.GetStatsRequest{
		StatsVar: statsVar,
		Place:    place,
	})
	t2 := time.Now()
	if err != nil {
		log.Fatalf("could not GetStats: %s", err)
	}
	diff := t2.Sub(t1)
	log.Println(diff)
	if len(resp.GetPayload()) > 1000 {
		log.Printf("%s", string(resp.GetPayload()[0:1000]))
	} else {
		log.Printf("%s", string(resp.GetPayload()))
	}
}
