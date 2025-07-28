This is a guided project with boot.dev.
# Gator_cli

this is a cli tool to aggregate rss info from blogs.

requires postgres and go to be installed

you can inatall this useing the command:

go inatall github.com/Hedonysym/gatorcli

to use this program yoou will first need to set up a .json file in your home directory: ~/.gatorconfig.json

that json file will be formatted like:
{
  "db_url": "database_connection_string_goes_here",
  "current_user_name": "this_will_be_overwritten_by_the_program"
}

once you do that you can run the program with:
gator <your> <arguments>

some commands to get started

register (username)
    - registers a username with the database

login (username)
    - login to an existing user

users
    - list registered users

reset
    - resets all database info

addfeed (url)
    - adds a blog feed to the database, the currentuser automatically follows it

follow (url)
    - current user follows an existing feed

unfollow (url)
    - current user unfollowes a feed, if no users follow a feed it will be deleted

getfeeds 
    - lists followed feeds by current user

browse (optional page limit)
    - shows latest posts in your feed

agg (time duration)
    - continously gathers posts from followed blogs, i suggest you use a duration at least 60s and leave this opwn in its own terminal window while browsing