package sequence

// Currently Ceres use the simple snowflake with the public IP of the machine which Ceres running on it
// TODO: next version need use the distributed senquence implementation

// Senquence interface to generate the unique senquence number
type Senquence interface {
	// Next will return the next senquence
	Next() uint64
}
