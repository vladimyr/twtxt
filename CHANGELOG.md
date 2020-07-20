
<a name="0.0.4"></a>
## [0.0.4](https://github.com/prologic/twtxt/compare/0.0.3...0.0.4) (2020-07-21)

### Bug Fixes

* Fix links opening in new window with target=_blank
* Fix typo on support page ([#5](https://github.com/prologic/twtxt/issues/5))
* Fix app versioning and add to base template so we can tell which version of twtxt is running
* Fix bug in TwtfileHandler with case sensitivity of nick param

### Features

* Add delete account support
* Add better layout of tweets so they stand out more
* Add support for Markdown formatting ([#10](https://github.com/prologic/twtxt/issues/10))
* Add pagination support ([#9](https://github.com/prologic/twtxt/issues/9))
* Add Follow/Unfollow to /discover view that understands the state of who a user follows or doesn't ([#8](https://github.com/prologic/twtxt/issues/8))

### Updates

* Update README.md
* Update README.md


<a name="0.0.3"></a>
## [0.0.3](https://github.com/prologic/twtxt/compare/0.0.2...0.0.3) (2020-07-19)

### Bug Fixes

* Fix bug with NormalizeURL() incorrectly translating https:// to http://
* Fix deps and cleanup unused ones
* Fix the layout of thee /settings view
* Fix a critical bug whereby users could re-register the same username and override someone else's account :/
* Fix username case sensitivity and normalize by forcing all usersnames to be lowercase and whitespace stripped
* Fix useability issue where some UI/UX would toggle between authenticated and unauthentiated state causing confusion

### Features

* Add support for configuring flags from the environment via the same long option names
* Add options to configure session cookie secret and expiry
* Add Contributing guideline and TODO
* Add additional logic to fix null/blank user account URL(s) to the FixUserAccountJob as well
* Add FixUserAccountsJob to fix existing broken accounts that might already exist
* Add new /discover view for convenience access to the global instance's timeline  with easy to use Follow links for discoverability


<a name="0.0.2"></a>
## [0.0.2](https://github.com/prologic/twtxt/compare/0.0.1...0.0.2) (2020-07-19)

### Bug Fixes

* Fix the  follow self behaviour to actually work
* Fix defaults to use the same  ones in twtxt's options
* Fix  URL() of User objects
* Fix import and hard-code no. of tweets to display

### Features

* Add feature whereby new registered users follow themselves by default
* Add support, stargazers and contributing info to READmE
* Add enhanced server capability for graceful/clean shutdowns
* Add /import feature to import multiple feeds at once ([#1](https://github.com/prologic/twtxt/issues/1))

### Updates

* Update feed update frequency to 5m by default


<a name="0.0.1"></a>
## 0.0.1 (2020-07-18)

### Bug Fixes

* Fix release tool
* Fix screenshots
* Fix broken links and incorrect text that hasn't happened yet
* Fix /login CTA link on /register page
* Fix /register CTA link on /login page
* Fix parsing store uri
* Fix bug ensuring feedsDir exists
* Fix Dockerfile

### Features

* Add theme-switcher and remove erroneous prism.js

### Updates

* Update README.md

