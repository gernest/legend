# legend of microservices [![Build Status](https://travis-ci.org/gernest/legend.svg)](https://travis-ci.org/gernest/legend)

Mircoservice architechture explained with golang.


__NOTE: This is still under heavy construction, stay safe and don't try this.__

# Synopsis
This is an interactive, fictional story of microservices architecture. There has
been a lot of buzzwords around microservices, and just like you I got courious
and had to go neck deep to unveil this myth.

I will be using the go(Golang) programming language to explain what I know about
microservices. Which means this text will cover interesting parts of the
language too.

To make things a bit interesting, I will design and build a bare bones twitter
clone in a microservice way.

The accounts below are totally my own and by now you should be aware that I am
no expert in anything academic, so take it with a grain of salt. I advice you
read this just like any other boring fiction art. 

# Intended audience
This is for technical people with a background in web development. The knowledge
of go(Golang) is optional as many of the things explained here might apply to
any programming language.

# Things to be covered

* microservice
* microservice archtecture
* service isolation
* service discovery
* service configuration
* service orchestration
* restful APIs
* service registry
* service scaling


# Prerequisite

To follow allong with the code present in this text, you need

* Go( Any version will just do fine)
* A brave heart ( This is "A bridge too far" )

# Here is where the story begins

## Microsevices

As the legends say, whenever you hear the term "Enterprise" know dark days are
ahead of you. Microservice is just a service the rest is buzzwords, and the
micro part of microservice is also a marketing thing. You bare with me, I will
tell you why instead of using microservices I prefer to use web service to imply
the microservice( as we know it).

Web services have been around for a while, and somewhere along the line someone
re branded them by giving them a shiny new name "Microservice". A web
service is just a service running on the web and using web technologies to serve
its purpose( i.e http,rpc,  json, html etc...).

### Microservice architechture

This is a design pattern for a web application where by the main application
relies on external web services to accomplish its intended purpose.


