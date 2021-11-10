# linksnitch

A GitHub Action that automatically check if some link in your README.md is broken or not responding.

## How to use
```yml
name: A workflow that check if links in README.md is working
on: push
jobs:
  build:
    name: Check if README.md links workds
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - uses: LeandroLS/linksnitch@main
```

If your README.md  has broken links, the workflow will fail and will show to you what links is broken.