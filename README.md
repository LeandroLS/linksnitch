# linksnitch 🔗

A GitHub Action that automatically check the status codes response of links in README.md. Useful for checking broken links.

## How to use
In your `.yml` workflow file:
```yml
name: A workflow that check if links in README.md is working
on: push
jobs:
  build:
    name: Check if README.md links work
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: LeandroLS/linksnitch@main
        with:
          allowedStatusCodes: '[200, 201]'
```
In `allowdStatusCodes` you can set all the status codes that your links can return. If your README.md has some link that don't return one of the status codes set in `allowedStatusCodes` the workflow will fail and will show to you what links is broken.

#### Badges
[![test](https://github.com/LeandroLS/linksnitch/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/LeandroLS/linksnitch/actions/workflows/test.yml)
