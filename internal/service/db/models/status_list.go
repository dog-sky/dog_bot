package models

import "google.golang.org/protobuf/types/known/timestamppb"

type FilterDate struct {
	From *timestamppb.Timestamp
	To   *timestamppb.Timestamp
}

type Filter struct {
	Date    FilterDate
	Actions []string
}
