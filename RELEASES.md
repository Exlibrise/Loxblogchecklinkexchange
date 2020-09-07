
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
-------------------

- Fixed: Performance decreased when using Lightning Charge and the amount of invoices in the Lightning Charge server increased (issue [#28](https://github.com/philippgille/ln-paywall/issues/28))
- Fixed: Since the introduction of the `ln.Invoice` struct the whole struct was logged instead of just the invoice string

### Breaking changes

> Note: The following breaking changes don't affect normal users of the package, but only those who use their own implementations of our interfaces.

- Changed: The struct `ln.Invoice` now has a field `ImplDepID string` which is required by the middlewares. It's an LN node implementation dependent ID (e.g. payment hash for lnd, some random string for Lightning Charge). (Required for issue [#28](https://github.com/philippgille/ln-paywall/issues/28).)
- Changed: `wall.LNclient` now requires the method `CheckInvoice(string) (bool, error)` to accept the LN node implementation dependent ID instead of the preimage hash as parameter. (Required for issue [#28](https://github.com/philippgille/ln-paywall/issues/28).)

v0.5.0 (2018-09-24)
-------------------

- Added: Support for [c-lightning](https://github.com/ElementsProject/lightning) with [Lightning Charge](https://github.com/ElementsProject/lightning-charge) (issue [#6](https://github.com/philippgille/ln-paywall/issues/6))
    - Note: The current implementation's performance decreases when the amount of invoices in the Lightning Charge server increases. This will be fixed in an upcoming release.
- Added: Package `pay` (issue [#20](https://github.com/philippgille/ln-paywall/issues/20))
    - Interface `pay.LNclient` - Abstraction of a Lightning Network node client for paying LN invoices. Enables developers to write their own implementations if the provided ones aren't enough.
    - Struct `pay.Client` - Replacement for a standard Go `http.Client`