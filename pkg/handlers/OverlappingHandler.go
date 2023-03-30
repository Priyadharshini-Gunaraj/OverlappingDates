package handlers

import (
	"log"
	"net/http"
	"time"

	pb "overlappingdates/proto"

	"github.com/golang/protobuf/jsonpb"
)

func OverlappingHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	req := &pb.OverlappingRequest{}
	err := jsonpb.Unmarshal(r.Body, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse each start and end dates
	_, err = parseISODate(req.Range1.StartDate)
	if err != nil {
		log.Fatal("invalid start date: ", err)
		return
	}
	_, err = parseISODate(req.Range1.EndDate)
	if err != nil {
		log.Fatal("invalid end date: ", err)
		return
	}

	_, err = parseISODate(req.Range2.StartDate)
	if err != nil {
		log.Fatal("invalid start date: ", err)
		return
	}
	_, err = parseISODate(req.Range2.EndDate)
	if err != nil {
		log.Fatal("invalid end date: ", err)
		return
	}

	// Check if date ranges overlap
	overlap := checkOverlap(req.Range1, req.Range2)

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

func parseISODate(dateStr string) (time.Time, error) {
	return time.Parse(time.RFC3339, dateStr)
}

func checkOverlap(range1, range2 *pb.DateRange) bool {
	// Compare start times
	if range1.StartDate > range2.EndDate || range2.StartDate > range1.EndDate {
		return false
	}

	return true
}
