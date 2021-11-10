# linksnitch

A GitHub Action that automatically check if some link in your README.md is broken or not responding.

## How to use

- uses: actions/checkout@v2
- uses: LeandroLS/linksnitch@main

If your README.md  has broken links, the workflow will fail and will show to you whats links is broken.