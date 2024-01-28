package updater

type UpdaterTransactor interface {
	SubmitRoot(latestEventTimestamp uint64, root [32]byte) error
}
