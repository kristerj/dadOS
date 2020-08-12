# dadOS

## Installation

Probably don't. Its not ready.

## What is this?

dadOS is not actually an OS. It aims to be a big keyboard-based desktop environment. It exists because I was looking for, and could not find, a very specific user experience.

## Design Goals

- 100% keyboard-driven
- Easy to hack on, especially for new programmers - emacs but without the baggage
- Single language
- Uniform UI
  - Uniform key bindings
  - Easily customizable color schemes
  - Must be responsive
    - Responsive buttons - show some state change on buttons before taking action
    - Snappy UI handling
- Run on old hardware
- Very basic user administration and auth for LANs 
  - Client autodiscover
  - User data stored in local database for each client, with updates replicated to other LAN clients
  - Start with dedicated master, hopefully move to a voting system to elect somewhere down the road
- VIM bindings wherever possible
- New services (tasks) must be trivial to add
  - One screen per service
  - Navigation and background service architecture already done - no thought required
    - One background worker thread spun up when dadOS starts per service
    - Each service defined by one file in services directory
    - One get content function where you plop your interface code
    - One background worker function where you plop your background task code
- Performs (my) common tasks:
  - Note-taking
  - MP3 playing
  - Wikipedia browsing
  - Mail client
  - IM client (with other dadOS terminals)
  - Calculator

## Why?

- I don't need a big bloated desktop GUI - there is no reason I shouldn't be able to be productive and do 90% of what I need on old hardware with a keyboard and console
- I like monochrome-ish text interfaces
- Mice are slow and I hate them
- I just want one big app with the stuff I use, navigable by keyboard
- I think its cool to boot into one app and automagically be able to communicate with the other devices on the network running that app. I don't want to have to install skype or teams or IRC or some dumb thing just to send a message to the computer in the kitchen
- To give my kids a stripped-down computing experience they can safely play with and learn from
- To turn all the x86 and ARM devices scattered around home networks useful lives as visual intercoms, music players, and information repositories
- I enjoy giving myself things to do that I will never actually accomplish
