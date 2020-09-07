
Releases
========

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/) and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

vNext
-----

v0.5.2 (2018-10-07)
-------------------

- Fixed: When using the Echo middleware the invoice response status code was 200 OK instead of 402 Payment Required (issue [#30](https://github.com/philippgille/ln-paywall/issues/30))
- Fixed: When using the Echo middleware error responses (including the invoice) were wrapped in JSON instead of just text (issue [#30](https://github.com/philippgille/ln-paywall/issues/30))
- Fixed: GoDoc for `storage.NewBoltClient(...)` contained usage suggestions that would lead to the possibility of clients cheating with reusing preimages

v0.5.1 (2018-10-02)