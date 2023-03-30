package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	pb "overlappingdates/proto"
	"testing"
	"time"
)

type requests []struct {
	r pb.OverlappingRequest
}

func TestOverlappingDates(t *testing.T) {
	// Start the service
	srv := http.Server{}
	defer srv.Shutdown(context.Background())

	// Create a client for the service
	client := &http.Client{Timeout: time.Second}

	// Create requests with the two date ranges
	req := []pb.OverlappingRequest{
		{
			Range1: &pb.DateRange{
				StartDate: "2022-01-01T00:00:00Z",
				EndDate:   "2022-02-01T00:00:00Z",
			},
			Range2: &pb.DateRange{
				StartDate: "2022-01-15T00:00:00Z",
				EndDate:   "2022-02-15T00:00:00Z",
			},
		},
		{
			Range1: &pb.DateRange{
				StartDate: "2022-01-15T00:00:00Z",
				EndDate:   "2022-02-15T00:00:00Z",
			},
			Range2: &pb.DateRange{
				StartDate: "2022-01-01T00:00:00Z",
				EndDate:   "2022-02-01T00:00:00Z",
			},
		},
		{
			Range1: &pb.DateRange{
				StartDate: "2022-01-01T00:00:00Z",
				EndDate:   "2022-01-30T00:00:00Z",
			},
			Range2: &pb.DateRange{
				StartDate: "2022-01-10T00:00:00Z",
				EndDate:   "2022-01-20T00:00:00Z",
			},
		},
	}

	for _, r := range req {
		jsonReq, err := json.Marshal(r)
		if err != nil {
			t.Fatalf("Failed to marshal request: %v", err)
		}

		// Send the request to the service
		resp, err := client.Post("http://localhost:8080/date-range/overlap", "application/json", bytes.NewBuffer(jsonReq))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		// Decode the JSON response
		var res pb.OverlappingResponse
		err = json.NewDecoder(resp.Body).Decode(&res)
		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		// Check that the response is correct
		if !res.Overlap {
			t.Errorf("Expected overlapping date ranges, but got non-overlapping date ranges")
		}
	}

}

func TestNonOverlappingDates(t *testing.T) {
	// Start the service
	srv := http.Server{}
	defer srv.Shutdown(context.Background())

	// Create a client for the service
	client := &http.Client{Timeout: time.Second}

	// Create requests with the two date ranges
	req := []pb.OverlappingRequest{
		{
			Range1: &pb.DateRange{
				StartDate: "2022-01-01T00:00:00Z",
				EndDate:   "2022-01-15T00:00:00Z",
			},
			Range2: &pb.DateRange{
				StartDate: "2023-02-01T00:00:00Z",
				EndDate:   "2023-02-15T00:00:00Z",
			},
		},
		{
			Range1: &pb.DateRange{
				StartDate: "2022-02-01T00:00:00Z",
				EndDate:   "2022-02-15T00:00:00Z",
			},
			Range2: &pb.DateRange{
				StartDate: "2023-01-01T00:00:00Z",
				EndDate:   "2023-01-15T00:00:00Z",
			},
		},
	}

	for _, r := range req {
		jsonReq, err := json.Marshal(r)
		if err != nil {
			t.Fatalf("Failed to marshal request: %v", err)
		}

		// Send the request to the service
		resp, err := client.Post("http://localhost:8080/date-range/overlap", "application/json", bytes.NewBuffer(jsonReq))
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer resp.Body.Close()

		// Decode the JSON response
		var res pb.OverlappingResponse
		err = json.NewDecoder(resp.Body).Decode(&res)

		if err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}

		// Check that the response is correct
		if res.Overlap {
			t.Errorf("Expected non-overlapping date ranges, but got overlapping date ranges")
		}
	}

}
