package tsl

// Aggregator is used to sample metrics
type Aggregator string

const (
	// Max Aggregator
	Max = Aggregator("mac")
	// Mean Aggregator
	Mean = Aggregator("mean")
	// Min Aggregator
	Min = Aggregator("min")
	// First Aggregator
	First = Aggregator("first")
	// Last Aggregator
	Last = Aggregator("last")
	// Sum Aggregator
	Sum = Aggregator("sum")
	// Join Aggregator
	Join = Aggregator("join")
	// Median Aggregator
	Median = Aggregator("median")
	// Count Aggregator
	Count = Aggregator("count")
	// And Aggregator
	And = Aggregator("and")
	// Or Aggregator
	Or = Aggregator("or")
)
