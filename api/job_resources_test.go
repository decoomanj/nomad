package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJobResource_PrefixList(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	c, s := makeClient(t, nil, nil)
	defer s.Stop()

	job := testJob()
	_, _, err := c.Jobs().Register(job, nil)
	assert.Nil(err)

	id := *job.ID
	prefix := id[:len(id)-2]
	resp, _, err := c.JobResources().PrefixList(prefix, "jobs")

	assert.Nil(err)
	assert.NotEqual(0, resp.Index)

	jobMatches := resp.Matches["jobs"]
	assert.Equal(1, len(jobMatches))
	assert.Equal(job.ID, jobMatches[0])
}
