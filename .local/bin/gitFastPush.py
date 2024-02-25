#! /usr/bin/python3

import os

os.system('git status')

answer = input("Do you want to commit and push? (y/n): ")

if answer == "y":
    os.system('git add .')
    os.system('git commit -m "update"')
    os.system('git push')
    os.system('git status')
else:
    print("Ok, exiting script.")
