# Shell Auto Completion

If you're using JFrog CLI from a bash, zsh, or fish shell, you can install JFrog CLI's auto-completion scripts to improve your command-line experience. Auto-completion helps save time and reduces errors by suggesting potential command options and arguments as you type.

Auto-completion allows you to:

1. **Increase Efficiency**: Quickly fill in commands and arguments without typing them out fully.
2. **Reduce Errors**: Minimize typographical errors in commands and options.
3. **Discover Commands**: Easily explore options for specific commands with in-line suggestions.

#### What is JFrog CLI? <a href="#what-is-jfrog-cli" id="what-is-jfrog-cli"></a>

JFrog CLI is a command-line interface for interacting with JFrog Artifactory and other JFrog products. It simplifies various functions, such as uploading or downloading files, managing repositories, and more. For more information, refer to [Jfrog CLI.](https://docs.jfrog-applications.jfrog.io/jfrog-applications/jfrog-cli)

#### How to Enable Auto-Completion? <a href="#how-to-enable-auto-completion" id="how-to-enable-auto-completion"></a>

The method of enabling auto-completion varies based on the shell you are using (bash, zsh, or fish).

#### Install JFrog CLI with Homebrew <a href="#install-jfrog-cli-with-homebrew" id="install-jfrog-cli-with-homebrew"></a>

If you're installing JFrog CLI using Homebrew, the bash, zsh, or fish auto-complete scripts are automatically installed. However, you need to ensure that your `.bash_profile` or `.zshrc` files are correctly configured. Refer to the [Homebrew Shell Completion documentation](https://docs.brew.sh/Shell-Completion) for specific instructions.

&#x20;

#### Using Oh My Zsh? <a href="#using-oh-my-zsh" id="using-oh-my-zsh"></a>

If you are using the Oh My Zsh framework, follow these steps to enable JFrog CLI auto-completion:

1.  Open your zsh configuration file, located at `$HOME/.zshrc`, with any text editor."

    `your-text-editor $HOME/.zshrc`
2. Locate the line starting with `plugins=`.
3.  Add `jfrog` to the list of plugins. For example:

    `plugins=(git mvn npm sdk jfrog)`
4. Save and close the file.
5.  Finally, apply the changes by running:

    `source $HOME/.zshrc`

&#x20;

#### Other Installation Methods <a href="#other-installation-methods" id="other-installation-methods"></a>

If you're not using Homebrew or Oh My Zsh, you can manually install the auto-completion scripts for your specific shell:

**For Bash**

To install auto-completion for bash, run the following command:

```
jf completion bash --install
```

Follow the on-screen instructions to complete the installation.

**For Zsh**

To install auto-completion for zsh, run the following command:

```
jf completion zsh --install
```

Again, follow the instructions provided during the installation process.

**For Fish**

To install auto-completion for fish, run the following command:

```
jf completion fish --install
```

Ensure you follow the relevant instructions to finalize the setup.

&#x20;

#### Verifying Installation <a href="#verifying-installation" id="verifying-installation"></a>

After installing the completion scripts, you can verify that auto-completion works by typing `jf` followed by pressing the `Tab` key. You should see a list of available commands and options.





