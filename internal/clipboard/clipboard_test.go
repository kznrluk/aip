package clipboard

import (
	"testing"
)

// Just a minimal test, as clipboard testing might be platform dependent
func TestWriteClipboard(t *testing.T) {
	err := WriteClipboard("test content")
	if err != nil {
		t.Errorf("WriteClipboard() error = %v", err)
	}
	// We won't assert reading back from clipboard in tests due to environment dependency.
}
