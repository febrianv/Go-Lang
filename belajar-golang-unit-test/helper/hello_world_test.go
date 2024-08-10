package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B){
	benchmarks := []struct{
		name 	string
		request string
	}{
		{
			name: "Edo", 
			request: "Edo",
		},
		{
			name: "Febrian", 
			request: "Febrian",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B){
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkHelloWorldSub(b *testing.B){
	b.Run("Edo", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			HelloWorld("Edo")
		}
	})

	b.Run("Febrian", func(b *testing.B){
		for i := 0; i < b.N; i++ {
			HelloWorld("Febrian")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B){
	for i := 0; i < b.N; i++ {
		HelloWorld("Edo")
	}
}

func BenchmarkHelloWorldFebrian(b *testing.B){
	for i := 0; i < b.N; i++ {
		HelloWorld("Febrian")
	}
}

func TestTableHelloWorld(t *testing.T){
	tests := []struct {
		name string
		request string
		expected string
	}{
		{
			name:     "Edo",
			request:  "Edo",
			expected: "Hello Edo",
		},
		{
			name:     "Febrian",
			request:  "Febrian",
			expected: "Hello Febrian",
		},
		{
			name:     "Vernando",
			request:  "Vernando",
			expected: "Hello Vernando",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Edo", func(t *testing.T){
		result := HelloWorld("Edo")
		require.Equal(t,"Hello Edo", result, "result must be 'Hello Edo'")
	})

	t.Run("Febrian", func(t *testing.T){
		result := HelloWorld("Febrian")
		require.Equal(t,"Hello Febrian", result, "result must be 'Hello Febrian'")
	})
}

func TestMain(m *testing.M){
	//before
	fmt.Println("Before unit test")

	m.Run()

	fmt.Println("After unit test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Tidak bisa run di windows OS")
	}

	result := HelloWorld("Vernando")

	require.Equal(t,"Hello Vernando", result, "result must be 'Hello Vernando'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("Vernando")

	require.Equal(t,"Hello Vernando", result, "result must be 'Hello Vernando'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("Vernando")

	assert.Equal(t,"Hello Vernando", result, "result must be 'Hello Vernando'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldEdo(t *testing.T) {
	result := HelloWorld("Edo")

	if result != "Hello Edo" {
		//error
		t.Error("Harusnya 'Hello Edo'")
	}

	fmt.Println("TestHelloWorldEdo Done")
}

func TestHelloWorldFebrian(t *testing.T) {
	result := HelloWorld("Febrian")

	if result != "Hello Febrian" {
		//error
		t.Fatal("Harusnya 'Hello Febrian'")
	}

	fmt.Println("TestHelloWorldFebrian Done")
}