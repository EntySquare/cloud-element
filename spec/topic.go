package spec

// topic definition
const (
	TOPIC_SEALER_DONE = "sealer.done"
	TOPIC_POST_DONE   = "post.done"
)

func MinerTopicSealerDone(minerID string) string {
	return minerID + "." + TOPIC_SEALER_DONE
}

func MinerTopicPostDone(minerID string) string {
	return minerID + "." + TOPIC_POST_DONE
}
