# Welcome to Bonno

## Getting Started

### Start the web server

    $ git clone https://github.com/bandoshintaro/bonno.git
    $ cd bonno
    $ docker-compose up -d

#### store data

    $ docker run -it -v .conf:/go/src/bonno/conf bando/bonno

#### use own movie data

    $ docker run -it -v /your/own/movie/dir:/go/src/bonno/public/douga bando/bonno
