![build](https://travis-ci.org/noyoshi/weatherparser.svg?branch=master)

To run in docker, make sure you have 
`WEATHER_KEY` set in your enviornment to a valid API key from openweathermap.org, and make sure you have build the docker image. 

Use

`$ ./run <IMAGE_ID>` 

to run it in docker.

If you need an API key, go to the website (its free). 

If you need to build the docker image, use 

`$ docker build .` 

from the root of the project directory to build it. You can get the image id from the output of this command, or from `$ docker images`.
