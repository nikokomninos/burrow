<div align="center" style="margin-bottom: 100px">
<img src="./assets/images/logo.png" style="max-width: 500px"></img>
<h1>Burrow</h1>
<p>A cross-platform content blocker that goes beyond static blocklists</p>
</div>

## Table of Contents

<!--toc:start-->
- [Table of Contents](#table-of-contents)
- [Introduction](#introduction)
  - [Motivation](#motivation)
    - [Limitations of a Static Blocklist](#limitations-of-a-static-blocklist)
    - [Limitations of the Current Market Options](#limitations-of-the-current-market-options)
- [Proposed Features](#proposed-features)
- [License](#license)
<!--toc:end-->

## Introduction

<strong>Burrow</strong> is a cross-platform content blocker written in Go. Its
purpose is to go beyond the traditional approach of blocking using static blocklists
and provide a complete toolkit to make sure you're always on track.

### Motivation

There are many content blockers to choose from currently, and they're all great
in their own ways. However, many of these blockers come with their own caveats.

#### Limitations of a Static Blocklist

The classic approach to content blocking is through the use of a static blocklist, which
can be enough for many people. However, static blocklists don't catch the following
edge cases:

- They do not catch intentional misspellings of blocked keywords to find alternative websites
not included in the blocklist
    - e.g: "keyword" -> "ke yword"
- The implementations usually only check the page URL and page title for blocking, which can lead
to access of content that does not have the keyword explicitly in the URL or page
title 
    - e.g: mirrors of websites the user intended to block

#### Limitations of the Current Market Options

While there are many viable options for a blocker available, they make your freedom
difficult to obtain:

- Blockers that offer a free version lock many of the most useful features behind
an (often expensive) paywall
- Blockers that offer a wider feature set are often locked behind a monthly
subscription
- Blockers are often closed-source which inhibits community contribution

## Proposed Features

- Support for Windows, macOS and Linux 
- Content blocking based on static blocklists, with additional detection measures:
    - Normalized input and typo matching to avoid intentional misspellings
    - Use of ML libraries including `sentence-transformers` to check for semantic
    similarities
- Safety measures to prevent uninstallation and prevent usage of browsers that
do not have the `burrow` extension installed
- Community contributed and maintained blocklists
- CLI support for automation and agent interaction
- Online sync for blocklists
- A _self-hostable_ accountability server that allows users to create groups that can:
    - Monitor when a blocked keyword or website is attempting to be accessed
    - Approve a one-time open of a blocked website

## License

This software is published under the GPLv3 License.
