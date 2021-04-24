# Site backup

A backup script that compresses, encrypts and uses Toml as config

Encryption is done with GPG.
Compression with Tar and Gzip.

This binary with create `bash` shells to run most commands.

## Dependencies

- gpg
- rsync
- ssh
- tar
- bash

> :information_source: **MySQL backups**
> MySQL or MariaDB backups require `mysqldump` on the webserver.

## Install

Download the `bk-linux-amd64` (or the binary conforming with your PC) binary from the releases, move it to your desired location.

## Usage

:pencil2: Copy `config-example.toml` to `config.toml` and change the values.

:zap: Run the binary from the location of the `config.toml`.

:tea: Drink some green tea while waiting for the script to finish executing.

:cloud: Copy the `archive-example-2021-04-04-130001.tar.gz.gpg` to your favourite untrusted cloud drive :speak_no_evil:.

:heart: Star this project.

## Security

A `.cache` directory is created to include an unencrypted copy of the last backup to decrease the MB's rsynced over next backup.

The encrypted archives can be safely move to any 3rd party cloud file hosting company.

## Decrypt files

To decrypt archives run:

```sh
gpg -d archive-<site_name>-<datetime>.tar.gz.gpg | tar -xvzf -
```

## Links

- https://www.maketecheasier.com/ssh-pipes-linux/
- https://simplebackups.io/blog/the-complete-mysqldump-guide-with-examples/
- https://andykdocs.de/development/Linux/2013-01-17+Rsync+over+SSH+with+Key+Authentication
- https://linuxconfig.org/how-to-create-compressed-encrypted-archives-with-tar-and-gpg
