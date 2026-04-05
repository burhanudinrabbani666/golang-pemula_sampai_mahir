package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Burhanudin",
			request: "Burhanudin",
		},
		{
			name:    "Rabbani",
			request: "Rabbani",
		},
		{
			name:    "Burhanudin Rabbani",
			request: "Burhanudin Rabbani",
		},
		{
			name:    "Budi Nugraha",
			request: "Budi Nugraha",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}

}

func BenchmarkSub(b *testing.B) {
	b.Run("Bani", func(b *testing.B) {
		for index := 0; index < b.N; index++ {
			HelloWorld("Bani")
		}
	})

	b.Run("Burhanudin", func(b *testing.B) {
		for index := 0; index < b.N; index++ {
			HelloWorld("Burhanudin")
		}
	})

}

func BenchmarkHelloWorld(b *testing.B) {
	for index := 0; index < b.N; index++ {
		HelloWorld("Bani")
	}
}

func BenchmarkHelloWorldBurhanudin(b *testing.B) {
	for index := 0; index < b.N; index++ {
		HelloWorld("Burhanudin")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld Bani",
			request:  "Bani",
			expected: "Hello Bani",
		},
		{
			name:     "HelloWorld Rian",
			request:  "Rian",
			expected: "Hello Rian",
		},
		{
			name:     "HelloWorld Heri",
			request:  "Heri",
			expected: "Hello Heri",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, "Result must be "+test.expected)
		})
	}

}

func TestSubTest(t *testing.T) {

	t.Run("Bani", func(t *testing.T) {

		result := HelloWorld("Bani")
		require.Equal(t, "Hello Bani", result, "Result must be Hello Bani")

	})

	t.Run("Udin", func(t *testing.T) {

		result := HelloWorld("Udin")
		require.Equal(t, "Hello Udin", result, "Result must be Hello Bani")

	})

}

func TestMain(m *testing.M) {

	fmt.Println("Before Test")

	m.Run()

	fmt.Println("After Test")

}

func TestSkip(t *testing.T) {

	if runtime.GOOS == "linux" {
		t.Skip("Cant run on mac OS")
	}

	result := HelloWorld("Bani")
	require.Equal(t, "Hello Bani", result, "Result must be Hello Bani")

}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Bani")
	require.Equal(t, "Hello Bani", result, "Result must be Hello Bani")

	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Bani")
	assert.Equal(t, "Hello Bani", result, "Result must be Hello Bani")

	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Bani")

	if result != "Hello Bani" {
		// unit test failed
		t.Fail()
	}

	fmt.Println("Ini test Hello World")
}

func TestHelloWorldBani(t *testing.T) {
	result := HelloWorld("Bani")

	if result != "Hello Bani" {
		// unit test failed
		t.Error("Result must be Hello Bani") // Ini akan lanjut eksekusi line berikutnya
	}

	fmt.Println("Dieksekusi walaupun Error")

}

func TestHelloWorldUdin(t *testing.T) {
	result := HelloWorld("Bani")

	if result != "Hello Bani" {
		// unit test failed
		t.Fatal("Result must be Hello Bani") // Ini tidak akan lanjut eksekusi line berikutnya
	}

	fmt.Println("Tidak dieksekusi ketika Error")
}
