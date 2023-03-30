package handlers

import (
	"net/http"
	"time"

	pb "overlappingdates/proto"

	"github.com/golang/protobuf/jsonpb"
)

type parsedDates struct {
	StartDate time.Time
	EndDate   time.Time
}

func OverlappingHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	req := &pb.OverlappingRequest{}
	err := jsonpb.Unmarshal(r.Body, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if date ranges overlap
	overlap, err := checkOverlap(req.Range1, req.Range2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create response
	res := &pb.OverlappingResponse{Overlap: overlap}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	m := jsonpb.Marshaler{EmitDefaults: true, OrigName: true}
	err = m.Marshal(w, res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func parseDateRange(r *pb.DateRange) (parsedDates, error) {
	var pd parsedDates
	StartDate, err := parseDate(r.StartDate)
	if err != nil {
		return pd, err
	}
	EndDate, err := parseDate(r.EndDate)
	if err != nil {
		return pd, err
	}
	pd.StartDate = StartDate
	pd.EndDate = EndDate
	return pd, nil
}

func parseDate(date string) (time.Time, error) {
	d, err := parseISODate(date)
	if err != nil {
		return time.Now(), err
	}
	return d, nil
}

func parseISODate(dateStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateStr)
}

func checkOverlap(range1, range2 *pb.DateRange) (bool, error) {
	// Parse each start and end dates
	parsedRange1, err := parseDateRange(range1)
	if err != nil {
		return false, err
	}
	parsedRange2, err := parseDateRange(range2)
	if err != nil {
		return false, err
	}

	// Compare start times
	if parsedRange1.StartDate.After(parsedRange2.EndDate) || parsedRange2.StartDate.After(parsedRange1.EndDate) {
		return false, nil
	}

	return true, nil
}
