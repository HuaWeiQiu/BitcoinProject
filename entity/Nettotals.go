package entity

type Nettotals struct {
	Totalbytesrecv int
	Totalbytessent int
	Timemillis     int64
	Uploadtarget   Uploadtarget
}
type Uploadtarget struct {
	Timeframe               int
	Target                  int
	Target_reached          bool
	Serve_historical_blocks bool
	Bytes_left_in_cycle     int
	Time_left_in_cycle      int
}
