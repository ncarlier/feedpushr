# [](https://github.com/ncarlier/feedpushr/compare/v3.1.1...v) (2021-12-03)



## [3.1.1](https://github.com/ncarlier/feedpushr/compare/v3.1.0...v3.1.1) (2021-12-03)


### Bug Fixes

* **ui:** fix external links ([6e6ba17](https://github.com/ncarlier/feedpushr/commit/6e6ba172756806f73bdcbdf433cd3cab2fd1c17c))


### Features

* add quota for feeds and outputs ([2da4102](https://github.com/ncarlier/feedpushr/commit/2da41027af988422561d1d1b8de4829894bbbbfc))
* **auth:** extract authorized username from multiple claims ([fd15078](https://github.com/ncarlier/feedpushr/commit/fd150786e681cc06af136f6a1fab066bf1f065c4))
* switch from statik to go:embed ([3bcc80e](https://github.com/ncarlier/feedpushr/commit/3bcc80eb69328b34f6fc10452c5e7ce114e7528d))



# [3.1.0](https://github.com/ncarlier/feedpushr/compare/v3.0.0...v3.1.0) (2021-06-13)


### Bug Fixes

* **api:** fix PSHB public access ([69c4921](https://github.com/ncarlier/feedpushr/commit/69c4921923be5a2f18399d18af0c264049016b71))
* **ui:** fix SSE API location ([1fbf422](https://github.com/ncarlier/feedpushr/commit/1fbf4223ffe6880afd91e5a2a58eb97dbe67d1a3))


### Features

* **authn:** JWT OIDC authentication support ([d01d3e9](https://github.com/ncarlier/feedpushr/commit/d01d3e92af870e536c4301ce9b80627027c6b35d))
* service registry support with Consul ([1cfbaf9](https://github.com/ncarlier/feedpushr/commit/1cfbaf9bdb81e8959a7e4b0f13bd46eda9c70203))
* **ui:** OIDC authentication ([59f9384](https://github.com/ncarlier/feedpushr/commit/59f93847add8c0a644ca86f4304e7c0b74c2c4c8))



# [3.0.0](https://github.com/ncarlier/feedpushr/compare/v2.2.0...v3.0.0) (2020-06-23)


### Bug Fixes

* **aggregator:** omit empty headers from the request ([e83c177](https://github.com/ncarlier/feedpushr/commit/e83c17700f1867dfff33c35e1fd5fb0fe604c69a))
* **api:** make props attribute response mandatory ([177614e](https://github.com/ncarlier/feedpushr/commit/177614e2ed764cd3518da2ae0b5b358c40362d01))
* **ci:** add missing package ([cc970d1](https://github.com/ncarlier/feedpushr/commit/cc970d1c91e21d1f4db382c8bb9a5477a21fc844))
* **filter:** add missing prop definition ([5a3c575](https://github.com/ncarlier/feedpushr/commit/5a3c575d2d72143924cf176a57d041139ae255dd))
* **filter:** only apply enabled filters ([e8a45c6](https://github.com/ncarlier/feedpushr/commit/e8a45c6b606f747578a8bd24daf8e6ee7bf838b5))
* **format:** improve tweet function security ([4e39bf0](https://github.com/ncarlier/feedpushr/commit/4e39bf08c15bb2fed9a04e7c6e7d8ce1f54c4bd5))
* init search index at startup ([83aca10](https://github.com/ncarlier/feedpushr/commit/83aca10bc5e636502ef5da694d5dcfab8697c341))
* **output:** fix readflow text attribute ([96ef893](https://github.com/ncarlier/feedpushr/commit/96ef893f7f28a3fdae351b7749674ccd0898b9eb))
* **output:** improve goroutine stop ([0c0179e](https://github.com/ncarlier/feedpushr/commit/0c0179e9f130d53d8e5d3ebfb0c8df99360616c9))
* **version:** fix version display ([07e46bf](https://github.com/ncarlier/feedpushr/commit/07e46bf630b89f16c98a6b9a00b3dda2f85bcd54))
* **websub:** fix callback URL ([1f33fcd](https://github.com/ncarlier/feedpushr/commit/1f33fcd3a56a6a52466998b4ef66729c6d7322cd))


### Features

* **contrib:** add Kafka plugin ([886f878](https://github.com/ncarlier/feedpushr/commit/886f8781a7fd9026c166ac3c3f0bcf78638d2fb7))
* **explorer:** add feed explorer ([aca9077](https://github.com/ncarlier/feedpushr/commit/aca90770a1c26c1483b0484f1314ffcbf68cf799))
* **feed:** deactivated by default ([e234dda](https://github.com/ncarlier/feedpushr/commit/e234dda0bb7393b084a6f2c4d1916a066ae77bfa))
* **feeds:** full-text search engine ([7b2cc08](https://github.com/ncarlier/feedpushr/commit/7b2cc08e25f13849258778d7e14c51ecf30aa114))
* **filter:** add HTTP filter plugin ([3afb377](https://github.com/ncarlier/feedpushr/commit/3afb3778156d0dcbf185c0880ee139f8b35ca352))
* **filter:** use external scraper for fetch filter ([d862502](https://github.com/ncarlier/feedpushr/commit/d86250291b6a761452505a5a30d844d85bdd9bb1))
* **format:** custom format functions ([bdbccd5](https://github.com/ncarlier/feedpushr/commit/bdbccd57d192360a76d6f552baa80fa6362ff916))
* **output:** improve readflow output ([6177e3b](https://github.com/ncarlier/feedpushr/commit/6177e3bd85fea9730fd64fb6c3752ac1f124e7f4))
* **ui:** autofocus on search box ([e7e746a](https://github.com/ncarlier/feedpushr/commit/e7e746a7ef17ff5da7ab5531c32eef68473735ee))
* **ui:** foldable spec description ([c9a3c02](https://github.com/ncarlier/feedpushr/commit/c9a3c02266b0ce2d5b80361f6ca848b02a83a59f))
* **ui:** improve base path support ([ddc4a75](https://github.com/ncarlier/feedpushr/commit/ddc4a7516689e41b5b4a1be281b55f7c9f6dbf8d))



# [2.2.0](https://github.com/ncarlier/feedpushr/compare/v2.1.0...v2.2.0) (2020-03-08)


### Bug Fixes

* **aggregator:** fix concurrent iteration and write ([0bb8124](https://github.com/ncarlier/feedpushr/commit/0bb8124c750844d720307d7869938ef714126583)), closes [#23](https://github.com/ncarlier/feedpushr/issues/23)
* **aggregator:** reset handler status on start ([298a149](https://github.com/ncarlier/feedpushr/commit/298a1490d81d35c52778b4f8c98488ec3102b204))
* fix bad package ref ([d95c4cf](https://github.com/ncarlier/feedpushr/commit/d95c4cff6782fdb2af6c659c18b54810d170620d))
* fix feed page attributes ([928cfa5](https://github.com/ncarlier/feedpushr/commit/928cfa5eeeb910aa47cf9fe15bee6a88275fe106))
* fix flag binding for slice ([a3f6a71](https://github.com/ncarlier/feedpushr/commit/a3f6a713fc48a66a0cdae941046331c1fd64cb72))
* **ui:** fix feed control state on pagination ([870eb31](https://github.com/ncarlier/feedpushr/commit/870eb3156a2f9412fb258879abd4901d8b54752d)), closes [#20](https://github.com/ncarlier/feedpushr/issues/20)


### Features

* **aggregator:** fan-out delay ([096d1a8](https://github.com/ncarlier/feedpushr/commit/096d1a8a0498604f94e8599b502ad9111938c1e9))
* **auth:** add Basic Auth support ([3975b90](https://github.com/ncarlier/feedpushr/commit/3975b90937c5939eb81c8c409401287ab9d5d505))
* configuration refactoring ([d7480be](https://github.com/ncarlier/feedpushr/commit/d7480becb5945a69a650f200c20d3f0be997384f))
* **doc:** fan-out delay configuration ([00faee4](https://github.com/ncarlier/feedpushr/commit/00faee440bdd92eb39e949e32cee4834f03f8cde))
* enable pagination for feeds ([e64b6b2](https://github.com/ncarlier/feedpushr/commit/e64b6b25d5d46d3a091679d7dea9cbf9e6f7cd9a))
* **plugins:** use formatter for Twitter and Mastodon plugins ([1562b88](https://github.com/ncarlier/feedpushr/commit/1562b88a9252400dfd62c4b21ccd50a9dbb1871b))
* redirect base URL to UI ([6c26daa](https://github.com/ncarlier/feedpushr/commit/6c26daae2ee7c97066e1dd6b2020e2f2d8036f9c))
* retrieve feed link from HTML URL ([922a783](https://github.com/ncarlier/feedpushr/commit/922a7830a33a36b2ff05d493de808f9292741046))
* **ui:** add feed HTML link ([b5a2ebe](https://github.com/ncarlier/feedpushr/commit/b5a2ebe0685cc78541b835f42a567566b1a85bf2)), closes [#21](https://github.com/ncarlier/feedpushr/issues/21)
* use output formatter ([b69484a](https://github.com/ncarlier/feedpushr/commit/b69484ac6c7263df928fa589b0a5f4cfce7888b3))



# [2.1.0](https://github.com/ncarlier/feedpushr/compare/v2.0.0...v2.1.0) (2020-01-06)


### Bug Fixes

* **contrib:** ignore Twitter error when duplicate ([f0099f2](https://github.com/ncarlier/feedpushr/commit/f0099f2709d9bbcb719271827132c8f6e3e8fbb8))
* remove 32 bit support ([b3053d5](https://github.com/ncarlier/feedpushr/commit/b3053d55d0a1408ef52273c4568a99943f363552))
* typos ([820ee03](https://github.com/ncarlier/feedpushr/commit/820ee03dc0ff31d399a69e4515499db5b6429e8c))


### Features

* add API info endpoint ([a234a23](https://github.com/ncarlier/feedpushr/commit/a234a23c8fcf0b22224931b7cb3b30e9397e3dbd))
* conditional expression for filters & outputs ([6b6d283](https://github.com/ncarlier/feedpushr/commit/6b6d283231e10ed77372facd6cfbbbcc4180df89)), closes [#8](https://github.com/ncarlier/feedpushr/issues/8)
* **launcher:** refactor agent into a launcher ([3cd5145](https://github.com/ncarlier/feedpushr/commit/3cd5145b378b52ed7b1bc0571db495ce50cb0f75))
* **output:** add configurable output format ([7cf0c42](https://github.com/ncarlier/feedpushr/commit/7cf0c42e37aa15a973e322df3e8e27d9e7016ad9))
* use a map for props options ([6eecb4d](https://github.com/ncarlier/feedpushr/commit/6eecb4d85d0203f352f4809d099052cdeaee12b2))
* use custom User-Agent for HTTP request ([514ae06](https://github.com/ncarlier/feedpushr/commit/514ae06a08d49694a1bfa7ee8a0db42435931cde))



# [2.0.0](https://github.com/ncarlier/feedpushr/compare/v1.2.0...v2.0.0) (2019-09-19)


### Bug Fixes

* **contrib:** fix plugin configuration ([7afd946](https://github.com/ncarlier/feedpushr/commit/7afd946d793f3ca169a1738ca8953f35095595ca))
* **contrib:** fix readflow plugin ([61f8c0d](https://github.com/ncarlier/feedpushr/commit/61f8c0d7e5efb00808f60d229fe086fa088ad7d6))
* **contrib:** fix twitter plugin configuration ([dfed7af](https://github.com/ncarlier/feedpushr/commit/dfed7af4112a771e8d871e4f356dccaaab338fc2))
* fix PSHB callback ([c5f01f0](https://github.com/ncarlier/feedpushr/commit/c5f01f02eeeb6b04300c9e74271444db8a519fe7))
* fix unit tests ([e228d47](https://github.com/ncarlier/feedpushr/commit/e228d47ce4b1b189196b8490d3ae4582a3e474ba))
* **pshb:** try to get feed link from other attribute ([b6ef0db](https://github.com/ncarlier/feedpushr/commit/b6ef0dbda1fce48410bb9f701893fb11cb9646b2))


### Features

* add agent ([212d399](https://github.com/ncarlier/feedpushr/commit/212d399773e5f2cc4962e1acb99f4d61915587b3))
* add alias for filters and outputs ([f76e3e6](https://github.com/ncarlier/feedpushr/commit/f76e3e6c9d5a76502a6d75e7429e92935a4fef64))
* add CLI attribute to clear configuration ([401390f](https://github.com/ncarlier/feedpushr/commit/401390f76f990497ea905a41e588802aef51dd79))
* **agent:** refactoring of the agent ([6cd7c39](https://github.com/ncarlier/feedpushr/commit/6cd7c391c33710362398d48dac7be854f5eda39c))
* **api:** WIP add CRUD API for filters and outputs ([a421877](https://github.com/ncarlier/feedpushr/commit/a421877073a2617391cf8ddd9fbcf72a878f1047))
* **api:** WIP add CRUD API for filters and outputs ([fbbbd78](https://github.com/ncarlier/feedpushr/commit/fbbbd78815e6e94d755702defa65ec90de38b51b))
* **api:** WIP add Spec API for filters and outputs ([7c00d24](https://github.com/ncarlier/feedpushr/commit/7c00d24efefac238354f8e215449f7f1263fc2bb))
* auto create all buckets ([779e897](https://github.com/ncarlier/feedpushr/commit/779e897e27182682f9fc79d64ae29b95b7b5998b))
* backport contrib repository inside the project ([238390b](https://github.com/ncarlier/feedpushr/commit/238390b372047458137c0ef3926edf1cd3538a97))
* **feed:** use dedicated forms ([18e9cd1](https://github.com/ncarlier/feedpushr/commit/18e9cd101278865c57b9f85d99b79b3e031ca170))
* map attribute types on HTML input types ([1c60d74](https://github.com/ncarlier/feedpushr/commit/1c60d74ca482ac2cccd605805ef88e47d88c564d))
* new UI foundation ([7307580](https://github.com/ncarlier/feedpushr/commit/7307580cedec6f44a87e61fad32ac26a6375c513))
* **output:** make readflow plugin as builtin ([e42c9fb](https://github.com/ncarlier/feedpushr/commit/e42c9fb32b8ce968ec8bd642df27e9509316f8f4))
* persist filter and output configuration ([c7bc4b3](https://github.com/ncarlier/feedpushr/commit/c7bc4b3512b6de136187dfa11b82fe765701c2d4))
* persist output configuration ([2ff1d20](https://github.com/ncarlier/feedpushr/commit/2ff1d209df6557e46e230bc9895c3287b7f519ea))
* **store:** add filter repository ([f28665e](https://github.com/ncarlier/feedpushr/commit/f28665ef583e712d28030365ae7f857482ae4c1a))
* **ui:** add about page ([ce341c8](https://github.com/ncarlier/feedpushr/commit/ce341c8c149e6662811a4d3e866beb4e31b65066))
* **ui:** add output pages ([f0e1e63](https://github.com/ncarlier/feedpushr/commit/f0e1e637c973a82dda1a6b026b499db1a1a3358b))
* **ui:** configure main theme ([464f049](https://github.com/ncarlier/feedpushr/commit/464f049bd48244d6be8e4aca38626a7eae5ea56e))
* **ui:** improve date rendering ([0ede79f](https://github.com/ncarlier/feedpushr/commit/0ede79ffcbbc33c4738013cffa6a5e72325cdedc))
* **ui:** new filter pages ([85eb0f6](https://github.com/ncarlier/feedpushr/commit/85eb0f659fdca1e0ddd4aa00526ab3cce22cde0a))
* **ui:** show total nb of feeds ([7ab8109](https://github.com/ncarlier/feedpushr/commit/7ab8109476a4f0222dad6e39b3b75f0ed06762c3))
* **ui:** switch to new UI ([495e410](https://github.com/ncarlier/feedpushr/commit/495e4102c035cda168ff86a16ab02b64e677fc6b))
* **ui:** WIP filter pages ([209add6](https://github.com/ncarlier/feedpushr/commit/209add64d6f0bd1f8c81eed0b41f8520fe5f2957))
* update LICENSE to GPLv3 ([bcb15a6](https://github.com/ncarlier/feedpushr/commit/bcb15a65ff7bbe62feb2805e8eeb6c0b55d5d1af))



# [1.2.0](https://github.com/ncarlier/feedpushr/compare/v1.1.0...v1.2.0) (2019-06-12)


### Bug Fixes

* **aggregator:** reset delay when manual start ([0d66cc9](https://github.com/ncarlier/feedpushr/commit/0d66cc969600ce3760b59197c0fa07bb3f1f9a79)), closes [#6](https://github.com/ncarlier/feedpushr/issues/6)
* disable agent for ARM architecture ([4faea98](https://github.com/ncarlier/feedpushr/commit/4faea985190baa40ff35e0eaa6fcb36e74f3081f))


### Features

* move agent from base code to contrib ([d4c55c3](https://github.com/ncarlier/feedpushr/commit/d4c55c3bf151b05b5a7edd92caf035a8358e864b))
* plugins autoload ([4af51b7](https://github.com/ncarlier/feedpushr/commit/4af51b78c3737a7c9cc51085c58521e452a019ad))
* use systray for desktop environment ([3826b84](https://github.com/ncarlier/feedpushr/commit/3826b844a286211aaacd26a8ddcadef504cd998d))



# [1.1.0](https://github.com/ncarlier/feedpushr/compare/v1.0.3...v1.1.0) (2019-05-14)


### Bug Fixes

* **opml:** fix import with inline categories ([1aebfef](https://github.com/ncarlier/feedpushr/commit/1aebfef0a22bf283f876fa898e8345d9e97090fa))


### Features

* **aggregator:** store aggregation status ([b1d00db](https://github.com/ncarlier/feedpushr/commit/b1d00dbdacb1e10a818340a997526e3bb618b4f0)), closes [#5](https://github.com/ncarlier/feedpushr/issues/5)
* **filter:** use readflow readability function for fetch filter ([7af532d](https://github.com/ncarlier/feedpushr/commit/7af532dfb897920929e8a71257b50d15165380f9))
* **tags:** add negative tag ([1692e5d](https://github.com/ncarlier/feedpushr/commit/1692e5d2df047dc9b085b68448f47a49608505ae))
* **ui:** add form micro help ([00b6487](https://github.com/ncarlier/feedpushr/commit/00b64878a78f92e85d52051a257cd5ce98fd4373)), closes [#3](https://github.com/ncarlier/feedpushr/issues/3)
* **ui:** loading screen on OPNML imports ([315d32a](https://github.com/ncarlier/feedpushr/commit/315d32a8f866d1f434fae431cf7829eca4cd5293))



## [1.0.3](https://github.com/ncarlier/feedpushr/compare/v1.0.2...v1.0.3) (2019-05-07)


### Bug Fixes

* **aggregator:** limit max delay between checks to h24 ([e0a64d1](https://github.com/ncarlier/feedpushr/commit/e0a64d14e3ab179d694a6dbd2d25b08df04255ad))


### Features

* **opml:** add category support ([060effe](https://github.com/ncarlier/feedpushr/commit/060effe480cf79823c0217e820fd11940b880a7e)), closes [#2](https://github.com/ncarlier/feedpushr/issues/2) [#4](https://github.com/ncarlier/feedpushr/issues/4)



## [1.0.2](https://github.com/ncarlier/feedpushr/compare/v1.0.1...v1.0.2) (2019-04-25)


### Bug Fixes

* **aggregator:** reduce error logs level ([9ddce0d](https://github.com/ncarlier/feedpushr/commit/9ddce0d22de2162edcfbf30efc65b46bceca0f55))
* fix articles counter ([28850d1](https://github.com/ncarlier/feedpushr/commit/28850d10c012cf6af3dd4cdd0692a180ad52ff04))


### Features

* add maskSecret function for plugins ([e6440f6](https://github.com/ncarlier/feedpushr/commit/e6440f655fc4497b7bb3488cced67c92c8d6fbe5))
* export configuration variables ([65217fc](https://github.com/ncarlier/feedpushr/commit/65217fc2b2fc694c75a2aa455ceefce5255f3b8a))



## [1.0.1](https://github.com/ncarlier/feedpushr/compare/v1.0.0...v1.0.1) (2019-04-10)


### Bug Fixes

* fix tags usage for filters and outputs ([78b92a4](https://github.com/ncarlier/feedpushr/commit/78b92a4f518071561feb2e6f8deed952afbe654b))



# [1.0.0](https://github.com/ncarlier/feedpushr/compare/4a64395ae891ef7fe87274d97d896d505fd90a06...v1.0.0) (2019-04-09)


### Bug Fixes

* **aggregator:** fix bad cache retention setup ([fe603bc](https://github.com/ncarlier/feedpushr/commit/fe603bc31d5436f65d99ccc98ec28c3bd81beb0c))
* **aggregator:** fix unit test ([9936ac0](https://github.com/ncarlier/feedpushr/commit/9936ac0c8be0f3646b2f2d9ec79257690c6540f7))
* **builder:** fix nil pointer ([eb74c80](https://github.com/ncarlier/feedpushr/commit/eb74c8015dca031219720cb52e301af33074933b))
* copy article link ([44c7da5](https://github.com/ncarlier/feedpushr/commit/44c7da556e4f8fdec0ea216f2b96af9e3529b6ef))
* **docker:** fix plugins copy ([047693a](https://github.com/ncarlier/feedpushr/commit/047693ac426039d30cac418475ff56749d13b405))
* **docker:** remove plugins from image ([d859025](https://github.com/ncarlier/feedpushr/commit/d8590252d31652fd1593c60eb4b8c2ca86496992))
* **filter:** minify filter should also filtering description ([087c4be](https://github.com/ncarlier/feedpushr/commit/087c4be4daa2d280723e88c3d3d2b76d55bba5c3))
* fix bad error format string ([74ab0a8](https://github.com/ncarlier/feedpushr/commit/74ab0a8564e74c23cc6eb20019938f16ef10b06c))
* **output:** fix error messages ([04787d3](https://github.com/ncarlier/feedpushr/commit/04787d3f6051a086ae88d7f258fec43ed80ec431))
* **plugin:** init output and filter plugin registry ([692c934](https://github.com/ncarlier/feedpushr/commit/692c934f5677dd2bdfd227acabdad918f97556ee))
* **pshb:** add support for (broken) Wordpress plugin ([e7334b3](https://github.com/ncarlier/feedpushr/commit/e7334b34ba0998c459653537442ba34e260dad0f))
* **store:** fix db uri ([aa58032](https://github.com/ncarlier/feedpushr/commit/aa58032fb5d6916fb44bff17391f496c3947ef32))
* **ui:** show error messages and pshb status ([fa366e7](https://github.com/ncarlier/feedpushr/commit/fa366e7012e4445db69a6a85d73438b3868defe8))


### Features

* **aggregator:** add timeout configuration ([45a887a](https://github.com/ncarlier/feedpushr/commit/45a887a717213b813a673fc55a1e03ef68ca1f1f))
* **api:** add feed title managment ([f837590](https://github.com/ncarlier/feedpushr/commit/f8375909a276626cd66d94a83e304be59f86c379))
* **api:** remove feed list limit ([a76d4c2](https://github.com/ncarlier/feedpushr/commit/a76d4c25884fa9a1c97ac59844ddeaebe3dc9bc4))
* configure CORS ([a6e53f1](https://github.com/ncarlier/feedpushr/commit/a6e53f1524a7c356e4e6087ac2eba604943106ac))
* **docker:** create an image with plugins ([20b9e49](https://github.com/ncarlier/feedpushr/commit/20b9e494f84f21db034fc16e8cc9e41e57cead3a))
* **feed:** add processed items counter ([a5d3645](https://github.com/ncarlier/feedpushr/commit/a5d36458c9c7ff43df959094b6ade3b2d4078f18))
* **feed:** add status attribute ([9b82c9b](https://github.com/ncarlier/feedpushr/commit/9b82c9b06122c2797cb206c2a46bc41db7fb908a))
* **filter:** add counter to title filter props ([69d2c42](https://github.com/ncarlier/feedpushr/commit/69d2c426d34bfcc121af10c9a95632a007191ecc))
* **filter:** add fetch filter ([0e72749](https://github.com/ncarlier/feedpushr/commit/0e72749e6bfd65dc5cb88db175bd04e6fbd59fb2))
* **filter:** add filter system (with plugins) ([8d6fd44](https://github.com/ncarlier/feedpushr/commit/8d6fd448d4d00a4471acade55072126acb5fbfc7))
* **filter:** add minify filter ([d37d104](https://github.com/ncarlier/feedpushr/commit/d37d104dbb6f0fb5af7fb9470c32248d76b09223))
* **filter:** use last readability lib for fetch filter ([0530fbc](https://github.com/ncarlier/feedpushr/commit/0530fbc5a92defe22256df78122d6d2eab742fca))
* improve exploitation logs ([8be34ae](https://github.com/ncarlier/feedpushr/commit/8be34aeed0b7cf9d8fb30c4c0394fd68b8769e7d))
* improve filter/output descriptions ([444cea7](https://github.com/ncarlier/feedpushr/commit/444cea7170109d0c6d8d1ad8ed6061ce252c77d1))
* **logging:** add Sentry for error recording ([f674543](https://github.com/ncarlier/feedpushr/commit/f67454395f0b81579d7a3fac6db0ea11e49a883f))
* **logging:** use async Sentry call ([55715e0](https://github.com/ncarlier/feedpushr/commit/55715e0c74282020b54dea62295f5dfb046d377b))
* multiple outputs support ([2052116](https://github.com/ncarlier/feedpushr/commit/20521166ae65bf5f53d0c41b53502e372895b751))
* **opml:** don't import existing feeds ([a804718](https://github.com/ncarlier/feedpushr/commit/a80471829057a748bb887d0eb278c7cf6f432acb))
* **output:** add external plugin support ([f2d4656](https://github.com/ncarlier/feedpushr/commit/f2d4656c5ff4510146e5a0fb370346f9bff31568))
* overwriting of the environmental configuration via parameters ([4a64395](https://github.com/ncarlier/feedpushr/commit/4a64395ae891ef7fe87274d97d896d505fd90a06))
* **plugin:** refactoring of the plugin system ([12c009f](https://github.com/ncarlier/feedpushr/commit/12c009f924bd4ab2da7c8ca6190c8f23298c8e86))
* **pshb:** add max TTL for a PSHB subscription ([cf6b34b](https://github.com/ncarlier/feedpushr/commit/cf6b34b12499f2f12f2213ad5a66641ee77d1cdb))
* **pshb:** compute subscribtion details URL ([83c1aae](https://github.com/ncarlier/feedpushr/commit/83c1aaeeff5479cd30a5cbae2a09e0d0f0324ae3))
* **tags:** add tags support ([a59998b](https://github.com/ncarlier/feedpushr/commit/a59998b9dd6bae1ea118c8bf194431078c6f3bb3))
* **ui:** add feed filter bar ([03314bc](https://github.com/ncarlier/feedpushr/commit/03314bc844c1addb9d577c03ca1b04236b1d4563))
* **ui:** add Web user interface ([aa8d50e](https://github.com/ncarlier/feedpushr/commit/aa8d50e74c77d84318a123cd9ba7cf89e7df49cb))
* **ui:** display filter and output descriptions as details ([9046f0a](https://github.com/ncarlier/feedpushr/commit/9046f0a22f338a17914cd831a9efdf833a8071de))
* **ui:** make feed table sortable ([cc319fc](https://github.com/ncarlier/feedpushr/commit/cc319fc78d3dba9f5e7387ee06753b655279bf80))
* **ui:** make Output a functional component ([f43d4c5](https://github.com/ncarlier/feedpushr/commit/f43d4c54825f4252751816c2ca515a1db5aa1f88))
* use URL to declare filters and outputs ([a77d74d](https://github.com/ncarlier/feedpushr/commit/a77d74d64fd294e2479b43f51f5e2417fd922f86))



