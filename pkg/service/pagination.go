package service

// Pagination ...
type Pagination struct {
	Limit  uint64 `json:"limit"`
	Offset uint64 `json:"offset"`
}

const (
	defaultPaginationLimit uint64 = 10
	maxPaginationLimit     uint64 = 100
)

// GetLimit ...
func (p Pagination) GetLimit() uint64 {
	return p.GetCustomLimit(defaultPaginationLimit, maxPaginationLimit)
}

// GetCustomLimit ...
func (p Pagination) GetCustomLimit(defaultLimit uint64, maxLimit uint64) uint64 {
	if p.Limit > maxLimit {
		return maxLimit
	}
	if p.Limit == 0 {
		return defaultLimit
	}
	return p.Limit
}

// GetOffset ...
func (p Pagination) GetOffset() uint64 {
	return p.Offset
}
