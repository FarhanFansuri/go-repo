package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	// "github.com/stretchr/testify/mock"
)

func Test_Hello_Jaka(t *testing.T) {

	if runtime.GOOS == "windows" {
		t.Skip("Cant run on windows")
	}
	result := Hello("Jak")

	assert.Equal(t, "Hello Jaka", result, "Result must Be Hello Jaka")
	fmt.Println("Dieksekusi")

}

func Test_Hello_Farhan(t *testing.T) {
	result := Hello("Farha")
	require.Equal(t, "Hello Farhan", result, "Result must Be Hello Farhan")
	fmt.Println("dieksekusi")
}

func TestSubTest(t *testing.T) {
	t.Run("Farhan", func(t *testing.T) {
		result := Hello("Farha")
		require.Equal(t, "Hello Farhan", result, "Result must Be Hello Farhan")
		fmt.Println("dieksekusi")
	})
	t.Run("Jaka", func(t *testing.T) {
		result := Hello("Jaka")
		require.Equal(t, "Hello Jaka", result, "Result must Be Hello Jaka")
		fmt.Println("dieksekusi")
	})
}

func TestMain(t *testing.M) {
	fmt.Println("Before Testing")
	t.Run()
	fmt.Println("After Testing")
}
