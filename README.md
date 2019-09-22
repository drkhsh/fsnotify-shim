# fsnotify shim

This is a no-op shim for [fsnotify](https://github.com/fsnotify/fsnotify)
for platforms which don't support it, such as Plan 9.

## TODO

- Use goroutines and polling to emulate fsnotify?
