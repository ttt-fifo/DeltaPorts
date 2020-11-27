--- newlib/libc/string/strchr.c.orig	2016-03-29 21:33:42 UTC
+++ newlib/libc/string/strchr.c
@@ -36,6 +36,14 @@ QUICKREF
 #include <string.h>
 #include <limits.h>
 
+#ifndef LONG_MAX
+#ifdef _LP64
+#define LONG_MAX 0x7fffffffffffffffL
+#else
+#define LONG_MAX 0x7fffffffUL
+#endif
+#endif
+
 /* Nonzero if X is not aligned on a "long" boundary.  */
 #define UNALIGNED(X) ((long)X & (sizeof (long) - 1))
 
