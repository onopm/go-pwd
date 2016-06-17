# go-pwd

gpwd - print name of current directory

# Description


## bash 利用時のカレントディレクトリを短縮(20文字)

````
PS1='${LOGNAME}@$(hostname -s):$(gpwd 20)$ '
````

深いディレクトリでのカレントディレクトリ表示は、ターミナルを圧迫するので短縮

````
ono@localhost:/home/ono/go/src/github.com/onopm/go-pwd$ pwd
/home/ono/go/src/github.com/onopm/go-pwd
ono@localhost:/home/air/go/src/github.com/onopm/go-pwd$ PS1='ono@localhost:$(gpwd 20)$ '
ono@localhost:~...rc/github.com/onopm/go-pwd$
````

