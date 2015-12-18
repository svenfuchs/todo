# TODO.txt [![Build Status](https://secure.travis-ci.org/svenfuchs/todo.png?branch=master)](https://travis-ci.org/svenfuchs/todo.txt)

My personal flavor of the [Todo.txt](https://github.com/ginatrapani/todo.txt-cli/wiki/The-Todo.txt-Format)
format. This is a Go port of my [Ruby CLI](https://github.com/svenfuchs/todo.txt).

## Assumptions

These are my personal opinions, and probably due to how I am used to use my
todo.txt file personally. Your mileage will most probably vary.

* Being able to embed todo.txt item lines into arbirary text files, and format
  a todo.txt file freely (adding section titles, comments, bullet point lists
  etc) is a good thing. A tool must play nice with this and must not rewrite
  any custom formatting.
* An item can belong to one or many projects.
* The concept of "contexts" (phone, work, home) seems like a stale relict from
  the GTD era to me. Does anybody actually use this? Also, the format `@name`
  indicates a person in most contexts nowadays.
* Assigning explicit priorities don't really work well. Re-ordering items on
  the todo.txt list works better.
* A concept of generic key/value pairs seems like a useful addition to make the
  format more flexible and adaptable. These also can be used for, e.g. due and
  done dates (as well as contexts and priorities, if they still seem useful).
* In order to integrate with services and other tools it seems useful to add
  the concept of an `id`. I'll use `[id]` for this. Typing ids is a hassle, so
  the tool should add them automatically.

## Installation

```
git clone https://github.com/svenfuchs/todo
cd todo
go build
# add ./build to your $PATH
```

## Usage

All dates can be given as `YYYY-mm-dd`, or as named date strings, e.g. `today`,
`yesterday`, `two weeks ago`, `last monday` etc.

List items:

```
$ todo list --help
usage: todo list [<flags>] [<input>]

Filter and list todo items.

Flags:
      --help               Show context-sensitive help (also try --help-long and --help-man).
  -f, --file=FILE          Todo.txt file to work with.
  -o, --format=FORMAT      Output format.
  -i, --id=ID              Filter by id.
  -s, --status=STATUS      Filter by status.
  -t, --text=TEXT          Filter by text.
  -p, --projects=PROJECTS  Filter by projects (comma separated).
  -a, --date=DATE          Filter by done date.
  -a, --after=AFTER        Filter by done after.
  -n, --since=SINCE        Filter by done since.
  -b, --before=BEFORE      Filter by done before.

Args:
  [<input>]  Filter by full line.
```

Toggle items:

```
$ todo toggle --help
usage: todo toggle [<flags>] [<input>]

Toggle todo items.

Flags:
      --help               Show context-sensitive help (also try --help-long and --help-man).
  -f, --file=FILE          Todo.txt file to work with.
  -i, --id=ID              Filter by id.
  -s, --status=STATUS      Filter by status.
  -t, --text=TEXT          Filter by text.
  -p, --projects=PROJECTS  Filter by projects (comma separated).
  -a, --date=DATE          Filter by done date.
  -a, --after=AFTER        Filter by done after.
  -n, --since=SINCE        Filter by done since.
  -b, --before=BEFORE      Filter by done before.
  -r, --report             Print a report to stderr.

Args:
  [<input>]  Filter by full line.
```

Archive items:

```
$ todo archive --help
usage: todo archive [<flags>]

Archive done todo items.

Flags:
      --help             Show context-sensitive help (also try --help-long and --help-man).
  -f, --file=FILE        Todo.txt file to work with.
  -a, --archive=ARCHIVE  File to archive to.
  -o, --format=FORMAT    Output format.
  -b, --before=BEFORE    Filter by done before.
  -r, --report           Print a report to stderr.
```

Push items to Idonethis:

```
$ todo push idonethis --help
usage: todo push idonethis [<flags>] [<input>]

Service Idonethis

Flags:
      --help                Show context-sensitive help (also try --help-long and --help-man).
  -f, --file=FILE           Todo.txt file to work with.
  -o, --format=FORMAT       Output format.
  -a, --date=DATE           Filter by done date.
  -a, --after=AFTER         Filter by done after.
  -n, --since=SINCE         Filter by done since.
  -r, --report              Print a report to stderr.
      --username=USERNAME   Idonethis username
      --token=TOKEN         Idonethis token
      --team=TEAM           Idonethis team

Args:
  [<input>]  Filter by full line.
```

## Examples

```
# Input files and stdin

$ cat todo.txt | todo toggle foo   # outputs to stdout
$ todo toggle foo --file todo.txt  # specify the file to work with
$ todo toggle foo                  # should assume ./TODO.txt but i can't figure
                                   # this out, see https://github.com/svenfuchs/todo.txt/blob/master/lib/todo/cli/cmd.rb#L29

# Filtering items

$ todo list --since 2015-12-01                     # by done date
$ todo list --after 2015-11-01 --before 2015-12-01 # by done date
$ todo list --due 2015-12-07                       # by due date
$ todo list --status pending                       # by status
$ todo list --status done                          # by status
$ todo list --project foo                          # by project
$ todo list --project foo --project bar            # by project
$ todo list --text foo                             # by text
$ todo list foo                                    # by text
$ todo list foo --since 2015-12-01 --status done   # by text, done date, and status

# Named dates

$ todo list --due tomorrow
$ todo list --since yesterday
$ todo list --since one.week.ago
$ todo list --since 'two weeks ago'
$ todo list --since 1_year_ago
$ todo list --since 'last friday'

# Formats

$ todo list --format short
$ todo list --format full
$ todo list --format id,text,tags

# Toggling

$ todo toggle foo
$ todo toggle --text foo
$ todo toggle -- '- foo' # passing a full item line with a leading `-`

# Archiving

$ todo archive --since 2015-12-01
$ todo archive # defaults to: two weeks ago
```


## Vim integration

* Vim mapping to a todo item status [toggle](https://github.com/svenfuchs/vim-todo.txt/blob/master/ftplugin/todo.vim#L1)
* Vim mapping to done items to idonethis [push](https://github.com/svenfuchs/vim-todo.txt/blob/master/ftplugin/todo.vim#L2)
* Vim mapping to done items to a separate file [archive](https://github.com/svenfuchs/vim-todo.txt/blob/master/ftplugin/todo.vim#L3)

## Mac OSX integration

* to `fswatch` the file, and [git changes to a Gist [launchctl](https://github.com/svenfuchs/todo.txt/blob/master/etc/me.todo-watch.plist) push](https://github.com/svenfuchs/todo.txt/blob/master/bin/push)
