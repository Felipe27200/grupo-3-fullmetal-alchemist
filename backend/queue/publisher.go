package queue

import (
	"encoding/json"
)

func PublishTransmutation(job any) error {
	data, err := json.Marshal(job)
	if err != nil {
		return err
	}

	return RedisClient.LPush(Ctx, "transmutation_queue", data).Err()
}
