package easyjson_test

import (
	"encoding/json"
	"testing"

	"easyjson_test/user"

	"github.com/mailru/easyjson"
)

func BenchmarkEasyJSONMarshal(b *testing.B) {
	user := user.User{
		ID:       1,
		Username: "john",
		Email:    "john@example.com",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = easyjson.Marshal(user)
	}
}

func BenchmarkJSONMarshal(b *testing.B) {
	user := user.User{
		ID:       1,
		Username: "john",
		Email:    "john@example.com",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(user)
	}
}

func BenchmarkEasyJSONUnmarshal(b *testing.B) {
	user := user.User{
		ID:       1,
		Username: "john",
		Email:    "john@example.com",
	}

	data, _ := easyjson.Marshal(user)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = easyjson.Unmarshal(data, &user)
	}
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	user := user.User{
		ID:       1,
		Username: "john",
		Email:    "john@example.com",
	}

	data, _ := json.Marshal(user)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = json.Unmarshal(data, &user)
	}
}
