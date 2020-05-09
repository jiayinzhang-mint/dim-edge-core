package influxdb

import (
	"context"
	"dim-edge/core/protocol"
)

// ListAllBuckets get bucket list
func (c *Client) ListAllBuckets(p *protocol.ListAllBucketsParams) (b []*protocol.Bucket, err error) {
	res, err := c.Store.ListAllBuckets(context.TODO(), p)
	if err != nil {
		return
	}

	b = res.Bucket
	return
}

// RetrieveBucket get one bucket
func (c *Client) RetrieveBucket(p *protocol.RetrieveBucketParams) (b *protocol.Bucket, err error) {
	b, err = c.Store.RetrieveBucket(context.TODO(), p)
	return
}

// RetrieveBucketLog get bucket log
func (c *Client) RetrieveBucketLog(p *protocol.RetreiveBucketLogParams) (l []*protocol.RetreiveBucketLogRes_Log, err error) {
	logs, err := c.Store.RetrieveBucketLog(context.TODO(), p)
	if err != nil {
		return
	}

	l = logs.Log
	return
}
