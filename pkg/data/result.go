package data

const (
	Noop = iota
	EmptyResult
	ConflictData
	ValidationError
	GenericError 
);

type ErrorType int;

type Result struct {
	Error ErrorType;
	Result interface{}
};