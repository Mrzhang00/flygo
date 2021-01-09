package session

type sessionListener func(session Session)

// Listener struct
type Listener struct {
	// Created Listener
	Created sessionListener
	// Destroyed Listener
	Destroyed sessionListener
	// Invalidated Listener
	Invalidated sessionListener
	// Refreshed Listener
	Refreshed sessionListener
}
