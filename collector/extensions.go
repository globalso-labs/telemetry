//go:build !windows && !darwin

/*
 * telemetry
 * extensions.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Tue, 24 Sep 2024 13:06:21 -0500 by nick.
 *
 * DISCLAIMER: This software is provided "as is" without warranty of any kind, either expressed or implied. The entire
 * risk as to the quality and performance of the software is with you. In no event will the author be liable for any
 * damages, including any general, special, incidental, or consequential damages arising out of the use or inability
 * to use the software (that includes, but not limited to, loss of data, data being rendered inaccurate, or losses
 * sustained by you or third parties, or a failure of the software to operate with any other programs), even if the
 * author has been advised of the possibility of such damages.
 * If a license file is provided with this software, all use of this software is governed by the terms and conditions
 * set forth in that license file. If no license file is provided, no rights are granted to use, modify, distribute,
 * or otherwise exploit this software.
 */

package collector

import (
	"context"

	"go.globalso.dev/x/telemetry/config"
)

func marshalExtensions(ctx context.Context, telemetry *config.Telemetry) (map[string]interface{}, error) {
	extensions := make(map[string]interface{})

	extensions["file_storage"] = marshalFileStorageExtension(ctx, telemetry)

	return extensions, nil
}

func marshalFileStorageExtension(_ context.Context, _ *config.Telemetry) map[string]interface{} {
	extension := make(map[string]interface{})

	extension["create_directory"] = true
	extension["directory"] = "/var/log/telemetry/file_storage"
	extension["compaction"] = map[string]interface{}{
		"on_start":  true,
		"directory": "/var/log/telemetry/file_storage/compaction",
	}

	return extension
}
