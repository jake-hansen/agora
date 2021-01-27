# Agora
[![Build Status](https://travis-ci.com/jake-hansen/agora.svg?token=HxiSsV9Yc7FqVgAY9svs&branch=main)](https://travis-ci.com/jake-hansen/agora)

Agora is a robust REST API that allows you to schedule meetings across various video conferencing platforms.

### Contributing
Getting your development enviroment setup is the first step to contributing to Agora. You will need to have Golang, Docker, and your favorite IDE installed.

Once you've cloned the repo, you'll need to checkout a feature branch from `develop` to start developing. Your feature branch should be named `feature\my-feature` where my-feature is the name of the feature you are working on. Ideally, your feature should be based on one or more open issues. If you are working on a new feature, open a new issue and tag it appropriately to describe what you are working on. Agora uses the [GitFlow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow) workflow which you can read about for more information.

Once your feature is completed, you should test your changes by writing unit tests. You can test your unit tests by running
```shell
$ make unit-test
```
You can also just run `go test`, but using the make target is recommended since the CI/CD platform uses this same Makefile to test and build images of Agora. This ensures that every developer and our pipelines are using the same exact environment.

Once you've confirmed your tests pass you can then build an image of Agora and start it. To do this, run the command
```shell
$ make build-image
```
This will build a Docker image containing an Agora executable. The image will be tagged with `latest` and also the SHA of the last commit on the branch you have checked out.

Once the image is built, you can run the image with
```shell
docker run agora
```
Depending on the changes you made, you might need to customize the Docker container that is running by exposing certain ports, creating volumes, etc. 

Once you are satisfied with the changes you've made - and you've ensured every unit test has passed, you can push your branch and open a Pull Request.

Before pushing your branch, its reccommended that you pull the lastest changes from the `develop` branch and attempt to locally merge your branch with `develop`. If your branch successfully merges, then you are good to push your branch. If your branch has merge conflicts, make sure you resolve those conflicts before you push.

If your PR is accepted and merged to develop, go ahead and delete your feature branch and start working on the next feature by starting the process over again!
