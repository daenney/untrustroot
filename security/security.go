// Copyright © 2016 Daniele Sluijters
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package security

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Execute macOS security CLI with the given arguments targetting the
// specified keychain.
func do(keychain string, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	args = append(args, keychain)
	cmd := exec.CommandContext(ctx, "security", args...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	if ctx.Err() == context.DeadlineExceeded {
		return nil, fmt.Errorf("Command 'security %v' took longer than 5s", args)
	}
	return out, nil
}

// ReadCertsFromKeychain will return any certificates found in the specified
// keychain in PEM-format.
func ReadCertsFromKeychain(path string) ([]byte, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("Cannot read keychain: %s", path)
	}
	return do(path, "find-certificate", "-a", "-p")
}
