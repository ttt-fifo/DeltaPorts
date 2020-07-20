--- components/cli/vendor/github.com/containerd/continuity/fs/stat_dragonfly.go.orig	1970-01-01 01:00:00.000000000 +0100
+++ components/cli/vendor/github.com/containerd/continuity/fs/stat_dragonfly.go	2020-07-16 15:03:17.477583000 +0200
@@ -0,0 +1,42 @@
+/*
+   Copyright The containerd Authors.
+
+   Licensed under the Apache License, Version 2.0 (the "License");
+   you may not use this file except in compliance with the License.
+   You may obtain a copy of the License at
+
+       http://www.apache.org/licenses/LICENSE-2.0
+
+   Unless required by applicable law or agreed to in writing, software
+   distributed under the License is distributed on an "AS IS" BASIS,
+   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
+   See the License for the specific language governing permissions and
+   limitations under the License.
+*/
+
+package fs
+
+import (
+	"syscall"
+	"time"
+)
+
+// StatAtime returns the access time from a stat struct
+func StatAtime(st *syscall.Stat_t) syscall.Timespec {
+	return st.Atim
+}
+
+// StatCtime returns the created time from a stat struct
+func StatCtime(st *syscall.Stat_t) syscall.Timespec {
+	return st.Ctim
+}
+
+// StatMtime returns the modified time from a stat struct
+func StatMtime(st *syscall.Stat_t) syscall.Timespec {
+	return st.Mtim
+}
+
+// StatATimeAsTime returns the access time as a time.Time
+func StatATimeAsTime(st *syscall.Stat_t) time.Time {
+	return time.Unix(int64(st.Atim.Sec), int64(st.Atim.Nsec)) // nolint: unconvert
+}
