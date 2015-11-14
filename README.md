# legend of microservices

Mircoservice architechture exaplained with golang.


# TL:DR
Things covered

* service isolation
* service discovery
* service configuration
* service orchestration
* restful APIs
* service registry

# Introduction
This is a very long and detailed document. I will try to explain how microservices work(Yup everything here is according to what I think, so take it with a grain of salt) with the aid of go programming language. 

# The microservice mindset
To think a given problem can be resolved by microservice is trivia. I will try to demostrate this by breaking down a simple problem that will be whole focus of this document. We are to build a barebone twitter clone that we shall call `legend`.

Our `legend` project should do the following.

* Allow mutliple users to use the service
* Allow users to tweet(phony tweetsnot real ones)
* Allow users to mention one another via @ tags


Our `legend` project should have the following

* User registration endpoint. This is where users will create accounts.
* tweeting endpoint. This is where users tweets.
* Home page endpoint. This is where we see the list of the users with their mentions


As you have seen above, I have only outlined "what" the legend project should do. Thes opens up differnd doors on "how" the legend project should be doing it.

# Service breakdown
We can breakdown the legend project into the following sevices.

* user accounts service. Whih offers user's account related functionality.
* storage serice
