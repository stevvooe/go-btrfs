/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package btrfs

import "os"

func getFileRootID(file *os.File) (objectID, error) {
	args := btrfs_ioctl_ino_lookup_args{
		objectid: firstFreeObjectid,
	}
	if err := iocInoLookup(file, &args); err != nil {
		return 0, err
	}
	return args.treeid, nil
}

func getPathRootID(path string) (objectID, error) {
	fs, err := Open(path, true)
	if err != nil {
		return 0, err
	}
	defer fs.Close()
	return getFileRootID(fs.f)
}
