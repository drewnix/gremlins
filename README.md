# Gremlins

## Introduction

Gremlins is a portfolio project to build a distributed application for orchestrating independent 
agents (called "gremlins"). 

The project currently consists of:

* Registry: services can dynamically register themselves as a particular service type on startup allowing service discovery by other services.
* LogServer: services can utilize a central logging service to log important messages.
* Keeper: the keeper is a UI and back-end API for interacting with individual gremlins.
* Gremlin: these are independent agent services that can take commands sent to them by the keeper service.

The eventual goal is to build this into a distributed CI/CD platform that runs on kubernetes, as both an educational project, a fun way to explore building distributed services, and also an excuse to learn Go and hopefully illustrate some of my good development practices to a larger audience.
