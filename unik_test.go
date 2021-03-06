package unik

import (
	"fmt"
	"github.com/plimble/unik/snowflake"
	"github.com/plimble/unik/uuid"
	"testing"
)

func TestGenerateID(t *testing.T) {
	sf := snowflake.New(0)
	fmt.Println("snowflake:", sf.Generate())
	u1 := uuid.NewV1()
	fmt.Println("uuid1:", u1.Generate())
	u4 := uuid.NewV4()
	fmt.Println("uuid4:", u4.Generate())
	u1b64 := uuid.NewV1Base64()
	fmt.Println("uuid1 base64:", u1b64.Generate())
}

func BenchmarkSnowflake(b *testing.B) {
	id := snowflake.New(0)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.Generate()
	}
}

func BenchmarkUUIDV1(b *testing.B) {
	id := uuid.NewV1()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.Generate()
	}
}

func BenchmarkUUIDV4(b *testing.B) {
	id := uuid.NewV4()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.Generate()
	}
}

func BenchmarkUUID1Base64(b *testing.B) {
	id := uuid.NewV1Base64()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.Generate()
	}
}
