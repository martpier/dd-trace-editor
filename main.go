package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	buf, err := os.ReadFile("traces")

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	payload := Payload{}
	_, err = payload.UnmarshalMsg(buf)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	offset := r.Uint64(); // an offset to be added to all IDs, so every ID is new but links between spans are preserved
	timeOffsetter := TimeOffsetter{}

	newPayload := Payload{}
	for _, trace := range payload {
		newTrace := []Span{}
		for _, span := range trace {
			span.Start = timeOffsetter.offsetTime(span.Start)
			span.TraceID += offset
			span.SpanID += offset
			if span.ParentID != 0 {
				span.ParentID +=offset
			}

			// apply custom transformation to spans below
			delete(span.Metrics, "_dd.agent_psr")
			// apply custom transformation to spans above

			newTrace = append(newTrace, span)
		}

	newPayload = append(newPayload, newTrace)
	}

	newRaw, err := newPayload.MarshalMsg(nil)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}


	err = os.WriteFile("new_traces", newRaw, 0o644)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
}

type TimeOffsetter struct {
	newReference int64
	originalReference int64
}

func (t* TimeOffsetter) offsetTime(original int64) int64 {
	if t.originalReference == 0 {
		t.originalReference = original
		t.newReference = time.Now().UnixNano()
	}

	return original - t.originalReference + t.newReference
}