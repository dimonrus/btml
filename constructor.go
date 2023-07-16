package btml

// NewCallback Constructor ITemplate type
type NewCallback func(name string) ITemplate

// Constructor for custom types
var Constructor NewCallback
