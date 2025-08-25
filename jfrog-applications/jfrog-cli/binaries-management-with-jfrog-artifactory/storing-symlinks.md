# Storing Symlinks in Artifactory

JFrog CLI supports uploading and downloading symlinks (soft links) between your local file system and Artifactory.

Symlinks are stored in Artifactory as zero-byte files with the following properties:

* `symlink.dest`: The path on the original file system to which the symlink points.
* `symlink.destsha1`: The SHA1 checksum of the value in the `symlink.dest` property.

To upload symlinks, run the `jf rt upload` command with the `--symlinks` option set to `true`.

When downloading symlinks stored in Artifactory, the CLI can verify that the file to which the symlink points exists and has the correct SHA1 checksum. To add this validation, use the `--validate-symlinks` option with the `jf rt download` command.
