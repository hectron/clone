# clone

A small command-line tool to clone repositories. This is a wrapper for `git clone`, but reduces the need to specify the
full url of the repository, and you don't need to type `git`.

## Usage

```zsh
git clone git@github.com:hectron/notes.git

# or the equivalent

clone hectron/notes

# if the CLI was compiled with the default owner as "hectron"

clone notes
```

## Installing

Compile the CLI, and move it to a location in your `$PATH`:

```zsh
# compile

make build

# Assuming that $HOME/.local/bin is in your $PATH

cp bin/clone ~/.local/bin/clone
```

If you'd like to make this work for GitLab, you can update the **ldflags** option in the `Makefile`'s **build** step:

```make
build: clean update-remote
  go build -o bin/clone -ldflags="-X 'main.GitRemoteUrl=gitlab.com'"
```

If you'd like to change the default owner, update the ldflags in the `Makefile`'s build step:

```make
build: clean update-remote
  go build -o bin/clone -ldflags="-X 'main.DefaultOwner=rails'"
```

## Testing

```zsh
make test
```
