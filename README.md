# JacquesFernandes/gopro

## About

CLI tool to connect to a GoPro, enumerate videos, group them based on date and move them off the GoPro

I'm trying to keep the dependencies of this project to an absolute minimum.

## Motivation

Made this tool while needing a way to automate moving recordings I make of my drives.
I needed to do the following:

1. Open the GoPro storage (which uses MTP)
2. Enumerate the files
3. Group them based on their created-on date
4. Move the files from the GoPro (based on their group) into a corresponding directory on my machine
