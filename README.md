# xpost

Post text from stdin to X via [xurl](https://github.com/xdevplatform/xurl).

[xurl](https://github.com/xdevplatform/xurl) is the official X API CLI, but
its `post` shortcut takes the text as an argument, not from stdin. xpost fills
that gap: it reads text from stdin, wraps it in the required JSON body, and
hands it off to `xurl`. If xurl ever accepts stdin directly, this tool will
no longer be needed.

## Usage

xpost reads text from stdin. It takes no arguments.

```
echo "hello from xpost" | xpost
xpost < post.txt
```

In Vim, any range can be piped to xpost via `:w !xpost`:
 
```
:'<,'>w !xpost     post the visual selection
:.w !xpost         post the current line
:%w !xpost         post the whole buffer
```

## Behavior

- Trailing whitespace and newlines are stripped. Internal newlines are
  preserved.
- Empty or whitespace-only input is rejected.
- Character count is not validated; X enforces its own limits and rejects
  oversized posts at the API level.
- xurl's stdout, stderr, and exit code are passed through unchanged.

## Exit codes

- `0`: posted successfully
- `1`: input validation failed, or xurl could not be started
- other: propagated from xurl

## License

This project is licensed under the [MIT License](./LICENSE).
