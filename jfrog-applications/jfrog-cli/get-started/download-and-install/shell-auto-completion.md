# Shell Auto-Completion

If you're using JFrog CLI from a bash, zsh, or fish shells, you can install JFrog CLI's auto-completion scripts.

## Install JFrog CLI with Homebrew?

If you're installing JFrog CLI using Homebrew, the bash, zsh, or fish auto-complete scripts are automatically installed by Homebrew. Please make sure that your `.bash_profile` or `.zshrc` are configured as described in the [Homebrew Shell Completion documentation](https://docs.brew.sh/Shell-Completion).

## Using Oh My Zsh?

With your favorite text editor, open `$HOME/.zshrc` and add jfrog to the plugin list.
For example:

```sh
plugins=(git mvn npm sdk jfrog)
```

## Other Installation Methods

To install auto-completion for bash, run the following command and follow the instructions to complete the installation:

```sh
jf completion bash --install
```

To install auto-completion for zsh, run the following command and follow the instructions to complete the installation:

```sh
jf completion zsh --install
```

To install auto-completion for fish, run the following command:

```sh
jf completion fish --install
```
