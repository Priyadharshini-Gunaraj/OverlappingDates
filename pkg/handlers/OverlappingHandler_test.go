package handlers

import (
	pb "overlappingdates/proto"
	"testing"
)

func TestCheckOverlap(t *testing.T) {
	tests := []struct {
		name           string
		range1         *pb.DateRange
		range2         *pb.DateRange
		expectedOutput bool
		expectedError  bool
	}{
		{
			name: "ranges overlap",
			range1: &pb.DateRange{
				StartDate: "2022-01-01T00:00:00Z",
				EndDate:   "2022-01-10T00:00:00Z",
			},
			range2: &pb.DateRange{
				StartDate: "2022-01-05T00:00:00Z",
				EndDate:   "2022-01-15T00:00:00Z",
			},
			expectedOutput: true,
		},
		{
			name: "ranges do not overlap",
			range1: &pb.DateRange{
				StartDate: "2022-01-01T00:00:00Z",
				EndDate:   "2022-01-10T00:00:00Z",
			},
			range2: &pb.DateRange{
				StartDate: "2022-01-15T00:00:00Z",
				EndDate:   "2022-01-20T00:00:00Z",
			},
			expectedOutput: false,
		},
		{
			name: "invalid start date format",
			range1: &pb.DateRange{
				StartDate: "2022-01-01",
				EndDate:   "2022-01-10T00:00:00Z",
			},
			range2: &pb.DateRange{
				StartDate: "2022-01-05T00:00:00Z",
				EndDate:   "2022-01-15T00:00:00Z",
			},
			expectedError: true,
		},
		{
			name: "invalid end date format",
			range1: &pb.DateRange{
				StartDate: "2022-01-01T00:00:00Z",
				EndDate:   "2022-01-10",
			},
			range2: &pb.DateRange{
				StartDate: "2022-01-05T00:00:00Z",
				EndDate:   "2022-01-15T00:00:00Z",
			},
			expectedError: true,
		},
	}

	for _, testRange := range tests {
		t.Run(testRange.name, func(t *testing.T) {
			result, err := checkOverlap(testRange.range1, testRange.range2)
			if testRange.expectedError {
				if err == nil {
					t.Errorf("Expected %v", testRange.expectedError)
				}
			} else if result != testRange.expectedOutput {
				t.Errorf("Expected %v, Actual %v", testRange.expectedOutput, result)
			}
		})
	}
}

func TestParseISODates(t *testing.T) {
	tests := []struct {
		name           string
		dateStr        string
		expectedOutput string
		expectedError  bool
	}{
		{
			name:           "valid date format",
			dateStr:        "2022-01-01T00:00:00Z",
			expectedOutput: "2022-01-01 00:00:00 +0000 UTC",
		},
		{
			name:          "invalid date format",
			dateStr:       "2022-01-01",
			expectedError: true,
		},
	}

	for _, testRange := range tests {
		t.Run(testRange.name, func(t *testing.T) {
			resp, err := parseISODate(testRange.dateStr)
			if testRange.expectedError {
				if err == nil {
					t.Errorf("expected %v, but got %v", testRange.expectedError, err)
				}

			} else if resp.String() != testRange.expectedOutput {
				t.Errorf("expected %v, but got %v", testRange.expectedOutput, resp.String())
			}
		})
	}
}
