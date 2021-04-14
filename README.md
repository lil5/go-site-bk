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

> :information_source: **MySQL backups**
> MySQL or MariaDB backups require `mysqldump` on the webserver.

## Setup

Copy `config-example.toml` to `config.toml` and change the values.

## Security

A `.cache` directory is created to include an unencrypted copy of the last backup to decrease the MB's rsynced over next backup.

The encrypted archives can be safely move to any 3rd party cloud file hosting company.

## Decrypt files

To decrypt archives run:

```sh
gpg -d folder.tar.gz.gpg | tar -xvzf -
```

## Links

- https://www.maketecheasier.com/ssh-pipes-linux/
- https://simplebackups.io/blog/the-complete-mysqldump-guide-with-examples/
- https://andykdocs.de/development/Linux/2013-01-17+Rsync+over+SSH+with+Key+Authentication
- https://linuxconfig.org/how-to-create-compressed-encrypted-archives-with-tar-and-gpg
