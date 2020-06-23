<a name=""></a>
# [](https://github.com/ncarlier/feedpushr/compare/v3.0.0-rc.3...v) (2020-06-23)


### Bug Fixes

* **aggregator:** omit empty headers from the request ([e83c177](https://github.com/ncarlier/feedpushr/commit/e83c177))
* **api:** make props attribute response mandatory ([177614e](https://github.com/ncarlier/feedpushr/commit/177614e))


### Features

* **filter:** use external scraper for fetch filter ([d862502](https://github.com/ncarlier/feedpushr/commit/d862502))
* **ui:** foldable spec description ([c9a3c02](https://github.com/ncarlier/feedpushr/commit/c9a3c02))



<a name="3.0.0-rc.3"></a>
# [3.0.0-rc.3](https://github.com/ncarlier/feedpushr/compare/v3.0.0-rc.2...v3.0.0-rc.3) (2020-06-08)


### Bug Fixes

* init search index at startup ([83aca10](https://github.com/ncarlier/feedpushr/commit/83aca10))
* **output:** fix readflow text attribute ([96ef893](https://github.com/ncarlier/feedpushr/commit/96ef893))


### Features

* **filter:** add HTTP filter plugin ([3afb377](https://github.com/ncarlier/feedpushr/commit/3afb377))
* **ui:** autofocus on search box ([e7e746a](https://github.com/ncarlier/feedpushr/commit/e7e746a))



<a name="3.0.0-rc.2"></a>
# [3.0.0-rc.2](https://github.com/ncarlier/feedpushr/compare/v3.0.0-rc.1...v3.0.0-rc.2) (2020-05-23)


### Bug Fixes

* **filter:** add missing prop definition ([5a3c575](https://github.com/ncarlier/feedpushr/commit/5a3c575))
* **filter:** only apply enabled filters ([e8a45c6](https://github.com/ncarlier/feedpushr/commit/e8a45c6))
* **output:** improve goroutine stop ([0c0179e](https://github.com/ncarlier/feedpushr/commit/0c0179e))
* **version:** fix version display ([07e46bf](https://github.com/ncarlier/feedpushr/commit/07e46bf))
* **websub:** fix callback URL ([1f33fcd](https://github.com/ncarlier/feedpushr/commit/1f33fcd))


### Features

* **contrib:** add Kafka plugin ([886f878](https://github.com/ncarlier/feedpushr/commit/886f878))
* **feed:** deactivated by default ([e234dda](https://github.com/ncarlier/feedpushr/commit/e234dda))
* **feeds:** full-text search engine ([7b2cc08](https://github.com/ncarlier/feedpushr/commit/7b2cc08))
* **output:** improve readflow output ([6177e3b](https://github.com/ncarlier/feedpushr/commit/6177e3b))
* **ui:** improve base path support ([ddc4a75](https://github.com/ncarlier/feedpushr/commit/ddc4a75))



<a name="3.0.0-rc.1"></a>
# [3.0.0-rc.1](https://github.com/ncarlier/feedpushr/compare/v2.2.0...v3.0.0-rc.1) (2020-04-11)


### Bug Fixes

* **ci:** add missing package ([cc970d1](https://github.com/ncarlier/feedpushr/commit/cc970d1))
* **format:** improve tweet function security ([4e39bf0](https://github.com/ncarlier/feedpushr/commit/4e39bf0))


### Features

* **explorer:** add feed explorer ([aca9077](https://github.com/ncarlier/feedpushr/commit/aca9077))
* **format:** custom format functions ([bdbccd5](https://github.com/ncarlier/feedpushr/commit/bdbccd5))



<a name="2.2.0"></a>
# [2.2.0](https://github.com/ncarlier/feedpushr/compare/v2.1.0...v2.2.0) (2020-03-08)


### Bug Fixes

* **aggregator:** fix concurrent iteration and write ([0bb8124](https://github.com/ncarlier/feedpushr/commit/0bb8124)), closes [#23](https://github.com/ncarlier/feedpushr/issues/23)
* **aggregator:** reset handler status on start ([298a149](https://github.com/ncarlier/feedpushr/commit/298a149))
* **ui:** fix feed control state on pagination ([870eb31](https://github.com/ncarlier/feedpushr/commit/870eb31)), closes [#20](https://github.com/ncarlier/feedpushr/issues/20)
* fix bad package ref ([d95c4cf](https://github.com/ncarlier/feedpushr/commit/d95c4cf))
* fix feed page attributes ([928cfa5](https://github.com/ncarlier/feedpushr/commit/928cfa5))
* fix flag binding for slice ([a3f6a71](https://github.com/ncarlier/feedpushr/commit/a3f6a71))


### Features

* **aggregator:** fan-out delay ([096d1a8](https://github.com/ncarlier/feedpushr/commit/096d1a8))
* **auth:** add Basic Auth support ([3975b90](https://github.com/ncarlier/feedpushr/commit/3975b90))
* **doc:** fan-out delay configuration ([00faee4](https://github.com/ncarlier/feedpushr/commit/00faee4))
* **plugins:** use formatter for Twitter and Mastodon plugins ([1562b88](https://github.com/ncarlier/feedpushr/commit/1562b88))
* **ui:** add feed HTML link ([b5a2ebe](https://github.com/ncarlier/feedpushr/commit/b5a2ebe)), closes [#21](https://github.com/ncarlier/feedpushr/issues/21)
* configuration refactoring ([d7480be](https://github.com/ncarlier/feedpushr/commit/d7480be))
* enable pagination for feeds ([e64b6b2](https://github.com/ncarlier/feedpushr/commit/e64b6b2))
* redirect base URL to UI ([6c26daa](https://github.com/ncarlier/feedpushr/commit/6c26daa))
* retrieve feed link from HTML URL ([922a783](https://github.com/ncarlier/feedpushr/commit/922a783))
* use output formatter ([b69484a](https://github.com/ncarlier/feedpushr/commit/b69484a))



<a name="2.1.0"></a>
# [2.1.0](https://github.com/ncarlier/feedpushr/compare/v2.0.0...v2.1.0) (2020-01-06)


### Bug Fixes

* remove 32 bit support ([b3053d5](https://github.com/ncarlier/feedpushr/commit/b3053d5))
* typos ([820ee03](https://github.com/ncarlier/feedpushr/commit/820ee03))
* **contrib:** ignore Twitter error when duplicate ([f0099f2](https://github.com/ncarlier/feedpushr/commit/f0099f2))


### Features

* add API info endpoint ([a234a23](https://github.com/ncarlier/feedpushr/commit/a234a23))
* conditional expression for filters & outputs ([6b6d283](https://github.com/ncarlier/feedpushr/commit/6b6d283)), closes [#8](https://github.com/ncarlier/feedpushr/issues/8)
* use a map for props options ([6eecb4d](https://github.com/ncarlier/feedpushr/commit/6eecb4d))
* use custom User-Agent for HTTP request ([514ae06](https://github.com/ncarlier/feedpushr/commit/514ae06))
* **launcher:** refactor agent into a launcher ([3cd5145](https://github.com/ncarlier/feedpushr/commit/3cd5145))
* **output:** add configurable output format ([7cf0c42](https://github.com/ncarlier/feedpushr/commit/7cf0c42))



<a name="2.0.0"></a>
# [2.0.0](https://github.com/ncarlier/feedpushr/compare/v2.0.0-rc.2...v2.0.0) (2019-09-19)


### Bug Fixes

* fix PSHB callback ([c5f01f0](https://github.com/ncarlier/feedpushr/commit/c5f01f0))
* **contrib:** fix plugin configuration ([7afd946](https://github.com/ncarlier/feedpushr/commit/7afd946))
* **contrib:** fix readflow plugin ([61f8c0d](https://github.com/ncarlier/feedpushr/commit/61f8c0d))
* **contrib:** fix twitter plugin configuration ([dfed7af](https://github.com/ncarlier/feedpushr/commit/dfed7af))
* **pshb:** try to get feed link from other attribute ([b6ef0db](https://github.com/ncarlier/feedpushr/commit/b6ef0db))


### Features

* update LICENSE to GPLv3 ([bcb15a6](https://github.com/ncarlier/feedpushr/commit/bcb15a6))
* **agent:** refactoring of the agent ([6cd7c39](https://github.com/ncarlier/feedpushr/commit/6cd7c39))
* **feed:** use dedicated forms ([18e9cd1](https://github.com/ncarlier/feedpushr/commit/18e9cd1))
* **output:** make readflow plugin as builtin ([e42c9fb](https://github.com/ncarlier/feedpushr/commit/e42c9fb))
* **ui:** improve date rendering ([0ede79f](https://github.com/ncarlier/feedpushr/commit/0ede79f))
* add agent ([212d399](https://github.com/ncarlier/feedpushr/commit/212d399))
* **ui:** show total nb of feeds ([7ab8109](https://github.com/ncarlier/feedpushr/commit/7ab8109))
* add alias for filters and outputs ([f76e3e6](https://github.com/ncarlier/feedpushr/commit/f76e3e6))
* add CLI attribute to clear configuration ([401390f](https://github.com/ncarlier/feedpushr/commit/401390f))
* backport contrib repository inside the project ([238390b](https://github.com/ncarlier/feedpushr/commit/238390b))



<a name="2.0.0-rc.2"></a>
# [2.0.0-rc.2](https://github.com/ncarlier/feedpushr/compare/v2.0.0-rc.1...v2.0.0-rc.2) (2019-09-04)


### Bug Fixes

* fix unit tests ([e228d47](https://github.com/ncarlier/feedpushr/commit/e228d47))


### Features

* map attribute types on HTML input types ([1c60d74](https://github.com/ncarlier/feedpushr/commit/1c60d74))



<a name="2.0.0-rc.1"></a>
# [2.0.0-rc.1](https://github.com/ncarlier/feedpushr/compare/v1.2.0...v2.0.0-rc.1) (2019-09-01)


### Features

* **ui:** add about page ([ce341c8](https://github.com/ncarlier/feedpushr/commit/ce341c8))
* **ui:** add output pages ([f0e1e63](https://github.com/ncarlier/feedpushr/commit/f0e1e63))
* **ui:** configure main theme ([464f049](https://github.com/ncarlier/feedpushr/commit/464f049))
* **ui:** new filter pages ([85eb0f6](https://github.com/ncarlier/feedpushr/commit/85eb0f6))
* **ui:** switch to new UI ([495e410](https://github.com/ncarlier/feedpushr/commit/495e410))
* auto create all buckets ([779e897](https://github.com/ncarlier/feedpushr/commit/779e897))
* new UI foundation ([7307580](https://github.com/ncarlier/feedpushr/commit/7307580))
* persist filter and output configuration ([c7bc4b3](https://github.com/ncarlier/feedpushr/commit/c7bc4b3))
* persist output configuration ([2ff1d20](https://github.com/ncarlier/feedpushr/commit/2ff1d20))
* **api:** WIP add CRUD API for filters and outputs ([a421877](https://github.com/ncarlier/feedpushr/commit/a421877))
* **api:** WIP add CRUD API for filters and outputs ([fbbbd78](https://github.com/ncarlier/feedpushr/commit/fbbbd78))
* **api:** WIP add Spec API for filters and outputs ([7c00d24](https://github.com/ncarlier/feedpushr/commit/7c00d24))
* **store:** add filter repository ([f28665e](https://github.com/ncarlier/feedpushr/commit/f28665e))
* **ui:** WIP filter pages ([209add6](https://github.com/ncarlier/feedpushr/commit/209add6))



<a name="1.2.0"></a>
# [1.2.0](https://github.com/ncarlier/feedpushr/compare/v1.1.0...v1.2.0) (2019-06-12)


### Bug Fixes

* disable agent for ARM architecture ([4faea98](https://github.com/ncarlier/feedpushr/commit/4faea98))
* **aggregator:** reset delay when manual start ([0d66cc9](https://github.com/ncarlier/feedpushr/commit/0d66cc9)), closes [#6](https://github.com/ncarlier/feedpushr/issues/6)


### Features

* move agent from base code to contrib ([d4c55c3](https://github.com/ncarlier/feedpushr/commit/d4c55c3))
* plugins autoload ([4af51b7](https://github.com/ncarlier/feedpushr/commit/4af51b7))
* use systray for desktop environment ([3826b84](https://github.com/ncarlier/feedpushr/commit/3826b84))



<a name="1.1.0"></a>
# [1.1.0](https://github.com/ncarlier/feedpushr/compare/v1.0.3...v1.1.0) (2019-05-14)


### Bug Fixes

* **opml:** fix import with inline categories ([1aebfef](https://github.com/ncarlier/feedpushr/commit/1aebfef))


### Features

* **aggregator:** store aggregation status ([b1d00db](https://github.com/ncarlier/feedpushr/commit/b1d00db)), closes [#5](https://github.com/ncarlier/feedpushr/issues/5)
* **filter:** use readflow readability function for fetch filter ([7af532d](https://github.com/ncarlier/feedpushr/commit/7af532d))
* **tags:** add negative tag ([1692e5d](https://github.com/ncarlier/feedpushr/commit/1692e5d))
* **ui:** add form micro help ([00b6487](https://github.com/ncarlier/feedpushr/commit/00b6487)), closes [#3](https://github.com/ncarlier/feedpushr/issues/3)
* **ui:** loading screen on OPNML imports ([315d32a](https://github.com/ncarlier/feedpushr/commit/315d32a))



<a name="1.0.3"></a>
## [1.0.3](https://github.com/ncarlier/feedpushr/compare/v1.0.2...v1.0.3) (2019-05-07)


### Bug Fixes

* **aggregator:** limit max delay between checks to h24 ([e0a64d1](https://github.com/ncarlier/feedpushr/commit/e0a64d1))


### Features

* **opml:** add category support ([060effe](https://github.com/ncarlier/feedpushr/commit/060effe)), closes [#2](https://github.com/ncarlier/feedpushr/issues/2) [#4](https://github.com/ncarlier/feedpushr/issues/4)



<a name="1.0.2"></a>
## [1.0.2](https://github.com/ncarlier/feedpushr/compare/v1.0.1...v1.0.2) (2019-04-25)


### Bug Fixes

* **aggregator:** reduce error logs level ([9ddce0d](https://github.com/ncarlier/feedpushr/commit/9ddce0d))
* fix articles counter ([28850d1](https://github.com/ncarlier/feedpushr/commit/28850d1))


### Features

* add maskSecret function for plugins ([e6440f6](https://github.com/ncarlier/feedpushr/commit/e6440f6))
* export configuration variables ([65217fc](https://github.com/ncarlier/feedpushr/commit/65217fc))



<a name="1.0.1"></a>
## [1.0.1](https://github.com/ncarlier/feedpushr/compare/v1.0.0...v1.0.1) (2019-04-10)


### Bug Fixes

* fix tags usage for filters and outputs ([78b92a4](https://github.com/ncarlier/feedpushr/commit/78b92a4))



<a name="1.0.0"></a>
# [1.0.0](https://github.com/ncarlier/feedpushr/compare/1.0.0...v1.0.0) (2019-04-09)


### Bug Fixes

* **aggregator:** fix bad cache retention setup ([fe603bc](https://github.com/ncarlier/feedpushr/commit/fe603bc))
* **aggregator:** fix unit test ([9936ac0](https://github.com/ncarlier/feedpushr/commit/9936ac0))
* **builder:** fix nil pointer ([eb74c80](https://github.com/ncarlier/feedpushr/commit/eb74c80))
* **docker:** fix plugins copy ([047693a](https://github.com/ncarlier/feedpushr/commit/047693a))
* **docker:** remove plugins from image ([d859025](https://github.com/ncarlier/feedpushr/commit/d859025))
* **filter:** minify filter should also filtering description ([087c4be](https://github.com/ncarlier/feedpushr/commit/087c4be))
* **output:** fix error messages ([04787d3](https://github.com/ncarlier/feedpushr/commit/04787d3))
* **plugin:** init output and filter plugin registry ([692c934](https://github.com/ncarlier/feedpushr/commit/692c934))
* **ui:** show error messages and pshb status ([fa366e7](https://github.com/ncarlier/feedpushr/commit/fa366e7))
* copy article link ([44c7da5](https://github.com/ncarlier/feedpushr/commit/44c7da5))
* fix bad error format string ([74ab0a8](https://github.com/ncarlier/feedpushr/commit/74ab0a8))
* **pshb:** add support for (broken) Wordpress plugin ([e7334b3](https://github.com/ncarlier/feedpushr/commit/e7334b3))
* **store:** fix db uri ([aa58032](https://github.com/ncarlier/feedpushr/commit/aa58032))


### Features

* multiple outputs support ([2052116](https://github.com/ncarlier/feedpushr/commit/2052116))
* **aggregator:** add timeout configuration ([45a887a](https://github.com/ncarlier/feedpushr/commit/45a887a))
* **api:** add feed title managment ([f837590](https://github.com/ncarlier/feedpushr/commit/f837590))
* **docker:** create an image with plugins ([20b9e49](https://github.com/ncarlier/feedpushr/commit/20b9e49))
* **feed:** add processed items counter ([a5d3645](https://github.com/ncarlier/feedpushr/commit/a5d3645))
* **feed:** add status attribute ([9b82c9b](https://github.com/ncarlier/feedpushr/commit/9b82c9b))
* **filter:** add counter to title filter props ([69d2c42](https://github.com/ncarlier/feedpushr/commit/69d2c42))
* **filter:** add fetch filter ([0e72749](https://github.com/ncarlier/feedpushr/commit/0e72749))
* **filter:** add filter system (with plugins) ([8d6fd44](https://github.com/ncarlier/feedpushr/commit/8d6fd44))
* **filter:** add minify filter ([d37d104](https://github.com/ncarlier/feedpushr/commit/d37d104))
* **filter:** use last readability lib for fetch filter ([0530fbc](https://github.com/ncarlier/feedpushr/commit/0530fbc))
* improve filter/output descriptions ([444cea7](https://github.com/ncarlier/feedpushr/commit/444cea7))
* **api:** remove feed list limit ([a76d4c2](https://github.com/ncarlier/feedpushr/commit/a76d4c2))
* **logging:** add Sentry for error recording ([f674543](https://github.com/ncarlier/feedpushr/commit/f674543))
* **logging:** use async Sentry call ([55715e0](https://github.com/ncarlier/feedpushr/commit/55715e0))
* **opml:** don't import existing feeds ([a804718](https://github.com/ncarlier/feedpushr/commit/a804718))
* **output:** add external plugin support ([f2d4656](https://github.com/ncarlier/feedpushr/commit/f2d4656))
* **plugin:** refactoring of the plugin system ([12c009f](https://github.com/ncarlier/feedpushr/commit/12c009f))
* **pshb:** add max TTL for a PSHB subscription ([cf6b34b](https://github.com/ncarlier/feedpushr/commit/cf6b34b))
* **pshb:** compute subscribtion details URL ([83c1aae](https://github.com/ncarlier/feedpushr/commit/83c1aae))
* **tags:** add tags support ([a59998b](https://github.com/ncarlier/feedpushr/commit/a59998b))
* **ui:** add feed filter bar ([03314bc](https://github.com/ncarlier/feedpushr/commit/03314bc))
* **ui:** add Web user interface ([aa8d50e](https://github.com/ncarlier/feedpushr/commit/aa8d50e))
* **ui:** display filter and output descriptions as details ([9046f0a](https://github.com/ncarlier/feedpushr/commit/9046f0a))
* **ui:** make feed table sortable ([cc319fc](https://github.com/ncarlier/feedpushr/commit/cc319fc))
* **ui:** make Output a functional component ([f43d4c5](https://github.com/ncarlier/feedpushr/commit/f43d4c5))
* configure CORS ([a6e53f1](https://github.com/ncarlier/feedpushr/commit/a6e53f1))
* improve exploitation logs ([8be34ae](https://github.com/ncarlier/feedpushr/commit/8be34ae))
* overwriting of the environmental configuration via parameters ([4a64395](https://github.com/ncarlier/feedpushr/commit/4a64395))
* use URL to declare filters and outputs ([a77d74d](https://github.com/ncarlier/feedpushr/commit/a77d74d))



