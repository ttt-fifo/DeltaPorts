--- newlib/libc/string/memccpy.c.orig	2016-03-29 21:33:42 UTC
+++ newlib/libc/string/memccpy.c
@@ -39,6 +39,14 @@ PORTABILITY
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
 /* Nonzero if either X or Y is not aligned on a "long" boundary.  */
 #define UNALIGNED(X, Y) \
   (((long)X & (sizeof (long) - 1)) | ((long)Y & (sizeof (long) - 1)))
