#!/bin/sh
echo "Committing"
git add .
echo "Enter commit message:"
read -r commitMessage
git commit -m "$commitMessage"
git push
git rev-parse HEAD
echo "Done"