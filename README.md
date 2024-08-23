# mini_wallet_exercise

<!-- ABOUT THE PROJECT -->
## About The Project

This is a technical exercise for software engineering candidates. 
For the exercise, use this documentation to build an API backend service, for managing a simple mini wallet. 

### Built With

This section should list any major frameworks that you built your project using. Leave any add-ons/plugins for the acknowledgements section. Here are a few examples.
* [Golang](https://golang.com)
* [Redis](https://redis.io/)

And supporting tools to easy managing redis caching like :
* [Medis](https://getmedis.com/)

<!-- GETTING STARTED -->
## Getting Started
Before we get started, it's important to know that  before you run this code you have to make sure that `Redis` is already exist and ready to run on your device. Than this code use a custom command to execute it with makefile to make more simple command like :
1. make update
2. make tidy
3. make start

So, let start it.
1. After clone this repository, just run `make update`.
2. Setup your `.env` file such as database connection and redis connection based on default setting on you device.
3. To make sure that all dependency is run well, than run `make tidy`.
4. Finally, you can run your project with command: `make start`.
5. Go to postman and set url like `http://localhost:8080/`, for information that port to run this project depend on configuratin on `.env`

And for additional information, i'm alredy setup unit-testing, just run `make test-service`.

## Afterword
Hopefully, it can be easily understood and useful. Thank you~
