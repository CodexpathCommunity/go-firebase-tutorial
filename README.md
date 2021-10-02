<p align="center"><img src="assets/logo_image.png" alt="Golang-Firebase-Storage" height="200px"></p>
<p align="center"><img src="assets/logo_text.png" alt="Golang-Firebase-Storage" height="31px"></p>

<p align="center">
<img  src="https://img.shields.io/badge/Go-29BEB0?style=for-the-badge&logo=go&logoColor=white">&nbsp;
<img  src="https://img.shields.io/badge/-FIREBASE-FFCA28?logo=firebase&style=for-the-badge&logoColor=white">&nbsp;
</p>
 
 # go-firebase-storage -Work in progress🛠️

 Simple [Golang]("link-to-golang-official-page") API that uses [Firebase]("link-to-firebase-official-page") as its backend to demonstrate various firebase services using Go such as uploading a simple post to [Firebase Firestore]("link-to-firebase-firestore-doc"), multipart/form-file upload to [Fireabase Storage]("link-to-firebase-storage-doc") and retrieving url of uploaded file, [Firebase/Social Authentication]("link-to-firebase-auth-doc") and [Firebase Cloud Messageing]("link-to-fcm-doc").

 I created this repository for a few reasons: 
 1. I was trying to do a fle upload to firebase storage and had a really hard time going through the poor documentation provided, so i decided to make this repo as a simple guide for anyone who wants to use Firebase services in their golang app.
 2. To practice making a [RESTful API]("link-to-restful-api-explaination") using golang.
 3. To have some sort of template for my personal side projects that use [firebase services]("firebase-services-page").

 ## Table of Content
- [Prerequisite](#prerequisite)
    - [Disclaimer](##disclaimer)
- [Firebase Services](#firebase-services)
    - [Firebase Firestore]("firebase-firestore")
    - [Firebase Storage]("firebase-storage")
    - [Firebase Cloud Messaging]("firebase-cloud-messaging")
- [Tech Stack](#techstack)
    - [Dependencies](##dependencies)
    - [CI/CD](##ci-cd)
- [Helpful Resources](#helpful-resources)
- [Work In Progress](#work-in-progress)

# Prerequisite
You need to have a firebase account in order to get started. A firebase account can easily be created from the [Firebase Console]("link-to-firebase-console") and they provide various payment plan but for the firebase services used in this repo they are all free by using the Spark Plan.

## Disclaimer

- Most of the firebase services used in this project such as Firebase Firestore and Firestore Storage have other alternatives that are mainly used at a production level such as [AWS]("link-to-aws-go"), [GCP]("link-to-gcp-go-doc") and [Heroku]("link-to-heroku-go-doc") etc.

- This repo is not meant to showcase any best practices or software architecture when developing a REST API using golang or when interacting with the Firebase Golang SDK.

- This repo is not meant to be any form of documentation for the Firebase Golang SDK but just a guide to anyone who might be having a difficult time navigating through the [official documentation]("link-to-offical-firebase-golang-sdk-doc").

Let's get started...

# Firebase Services
## Firebase Storage
