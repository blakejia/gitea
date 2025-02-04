// Copyright 2023 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package storage

import (
	"os"
	"testing"
)

func TestMinioStorageIterator(t *testing.T) {
	if os.Getenv("CI") == "" {
		t.Skip("minioStorage not present outside of CI")
		return
	}
	testStorageIterator(t, string(MinioStorageType), MinioStorageConfig{
		Endpoint:        "127.0.0.1:9000",
		AccessKeyID:     "123456",
		SecretAccessKey: "12345678",
		Bucket:          "gitea",
		Location:        "us-east-1",
	})
}
