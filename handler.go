package main

import (
	"context"
	"encoding/json"
	"fmt"
)

func handle(ctx context.Context, payload map[string]any) (any, error) {
	// VTLの$contextにアサインされているオブジェクトが返る
	fmt.Println(marshal(payload))
	return nil, nil
}

func marshal(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}
