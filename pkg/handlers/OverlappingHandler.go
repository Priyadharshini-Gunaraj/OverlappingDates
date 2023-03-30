package handlers

import (
	"net/http"

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

func checkOverlap(range1, range2 *pb.DateRange) bool {
	// Compare start times
	if range1.StartDate > range2.EndDate || range2.StartDate > range1.EndDate {
		return false
	}

	return true
}
