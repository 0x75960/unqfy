unqfy
=====

A command line tool.

Copy unique files in directory.

* uniqify by sha256sum of file

Setup
-----

```sh
go get -u github.com/0x75960/unqfy
```

Usage
-----

```sh
# unqfy create dst_dir if direcotry does not exist
unqfy src_dir dst_dir
```

Save files as sha256sum

* e.g. "a.txt" => "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
