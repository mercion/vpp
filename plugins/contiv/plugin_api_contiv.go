package contiv

// API for other plugins to query network-related information.
type API interface {
	// GetIfName looks up logical interface name that corresponds to the interface
	// associated with the given pod.
	GetIfName(podNamespace string, podName string) (name string, exists bool)

	// GetNsIndex returns the index of the VPP session namespace associated
	// with the given pod.
	GetNsIndex(podNamespace string, podName string) (nsIndex int, exists bool)
}
