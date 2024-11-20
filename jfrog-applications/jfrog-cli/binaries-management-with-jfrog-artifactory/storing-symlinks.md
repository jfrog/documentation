# Storing Symlinks in Artifactory

JFrog CLI lets you upload and download artifacts from your local file system to Artifactory, this also includes uploading symlinks (soft links).

Symlinks are stored in Artifactory as files with a zero size, with the following properties:\
**symlink.dest** - The actual path on the original filesystem to which the symlink points\
**symlink.destsha1** - the SHA1 checksum of the value in the **symlink.dest** property

To upload symlinks, the `jf rt upload` command should be executed with the `--symlinks` option set to true.

When downloading symlinks stored in Artifactory, the CLI can verify that the file to which the symlink points actually exists and that it has the correct SHA1 checksum. To add this validation, you should use the `--validate-symlinks` option with the `jf rt download` command.
