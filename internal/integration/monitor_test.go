package integration_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/buildkite/agent-stack-k8s/v2/internal/controller/monitor"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes/fake"
)

func TestInvalidOrg(t *testing.T) {
	m, err := monitor.New(zap.Must(zap.NewDevelopment()), fake.NewSimpleClientset(), monitor.Config{
		Token:        os.Getenv("BUILDKITE_TOKEN"),
		MaxInFlight:  1,
		PollInterval: time.Second,
		Org:          "foo",
		Tags:         []string{"queue=default", "foo=bar"},
	})
	require.NoError(t, err)

	require.ErrorContains(t, <-m.Start(context.Background(), nil), "invalid organization")
}
