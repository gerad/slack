package slack

// Args can be used to create ad hoc lists of arguments to pass to Exec
type Args map[string]string

// Args allows Args to conform to the Argsable interface used by Exec
func (a Args) Args() map[string]string { return a }
