# CHATTER

Chatter is a simple peer-to-peer UDP chat application. 

<div style="text-align:center"><img align="center" src="https://i.chzbgr.com/original/2734699264/hAED1D724/cheezburger-image-2734699264" alt="ALAN!ALAN!!! STEVE!STEVE!!"></div>

---

## Requirements

- [Go](https://golang.org/dl/) installed (> Go 1.10)
- Gone through [A Tour of Go](https://tour.golang.org/)

---

## Goals

The goal of this workshop session is to understand how to:

- implement a peer discovery service using [UDP multicast](https://en.wikipedia.org/wiki/Multicast).
- keep track of discovered peers.
- broadcast messages to discovered peers.
- implement a basic UDP server to handle incoming messages from discovered peers.
- interact with an existing terminal ui instance to display incoming messages and possibly, `INFO` messages about the application state.

---

## Advices

Use the code in the directory `start` as scaffold. *Comments are placed to guide you through so don't feel worried ^_^*

You can refer to the code in the directory `final` as reference if you ever get stuck ;)

---

## Gotchas

Code in the directory `final` is **NOT** a state-of-the-art implementation of Chatter. Some details have been left out intentionally to give you room for creativity and some debugging ;)

---

**GOOD LUCK AND HAVE FUN!**
