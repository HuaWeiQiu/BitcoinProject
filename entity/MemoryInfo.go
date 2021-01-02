package entity

type MemoryInfo struct {
	Locked Locked
}
type Locked struct {
	Used        int64
	Free        int64
	total       int64
	Locked      int
	Chunks_used int
	Chunks_free int
}
