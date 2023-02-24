package app

type CliOptions interface {
	Flags() (fss NamedFlagSets)
	Validate() []error
}

// CompletableOptions abstracts options which can be completed.
type CompletableOptions interface {
	Complete() error
}

// PrintableOptions abstracts options which can be printed.
type PrintableOptions interface {
	String() string
}
