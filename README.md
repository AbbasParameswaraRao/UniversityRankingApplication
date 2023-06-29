# UniversityRankingApplication
There is a Univerisity rank json file have rank and uneveristy details.

Created a REST Api using Go and Gin

Implemented the Cache,logging and Viper.

there is two GET functions

http://localhost:8080/universities

getting response all universites with their respective ranks

http://localhost:8080/university/:rank

Getting responce for given rank it also specifies the wether data retrieved from databae or cache.

http://localhost:8080/deluniversity/:rank

Deleted university for given rank


http://localhost:8080/updateuniversity/:rank

updated university with given rank
