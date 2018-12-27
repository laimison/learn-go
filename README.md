# learn-go

This is a repository to learn Go from the beginning or get back if something is forgotten in the future

## Some tips on Mac OS or Windows GIT Bash

### Backup your original bash profile

`cp ~/.bash_profile ~/.bash_profile.backup 2>/dev/null`

### See full path in your terminal

`grep -qE 'export.*PS1=' ~/.bash_profile || echo "export PS1='\e[36;1m\u:\e[0m\w$ '" >> ~/.bash_profile`

### Loop My Script

Add this function to your user's profile

`grep -qE 'loop.*()' ~/.bash_profile || echo 'loop () { while true; do $*; sleep 4; echo; done; }' >> ~/.bash_profile`

Reopen your terminal and try it

`cd learn-go`

`loop ./hello_world.go`

### Git push in one line

Add this function to your user's profile

``grep -qE 'git_push.*().*{' ~/.bash_profile || echo 'git_push () { git status && echo && echo $* | grep [a-zA-Z] && echo "Pushing as `git config user.name` in 5 seconds (CTRL+C to decline) ..." && sleep 5 && git config --global push.default current && git add --all && git commit -m "$*" && git push ; }' >> ~/.bash_profile``

Reopen your terminal and try it

`cd your_git_project`

`git_push updated array \& simplified a function`

### VIM colors for Mac only

`grep -e 'filetype plugin indent on' -e 'syntax on' -e 'color desert' ~/.vimrc | wc -l | awk -F ' ' '{print $1}' | grep -Evq "^(0|1|2)$" || printf "\nfiletype plugin indent on\nsyntax on\ncolor desert\n" >> ~/.vimrc`

## References

Parsing json in Golang includes multiple cases (for example JSON Arrays, etc.) [https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/](https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/)

Learning Go from zero to hero [https://medium.freecodecamp.org/learning-go-from-zero-to-hero-d2a3223b3d86](https://medium.freecodecamp.org/learning-go-from-zero-to-hero-d2a3223b3d86)
