# Welcome to Bonno 

[![Join the chat at https://gitter.im/bando_bonno/Lobby](https://badges.gitter.im/bando_bonno/Lobby.svg)](https://gitter.im/bando_bonno/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Build Status](https://drone.io/github.com/bandoshintaro/bonno/status.png)](https://drone.io/github.com/bandoshintaro/bonno/latest)

## Getting Started

### Start the web server:

    $ git clone https://github.com/bandoshintaro/bonno.git
	$ cd bonno
    $ docker-compose up -d

#### store data
    $ docker run -it -v .conf:/go/src/bonno/conf bando/bonno
#### use own movie data
    $ docker run -it -v /your/own/movie/dir:/go/src/bonno/public/douga bando/bonno
