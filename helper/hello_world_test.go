package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)
func BenchmarkTable(b *testing.B) {
	benchmarks:= []struct {
		name string
		request string
	}{
		{
			name : "Kurao",
			request: "Kurao",
		},
		{
			name : "Hikari",
			request: "Hikari",
		},
		{
			name : "Indra",
			request: "Indra",
		},
	}

	for _, benchmark :=range benchmarks {
		b.Run(benchmark.name, func(b *testing.B){
			for i := 0 ; i < b.N; i++{
				HelloWorld(benchmark.request)
			}
		})
	}
}
//go test -v -run=tidakAda -bench=BenchmarkSub
func BenchmarkSub(b *testing.B) {
	b.Run("Eko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Eko")
		}
	})
	b.Run("KuraoHikari", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("KuraoHikari")
		}
	})
}

//go test -v -run=tidakAda -bench=.
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Eko")
	}
}
func BenchmarkHelloWorldKurao(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("KuraoHikari")
	}
} 
func TestTableHelloWorld(t *testing.T){
	tests:= []struct {
		name string
		request string
		expected string
		errmessage string
	}{
		{
			name : "Kurao",
			request : "Kurao",
			expected : "Hello Kuro",
			errmessage : "result must be hello Kurao",
		},
		{
			name : "Hikari",
			request : "Hikari",
			expected : "Hello Hikari",
			errmessage : "result must be hello Hikari",
		},
		{
			name : "Wibu",
			request : "Wibu",
			expected : "Hello Wibu",
			errmessage : "result must be hello Wibu",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T){
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result, test.errmessage)
		})
	}
} 

func TestSubTest(t *testing.T){
	t.Run("Kurao", func(t *testing.T) { // go test -v -run=TestSubTest/Kurao
		result := HelloWorld("Kurao")
		require.Equal(t, "Hello Kurao", result, "result must be hello Kurao")
	})
	t.Run("Hikari", func(t *testing.T) {
		result := HelloWorld("Hikari")
		require.Equal(t, "Hello Hikari", result, "result must be hello Hikari")
	})
}

func TestMain(m *testing.M){
	fmt.Println("----------Before Unitest------------")

	m.Run() // eksekusi semua test

	fmt.Println("----------after Unitest------------")
}
func TestSkip(t *testing.T){
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on Windows")
	}
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on Darwin to :v")
	}

	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "result must be hello Eko")
}
func TestHelloWorldRequire(t *testing.T){
	result := HelloWorld("Eko")
	require.Equal(t, "Hello Eko", result, "result must be hello Eko")
	fmt.Println("Test HelloWorld with assertion Done")
}
func TestHelloWorldAssert(t *testing.T){
	result := HelloWorld("Eko")
	assert.Equal(t, "Hello Eko", result, "result must be hello Eko")
	fmt.Println("Test HelloWorld with assertion Done")
}
func TestHelloWorldEko(t *testing.T){
	result := HelloWorld("Eko")
	if result != "Hello Eko" {
		// unit test failed
		// panic("Result is not Hello Eko")
		// t.Fail()
		t.Error("harusnya Hello Eko")
	}
	fmt.Println("Test HelloWorldEko Done")
}

func TestHelloWorldKurao(t *testing.T){
	result := HelloWorld("Kurao")
	if result != "Hello Kurao" {
		// unit test failed
		// t.FailNow()
		t.Fatal("harusnya Hello Kurao YGY")
	}
	fmt.Println("Test HelloWorldKurao Done")
}