package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNsName(t *testing.T) {
	t.Parallel()

	namespace, configMapName, _ := ParseNsName("default/test")
	assert.Equal(t, "default", namespace, "they should be equal")
	assert.Equal(t, "test", configMapName, "they should be equal")

	input := "default_test"
	_, _, err := ParseNsName("default_test")
	expectError := fmt.Errorf("invalid format (namespace/name) found in '%v'", input)

	assert.Error(t, err, expectError, "an error was expected")
}

func TestCheckK8sTag(t *testing.T) {
	t.Parallel()

	var tags []string
	tags = append(tags, "test")
	tags = append(tags, "test1")

	assert.False(t, CheckK8sTag(tags, "kubernetes"), "CheckK8sTag should be false")

	tags = append(tags, "kubernetes")

	assert.True(t, CheckK8sTag(tags, "kubernetes"), "CheckK8sTag should be true")
}

func TestHasLabel(t *testing.T) {
	t.Parallel()

	labels := make(map[string]string)
	labels["pod"] = "selector"

	assert.False(t, HasLabel(labels, "pod=random"), "HasLabel should be false")
	assert.True(t, HasLabel(labels, "pod=selector"), "HasLabel should be true")
	assert.False(t, HasLabel(labels, ""), "HasLabel should be false")
}
