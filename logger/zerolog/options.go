/*
 * telemetry
 * options_generic.go
 * This file is part of telemetry.
 * Copyright (c) 2024.
 * Last modified at Wed, 10 Jul 2024 00:23:06 -0500 by nick.
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

package zerolog

import (
	"strconv"
)

var MaxPathNumber = 0

func countParentDirectories(filePath string, maxPath int) int {
	parentsCounter := -1
	for i := len(filePath) - 1; i > 0; i-- {
		if filePath[i] == '/' {
			parentsCounter++
		}
		if parentsCounter == maxPath {
			return i + 1
		}
	}

	return 0
}

var CallerMarshalFunc = func(_ uintptr, file string, line int) string {
	file = file[countParentDirectories(file, MaxPathNumber):]
	return file + ":" + strconv.Itoa(line)
}
