# telus-VSG1432-router-password

## What is this?
A simple utility to dump admin, tech, and root, passwords from VSG1432 routers used by Telus.

## Who needs this?
Anybody looking for the randomly generated passwords to access the advanced features of their Telus VSG1432 router (and possibly other ISPs as well).

## Where?
Earth, as far as one can tell.

## When?
Christmas time, shortly after I watched Rogue One.

## Why?
I needed to fix my father-in-laws WiFi while spending a few days there over Christmas holidays (after watching Rogue One).

The router was placed into bridge mode years ago and the need arose to do a factory reset on it in order to troubleshoot what turned out to be a bad Apple Airport (don't place them under pipes that get condensation in extended cold weather, who knew?). The reset caused the Telus router to be updated to a new firmware version that didn't contain hardcoded root/tech passwords like it used to. I guess that is a step forward until one realizes Telus switched to randomly generated passwords that are published via a UPNP service.

This program will discover and print those passwords if you are connected to the same network the router is on.

## How?

Checkout the repository, set your GOPATH to the checked out repository, and run the following from the repository directory while connected to the same network the VSG1432 is on.

```
$ go run src/github.com/stancarney/tvrp/tvrp.go
```
