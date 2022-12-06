# ar
ag-like text replacement tool.

* few arguments; Replace files under the current directory (recursive).
* human readable; Output replacement result as an Unified diff format.

## usage
```
$ ar <before> <after>
```

## in-place mode (default)
```script
$ cat txtfile
A test text file
```

```diff
$ ar test replaced
--- a/txtfile
+++ b/txtfile
@@ -1,1 +1,1 @@
-A test text file
+A replaced text file
```

```script
$ cat txtfile
A replaced text file
```

## dry-run mode
```script
$ cat txtfile
A test text file
```

```diff
$ ar test replaced -dry | tee replace.patch
--- a/txtfile
+++ b/txtfile
@@ -1,1 +1,1 @@
-A test text file
+A replaced text file
```

```script
$ cat txtfile
A test text file

$ patch -p1  replace.patch
patching file txtfile

$ cat txtfile
A replaced text file
```
