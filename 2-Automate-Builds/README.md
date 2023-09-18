# Automating Builds

## Task
Create a workflow file in .github/workflow that will take our app we dockerized in the last section, and automate the process of building the dockerfile and pushing the image to dockerhub.

This workflow file should launch whenever a change is made to the folder containing the app. For the sake of separation, the app has been copied into this folder. 
### Prerequisites
* Fork this repo
* A github account
* A dockerhub account

### Success Criteria

* The workflow starts a github action whenever a change is made to the section 2 folder.
* The action can also be started manually
* The action builds the dockerfile, and pushes the created image to dockerhub
* Bonus - tag the image with todays date, in addition to the default tag of latest.
