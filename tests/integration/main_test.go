package integration

import (
	"context"
	"os"
	"testing"

	"github.com/mmadfox/testcontainers/infra"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestIntegration(t *testing.T) {
	if os.Getenv("GPSGEND_INTEGRATION_TESTS_ENABLED") != "true" {
		t.Skip()
	}

	ctx := context.Background()
	infrastructure := infra.NewSets()
	infrastructure.SetupMongo(ctx)
	defer infrastructure.Close()
	require.NoError(t, infrastructure.Err())

	suite.Run(t, &storageMongoSuite{infra: infrastructure})
}
