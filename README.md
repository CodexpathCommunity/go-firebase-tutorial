<p align="center"><img src="assets/logo_image.png" alt="Golang-Firebase-Storage" height="200px"></p>
<p align="center"><img src="assets/logo_text.png" alt="Golang-Firebase-Storage" height="31px"></p>

<p align="center">
<img  src="https://img.shields.io/badge/Go-29BEB0?style=for-the-badge&logo=go&logoColor=white">&nbsp;
<img  src="https://img.shields.io/badge/-FIREBASE-FFCA28?logo=firebase&style=for-the-badge&logoColor=white">&nbsp;
</p>
 
 # go-firebase-storage -Work in progress🛠️

 Simple [Golang](https://golang.org/) API that uses [Firebase](https://firebase.google.com/) as its backend to demonstrate various firebase services using Go such as uploading a simple post to [Firebase Firestore](https://firebase.google.com/docs/firestore), multipart/form-file upload to [Fireabase Storage](https://firebase.google.com/docs/storage) and retrieving url of uploaded file, [Firebase/Social Authentication](https://firebase.google.com/docs/auth) and [Firebase Cloud Messaging](https://firebase.google.com/docs/cloud-messaging/).

 I created this repository for a few reasons: 
 1. I was trying to do a fle upload to firebase storage and had a really hard time going through the poor documentation provided, so i decided to make this repo as a simple guide for anyone who wants to use Firebase services in their golang app.
 2. To practice making a [RESTful API](https://www.redhat.com/en/topics/api/what-is-a-rest-api) using golang.
 3. To have some sort of template for my personal side projects that use [firebase services](https://firebase.google.com/products-build).

 ## Table of Content
- [Prerequisite](#prerequisite)
    - [Disclaimer](##disclaimer)
- [Firebase Services](#firebase-services)
    - [Firebase Firestore](##firebase-firestore)
    - [Firebase Storage](##firebase-storage)
    - [Firebase Cloud Messaging](##firebase-cloud-messaging)
- [Tech Stack](#techstack)
    - [Dependencies](##dependencies)
    - [CI/CD](##ci-cd)
- [Helpful Resources](#helpful-resources)
- [Work In Progress](#work-in-progress)

# Prerequisite
You need to have a firebase account in order to get started. A firebase account can easily be created from the [Firebase Console]("link-to-firebase-console") and they provide various payment plan but for the firebase services used in this repo they are all free by using the Spark Plan.

## Disclaimer

- Most of the firebase services used in this project such as Firebase Firestore and Firestore Storage have other alternatives that are mainly used at a production level such as [AWS](https://aws.amazon.com/sdk-for-go/), [GCP](https://cloud.google.com/go) and [Heroku](https://devcenter.heroku.com/articles/getting-started-with-go) etc.

- This repo is not meant to showcase any best practices or software architecture when developing a REST API using golang or when interacting with the Firebase Golang SDK.

- This repo is not meant to be any form of documentation for the Firebase Golang SDK but just a guide to anyone who might be having a difficult time navigating through the [official documentation](https://github.com/firebase/firebase-admin-go).

# Helpful Resources
- The official documentation for the [Golang Firebase SDK](https://firebase.googleblog.com/2017/08/introducing-firebase-admin-sdk-for-go.html)
- [CRUD RESTful API](https://levelup.gitconnected.com/crud-restful-api-with-go-gorm-jwt-postgres-mysql-and-testing-460a85ab7121) blog post by [Victor Stevens]
- [Uploading files in Go](https://tutorialedge.net/golang/go-file-upload-tutorial/) blog post by [Tutorial Edge](https://tutorialedge.net/)
