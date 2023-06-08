# Release Notes

A Forgejo release is published shortly after a Gitea release is published and they have [matching release numbers](https://codeberg.org/forgejo/forgejo/src/branch/forgejo/CONTRIBUTING/RELEASE.md#release-numbering). Additional Forgejo releases may be published to address urgent security issues or bug fixes. Forgejo release notes include all Gitea release notes.

The Forgejo admin should carefully read the required manual actions before upgrading. A point release (e.g. v1.19.1 or v1.19.2) does not require manual actions but others might (e.g. v1.18.0, v1.19.0).

## DRAFT 1.20.0-0

The [complete list of commits](https://codeberg.org/forgejo/forgejo/commits/branch/v1.20/forgejo) included in the `Forgejo v1.20.0-?` release can be reviewed from the command line with:

```shell
$ git clone https://codeberg.org/forgejo/forgejo/
$ git -C forgejo log --oneline --no-merges origin/v1.19/forgejo..origin/v1.20/forgejo
```

- [Forgejo Semantic Version](https://forgejo.org/docs/v1.20/user/semver/)
  The semantic version was updated to `5.0.0+0-gitea-1.20.0` because it contains breaking changes.
- [CI]
  - Workflows are now [available to run tests](https://codeberg.org/forgejo/forgejo/src/branch/forgejo/.forgejo/workflows) on `Forgejo` itself. It is not enabled yet on Codeberg but will work if the repository is mirrored on an instance where [Forgejo Actions](https://forgejo.org/docs/v1.20/user/actions/) is enabled.
- [MODERATION]
  - Blocking another user is desirable if they are acting maliciously or are spamming your repository. When you block a user, Forgejo does not explicitly notify them, but they may learn through an interaction with you that is blocked. [Read more about blocking users](https://forgejo.org/docs/v1.20/user/blocking-user/).
- [PACKAGES]
  - [PACKAGES SWIFT] commit c709fa17a77eae391cafbe72d6b2594f74d86a60 Add Swift package registry [22404](https://github.com/go-gitea/gitea/pull/22404)
  - [PACKAGE debian] commit bf999e406994ab34420fb62e0de7948c8c2116c1 Add Debian package registry [24426](https://github.com/go-gitea/gitea/pull/24426)
  - [PACKAGES RPM] commit 05209f0d1d4b996b8beb6633880b8fe12c15932b Add RPM registry [23380](https://github.com/go-gitea/gitea/pull/23380)
  - [PACKAGE alpine] commit 9173e079ae9ddf18685216fd846ca1727297393c Add Alpine package registry [23714](https://github.com/go-gitea/gitea/pull/23714)
  - [PACKAGE go] commit 5968c63a11c94b0fdde0485af194bebb2ea1b8e7 Add Go package registry [24687](https://github.com/go-gitea/gitea/pull/24687)
  - [PACKAGES CRAN] commit cdb088cec288a20e14240f86a689dd14f4cd603b Add CRAN package registry [22331](https://github.com/go-gitea/gitea/pull/22331)
  - [PACKAGES cargo]] commit 723598b803919bfc074fee05f830421a99881c3b Implement Cargo HTTP index [24452](https://github.com/go-gitea/gitea/pull/24452)
- [A11Y]
  - [A11Y] commit 6c354546547cd3a9595a7db119a6480d9cd506a7 Improve accessibility for issue comments [22612](https://github.com/go-gitea/gitea/pull/22612)
  - [A11Y] commit a78e0b7dade16bc6509b943fe86e74962f1b95b6 Add accessibility to the menu on the navbar [23059](https://github.com/go-gitea/gitea/pull/23059)
  - [A11Y] commit e8935606f5f1fff3c59222ebca6d4615ab06fb0b Scoped labels: set aria-disabled on muted Exclusive option for a11y [23306](https://github.com/go-gitea/gitea/pull/23306)
  - [A11Y] commit d4f35bd681af0632da988e15306f330e020422b2 Use a general approch to improve a11y for all checkboxes and dropdowns. [23542](https://github.com/go-gitea/gitea/pull/23542)
  - [A11Y RTL] commit 32d9c47ec7706d8f06e09b42e09a28d7a0e3c526 Add RTL rendering support to Markdown [24816](https://github.com/go-gitea/gitea/pull/24816)
  - [A11Y] commit e95b42e187cde9ac4bd541cd714bdb4f5c1fd8bc Improve accessibility when (re-)viewing files [24817](https://github.com/go-gitea/gitea/pull/24817)
  - [A11Y] commit 87f0f7e670c6c0e6aeab8c4458bfdb9d954eacec Add aria attributes to interactive time tooltips. [23661](https://github.com/go-gitea/gitea/pull/23661)
- [TIME]
  - [TIME] commit b7b58348317cbe0145dc453d45c886b8e2764b4c Use auto-updating, natively hoverable, localized time elements [23988](https://github.com/go-gitea/gitea/pull/23988)
  - [TIME] commit 25faee3c5f5be23c99b3b7e50418fc0dbad7a41b Fix date  display bug [24047](https://github.com/go-gitea/gitea/pull/24047)
  - [TIME] commit 97176754beb4de23fa0f68df715c4737919c93b0 Localize milestone related time strings [24051](https://github.com/go-gitea/gitea/pull/24051)
  - [TIME] commit 70bb4984cdad9a15d676708bd345b590aa42d72a Allow using localized absolute date times within phrases with place holders and localize issue due date events [24275](https://github.com/go-gitea/gitea/pull/24275)
  - [TIME] commit 5bc9f7fcf9aece92c3fa2a0ea56e5585261a7f28 Improve commit date in commit graph [24399](https://github.com/go-gitea/gitea/pull/24399)
  - [TIME] commit 62ca5825f73ad5a25ffeb6c3ef66f0eaf5d30cdf Fix incorrect last online time in runner_edit.tmpl [24376](https://github.com/go-gitea/gitea/pull/24376)
  - [TIME] commit dbb37367854d108ebfffcac27837c0afac199a8e Fix incorrect webhook time and use relative-time to display it [24477](https://github.com/go-gitea/gitea/pull/24477)
  - [TIME] commit 3d266dd0f3dbae7e417c0e790e266aebc0078814 In TestViewRepo2, convert computed timezones to local time [24579](https://github.com/go-gitea/gitea/pull/24579)
- [WIKI]
  - [WIKI] commit c0246677a692de804ffe1bb5f7d630fb002dd128 Fix markup background, improve wiki rendering [23750](https://github.com/go-gitea/gitea/pull/23750)
  - [WIKI] commit 2f468381205f5f7e279791aa71e5288710a6476c Re-add initial wiki page text when editing the page [23984](https://github.com/go-gitea/gitea/pull/23984)
  - [WIKI] commit 1ab16e48cccc086e7f97fb3ae8a293fe47a3a452 Improve Wiki TOC [24137](https://github.com/go-gitea/gitea/pull/24137)
  - [WIKI] commit 284b41f45244bbe46fc8feee15bbfdf66d150e79 Fix bug when deleting wiki with no code write permission [24274](https://github.com/go-gitea/gitea/pull/24274)
  - [WIKI] commit d347208114966166ffa9655adc5b202676546c31 Improve External Wiki in Repo Header [24304](https://github.com/go-gitea/gitea/pull/24304)
  - [WIKI] commit db582d97ef6cd7d9f73a63c99639f6d00f40dc5a Improve wiki user title test [24559](https://github.com/go-gitea/gitea/pull/24559)
  - [WIKI] commit 60e7963141681895dcc81da944192c4292c6a20a Fix inconsistent wiki path converting. [24277](https://github.com/go-gitea/gitea/pull/24277)
  - [WIKI] commit b39a5bbbd610ba30651218658caaec1c86d6bca1 Make wiki title supports dashes and improve wiki name related features [24143](https://github.com/go-gitea/gitea/pull/24143)
- [UI / UX]
  - [BREAKING UX preview render] commit 84daddc2fa74393cdc13371b0cc44f0444cfdae0 Editor preview support for external renderers [23333](https://github.com/go-gitea/gitea/pull/23333)
  - [BREAKING branding] commit d44e1565dadd09b4cdbb924479bf6e59a4d3c403 Refactor `setting.Other` and remove unused `SHOW_FOOTER_BRANDING` (#24270)
  - [BREAKING theme tags] commit c7612d178c5b954d4846cd27a65a7fa15fd1ba65 Remove meta tags `theme-color` and `default-theme` [24960](https://github.com/go-gitea/gitea/pull/24960)
  - [BREAKING UI] commit 520eb57d7642a5fca3df319e5b5d1c7c9018087c Use a separate admin page to show global stats, remove `actions` stat [25062](https://github.com/go-gitea/gitea/pull/25062)
  - [UI] commit 6e90a1459b980f68a0f43c60c04fbb857e6a105b Add word-break to sidebar-item-link [23146](https://github.com/go-gitea/gitea/pull/23146)
  - [UI] commit 303b72c2d12bba44dc3def5fb3dfc1e5418a83ab Fix Fomantic UI's `touchstart` fastclick, always use `click` for click events [23065](https://github.com/go-gitea/gitea/pull/23065)
  - [UI] commit 10cdcb9ea8077098921d72720f9f36fcfd950452 Add "Reviewed by you" filter for pull requests [22927](https://github.com/go-gitea/gitea/pull/22927)
  - [UI] commit 843f81113ebe71fd725210c5a382268333865cc7 Projects: rename Board to Column in interface and improve consistency [22767](https://github.com/go-gitea/gitea/pull/22767)
  - [UI] commit f4920c9c7f5947d3b6476610f39bc3492ab4ef3b Add pagination for dashboard and user activity feeds [22937](https://github.com/go-gitea/gitea/pull/22937)
  - [UI] commit d20b29d7cea0fcba5e423dcfb7bbb7d2c15959d6 Fix height for sticky head on large screen on PR page [23111](https://github.com/go-gitea/gitea/pull/23111)
  - [ACTIONS] commit edf98a2dc30956c8e04b778bb7f1ce55c14ba963 Require approval to run actions for fork pull request [22803](https://github.com/go-gitea/gitea/pull/22803)
  - [UI] commit 0bc8bb3cc4f003e70bfee75863b74c2243c6d23c Make issue meta dropdown support Enter, confirm before reloading [23014](https://github.com/go-gitea/gitea/pull/23014)
  - [UI] commit 403f3e9208b2d2564d67bdf87be758c487083e28 Use the correct selector to hide the checkmark of selected labels on clear [23224](https://github.com/go-gitea/gitea/pull/23224)
  - [UI] commit 7a5af25592003ddc3017fcd7b822a3e02fc40ef6 Fix incorrect checkbox behaviors in the dashboard repolist's filter [23147](https://github.com/go-gitea/gitea/pull/23147)
  - [UI] commit 188c8c12c290e131fb342e3203634828652b0af5 Make Ctrl+Enter submit a pending comment (starting review) instead of submitting a single comment [23245](https://github.com/go-gitea/gitea/pull/23245)
  - [UI BIG] commit 7f9d58fab8a3c4fd1a8f18d58e36fbfab7b30f33 Support paste treepath when creating a new file or updating the file name [23209](https://github.com/go-gitea/gitea/pull/23209)
  - [UI] commit ea1d09718ce2e3bf043c60bae76dd6bd7e84e9fe Fix commit retrieval by tag [21804](https://github.com/go-gitea/gitea/pull/21804)
  - [UI] commit 0945bf63d3e0784c06a9d85504d42458b091d6b8 Fix missed `.hide` class [23208](https://github.com/go-gitea/gitea/pull/23208)
  - [UI BIG] commit de6c718b46ebd3b7f6362c766eed328044d95ec7 Allow `<video>` in MarkDown [22892](https://github.com/go-gitea/gitea/pull/22892)
  - [UI BIG] commit 545495dcb0a4cb9d820132dde4f1127f7fe91aa4 Pull Requests: add button to compare force pushed commits [22857](https://github.com/go-gitea/gitea/pull/22857)
  - [UI] commit ea7f0d6fcfe9567cac8151536b36450de8645e88 Change interactiveBorder to fix popup preview  [23169](https://github.com/go-gitea/gitea/pull/23169)
  - [UI] commit d949d8e074407a96dbcfa98a71ccd80527b5ad78 add user visibility in dashboard navbar [22747](https://github.com/go-gitea/gitea/pull/22747)
  - [UX] commit dad057b6393548ad389ead07c2cce5b3ac2811e0 Handle OpenID discovery URL errors a little nicer when creating/editing sources [23397](https://github.com/go-gitea/gitea/pull/23397)
  - [UX] commit d647e74502fdf734c89b3e6592a9ad88c3005971 Improve squash merge commit author and co-author with private emails [22977](https://github.com/go-gitea/gitea/pull/22977)
  - [UI] commit 17c8a0523a9f1bf717046823d4c0f69d606fe90a Fix and move "Use this template" button [23398](https://github.com/go-gitea/gitea/pull/23398)
  - [UI] commit a04eeb2a548d4fd2b63873fc2acff49c52e19723 Show edit/close/delete button on organization wide repositories [23388](https://github.com/go-gitea/gitea/pull/23388)
  - [UI] commit e72290fd9aeb77a47311483d1d565e428ce40cd9 Sync the class change of Edit Column Button to JS code [23400](https://github.com/go-gitea/gitea/pull/23400)
  - [UI] commit 75022f8b1a513ca2fd7ca66a2f05ecc49e2f1460 Refactor branch/tag selector dropdown (first step) [23394](https://github.com/go-gitea/gitea/pull/23394)
  - [UX] commit 3de9e63fd04d61e08fcbdec035c9f138347d9f37 Hide target selector if tag exists when creating new release [23171](https://github.com/go-gitea/gitea/pull/23171)
  - [UI] commit cf29ee6dd290525635a0e1b823506e81f845b978 Add missing tabs to org projects page [22705](https://github.com/go-gitea/gitea/pull/22705)
  - [UI] commit bf730528cadc4727eab8844934b6a0716d327243 Fix 'View File' button in code search [23478](https://github.com/go-gitea/gitea/pull/23478)
  - [UI] commit aac07d010f261c00fb3bd9644c71dc108c668c11 Add workflow error notification in ui [23404](https://github.com/go-gitea/gitea/pull/23404)
  - [UI] commit 6ff5400af91aefb02cbc7dd59f6be23cc2bf7865 Make branches list page operations remember current page [23420](https://github.com/go-gitea/gitea/pull/23420)
  - [UI] commit e82f1b15c7120ad13fd3b67cf7e2c6cb9915c22d Refactor dashboard repo list to Vue SFC [23405](https://github.com/go-gitea/gitea/pull/23405)
  - [UI] commit 81fe5d61851c0e586af7d32c29171ceff9a571bb Convert `<div class="button">` to `<button class="button">` [23337](https://github.com/go-gitea/gitea/pull/23337)
  - [UX] commit 5eea61dbc8f8e82e0dd05addf76751ee517459a0 Fix missing commit status in PR which from forked repo [23351](https://github.com/go-gitea/gitea/pull/23351)
  - [UX UI search] commit 661e78bed5c0879c32c53eb60f3d6898b93e1f08 Allow both fullname and username search when `DEFAULT_SHOW_FULL_NAME` is true [23463](https://github.com/go-gitea/gitea/pull/23463)
  - [UI] commit 39d3711f3036db42d7ddf73dbdb125be611bcbba Change `Close` to either `Close issue` or `Close pull request` [23506](https://github.com/go-gitea/gitea/pull/23506)
  - [UX review] commit a8c30a45fa49a3a551b1dca882960008c254bb3d `Publish Review` buttons should indicate why they are disabled [23598](https://github.com/go-gitea/gitea/pull/23598)
  - [UI] commit 529bac1950adf4fb72433cb67f66b0eec49224fe Polyfill the window.customElements [23592](https://github.com/go-gitea/gitea/pull/23592)
  - [UI GPG] commit 12ddc48c5c02123b1e6dab9d2d38b03f027d196e Use octicon-verified for gpg signatures [23529](https://github.com/go-gitea/gitea/pull/23529)
  - [UI stars] commit 06c067bb0f9eeb8873ddc298819b30fc5913943f Remove stars in dashboard repo list [23530](https://github.com/go-gitea/gitea/pull/23530)
  - [UI] commit 272cf6a2a976d4e22cccdaea720f48861d9b200e Make time tooltips interactive [23526](https://github.com/go-gitea/gitea/pull/23526)
  - [UI] commit 389e83f7eb68c43f6f0313b20acde547aef12442 Improve `<SvgIcon>` to make it output `svg` node and optimize performance [23570](https://github.com/go-gitea/gitea/pull/23570)
  - [UX issue config] commit f384b13f1cd44be3a87df5553a0099390dacd010 Implement Issue Config [20956](https://github.com/go-gitea/gitea/pull/20956)
  - [UI] commit 2c585d62a4ebbb52175b8fd8697458ae1c3b2937 User/Org Feed render description as per web [23887](https://github.com/go-gitea/gitea/pull/23887)
  - [UI TAGS] commit b78c955958301dde72d8caf189531f6e53c496b4 Fix tags view [23243](https://github.com/go-gitea/gitea/pull/23243)
  - [UI] commit 9cefb7be737c6564668d95d8a43b4425e9a03d13 Fix new issue/pull request btn margin when it is next to sort [23647](https://github.com/go-gitea/gitea/pull/23647)
  - [UX preview] commit ac64c8297444ade63a2a364c4afb7e6c1de5a75f Allow new file and edit file preview if it has editable extension [23624](https://github.com/go-gitea/gitea/pull/23624)
  - [UI] commit ca905b82df7f1d2a823d8df4448d485e5902876d Append `(comment)` when a link points at a comment rather than the whole issue [23734](https://github.com/go-gitea/gitea/pull/23734)
  - [UX diff] commit aa4d1d94f79e8edd9fa9ff2813fea12b085b2cae Diff improvements [23553](https://github.com/go-gitea/gitea/pull/23553)
  - [UX ONLY_SHOW_RELEVANT_REPOS] commit e57e1144c5ae7a2995e6818c6ae32139e563add7 Add ONLY_SHOW_RELEVANT_REPOS back, fix explore page bug, make code more strict [23766](https://github.com/go-gitea/gitea/pull/23766)
  - commit ed5e7d03c6c44666c6fe97a15e8ce33d223c4466 Don't apply the group filter when listing LDAP group membership if it is empty [23745](https://github.com/go-gitea/gitea/pull/23745)
  - [UX allow . in name] commit 88033438aa8214569913899a17b19b57bd609d97 Support "." char as user name for User/Orgs in RSS/ATOM/GPG/KEYS path ... [23874](https://github.com/go-gitea/gitea/pull/23874)
  - [UI] commit ca5722a0fae6cc16dc99021176596970bbf29caf Ensure RSS icon is present on all repo tabs [23904](https://github.com/go-gitea/gitea/pull/23904)
  - [UI] commit 6eb678374b583079a0a08b7ed0c9ca220c0c0434 Refactor authors dropdown (send get request from frontend to avoid long wait time) [23890](https://github.com/go-gitea/gitea/pull/23890)
  - [UX RELEASE permalink] commit 42919ccb7cd32ab67d0878baf2bac6cd007899a8 Make Release Download URLs predictable [23891](https://github.com/go-gitea/gitea/pull/23891)
  - [UX project] commit 6a4be2cb6a6193a3f41d5e08d05044e3c54efc38 Add cardtype to org/user level project on creation, edit and view [24043](https://github.com/go-gitea/gitea/pull/24043)
  - [UX] commit 52b17bfa07fea29441cd961da4edaf1ea97fe348 Add repository counter badge to repository tab [24205](https://github.com/go-gitea/gitea/pull/24205)
  - [UX dump] commit cb1536471bcef4d78a3fe5cbd738b9f60fabbcc2 Add --quiet option to gitea dump [22969](https://github.com/go-gitea/gitea/pull/22969)
  - [UI] commit 774d1a0fbdadd1136b6af895f8d449b0c8db54cb Tweak pull request branch delete ui [23951](https://github.com/go-gitea/gitea/pull/23951)
  - [UI] commit 9c33cbd3441fc866c3056d53b62c5420588a538a Fix no edit/close/delete button in org repo project view page  [24301](https://github.com/go-gitea/gitea/pull/24301)
  - [UX] commit c41bc4f1279c9e1e6e11d7b5fcfe7ef089fc7577 Display when a repo was archived [22664](https://github.com/go-gitea/gitea/pull/22664)
  - [UI] commit 83022013c83feb5488952baea3ef0be818dfce21 Fix layouts of admin table / adapt repo / email test  [24370](https://github.com/go-gitea/gitea/pull/24370)
  - [UX] commit e9b39250b285f1b9cbf9739f33c06fc57401f314 Improve pull request merge box when pull request merged and branch deleted. [24397](https://  - [UI] commit 94d6b5b09d49b2622c2164a03cfae45dced96c74 Add "Updated" column for admin repositories list [24429](https://github.com/go-gitea/gitea/pull/24429)
github.com/go-gitea/gitea/pull/24397)
  - [UI] commit 72e956b79a3b2e055bb5d4d5e20e88eaa2eeec96 Improve protected branch setting page [24379](https://github.com/go-gitea/gitea/pull/24379)
  - [UX goto issue] commit 1144b1d129de530b2c07dfdfaf55de383cd82212 Add goto issue id function [24479](https://github.com/go-gitea/gitea/pull/24479)
  - [UI] commit 97b70a0cd40e8f73cdf6ba4397087b45061de3d8 Add org visibility label to non-organization's dashboard [24558](https://github.com/go-gitea/gitea/pull/24558)
  - [UX] commit 4daf40505a5f89747982ddd2f1df2a4001720846 Sort users and orgs on explore by recency by default [24279](https://github.com/go-gitea/gitea/pull/24279)
  - [UX graceful restart] commit 7565e5c3de051400a9e3703f707049cbb9054cf3 Implement systemd-notify protocol [21151](https://github.com/go-gitea/gitea/pull/21151)
  - [UX] commit 4810fe55e3e73edb962052df46bef125eb1817b3 Add status indicator on main home screen for each repo [24638](https://github.com/go-gitea/gitea/pull/24638)
  - [UX] commit b5c26fa825e08122843ad6d27191d399a9af1c37 Add markdown preview to Submit Review Textarea [24672](https://github.com/go-gitea/gitea/pull/24672)
  - [UX issue template] commit c4303efc23ea19f16ee826809f43888ee4583ebb Support markdown editor for issue template [24400](https://github.com/go-gitea/gitea/pull/24400)
  - [UI] commit 4aec1f87a4affe606e96e27c2e8660092aac3afb Remove highlight in repo list [24675](https://github.com/go-gitea/gitea/pull/24675)
  - [UI] commit 8251b317f7b7a2b5f626a02fa3bb540a1495e81d Improve empty notifications display [24668](https://github.com/go-gitea/gitea/pull/24668)
  - [UX] commit f6e029e6c7849d4361abf7f1d749b5d528364ac4 Make repo migration cancelable and fix various bugs [24605](https://github.com/go-gitea/gitea/pull/24605)
  - [UI] commit b3af7484bc821d71cb20f6289f767119494bc81e Fix missing badges in org settings page [24654](https://github.com/go-gitea/gitea/pull/24654)
  - [UI RSS] commit 67db6b697636221e09536e89ac8600a47f79b5cb RSS icon fixes [24476](https://github.com/go-gitea/gitea/pull/24476)
  - [UX notification list] commit f7ede92f82f7f3ec7bb31a1249f9524e5b728f34 Notification list enhancements, fix striped tables on dark theme [24639](https://github.com/go-gitea/gitea/pull/24639)
  - [UI] commit ea7954f069bf8bcb87d520f8aab0a80b0768590d Modify luminance calculation and extract related functions into single files [24586](https://github.com/go-gitea/gitea/pull/24586)
  - [UX review] commit ae0fa64ef6261bc99b9b7f6af2047c017399f509 Review fixes and enhancements [24526](https://github.com/go-gitea/gitea/pull/24526)
  - [UI] commit df00ccacc9a4840fe86bed75a77841f8801d11d2 Fix invite display [24447](https://github.com/go-gitea/gitea/pull/24447)
  - [UX] commit e8173c2c33f1dd5b0a2c044255434d414cab62d2 Move `Rename branch` from repo settings page to the page of branches list [24380](https://github.com/go-gitea/gitea/pull/24380)
  - [UX] commit 3f0651d4d61d62a16e1bb672056014ab02db5746 Improve milestone filter on issues page [22423](https://github.com/go-gitea/gitea/pull/22423)
  - [UI] commit 8f4dafcd4e6b0b5d307c3e060ffe908c2a96f047 Rework header bar on issue, pull requests and milestone [24420](https://github.com/go-gitea/gitea/pull/24420)
  - [UI] commit 8bbbf7e6b8af072e8c924982019c1fc544403196 Remove fluid on compare diff page [24627](https://github.com/go-gitea/gitea/pull/24627)
  - [UI avatar] commit 82224c54e0488738dbd3b7eccf56ab08b6790627 Improve avatar uploading / resizing / compressing, remove Fomantic card module [24653](https://github.com/go-gitea/gitea/pull/24653)
  - [UI] commit b9fad73e9fcf40e81cde3304198105af6c668421 Unification of registration fields order [24737](https://github.com/go-gitea/gitea/pull/24737)
  - [UI] commit 6a3a54cf484bf5137e2af5bc93294b783feb23a4 Remove background on user dashboard filter bar [24779](https://github.com/go-gitea/gitea/pull/24779)
  - [UX] commit b807d2f6205bf1ba60d3a543e8e1a16f7be956df Support no label/assignee filter and batch clearing labels/assignees [24707](https://github.com/go-gitea/gitea/pull/24707)
  - [UI] commit 5c0745c0349f0709d0fc36fd8a97dcab86bce28a Add validations.required check to dropdown field [24849](https://github.com/go-gitea/gitea/pull/24849)
  - [UX notifications list] commit 27c221aa5db116e2cc90afbfa9b92f2a220af853 Rework notifications list [24812](https://github.com/go-gitea/gitea/pull/24812)
  - [UI] commit 35ce7ca25b5756441949312d79aa6382f98ce8d6 Hide 'Mirror Settings' when unneeded, improve hints [24433](https://github.com/go-gitea/gitea/pull/24433)
  - [UX] commit a70d853d064a97f0be1d3702a9c3912494b546ec Consolidate the two review boxes into one [24738](https://github.com/go-gitea/gitea/pull/24738)
  - [UI] commit e3897148f9e612e363854d790a1e77807dac8d0d Minor UI improvements: logo alignment, auth map editor, auth name display [25043](https://github.com/go-gitea/gitea/pull/25043)
  - [UX tree view] commit 72eedfb91584720da774909d3f078b7d515c9fdd Show file tree by default [25052](https://github.com/go-gitea/gitea/pull/25052)
  - [UX diff copy] commit c5ede35124c8d5280219c24049bb0ad7da9f02ed Add button on diff header to copy file name, misc diff header tweaks [24986](https://github.com/go-gitea/gitea/pull/24986)
  - [UI] commit 58536093b3112841bc69edb542189893b57e7a47 Add details summary for vertical menus in settings to allow toggling [25098](https://github.com/go-gitea/gitea/pull/25098)
  - [UI] commit 7d192cb674bc475b123c84b205aca821247c5dd1 Add Progressbar to Milestone Page [25050](https://github.com/go-gitea/gitea/pull/25050)
  - [UI] commit 7abe958f5b507efa676fb3b2e27d30517f6d1908 Fix color for transfer related buttons when having no permission to act [24510](https://github.com/go-gitea/gitea/pull/24510)
  - [UI] commit 4a722c9a45659e7732258397bbb3dd1039ea1952 Make Issue/PR/projects more compact, misc CSS tweaks [24459](https://github.com/go-gitea/gitea/pull/24459)
- [PERF]
  - [PERF] commit 1319ba6742a8562453646763adad22379674bab5 Use minio/sha256-simd for accelerated SHA256 [23052](https://github.com/go-gitea/gitea/pull/23052)
  - [PERF] commit ef4fc302468cc8a9fd8f65c4ebdc6f55138450d1 Speed up HasUserStopwatch & GetActiveStopwatch [23051](https://github.com/go-gitea/gitea/pull/23051)
  - [PERF] commit 0268ee5c37b8ad733678f02bc15ec8642da62c10 Do not create commit graph for temporary repos [23219](https://github.com/go-gitea/gitea/pull/23219)
  - [PERF] commit 75ea0d5dba5dbf2f84cef2d12460fdd566d43e62 Faster git.GetDivergingCommits [24482](https://github.com/go-gitea/gitea/pull/24482)
  - [PERF] commit df48af22296ccce8e9bd18e5d35c9a3cdf5acb0f Order pull request conflict checking by recently updated, for each push [23220](https://github.com/go-gitea/gitea/pull/23220)
- [AUTH]
  - [MAIL smtp auth] commit 8be6da3e2fd0b685aeb6b9e7fd9dee5a4571163a Add ntlm authentication support for mail [23811](https://github.com/go-gitea/gitea/pull/23811)
  - [AUTH LDAP] commit b8c19e7a11525da4174b6f80f87ff3e844d03d8a Update LDAP filters to include both username and email address [24547](https://github.com/go-gitea/gitea/pull/24547)
  - [AUTH PKCE] commit 7d855efb1fe6b97c5d87492f67ed6aefd31b2474 Allow for PKCE flow without client secret + add docs [25033](https://github.com/go-gitea/gitea/pull/25033)
  - [AUTH OAuth redirect] commit ca35dec18b3d3d7dd5cde4c69a10ae830961faf7 Add ability to set multiple redirect URIs in OAuth application UI [25072](https://github.com/go-gitea/gitea/pull/25072)
- [REFACTOR]
  - [BREAKING REFACTOR logger] commit 4647660776436f0a83129b4ceb8426b1fb0599bb Rewrite logger system [24726](https://github.com/go-gitea/gitea/pull/24726)
  - [BREAKING REFACTOR queue] commit 6f9c278559789066aa831c1df25b0d866103d02d Rewrite queue [24505](https://github.com/go-gitea/gitea/pull/24505)
  - [REFACTOR pull mirror] commit 99283415bcbaa8acfe4d249ce3040de2f3a8b006 Refactor Pull Mirror and fix out-of-sync bugs [24732](https://github.com/go-gitea/gitea/pull/24732)
  - [REFACTOR git] commit f4538791f5fc82b173608fcf9c30e36ec01dc9d3 Refactor internal API for git commands, use meaningful messages instead of "Internal Server Error" [23687](https://github.com/go-gitea/gitea/pull/23687)
  - [REFACTOR route] commit 92fd3fc4fd369b6a8c0a022a32a80dec2340223a Refactor "route" related code, fix Safari cookie bug [24330](https://github.com/go-gitea/gitea/pull/24330)
  - [REFACTOR] commit 8598356df1eb21b6e33ecb9f9268ba36c5488e7c Refactor and tidy-up the merge/update branch code [22568](https://github.com/go-gitea/gitea/pull/22568)
  - [REFACTOR] commit 542cec98f8c07e0f046a35f1d516807416536e74 Refactor merge/update git command calls [23366](https://github.com/go-gitea/gitea/pull/23366)
  - [REFACTOR] commit ec261b63e14f84da3e2d9a6e27c8b831a7750677 Refactor repo commit list [23690](https://github.com/go-gitea/gitea/pull/23690)
  - [REFACTOR cookie] commit 5b9557aef59b190c55de9ea218bf51152bc04786 Refactor cookie [24107](https://github.com/go-gitea/gitea/pull/24107)
  - [REFACTOR web route] commit b9a97ccd0ea1ee44db85b0fbb80b75255af7c742 Refactor web route [24080](https://github.com/go-gitea/gitea/pull/24080)
  - [REFACTOR issue stats] commit 38cf43d0606c13c38f459659f38e26cf31dceccb Some refactors for issues stats [24793](https://github.com/go-gitea/gitea/pull/24793)
  - [REFACTOR] commit c59a057297c782f44a81a3e630b5094a58099edb Refactor rename user and rename organization [24052](https://github.com/go-gitea/gitea/pull/24052)
  - [REWORK logger] commit 0d54395fb544d52585046bf0424659cec0626e31 Improve logger Pause handling [24946](https://github.com/go-gitea/gitea/pull/24946)
  - [REWORK queue / logger] commit 18f26cfbf7f9b36b838c0e8762bfba98c89b9797 Improve queue and logger context [24924](https://github.com/go-gitea/gitea/pull/24924)
  - [REFACTOR scoped token] commit 18de83b2a3fc120922096b7348d6375094ae1532 Redesign Scoped Access Tokens [24767](https://github.com/go-gitea/gitea/pull/24767)
  - [REFACTOR ini] commit de4a21fcb4476772c69c36d086549e89ed4dcf6c Refactor INI package (first step) [25024](https://github.com/go-gitea/gitea/pull/25024)
  - [REFACTOR diffFileInfo] commit ee99cf6313ba565523b3c43f61ffda4b71e2c39b Refactor diffFileInfo / DiffTreeStore  [24998](https://github.com/go-gitea/gitea/pull/24998)
- [TEMPLATES]
  - [TEMPLATES expressions] commit 5b89670a318e52e271f65d96bfe1116d85d20988 Use a general Eval function for expressions in templates. [23927](https://github.com/go-gitea/gitea/pull/23927)
  - [CMD reload templates] commit 3588edbb08f93aaa56defa82dffdbb202cd9aa4a Add gitea manager reload-templates command [24843](https://github.com/go-gitea/gitea/pull/24843)
- [RSS]
  - [RSS feed] commit 59d060622d375c4123ea88e2fa6c4f34d4fea4d3 Improve RSS [24335](https://github.com/go-gitea/gitea/pull/24335)
  - [RSS feed] commit 56d4893b2a996da6388801c9c8ff16b9b588ad55 Add RSS Feeds for branches and files [22719](https://github.com/go-gitea/gitea/pull/22719)
- [API]
  - [API EMAIL] commit d56bb7420184c0c2f451f4bcaa96c9b3b00c393d add admin API email endpoints [22792](https://github.com/go-gitea/gitea/pull/22792)
  - [API USER RENAME] commit 03591f0f95823a0b1dcca969d2a3ed505c7e6d73 add user rename endpoint to admin api [22789](https://github.com/go-gitea/gitea/pull/22789)
  - [API admin search] commit 6f9cc617fcc42477dec5ccab83d06f0a96544403 Add login name and source id for admin user searching API [23376](https://github.com/go-gitea/gitea/pull/23376)
  - [API] commit 574d8fe6d6675c8aa05e2b75fdbc01c009efd8be Add absent repounits to create/edit repo API [23500](https://github.com/go-gitea/gitea/pull/23500)
  - [API issue dependencies] commit 3cab9c6b0c050bfcb9f2f067e7dc1b0242875254 Add API to manage issue dependencies [17935](https://github.com/go-gitea/gitea/pull/17935)
  - [API activity feeds] commit 6b0df6d8da76d77a9b5c42dcfa78dbfe197fd56d Add activity feeds API [23494](https://github.com/go-gitea/gitea/pull/23494)
  - [API license] commit fb37eefa282543fd8ce63c361cd4cf0dfac9943c Add API for License templates [23009](https://github.com/go-gitea/gitea/pull/23009)
  - [API gitignore] commit 36a5d4c2f3b5670e5e921034cd5d25817534a6d4 Add API for gitignore templates [22783](https://github.com/go-gitea/gitea/pull/22783)
  - [API upload empty repo] commit cf465b472166ccf6d3e001e3043e4bf43e16e6b3 Support uploading file to empty repo by API [24357](https://github.com/go-gitea/gitea/pull/24357)
  - [API COMMIT --not] commit f766b002938b5c81e343c81fda3c0669fa09809f Add ability to specify '--not' from GetAllCommits [24409](https://github.com/go-gitea/gitea/pull/24409)
  - [API GetAllCommits] commit 1dd83dbb917d55bd253001646d6743f247a4d98b Filters for GetAllCommits [24568](https://github.com/go-gitea/gitea/pull/24568)
  - [API get single commit] commit 5930ab5fdf7a970fcca3cd50b44cf1cacb615a54 Filter get single commit [24613](https://github.com/go-gitea/gitea/pull/24613)
  - [API create branch] commit cd9a13ebb47d32f46b38439a524e3b2e0c619490 Create a branch directly from commit on the create branch API [22956](https://github.com/go-gitea/gitea/pull/22956)
  - [BREAKING API team] commit 0a3c4d4a595cc7e12462dde393ed64186260f26b Fix team members API endpoint pagination [24754](https://github.com/go-gitea/gitea/pull/24754)
  - [API label templates] commit 25dc1556cd70b567a4920beb002a0addfbfd6ef2 Add API for Label templates [24602](https://github.com/go-gitea/gitea/pull/24602)
  - [API changing/creating/deleting multiple files] commit 275d4b7e3f4595206e5c4b1657d4f6d6969d9ce2 API endpoint for changing/creating/deleting multiple files [24887](https://github.com/go-gitea/gitea/pull/24887)
- [FEATURES]
  - [BREAKING] (maybe) commit f5987c24e2b561952ebf9a2485b863325c16ee48 Make `gitea serv` respect git binary home [23138](https://github.com/go-gitea/gitea/pull/23138)
  - [README] commit 52e24167e5ebe0297f7630e9daecd6ffc9570a99 Test renderReadmeFile [23185](https://github.com/go-gitea/gitea/pull/23185)
  - [REFLOGS] commit 757b4c17e900f1d11a81bc9467d90e6c245ee8f2 Support reflogs [22451](https://github.com/go-gitea/gitea/pull/22451)
  - [DOCTOR] commit df411819ebe4d3e6852997ce41fadf837d5d4ea0 Check LFS/Packages settings in dump and doctor command [23631](https://github.com/go-gitea/gitea/pull/23631)
  - [MINIO] commit 0e7bec1849d2d7a87713abe494b4d3ef416180d4 Add InsecureSkipVerify to Minio Client for Storage [23166](https://github.com/go-gitea/gitea/pull/23166)
  - [MINIO MD5 checksum] commit 5727056ea109eb04ee535b981349cdfb44df9fae Make minio package support legacy MD5 checksum [23768](https://github.com/go-gitea/gitea/pull/23768)
  - [PRIVACY email display] commit 6706ac2a0f5f2fe4f8e2555be7e2a8b4d5946398 Fix profile page email display, respect settings [23747](https://github.com/go-gitea/gitea/pull/23747)
  - [INDEX meilisearch] commit 92c160d8e716cb3d05215a97cf521e843596f562 Add meilisearch support [23136](https://github.com/go-gitea/gitea/pull/23136)
  - [PRIVACY email] commit 5e1bd8af5f16f9db88cfeb5b80bdf731435cacfb Show visibility status of email in own profile [23900](https://github.com/go-gitea/gitea/pull/23900)
  - [BREAKING SSH key parsing] commit 7a8a4f54321f208ebbb0f708a5f0e49c4cd4cc04 Prefer native parser for SSH public key parsing [23798](https://github.com/go-gitea/gitea/pull/23798)
  - [REDIS] commit 985f76dc4b0692c4d6c6f37e82500ef859557c16 Update redis library to support redis v7 [24114](https://github.com/go-gitea/gitea/pull/24114)
  - [RESERVED users] commit 1819c4b59b81ba4db2a38d3b3dc81f29102fde51 Add new user types `reserved`, `bot`, and `remote` [24026](https://github.com/go-gitea/gitea/pull/24026)
  - [NEW files to empty repo] commit e422342eebc18034ef586ec58f1e2fff0340091d Allow adding new files to an empty repo [24164](https://github.com/go-gitea/gitea/pull/24164)
  - [WEBP avatars] commit 65fe0fb22cfb264f0b756065d0c3ce7a17d7e55b Allow `webp` images as avatars [24248](https://github.com/go-gitea/gitea/pull/24248)
  - [MARKDOWN livemd] commit 58caf422e67c78f87327bc9b00f89083a2432940 Add .livemd as a markdown extension [22730](https://github.com/go-gitea/gitea/pull/22730)
  - [FOLLOW org] commit cc64a925602d54f3439dd19f16b5280bd0377a7a Add follow organization and fix the logic of following page [24345](https://github.com/go-gitea/gitea/pull/24345)
  - [PROFILE README] commit c090f87a8db5b51e0aa9c7278b38ddc862c048ac Add Gitea Profile Readmes [23260](https://github.com/go-gitea/gitea/pull/23260)
  - [HTTP RANGE] commit 023a048f52b5bf8c4b715285245a129f04e05a8c Make repository response support HTTP range request [24592](https://github.com/go-gitea/gitea/pull/24592)
  - [status check pattern] commit e7c2231dee356df5cbe5a47c07e31e3a8d090a6f Support for status check pattern [24633](https://github.com/go-gitea/gitea/pull/24633)
  - [EMAIL allow/block] commit 2cb66fff60c95efbd58b797f1197f2421f4687ce Support wildcard in email domain allow/block list [24831](https://github.com/go-gitea/gitea/pull/24831)
  - [INSTALL page] commit abcf5a7b5e2c3df951b8048317a99a89b040b489 Fix install page context, make the install page tests really test [24858](https://github.com/go-gitea/gitea/pull/24858)
  - [environment-to-ini FILE] commit c21605951b581440bb08b65d5907b1cd4e0ab6c5 Make environment-to-ini  support loading key value from file [24832](https://github.com/go-gitea/gitea/pull/24832)
  - [APP ini git config] commit 8080ace6fcf73a5fbe4a0dd71881228abd0c68b9 Support changing git config through `app.ini`, use `diff.algorithm=histogram` by default [24860](https://github.com/go-gitea/gitea/pull/24860)
  - [PIN issues] commit aaa109466350c531b9238a61115b2877daca57d3 Add the ability to pin Issues [24406](https://github.com/go-gitea/gitea/pull/24406)
  - [BREAKING reflog / config] commit 2f149c5c9db97f20fbbc65e32d1f3133048b11a2 Use `[git.config]` for reflog cleaning up [24958](https://github.com/go-gitea/gitea/pull/24958)
  - [SEARCH skip forks mirrors] commit 033d92997fc16baee097d2b25f08e0984e628abd Allow skipping forks and mirrors from being indexed [23187](https://github.com/go-gitea/gitea/pull/23187)
- [WEBHOOK]
  - [WEBHOOKS] commit 2173f14708ff3b35d7821fc9b6dcb5fcd06b8494 Add user webhooks [21563](https://github.com/go-gitea/gitea/pull/21563)
  - [WEBHOOK] commit 9e04627acaaa853e5269f98f53f2615077cfb028 Fix incorrect `HookEventType` of pull request review comments [23650](https://github.com/go-gitea/gitea/pull/23650)
  - [WEBHOOK review request] commit 309354c70ee994a1e8f261d7bc24e7473e601d02 New webhook trigger for receiving Pull Request review requests [24481](https://github.com/go-gitea/gitea/pull/24481)
- [DISCARDED]
  - [GITEA only BREAKING service worker] commit 50bd7d0b24016b0cf48dfbafe84b5953fe20c34f Remove the service worker [25010](https://github.com/go-gitea/gitea/pull/25010)

* Container images upgraded to Alpine 3.18

  The Forgejo container images are now based on [Alpine 3.18](https://alpinelinux.org/posts/Alpine-3.18.0-released.html) instead of Alpine 3.1.17 It includes an upgrade from git ...

## 1.19.3-0

The [complete list of commits](https://codeberg.org/forgejo/forgejo/commits/branch/v1.19/forgejo) included in the `Forgejo v1.19.3-0` release can be reviewed from the command line with:

```shell
$ git clone https://codeberg.org/forgejo/forgejo/
$ git -C forgejo log --oneline --no-merges v1.19.2-0..v1.19.3-0
```

This stable release contains security fixes.

* Recommended Action

  We recommend that all Forgejo installations are upgraded to the latest version.

* [Forgejo Semantic Version](https://forgejo.org/docs/v1.19/user/semver/)

  The semantic version was updated from `4.2.0+0-gitea-1.19.2` to `4.2.1+0-gitea-1.19.3` because of the rebuild with [Go version 1.20.4](https://github.com/golang/go/issues?q=milestone%3AGo1.20.4+label%3ACherryPickApproved).

* Security fixes

  * Forgejo was recompiled with Go version v1.20.4 published 2 May 2023. It fixes [three vulnerabilities](https://github.com/golang/go/issues?q=milestone%3AGo1.20.4+label%3ACherryPickApproved) ([CVE-2023-29400](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-29400), [CVE-2023-24540](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-24540), [CVE-2023-24539](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-24539)) related to the [html/template](https://pkg.go.dev/html/template) package. The [Forgejo security team](https://forgejo.org/.well-known/security.txt) analyzed the security fixes it contains and concluded that Forgejo is not affected but recommended a rebuild as a precaution.

* Bug fixes

  The most prominent one is described here, others can be found in the list of commits included in the release as described above.

  * [Allow users that are not signed in to browse the tag list](https://codeberg.org/forgejo/forgejo/commit/e76b3f72b23bf778a562682d829744451b27d615). Requiring users to be signed in to view the tag list was a regression introduced in Forgejo v1.19.2-0.

## 1.19.2-0

The [complete list of commits](https://codeberg.org/forgejo/forgejo/commits/branch/v1.19/forgejo) included in the `Forgejo v1.19.2-0` release can be reviewed from the command line with:

```shell
$ git clone https://codeberg.org/forgejo/forgejo/
$ git -C forgejo log --oneline --no-merges v1.19.1-0..v1.19.2-0
```

This stable release contains **important security fixes**.

* Recommended Action

  We **strongly recommend** that all Forgejo installations are upgraded to the latest version as soon as possible.

* [Forgejo Semantic Version](https://forgejo.org/docs/v1.19/user/semver/)

  The semantic version was updated from `4.1.0+0-gitea-1.19.1` to `4.2.0+0-gitea-1.19.2` because of the changes introduced in the internal CI.

* Security fixes

  * Token scopes were not enforced in some cases ([patch 1](https://codeberg.org/forgejo/forgejo/commit/7c3ac69c0) and [patch 2](https://codeberg.org/forgejo/forgejo/commit/10d3ed53f1cc6d383b52637bedd7bc3679476eb4)). The [scoped token](https://forgejo.org/docs/v1.19/user/oauth2-provider/#scoped-tokens) were introduced in Forgejo v1.19 allow for the creation of application tokens that only have limited permissions, such as creating packages or accessing repositories. Prior to Forgejo v1.19 tokens could be used to perform any operation the user issuing the token could.
  * [Permissions to delete secrets was not enforced](https://codeberg.org/forgejo/forgejo/commit/68d80eb56). The experimental internal CI relies on secrets managed via the web interface, for instance to communicate credentials to a job. Secrets are only used in the context of the experimental internal CI.

* Bug fixes

  The most prominent ones are described here, others can be found in the list of commits included in the release as described above.

  * [Restore public access to some API endpoints](https://codeberg.org/forgejo/forgejo/commit/b00f7c3c545c6a00a747a5aea7596f45c50157ac). When [scoped token](https://forgejo.org/docs/v1.19/user/oauth2-provider/#scoped-tokens) introduced in Forgejo v1.19, some API endpoints that were previously accessible anonymously became restricted: `/orgs`, `/orgs/{org}`, `/orgs/{org}/repos`, `/orgs/{org}/public_members`, `/orgs/{org}/public_members/{username}`, `/orgs/{org}/labels`.
  * [Fix 2-dot direct compare to use the right base commit](https://codeberg.org/forgejo/forgejo/commit/494e373292962de34b7ea7efd3f4a8d2f27daa26). For 2-dot direct compare, the base commit should be used in the title and templates, as is used elsewhere, not the common ancestor which is used for 3-dot compare.
  * [Make CORS work](https://codeberg.org/forgejo/forgejo/commit/2e6e5bc9c96ebb760f28c08423bb0c244ca7e01c). No [CORS](https://en.wikipedia.org/wiki/Cross-origin_resource_sharing) headers were set, even if CORS was enabled in the configuration.
  * [Fix issue attachment removal](https://codeberg.org/forgejo/forgejo/commit/d5f2c9d74d63443cc2abbcabc268cf1121f58e8b). When an attachment was removed from an issue or review comment, all of the image/attachment links were broken.
  * [Fix wiki write permissions for users who do not have repository write permissions](https://codeberg.org/forgejo/forgejo/commit/8c465206e2fea27076fdb986ea0478729653f0b5). When a team member had write access to the wiki but not to the code repository, some operations (deleting a page for instance) were denied.
  * [Respect the REGISTER_MANUAL_CONFIRM setting when registering via OAuth](https://codeberg.org/forgejo/forgejo/commit/116b6d5b27c40b248281f5fd543f7aa8df0d59d3). Contrary to the local registration, the OAuth registration flow activated a newly registered user regardless of the value of `REGISTER_MANUAL_CONFIRM`.
  * [Fix tags list for repos whose release setting is disabled](https://codeberg.org/forgejo/forgejo/commit/eeee32cdc3aab4d2086b24aae670a39501c9ea99). When releases was disabled the "tags" button led to a `Not Found` page, even when tags existed.

* Container image upgrades

  In the Forgejo container images the Git version was upgraded to [2.38.5](https://github.com/git/git/blob/master/Documentation/RelNotes/2.38.5.txt) as a precaution. The [Forgejo security team](https://forgejo.org/.well-known/security.txt) analyzed the security fixes it contains and concluded that Forgejo is not affected.

## 1.19.1-0

The [complete list of commits](https://codeberg.org/forgejo/forgejo/commits/branch/v1.19/forgejo) included in the `Forgejo v1.19.1-0` release can be reviewed from the command line with:

```shell
$ git clone https://codeberg.org/forgejo/forgejo/
$ git -C forgejo log --oneline --no-merges v1.19.0-3..v1.19.1-0
```

This stable release includes bug fixes. Functional changes related to the experimental CI have also been backported.

* Recommended Action

  We recommend that all installations are upgraded to the latest version.

* [Forgejo Semantic Version](https://forgejo.org/docs/v1.19/user/semver/)

  The semantic version was updated from `4.0.0+0-gitea-1.19.0` to `4.1.0+0-gitea-1.19.1` because of the changes introduced in the internal CI.

* Bug fixes

  The most prominent ones are described here, others can be found in the list of commits included in the release as described above.

  * [Fix RSS/ATOM/GPG/KEYS path for users (or orgs) with a dot in their name](https://codeberg.org/forgejo/forgejo/commit/085b56302cfd9a949319a3a1e32e008b4a0d0772). It is allowed for a user (or an organization) to have a dot in their name, for instance `user.name`. Because of a [bug in Chi](https://codeberg.org/forgejo/forgejo/issues/652) it was not possible to access `/user.name.png`, `/user.name.gpg`, etc. A workaround was implemented while a [proper fix is being discussed](https://github.com/go-chi/chi/pull/811).
  * [Creating a tag via the web interface no longer requires a title](https://codeberg.org/forgejo/forgejo/commit/1b8ecd179bdb58427b99c2c2eb9ad5a45abf7055).
  * [Use fully qualified URLs in Dockerfile](https://codeberg.org/forgejo/forgejo/commit/833a4b177596debc138e5723219fd063d067bd5b). The Dockerfile to create the Forgejo container image now uses the fully qualified image `docker.io/library/golang:1.20-alpine3.17` instead of `golang:1.20-alpine3.17`. This allows for building on platforms that don't have docker hub as the default container registry.
  * [Redis use Get/Set instead of Rename when Regenerate session id](https://codeberg.org/forgejo/forgejo/commit/3a7cb1a83b4ecd89421b5656b8caeb30c2b13c7c). The old sid and new sid may be in different redis cluster slot.
  * [Do not escape space between PyPI repository url and package name](https://codeberg.org/forgejo/forgejo/commit/cfde557e231417b7fb3cde3e9bab70d05b7d182f). It also adds a trailing slash to the PyPI repository URL in accordance to [Python PEP-503](https://peps.python.org/pep-0503/).
  * [Fix failure when using the API and an empty rule_name to protect a branch](https://codeberg.org/forgejo/forgejo/commit/abf0386e2ef6b56c048c04cd3d6913f453c87cb1). The `rule_name` parameter for the [/repos/{owner}/{repo}/branch_protections](https://code.forgejo.org/api/swagger#/repository/repoCreateBranchProtection) API now defaults to the branch name instead of being empty.
  * [Fix branch protection priority](https://codeberg.org/forgejo/forgejo/commit/580da8f35320dbd15b168bf8ccfaff6187ff87e0). Contrary to [the documentation](https://forgejo.org/docs/v1.19/user/protection/#protected-branches) it was possible for a glob rule to take precedence over a non-glob rule.
  * [Fix deleting an issue when the git repo does not exist](https://codeberg.org/forgejo/forgejo/commit/1d8ae34e57e46b84a885b4f072d949344c5977c4). If a project had an issue tracker (such as the [Forgejo discussion](https://codeberg.org/forgejo/discussions/issues)) but [no git repository](https://codeberg.org/forgejo/discussions/), trying to delete an issue would fail.
  * [Fix accidental overwriting of LDAP team memberships](https://codeberg.org/forgejo/forgejo/commit/66aa85429684aca45753ac9578492ed3f7507ea3). If an LDAP user is a member of two groups, the LDAP group sync only matched the last group.

## 1.19.0-3

The [complete list of commits](https://codeberg.org/forgejo/forgejo/commits/branch/v1.19/forgejo) included in the `Forgejo v1.19.0-3` release can be reviewed from the command line with:

```shell
$ git clone https://codeberg.org/forgejo/forgejo/
$ git -C forgejo log --oneline --no-merges v1.19.0-2..v1.19.0-3
```

This stable release includes security updates and bug fixes.

* Recommended Action

  We recommend that all installations are upgraded to the latest version.

* Security

  The [Forgejo security team](https://forgejo.org/.well-known/security.txt) analyzed the vulnerabilities fixed in the latest [Go 1.20.3 packages](https://go.dev/doc/devel/release#go1.20.minor) and [Alpine 3.17.3](https://alpinelinux.org/posts/Alpine-3.17.3-released.html) and concluded that Forgejo is not affected.

  As a precaution the Forgejo v1.19.0-3 binaries were compiled with [Go 1.20.3 packages](https://go.dev/doc/devel/release#go1.20.minor) as published on 4 April 2023 and the container images were built with [Alpine 3.17.3](https://alpinelinux.org/posts/Alpine-3.17.3-released.html) as published on 29 March 2023.

* [Forgejo Semantic Version](https://forgejo.org/docs/v1.19/user/semver/)

  The semantic version was updated from `3.0.0+0-gitea-1.19.0` to `4.0.0+0-gitea-1.19.0` because of the breaking changes described below.

* Breaking changes

  They should not have a significant impact because they are related to experimental features (federation and CI).

  * [Use User.ID instead of User.Name in ActivityPub API for Person IRI](https://codeberg.org/forgejo/forgejo/commit/2fcd57d5ae5b5926e5b0b87e46f78ad4ac83cbbd)

    The ActivityPub id is an HTTPS URI that should remain constant, even if
the user changes their name.

  * [Actions unit is repo.actions instead of actions.actions](https://codeberg.org/forgejo/forgejo/commit/9596bd3712caec440859fce93d05e19cf95e5330)

    All instances of `actions.actions` in the `DISABLED_REPO_UNITS` or `DEFAULT_REPO_UNITS` configuration variables must be replaced with `repo.actions`.

* Bug fixes

  They are for the most part about user interface and actions. The most prominent ones are:

  * [Do not filter repositories by default on the explore page](https://codeberg.org/forgejo/forgejo/commit/d15f20b2d2ce613cc8b36536995f29f81797c002). The behavior of the explore page is back to what it was in Forgejo v1.18. Changing it was confusing.
  * [Skip LFS when disabled in dump and doctor](https://codeberg.org/forgejo/forgejo/commit/b6a2323981a7a89205a382ddf0542e205e292d3d).
  * [Do not display own email on the profile](https://codeberg.org/forgejo/forgejo/commit/1fed0e1adc8dd2d27d2d7e34dda29c8e79e5e6e8).
  * [Make minio package support legacy MD5 checksum](https://codeberg.org/forgejo/forgejo/commit/b73d1ac1eb7d5c985749dc721bbea7ebd14f9c83).
  * [Do not triggers Webhooks and actions on closed PR](https://codeberg.org/forgejo/forgejo/commit/a04535e212b04c0f6643a4f36904a3d1bf30c63f).

## 1.19.0-2

The [complete list of commits](https://codeberg.org/forgejo/forgejo/commits/branch/v1.19/forgejo) included in the `Forgejo v1.19.0-2` release can be reviewed from the command line with:

```shell
$ git clone https://codeberg.org/forgejo/forgejo/
$ git -C forgejo log --oneline --no-merges origin/v1.18/forgejo..origin/v1.19/forgejo
```

* Breaking changes
  * [Scoped access tokens](https://codeberg.org/forgejo/forgejo/commit/de484e86bc)

    Forgejo access token, used with the [API](https://forgejo.org/docs/v1.19/admin/api-usage/) can now have a "scope" that limits what it can access. Existing tokens stored in the database and created before Forgejo v1.19 had unlimited access. For backward compatibility, their access will remain the same and they will continue to work as before. However, **newly created token that do not specify a scope will now only have read-only access to public user profile and public repositories**.

    For instance, the `/users/{username}/tokens` API endpoint will require the `scopes: ['all', 'sudo']` parameter and the `forgejo admin user generate-access-token` will require the `--scopes all,sudo` argument obtain tokens with ulimited access as before for admin users.

    [Read more about the scoped tokens](https://forgejo.org/docs/v1.19/user/oauth2-provider/#scoped-tokens).

  * [Disable all units except code and pulls on forks](https://codeberg.org/forgejo/forgejo/commit/2741546be)

    When forking a repository, the fork will now have issues, projects, releases, packages and wiki disabled. These can be enabled in the repository settings afterwards. To change back to the previous default behavior, configure `DEFAULT_FORK_REPO_UNITS` to be the same value as `DEFAULT_REPO_UNITS`.

  * [Filter repositories by default on the explore page](https://codeberg.org/forgejo/forgejo/commit/4d20a4a1b)

    The explore page now always filters out repositories that are considered not relevant because they are either forks or have no topic and not description and no icon. A link is shown to display all repositories, unfiltered.

    <img src="./releases/images/forgejo-v1.19-relevant.png" alt="Explore repositories" width="600" />

  * [Remove deprecated DSA host key from Docker Container](https://codeberg.org/forgejo/forgejo/commit/f17edfaf5a31ea3f4e9152424b75c2c4986acbe3)
    Since OpenSSH 7.0 and greater similarly disable the ssh-dss (DSA) public key algorithm, and recommend against its use. http://www.openssh.com/legacy.html

  * Additional restrictions on valid user names

    The algorithm for validating user names was modified and some users may have invalid names. The command `forgejo doctor --run check-user-names` will list all of them so they can be renamed.

    If a Forgejo instance has users or organizations named `forgejo-actions` and `gitea-actions`, they will also need to be renamed before the upgrade. They are now reserved names for the experimental internal CI/CD named `Actions`.

  * [Semantic version](https://forgejo.org/docs/latest/user/semver)

    Since v1.18.5, in addition to the Forgejo release number, a [semantic version](https://semver.org/#semantic-versioning-200) number (e.g. `v3.0.0`) can be obtained from the `number` key of a new `/api/forgejo/v1/version` endpoint.

    Now, it reflects the Gitea version that Forgejo depends on, is no longer prefixed with `v` (e.g. `3.0.0+0-gitea-1.19.0`), and can be obtained from the `version` key of the same endpoint.
* Features

  * [Documentation](https://forgejo.org/docs/latest/)
    The first version of the [Forgejo documentation](https://forgejo.org/docs/latest/) is available and covers the administration of Forgejo, from installation to troubleshooting.

    [Read more about semantic versions](https://forgejo.codeberg.page/docs/v1.19/user/semver)

  * [Webhook authorization header](https://codeberg.org/forgejo/forgejo/commit/b6e81357bd6fb80f8ba94c513f89a210beb05313)
    Forgejo webhooks can be configured to send an [authorization header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization) to the target.

    [Read more about the webhook authorization header](https://forgejo.codeberg.page/docs/v1.19/user/webhooks/#authorization-header)

  * [Incoming emails](https://codeberg.org/forgejo/forgejo/commit/fc037b4b825f0501a1489e10d7c822435d825cb7)
    You can now set up Forgejo to receive incoming email. When enabled, it is now possible to reply to an email notification from Forgejo and:
    * Add a comment to an issue or a pull request
    * Unsubscribe to the notifications

    [Read more about incoming emails](https://forgejo.org/docs/v1.19/admin/incoming-email/)

  * Packages registries
    * Support for [Cargo](https://forgejo.org/docs/v1.19/admin/packages/cargo/), [Conda](https://forgejo.org/docs/v1.19/admin/packages/conda/) and [Chef](https://forgejo.org/docs/v1.19/admin/packages/chef/)
    * [Cleanup rules](https://codeberg.org/forgejo/forgejo/commit/32db62515)
    * [Quota limits](https://codeberg.org/forgejo/forgejo/commit/20674dd05)

  * [Option to prohibit fork if user reached maximum limit of repositories](https://codeberg.org/forgejo/forgejo/commit/7cc7db73b)
    It is possible for a user to create as many fork as they want, even when a quota on the number of repositories is imposed. The new `ALLOW_FORK_WITHOUT_MAXIMUM_LIMIT` setting can now be set to `false` so forks are prohibited if that means exceeding the quota.

    [Read more about repository configurations](https://forgejo.org/docs/v1.19/admin/config-cheat-sheet/#repository-repository)

  * [Scoped labels](https://codeberg.org/forgejo/forgejo/commit/6221a6fd5)
    Labels that contain a forward slash (**/**) separator are displayed with a slightly different color before and after the separator, as a visual aid. The first part of the label defines its "scope".

    [Read more about scoped labels](https://forgejo.org/docs/v1.19/user/labels/).

  * [Support org/user level projects](https://codeberg.org/forgejo/forgejo/commit/6fe3c8b39)
    It is now possible to create projects (kanban boards) for an organization or a user, in the same way it was possible for an individual repository.

  * [Map OIDC groups to Orgs/Teams](https://codeberg.org/forgejo/forgejo/commit/e8186f1c0)
    When a user logs in Forgejo using an provider such as [Keycloak](https://www.keycloak.org/), they can now automatically be part of a Forgejo team, depending on the OIDC group they belong to. For instance:

    ```json
    {"Developer": {"MyForgejoOrganization": ["MyForgejoTeam1", "MyForgejoTeam2"]}}
    ```

    Means that the user who is in the OIDC group `Developer` will automatically be a member of the `MyForgejoTeam1` and `MyForgejoTeam2` teams in the `MyForgejoOrganization` organization.
    This mapping is set when adding a new `Authentication Source` in the `Site Administration` panel.

    <img src="./releases/images/forgejo-v1.19-oidc-part1.png" alt="OIDC Group mapping part1" width="500" />

    ...

    <img src="./releases/images/forgejo-v1.19-oidc-part2.png" alt="OIDC Group mapping part2" width="500" />

    [Read more about OIDC groups mapping](https://forgejo.org/docs/v1.19/user/oauth2-provider/#endpoints)

  * [RSS feed for releases and tags](https://codeberg.org/forgejo/forgejo/commit/48d71b7d6)

    A RSS feed is now available for releases at `/{owner}/{repo}/releases.rss` and tags at `/{owner}/{repo}/tags.rss`.

  * [Supports wildcard protected branch](https://codeberg.org/forgejo/forgejo/commit/2782c1439)

    Instead of selecting a branch to be protected, the name of the branch must be specified and can be a pattern such as `precious*`.

    [Read more about branch protection](https://forgejo.org/docs/v1.19/user/protection/#protected-branches).

  * [Garbage collect LFS](https://codeberg.org/forgejo/forgejo/commit/651fe4bb7)
    Add a doctor command for full garbage collection of LFS: `forgejo doctor --run gc-lfs`.

  * Additions to the API

    * [Management for issue/pull and comment attachments](https://codeberg.org/forgejo/forgejo/commit/3c59d31bc)
    * [Get latest release](https://codeberg.org/forgejo/forgejo/commit/4d072a4c4)
    * [System hook](https://codeberg.org/forgejo/forgejo/commit/c0015979a)

  * [Option to disable releases on a repository](https://codeberg.org/forgejo/forgejo/commit/faa96553d)

    It is now possible to disable releases on a repository, in the same way it is possible to disable issues or packages.

  * [Git reflog support](https://codeberg.org/forgejo/forgejo/commit/757b4c17e)
    The [git reflog](https://git-scm.com/docs/git-reflog) are now active by default on all repositories and
    kept around for 90 days. It allows the Forgejo admin to recover the previous tip of a branch after an
    accidental force push.

    [Read more about reflog](https://forgejo.org/docs/v1.19/admin/config-cheat-sheet/#git---reflog-settings-gitreflog)

  * [Actions](https://codeberg.org/forgejo/forgejo/commit/4011821c946e8db032be86266dd9364ccb204118): an experimental CI/CD

    It appears for the first time in this Forgejo release but is not yet fit for production. It is not fully implemented and may be insecure. However, as long as it is not enabled, it presents no risk to existing Forgejo instances.

    If a repository has a file such as `.forgejo/workflows/test.yml`, it will be interpreted, for instance to run tests and verify the code in the repository works as expected (Continuous Integration). It can also be used to create HTML pages for a website and publish them (Continous Deployment). The syntax is similar to GitHub Actions and the jobs can be controled from the Forgejo web interface.

    [Read more about Forgejo Actions](https://forgejo.codeberg.page/2023-02-27-forgejo-actions/)

    <img src="./releases/images/forgejo-v1.19.0-0-rc0.png" alt="Actions" width="600" />

* User Interface improvements

  * [Review box on small screens](https://codeberg.org/forgejo/forgejo/commit/1fcf96ad0)
    The rendering of the review box is improved on small screens.

  * [Video element enabled in markdown](https://codeberg.org/forgejo/forgejo/commit/f8a40dafb)
    The `<video>` HTML tag can now be used in MarkDown, with the `src`, `autoplay`, and `controls` attributes.

  * [Copy citation file content in APA and BibTex format](https://codeberg.org/forgejo/forgejo/commit/9f8e77891)
    If a [BibTeX](https://fr.wikipedia.org/wiki/BibTeX) file named `CITATION.bib` is at the root of the repository, it can be conveniently copied and converted in APA by following the `Cite this repository` link.

    <img src="./releases/images/forgejo-v1.19-citation-link.png" alt="Citation link" width="500" />

    It will open a dialog box with the available formats and a preview of the content.

    <img src="./releases/images/forgejo-v1.19-citation-dialog.png" alt="Citation dialog" width="500" />

    The CFF format is also supported when a `CITATION.cff` file used instead.

  * [Display asciicast](https://codeberg.org/forgejo/forgejo/commit/d9f748a70)

    Files with the `.cast` extension are displayed in the Forgejo web interface as [asciicast v2](https://github.com/asciinema/asciinema/blob/develop/doc/asciicast-v2.md) using [asciinema-player](https://github.com/asciinema/asciinema-player).

  * [Attention blocks Note and Warning](https://codeberg.org/forgejo/forgejo/commit/cb8328853)

    For each quote block, the first `**Note**` or `**Warning**` gets an icon prepended to it and its text is colored accordingly.

    <img src="./releases/images/forgejo-v1.19-note-warning.png" alt="Attention block" width="400" />

  * [Support for commit cross references](https://codeberg.org/forgejo/forgejo/commit/d0d257b24)

    A commit hash can now be prefixed by the repository to be referenced from a comment in another repository: `owner/repo@commit`.

  * [Preview images for Issue cards in Project Board view](https://codeberg.org/forgejo/forgejo/commit/fb1a2a13f)

    If the card preview in the project is set to **Images and Text**, it displays images found in the corresponding issue. The most recent is displayed first, up to five images.

    [Read more about card preview images](https://forgejo.org/docs/v1.19/user/project/#card-previews-images).

  * [Add "Copy" button to file view of raw text](https://codeberg.org/forgejo/forgejo/commit/e3a7f1579)

    If a raw text file is displayed, a copy button of the text is enabled.

    **Before**

    <img src="./releases/images/forgejo-v1.19-raw-copy-before.png" alt="Raw copy before" width="500" />

    **After**

    <img src="./releases/images/forgejo-v1.19-raw-copy-after.png" alt="Raw copy after" width="500" />

  * [Setting to allow edits on PRs by maintainers](https://codeberg.org/forgejo/forgejo/commit/49919c636)

    Add setting to allow edits by maintainers by default, to avoid having to often ask contributors to enable this.

* Container images upgraded to Alpine 3.17

  The Forgejo container images are now based on [Alpine 3.17](https://alpinelinux.org/posts/Alpine-3.17.0-released.html) instead of Alpine 3.16. It includes an upgrade from git 2.36.5 to git 2.38.4 and from openssh 9.0p1 to openssh 9.1p1.

## 1.18.5-0

This stable release contains an **important security fix** for Forgejo to raise the protection against brute force attack on hashed passwords stored in the database to match industry standards, [as described in detail in a companion blog post](https://forgejo.org/2023-02-23-release-v1/).

### Recommended Action

We **strongly recommend** that all Forgejo installations are upgraded to the latest version as soon as possible.

If `PASSWORD_HASH_ALGO` is explicitly set in `app.ini`, comment it out so that the stronger algorithm is used instead.

All password hashes stored with another algorithm will be updated to the new algorithm on the next usage of this password (e.g. a user provides the password to the Forgejo server when they login). It does not require manual intervention.

### Forgejo

* SECURITY
  * Upgrade the default password hash algorithm to pbkdf2 with 320,000 iterations (https://codeberg.org/forgejo/forgejo/pulls/407)
* BUGFIXES
  * Return the Forgejo semantic version instead of "development" (https://codeberg.org/forgejo/forgejo/pulls/381)

### Gitea

* SECURITY
  * Provide the ability to set password hash algorithm parameters (https://github.com/go-gitea/gitea/pull/22942) (https://github.com/go-gitea/gitea/pull/22943)
* BUGFIXES
  * Use `--message=%s` for git commit message (https://github.com/go-gitea/gitea/pull/23028) (https://github.com/go-gitea/gitea/pull/23029)
  * Render access log template as text instead of HTML (https://github.com/go-gitea/gitea/pull/23013) (https://github.com/go-gitea/gitea/pull/23025)
  * Fix the Manually Merged form (https://github.com/go-gitea/gitea/pull/23015) (https://github.com/go-gitea/gitea/pull/23017)
  * Use beforeCommit instead of baseCommit (https://github.com/go-gitea/gitea/pull/22949) (https://github.com/go-gitea/gitea/pull/22996)
  * Display attachments of review comment when comment content is blank (https://github.com/go-gitea/gitea/pull/23035) (https://github.com/go-gitea/gitea/pull/23046)
  * Return empty url for submodule tree entries (https://github.com/go-gitea/gitea/pull/23043) (https://github.com/go-gitea/gitea/pull/23048)
  * Notify on container image create (https://github.com/go-gitea/gitea/pull/22806) (https://github.com/go-gitea/gitea/pull/22965)
  * Some refactor about code comments(https://github.com/go-gitea/gitea/pull/20821) (https://github.com/go-gitea/gitea/pull/22707)

Note that there is no Forgejo v1.18.4-N because Gitea v1.18.4 was replaced by Gitea v1.18.5 a few days after its release because of a regression. Forgejo was not affected.

## 1.18.3-2

This stable release includes a security fix for `git` and bug fixes.

### Git

Git [recently announced](https://github.blog/2023-02-14-git-security-vulnerabilities-announced-3/) new versions to address two CVEs ([CVE-2023-22490](https://cve.circl.lu/cve/CVE-2023-22490), [CVE-2023-23946](https://cve.circl.lu/cve/CVE-2023-23946)). On 14 Februrary 2023, Git published the maintenance release v2.39.2, together with releases for older maintenance tracks v2.38.4, v2.37.6, v2.36.5, v2.35.7, v2.34.7, v2.33.7, v2.32.6, v2.31.7, and v2.30.8. All major GNU/Linux distributions also provide updated packages via their security update channels.

We recommend that all installations running a version affected by the issues described below are upgraded to the latest version as soon as possible.

* When using a Forgejo binary: upgrade the `git` package to a version greater or equal to v2.39.2, v2.38.4, v2.37.6, v2.36.5, v2.35.7, v2.34.7, v2.33.7, v2.32.6, v2.31.7 or v2.30.8
* When using a Forgejo container image: `docker pull codeberg.org/forgejo/forgejo:1.18.3-2`

### Forgejo

* BUGFIXES
  * Use proxy for pull mirror (https://github.com/go-gitea/gitea/pull/22771) (https://github.com/go-gitea/gitea/pull/22772)
  * Revert "Fixes accessibility of empty repository commit status" (https://github.com/go-gitea/gitea/pull/22632)
    * A regression introduced in 1.18.3-1 prevented the CI status from displaying for commits with more than one pipeline
* FORGEJO RELEASE PROCESS BUGFIXES
  * The tag SHA in the uploaded repository must match (https://codeberg.org/forgejo/forgejo/pulls/345) [Read more about the consequences of this on the Forgejo blog](https://forgejo.org/2023-02-12-tags/)

### Gitea

* BUGFIXES
  * Load issue before accessing index in merge message (https://github.com/go-gitea/gitea/pull/22822) (https://github.com/go-gitea/gitea/pull/22830)
  * Fix isAllowed of escapeStreamer (https://github.com/go-gitea/gitea/pull/22814) (https://github.com/go-gitea/gitea/pull/22837)
  * Escape filename when assemble URL (https://github.com/go-gitea/gitea/pull/22850) (https://github.com/go-gitea/gitea/pull/22871)
  * Fix PR file tree folders no longer collapsing (https://github.com/go-gitea/gitea/pull/22864) (https://github.com/go-gitea/gitea/pull/22872)
  * Fix incorrect role labels for migrated issues and comments (https://github.com/go-gitea/gitea/pull/22914) (https://github.com/go-gitea/gitea/pull/22923)
  * Fix blame view missing lines (https://github.com/go-gitea/gitea/pull/22826) (https://github.com/go-gitea/gitea/pull/22929)
  * Fix 404 error viewing the LFS file (https://github.com/go-gitea/gitea/pull/22945) (https://github.com/go-gitea/gitea/pull/22948)
* FEATURES
  * Add command to bulk set must-change-password (https://github.com/go-gitea/gitea/pull/22823) (https://github.com/go-gitea/gitea/pull/22928)

## 1.18.3-1

This stable release includes bug fixes.

### Forgejo

* ACCESSIBILITY
  * Add ARIA support for Fomantic UI checkboxes (https://github.com/go-gitea/gitea/pull/22599)
  * Fixes accessibility behavior of Watching, Staring and Fork buttons (https://github.com/go-gitea/gitea/pull/22634)
  * Add main landmark to templates and adjust titles (https://github.com/go-gitea/gitea/pull/22670)
  * Improve checkbox accessibility a bit by adding the title attribute (https://github.com/go-gitea/gitea/pull/22593)
  * Improve accessibility of navigation bar and footer (https://github.com/go-gitea/gitea/pull/22635)
* PRIVACY
  * Use DNS queries to figure out the latest Forgejo version (https://codeberg.org/forgejo/forgejo/pulls/278)
* BRANDING
  * Change the values for the nodeinfo API to correctly identify the software as Forgejo (https://codeberg.org/forgejo/forgejo/pulls/313)
* CI
  * Use tagged test environment for stable branches (https://codeberg.org/forgejo/forgejo/pulls/318)

### Gitea

* BUGFIXES
  * Fix missing message in git hook when pull requests disabled on fork (https://github.com/go-gitea/gitea/pull/22625) (https://github.com/go-gitea/gitea/pull/22658)
  * add default user visibility to cli command "admin user create" (https://github.com/go-gitea/gitea/pull/22750) (https://github.com/go-gitea/gitea/pull/22760)
  * Fix color of tertiary button on dark theme (https://github.com/go-gitea/gitea/pull/22739) (https://github.com/go-gitea/gitea/pull/22744)
  * Fix restore repo bug, clarify the problem of ForeignIndex (https://github.com/go-gitea/gitea/pull/22776) (https://github.com/go-gitea/gitea/pull/22794)
  * Escape path for the file list (https://github.com/go-gitea/gitea/pull/22741) (https://github.com/go-gitea/gitea/pull/22757)
  * Fix bugs with WebAuthn preventing sign in and registration. (https://github.com/go-gitea/gitea/pull/22651) (https://github.com/go-gitea/gitea/pull/22721)
* PERFORMANCES
  * Improve checkIfPRContentChanged (https://github.com/go-gitea/gitea/pull/22611) (https://github.com/go-gitea/gitea/pull/22644)

## 1.18.3-0

This stable release includes bug fixes.

### Forgejo

* BUGFIXES
  * Fix line spacing for plaintext previews (https://github.com/go-gitea/gitea/pull/22699) (https://github.com/go-gitea/gitea/pull/22701)
  * Fix README TOC links (https://github.com/go-gitea/gitea/pull/22577) (https://github.com/go-gitea/gitea/pull/22677)
  * Don't return duplicated users who can create org repo (https://github.com/go-gitea/gitea/pull/22560) (https://github.com/go-gitea/gitea/pull/22562)
  * Link issue and pull requests status change in UI notifications directly to their event in the timelined view. (https://github.com/go-gitea/gitea/pull/22627) (https://github.com/go-gitea/gitea/pull/22642)

### Gitea

* BUGFIXES
  * Add missing close bracket in imagediff (https://github.com/go-gitea/gitea/pull/22710) (https://github.com/go-gitea/gitea/pull/22712)
  * Fix wrong hint when deleting a branch successfully from pull request UI (https://github.com/go-gitea/gitea/pull/22673) (https://github.com/go-gitea/gitea/pull/22698)
  * Fix missing message in git hook when pull requests disabled on fork (https://github.com/go-gitea/gitea/pull/22625) (https://github.com/go-gitea/gitea/pull/22658)

## 1.18.2-1

This stable release includes a security fix. It was possible to reveal a user's email address, which is problematic because users can choose to hide their email address from everyone. This was possible because the notification email for a repository transfer request to an organization included every user's email address in the owner team. This has been fixed by sending individual emails instead and the code was refactored to prevent it from happening again.

We **strongly recommend** that all installations are upgraded to the latest version as soon as possible.

### Gitea

* BUGFIXES
  * When updating by rebase we need to set the environment for head repo (https://github.com/go-gitea/gitea/pull/22535) (https://github.com/go-gitea/gitea/pull/22536)
  * Mute all links in issue timeline (https://github.com/go-gitea/gitea/pull/22534)
  * Truncate commit summary on repo files table. (https://github.com/go-gitea/gitea/pull/22551) (https://github.com/go-gitea/gitea/pull/22552)
  * Prevent multiple `To` recipients (https://github.com/go-gitea/gitea/pull/22566) (https://github.com/go-gitea/gitea/pull/22569)

## 1.18.2-0

This stable release includes bug fixes.

### Gitea

* BUGFIXES
  * Fix issue not auto-closing when it includes a reference to a branch (https://github.com/go-gitea/gitea/pull/22514) (https://github.com/go-gitea/gitea/pull/22521)
  * Fix invalid issue branch reference if not specified in template (https://github.com/go-gitea/gitea/pull/22513) (https://github.com/go-gitea/gitea/pull/22520)
  * Fix 500 error viewing pull request when fork has pull requests disabled (https://github.com/go-gitea/gitea/pull/22512) (https://github.com/go-gitea/gitea/pull/22515)
  * Reliable selection of admin user (https://github.com/go-gitea/gitea/pull/22509) (https://github.com/go-gitea/gitea/pull/22511)

## 1.18.1-0

This is the first Forgejo stable point release.

### Forgejo

### Critical security update for Git

Git [recently announced](https://github.blog/2023-01-17-git-security-vulnerabilities-announced-2/) new versions to address two CVEs ([CVE-2022-23521](https://cve.circl.lu/cve/CVE-2022-23521), [CVE-2022-41903](https://cve.circl.lu/cve/CVE-2022-41903)). On 17 January 2023, Git published the maintenance release v2.39.1, together with releases for older maintenance tracks v2.38.3, v2.37.5, v2.36.4, v2.35.6, v2.34.6, v2.33.6, v2.32.5, v2.31.6, and v2.30.7. All major GNU/Linux distributions also provide updated packages via their security update channels.

We **strongly recommend** that all installations running a version affected by the issues described below are upgraded to the latest version as soon as possible.

* When using a Forgejo binary: upgrade the `git` package to a version greater or equal to v2.39.1, v2.38.3, v2.37.5, v2.36.4, v2.35.6, v2.34.6, v2.33.6, v2.32.5, v2.31.6, or v2.30.7
* When using a Forgejo container image: `docker pull codeberg.org/forgejo/forgejo:1.18.1-0`

Read more in the [Forgejo blog](https://forgejo.org/2023-01-18-release-v1-18-1-0/).

#### Release process stability

The [release process](https://codeberg.org/forgejo/forgejo/src/branch/v1.18/forgejo-ci) based on [Woodpecker CI](https://woodpecker-ci.org/) was entirely reworked to be more resilient to transient errors. A new release is first uploaded into the new [Forgejo experimental](https://codeberg.org/forgejo-experimental/) organization for testing purposes.

Automated end to end testing of releases was implemented with a full development cycle including the creation of a new repository and a run of CI. It relieves the user and developer from the burden of tedious manual testing.

#### Container environment variables

When running a container, all environment variables starting with `FORGEJO__` can be used instead of `GITEA__`. For backward compatibility with existing scripts, it is still possible to use `GITEA__` instead of `FORGEJO__`. For instance:

```
docker run --name forgejo -e FORGEJO__security__INSTALL_LOCK=true codeberg.org/forgejo/forgejo:1.18.1-0
```

#### Forgejo hook types

A new `forgejo` hook type is available and behaves exactly the same as the existing `gitea` hook type. It will be used to implement additional features specific to Forgejo in a way that will be backward compatible with Gitea.

#### X-Forgejo headers

Wherever a `X-Gitea` header is received or sent, an identical `X-Forgejo` is added. For instance when a notification mail is sent, the `X-Forgejo-Reason` header is set to explain why. Or when a webhook is sent, the `X-Forgejo-Event` header is set with `push`, `tag`, etc. for Woodpecker CI to decide on an action.

#### Look and feel fixes

The Forgejo theme was [modified](https://codeberg.org/forgejo/forgejo/src/branch/v1.18/forgejo-branding) to take into account user feedback.

### Gitea

* API
  * Add `sync_on_commit` option for push mirrors api (https://github.com/go-gitea/gitea/pull/22271) (https://github.com/go-gitea/gitea/pull/22292)
* BUGFIXES
  * Update `github.com/zeripath/zapx/v15` (https://github.com/go-gitea/gitea/pull/22485)
  * Fix pull request API field `closed_at` always being `null` (https://github.com/go-gitea/gitea/pull/22482) (https://github.com/go-gitea/gitea/pull/22483)
  * Fix container blob mount (https://github.com/go-gitea/gitea/pull/22226) (https://github.com/go-gitea/gitea/pull/22476)
  * Fix error when calculating repository size (https://github.com/go-gitea/gitea/pull/22392) (https://github.com/go-gitea/gitea/pull/22474)
  * Fix Operator does not exist bug on explore page with ONLY_SHOW_RELEVANT_REPOS (https://github.com/go-gitea/gitea/pull/22454) (https://github.com/go-gitea/gitea/pull/22472)
  * Fix environments for KaTeX and error reporting (https://github.com/go-gitea/gitea/pull/22453) (https://github.com/go-gitea/gitea/pull/22473)
  * Remove the netgo tag for Windows build (https://github.com/go-gitea/gitea/pull/22467) (https://github.com/go-gitea/gitea/pull/22468)
  * Fix migration from GitBucket (https://github.com/go-gitea/gitea/pull/22477) (https://github.com/go-gitea/gitea/pull/22465)
  * Prevent panic on looking at api "git" endpoints for empty repos (https://github.com/go-gitea/gitea/pull/22457) (https://github.com/go-gitea/gitea/pull/22458)
  * Fix PR status layout on mobile (https://github.com/go-gitea/gitea/pull/21547) (https://github.com/go-gitea/gitea/pull/22441)
  * Fix wechatwork webhook sends empty content in PR review (https://github.com/go-gitea/gitea/pull/21762) (https://github.com/go-gitea/gitea/pull/22440)
  * Remove duplicate "Actions" label in mobile view (https://github.com/go-gitea/gitea/pull/21974) (https://github.com/go-gitea/gitea/pull/22439)
  * Fix leaving organization bug on user settings -> orgs (https://github.com/go-gitea/gitea/pull/21983) (https://github.com/go-gitea/gitea/pull/22438)
  * Fixed colour transparency regex matching in project board sorting (https://github.com/go-gitea/gitea/pull/22092) (https://github.com/go-gitea/gitea/pull/22437)
  * Correctly handle select on multiple channels in Queues (https://github.com/go-gitea/gitea/pull/22146) (https://github.com/go-gitea/gitea/pull/22428)
  * Prepend refs/heads/ to issue template refs (https://github.com/go-gitea/gitea/pull/20461) (https://github.com/go-gitea/gitea/pull/22427)
  * Restore function to "Show more" buttons (https://github.com/go-gitea/gitea/pull/22399) (https://github.com/go-gitea/gitea/pull/22426)
  * Continue GCing other repos on error in one repo (https://github.com/go-gitea/gitea/pull/22422) (https://github.com/go-gitea/gitea/pull/22425)
  * Allow HOST has no port (https://github.com/go-gitea/gitea/pull/22280) (https://github.com/go-gitea/gitea/pull/22409)
  * Fix omit avatar_url in discord payload when empty (https://github.com/go-gitea/gitea/pull/22393) (https://github.com/go-gitea/gitea/pull/22394)
  * Don't display stop watch top bar icon when disabled and hidden when click other place (https://github.com/go-gitea/gitea/pull/22374) (https://github.com/go-gitea/gitea/pull/22387)
  * Don't lookup mail server when using sendmail (https://github.com/go-gitea/gitea/pull/22300) (https://github.com/go-gitea/gitea/pull/22383)
  * Fix gravatar disable bug (https://github.com/go-gitea/gitea/pull/22337)
  * Fix update settings table on install (https://github.com/go-gitea/gitea/pull/22326) (https://github.com/go-gitea/gitea/pull/22327)
  * Fix sitemap (https://github.com/go-gitea/gitea/pull/22272) (https://github.com/go-gitea/gitea/pull/22320)
  * Fix code search title translation (https://github.com/go-gitea/gitea/pull/22285) (https://github.com/go-gitea/gitea/pull/22316)
  * Fix due date rendering the wrong date in issue (https://github.com/go-gitea/gitea/pull/22302) (https://github.com/go-gitea/gitea/pull/22306)
  * Fix get system setting bug when enabled redis cache (https://github.com/go-gitea/gitea/pull/22298)
  * Fix bug of DisableGravatar default value (https://github.com/go-gitea/gitea/pull/22297)
  * Fix key signature error page (https://github.com/go-gitea/gitea/pull/22229) (https://github.com/go-gitea/gitea/pull/22230)
* TESTING
  * Remove test session cache to reduce possible concurrent problem (https://github.com/go-gitea/gitea/pull/22199) (https://github.com/go-gitea/gitea/pull/22429)
* MISC
  * Restore previous official review when an official review is deleted (https://github.com/go-gitea/gitea/pull/22449) (https://github.com/go-gitea/gitea/pull/22460)
  * Log STDERR of external renderer when it fails (https://github.com/go-gitea/gitea/pull/22442) (https://github.com/go-gitea/gitea/pull/22444)

## 1.18.0-1

This is the first Forgejo release.

### Forgejo improvements

#### Woodpecker CI

A new [CI configuration](https://codeberg.org/forgejo/forgejo/src/branch/v1.18/forgejo-ci) based on [Woodpecker CI](https://woodpecker-ci.org/) was created. It is used to:

* run tests on every Forgejo pull request ([compliance](https://codeberg.org/forgejo/forgejo/src/tag/v1.18.0-1/.woodpecker/compliance.yml), [unit tests and integration tests](https://codeberg.org/forgejo/forgejo/src/tag/v1.18.0-1/.woodpecker/testing-amd64.yml))
* publish the Forgejo v1.18.0-1 release, [as binary packages](https://codeberg.org/forgejo/forgejo/releases/tag/v1.18.0-1) for amd64, arm64 and armv6 and [container images](https://codeberg.org/forgejo/-/packages/container/forgejo/1.18.0-1) for amd64 and arm64, root and rootless

#### Look and feel

The default themes were replaced by Forgejo themes and the landing page was [modified](https://codeberg.org/forgejo/forgejo/src/branch/v1.18/forgejo-branding) to display the Forgejo logo and names but the look and feel remains otherwise identical to Gitea.

<img src="./releases/images/forgejo-v1.18.0-rc1-2-landing.jpg" alt="Landing page" width="600" />

#### Privacy

Gitea instances fetch https://dl.gitea.io/gitea/version.json weekly by default, which raises privacy concerns. In Forgejo [this feature needs to be explicitly activated](https://codeberg.org/forgejo/forgejo/src/branch/v1.18/forgejo-privacy) at installation time or by modifying the configuration file. Forgejo also provides an alternative [RSS feed](https://forgejo.org/releases/) to be informed when a new release is published.

### Gitea

* SECURITY
  * Remove ReverseProxy authentication from the API (https://github.com/go-gitea/gitea/pull/22219) (https://github.com/go-gitea/gitea/pull/22251)
  * Support Go Vulnerability Management (https://github.com/go-gitea/gitea/pull/21139)
  * Forbid HTML string tooltips (https://github.com/go-gitea/gitea/pull/20935)
* BREAKING
  * Rework mailer settings (https://github.com/go-gitea/gitea/pull/18982)
  * Remove U2F support (https://github.com/go-gitea/gitea/pull/20141)
  * Refactor `i18n` to `locale` (https://github.com/go-gitea/gitea/pull/20153)
  * Enable contenthash in filename for dynamic assets (https://github.com/go-gitea/gitea/pull/20813)
* FEATURES
  * Add color previews in markdown (https://github.com/go-gitea/gitea/pull/21474)
  * Allow package version sorting (https://github.com/go-gitea/gitea/pull/21453)
  * Add support for Chocolatey/NuGet v2 API (https://github.com/go-gitea/gitea/pull/21393)
  * Add API endpoint to get changed files of a PR (https://github.com/go-gitea/gitea/pull/21177)
  * Add filetree on left of diff view (https://github.com/go-gitea/gitea/pull/21012)
  * Support Issue forms and PR forms (https://github.com/go-gitea/gitea/pull/20987)
  * Add support for Vagrant packages (https://github.com/go-gitea/gitea/pull/20930)
  * Add support for `npm unpublish` (https://github.com/go-gitea/gitea/pull/20688)
  * Add badge capabilities to users (https://github.com/go-gitea/gitea/pull/20607)
  * Add issue filter for Author (https://github.com/go-gitea/gitea/pull/20578)
  * Add KaTeX rendering to Markdown. (https://github.com/go-gitea/gitea/pull/20571)
  * Add support for Pub packages (https://github.com/go-gitea/gitea/pull/20560)
  * Support localized README (https://github.com/go-gitea/gitea/pull/20508)
  * Add support mCaptcha as captcha provider (https://github.com/go-gitea/gitea/pull/20458)
  * Add team member invite by email (https://github.com/go-gitea/gitea/pull/20307)
  * Added email notification option to receive all own messages (https://github.com/go-gitea/gitea/pull/20179)
  * Switch Unicode Escaping to a VSCode-like system (https://github.com/go-gitea/gitea/pull/19990)
  * Add user/organization code search (https://github.com/go-gitea/gitea/pull/19977)
  * Only show relevant repositories on explore page (https://github.com/go-gitea/gitea/pull/19361)
  * User keypairs and HTTP signatures for ActivityPub federation using go-ap (https://github.com/go-gitea/gitea/pull/19133)
  * Add sitemap support (https://github.com/go-gitea/gitea/pull/18407)
  * Allow creation of OAuth2 applications for orgs (https://github.com/go-gitea/gitea/pull/18084)
  * Add system setting table with cache and also add cache supports for user setting (https://github.com/go-gitea/gitea/pull/18058)
  * Add pages to view watched repos and subscribed issues/PRs (https://github.com/go-gitea/gitea/pull/17156)
  * Support Proxy protocol (https://github.com/go-gitea/gitea/pull/12527)
  * Implement sync push mirror on commit (https://github.com/go-gitea/gitea/pull/19411)
* API
  * Allow empty assignees on pull request edit (https://github.com/go-gitea/gitea/pull/22150) (https://github.com/go-gitea/gitea/pull/22214)
  * Make external issue tracker regexp configurable via API (https://github.com/go-gitea/gitea/pull/21338)
  * Add name field for org api (https://github.com/go-gitea/gitea/pull/21270)
  * Show teams with no members if user is admin (https://github.com/go-gitea/gitea/pull/21204)
  * Add latest commit's SHA to content response (https://github.com/go-gitea/gitea/pull/20398)
  * Add allow_rebase_update, default_delete_branch_after_merge to repository api response (https://github.com/go-gitea/gitea/pull/20079)
  * Add new endpoints for push mirrors management (https://github.com/go-gitea/gitea/pull/19841)
* ENHANCEMENTS
  * Add setting to disable the git apply step in test patch (https://github.com/go-gitea/gitea/pull/22130) (https://github.com/go-gitea/gitea/pull/22170)
  * Multiple improvements for comment edit diff (https://github.com/go-gitea/gitea/pull/21990) (https://github.com/go-gitea/gitea/pull/22007)
  * Fix button in branch list, avoid unexpected page jump before restore branch actually done (https://github.com/go-gitea/gitea/pull/21562) (https://github.com/go-gitea/gitea/pull/21928)
  * Fix flex layout for repo list icons (https://github.com/go-gitea/gitea/pull/21896) (https://github.com/go-gitea/gitea/pull/21920)
  * Fix vertical align of committer avatar rendered by email address (https://github.com/go-gitea/gitea/pull/21884) (https://github.com/go-gitea/gitea/pull/21918)
  * Fix setting HTTP headers after write (https://github.com/go-gitea/gitea/pull/21833) (https://github.com/go-gitea/gitea/pull/21877)
  * Color and Style enhancements (https://github.com/go-gitea/gitea/pull/21784, #21799) (https://github.com/go-gitea/gitea/pull/21868)
  * Ignore line anchor links with leading zeroes (https://github.com/go-gitea/gitea/pull/21728) (https://github.com/go-gitea/gitea/pull/21776)
  * Quick fixes monaco-editor error: "vs.editor.nullLanguage" (https://github.com/go-gitea/gitea/pull/21734) (https://github.com/go-gitea/gitea/pull/21738)
  * Use CSS color-scheme instead of invert (https://github.com/go-gitea/gitea/pull/21616) (https://github.com/go-gitea/gitea/pull/21623)
  * Respect user's locale when rendering the date range in the repo activity page (https://github.com/go-gitea/gitea/pull/21410)
  * Change `commits-table` column width (https://github.com/go-gitea/gitea/pull/21564)
  * Refactor git command arguments and make all arguments to be safe to be used (https://github.com/go-gitea/gitea/pull/21535)
  * CSS color enhancements (https://github.com/go-gitea/gitea/pull/21534)
  * Add link to user profile in markdown mention only if user exists (https://github.com/go-gitea/gitea/pull/21533, #21554)
  * Add option to skip index dirs (https://github.com/go-gitea/gitea/pull/21501)
  * Diff file tree tweaks (https://github.com/go-gitea/gitea/pull/21446)
  * Localize all timestamps (https://github.com/go-gitea/gitea/pull/21440)
  * Add `code` highlighting in issue titles (https://github.com/go-gitea/gitea/pull/21432)
  * Use Name instead of DisplayName in LFS Lock (https://github.com/go-gitea/gitea/pull/21415)
  * Consolidate more CSS colors into variables (https://github.com/go-gitea/gitea/pull/21402)
  * Redirect to new repository owner (https://github.com/go-gitea/gitea/pull/21398)
  * Use ISO date format instead of hard-coded English date format for date range in repo activity page (https://github.com/go-gitea/gitea/pull/21396)
  * Use weighted algorithm for string matching when finding files in repo (https://github.com/go-gitea/gitea/pull/21370)
  * Show private data in feeds (https://github.com/go-gitea/gitea/pull/21369)
  * Refactor parseTreeEntries, speed up tree list (https://github.com/go-gitea/gitea/pull/21368)
  * Add GET and DELETE endpoints for Docker blob uploads (https://github.com/go-gitea/gitea/pull/21367)
  * Add nicer error handling on template compile errors (https://github.com/go-gitea/gitea/pull/21350)
  * Add `stat` to `ToCommit` function for speed (https://github.com/go-gitea/gitea/pull/21337)
  * Support instance-wide OAuth2 applications (https://github.com/go-gitea/gitea/pull/21335)
  * Record OAuth client type at registration (https://github.com/go-gitea/gitea/pull/21316)
  * Add new CSS variables --color-accent and --color-small-accent (https://github.com/go-gitea/gitea/pull/21305)
  * Improve error descriptions for unauthorized_client (https://github.com/go-gitea/gitea/pull/21292)
  * Case-insensitive "find files in repo" (https://github.com/go-gitea/gitea/pull/21269)
  * Consolidate more CSS rules, fix inline code on arc-green (https://github.com/go-gitea/gitea/pull/21260)
  * Log real ip of requests from ssh (https://github.com/go-gitea/gitea/pull/21216)
  * Save files in local storage as group readable (https://github.com/go-gitea/gitea/pull/21198)
  * Enable fluid page layout on medium size viewports (https://github.com/go-gitea/gitea/pull/21178)
  * File header tweaks (https://github.com/go-gitea/gitea/pull/21175)
  * Added missing headers on user packages page (https://github.com/go-gitea/gitea/pull/21172)
  * Display image digest for container packages (https://github.com/go-gitea/gitea/pull/21170)
  * Skip dirty check for team forms (https://github.com/go-gitea/gitea/pull/21154)
  * Keep path when creating a new branch (https://github.com/go-gitea/gitea/pull/21153)
  * Remove fomantic image module (https://github.com/go-gitea/gitea/pull/21145)
  * Make labels clickable in the comments section. (https://github.com/go-gitea/gitea/pull/21137)
  * Sort branches and tags by date descending (https://github.com/go-gitea/gitea/pull/21136)
  * Better repo API unit checks (https://github.com/go-gitea/gitea/pull/21130)
  * Improve commit status icons (https://github.com/go-gitea/gitea/pull/21124)
  * Limit length of repo description and repo url input fields (https://github.com/go-gitea/gitea/pull/21119)
  * Show .editorconfig errors in frontend (https://github.com/go-gitea/gitea/pull/21088)
  * Allow poster to choose reviewers (https://github.com/go-gitea/gitea/pull/21084)
  * Remove black labels and CSS cleanup (https://github.com/go-gitea/gitea/pull/21003)
  * Make e-mail sanity check more precise (https://github.com/go-gitea/gitea/pull/20991)
  * Use native inputs in whitespace dropdown (https://github.com/go-gitea/gitea/pull/20980)
  * Enhance package date display (https://github.com/go-gitea/gitea/pull/20928)
  * Display total blob size of a package version (https://github.com/go-gitea/gitea/pull/20927)
  * Show language name on hover (https://github.com/go-gitea/gitea/pull/20923)
  * Show instructions for all generic package files (https://github.com/go-gitea/gitea/pull/20917)
  * Refactor AssertExistsAndLoadBean to use generics (https://github.com/go-gitea/gitea/pull/20797)
  * Move the official website link at the footer of gitea (https://github.com/go-gitea/gitea/pull/20777)
  * Add support for full name in reverse proxy auth (https://github.com/go-gitea/gitea/pull/20776)
  * Remove useless JS operation for relative time tooltips (https://github.com/go-gitea/gitea/pull/20756)
  * Replace some icons with SVG (https://github.com/go-gitea/gitea/pull/20741)
  * Change commit status icons to SVG (https://github.com/go-gitea/gitea/pull/20736)
  * Improve single repo action for issue and pull requests (https://github.com/go-gitea/gitea/pull/20730)
  * Allow multiple files in generic packages (https://github.com/go-gitea/gitea/pull/20661)
  * Add option to create new issue from /issues page (https://github.com/go-gitea/gitea/pull/20650)
  * Background color of private list-items updated (https://github.com/go-gitea/gitea/pull/20630)
  * Added search input field to issue filter (https://github.com/go-gitea/gitea/pull/20623)
  * Increase default item listing size `ISSUE_PAGING_NUM` to 20 (https://github.com/go-gitea/gitea/pull/20547)
  * Modify milestone search keywords to be case insensitive again (https://github.com/go-gitea/gitea/pull/20513)
  * Show hint to link package to repo when viewing empty repo package list (https://github.com/go-gitea/gitea/pull/20504)
  * Add Tar ZSTD support (https://github.com/go-gitea/gitea/pull/20493)
  * Make code review checkboxes clickable (https://github.com/go-gitea/gitea/pull/20481)
  * Add "X-Gitea-Object-Type" header for GET `/raw/` & `/media/` API (https://github.com/go-gitea/gitea/pull/20438)
  * Display project in issue list (https://github.com/go-gitea/gitea/pull/20434)
  * Prepend commit message to template content when opening a new PR (https://github.com/go-gitea/gitea/pull/20429)
  * Replace fomantic popup module with tippy.js (https://github.com/go-gitea/gitea/pull/20428)
  * Allow to specify colors for text in markup (https://github.com/go-gitea/gitea/pull/20363)
  * Allow access to the Public Organization Member lists with minimal permissions (https://github.com/go-gitea/gitea/pull/20330)
  * Use default values when provided values are empty (https://github.com/go-gitea/gitea/pull/20318)
  * Vertical align navbar avatar at middle (https://github.com/go-gitea/gitea/pull/20302)
  * Delete cancel button in repo creation page (https://github.com/go-gitea/gitea/pull/21381)
  * Include login_name in adminCreateUser response (https://github.com/go-gitea/gitea/pull/20283)
  * fix: icon margin in user/settings/repos (https://github.com/go-gitea/gitea/pull/20281)
  * Remove blue text on migrate page (https://github.com/go-gitea/gitea/pull/20273)
  * Modify milestone search keywords to be case insensitive (https://github.com/go-gitea/gitea/pull/20266)
  * Move some files into models' sub packages (https://github.com/go-gitea/gitea/pull/20262)
  * Add tooltip to repo icons in explore page (https://github.com/go-gitea/gitea/pull/20241)
  * Remove deprecated licenses (https://github.com/go-gitea/gitea/pull/20222)
  * Webhook for Wiki changes (https://github.com/go-gitea/gitea/pull/20219)
  * Share HTML template renderers and create a watcher framework (https://github.com/go-gitea/gitea/pull/20218)
  * Allow enable LDAP source and disable user sync via CLI (https://github.com/go-gitea/gitea/pull/20206)
  * Adds a checkbox to select all issues/PRs (https://github.com/go-gitea/gitea/pull/20177)
  * Refactor `i18n` to `locale` (https://github.com/go-gitea/gitea/pull/20153)
  * Disable status checks in template if none found (https://github.com/go-gitea/gitea/pull/20088)
  * Allow manager logging to set SQL (https://github.com/go-gitea/gitea/pull/20064)
  * Add order by for assignee no sort issue (https://github.com/go-gitea/gitea/pull/20053)
  * Take a stab at porting existing components to Vue3 (https://github.com/go-gitea/gitea/pull/20044)
  * Add doctor command to write commit-graphs (https://github.com/go-gitea/gitea/pull/20007)
  * Add support for authentication based on reverse proxy email (https://github.com/go-gitea/gitea/pull/19949)
  * Enable spellcheck for EasyMDE, use contenteditable mode (https://github.com/go-gitea/gitea/pull/19776)
  * Allow specifying SECRET_KEY_URI, similar to INTERNAL_TOKEN_URI (https://github.com/go-gitea/gitea/pull/19663)
  * Rework mailer settings (https://github.com/go-gitea/gitea/pull/18982)
  * Add option to purge users (https://github.com/go-gitea/gitea/pull/18064)
  * Add author search input (https://github.com/go-gitea/gitea/pull/21246)
  * Make rss/atom identifier globally unique (https://github.com/go-gitea/gitea/pull/21550)
* BUGFIXES
  * Auth interface return error when verify failure (https://github.com/go-gitea/gitea/pull/22119) (https://github.com/go-gitea/gitea/pull/22259)
  * Use complete SHA to create and query commit status (https://github.com/go-gitea/gitea/pull/22244) (https://github.com/go-gitea/gitea/pull/22257)
  * Update bleve and zapx to fix unaligned atomic (https://github.com/go-gitea/gitea/pull/22031) (https://github.com/go-gitea/gitea/pull/22218)
  * Prevent panic in doctor command when running default checks (https://github.com/go-gitea/gitea/pull/21791) (https://github.com/go-gitea/gitea/pull/21807)
  * Load GitRepo in API before deleting issue (https://github.com/go-gitea/gitea/pull/21720) (https://github.com/go-gitea/gitea/pull/21796)
  * Ignore line anchor links with leading zeroes (https://github.com/go-gitea/gitea/pull/21728) (https://github.com/go-gitea/gitea/pull/21776)
  * Set last login when activating account (https://github.com/go-gitea/gitea/pull/21731) (https://github.com/go-gitea/gitea/pull/21755)
  * Fix UI language switching bug (https://github.com/go-gitea/gitea/pull/21597) (https://github.com/go-gitea/gitea/pull/21749)
  * Quick fixes monaco-editor error: "vs.editor.nullLanguage" (https://github.com/go-gitea/gitea/pull/21734) (https://github.com/go-gitea/gitea/pull/21738)
  * Allow local package identifiers for PyPI packages (https://github.com/go-gitea/gitea/pull/21690) (https://github.com/go-gitea/gitea/pull/21727)
  * Deal with markdown template without metadata (https://github.com/go-gitea/gitea/pull/21639) (https://github.com/go-gitea/gitea/pull/21654)
  * Fix opaque background on mermaid diagrams (https://github.com/go-gitea/gitea/pull/21642) (https://github.com/go-gitea/gitea/pull/21652)
  * Fix repository adoption on Windows (https://github.com/go-gitea/gitea/pull/21646) (https://github.com/go-gitea/gitea/pull/21650)
  * Sync git hooks when config file path changed (https://github.com/go-gitea/gitea/pull/21619) (https://github.com/go-gitea/gitea/pull/21626)
  * Fix 500 on PR files API (https://github.com/go-gitea/gitea/pull/21602) (https://github.com/go-gitea/gitea/pull/21607)
  * Fix `Timestamp.IsZero` (https://github.com/go-gitea/gitea/pull/21593) (https://github.com/go-gitea/gitea/pull/21603)
  * Fix viewing user subscriptions (https://github.com/go-gitea/gitea/pull/21482)
  * Fix mermaid-related bugs (https://github.com/go-gitea/gitea/pull/21431)
  * Fix branch dropdown shifting on page load (https://github.com/go-gitea/gitea/pull/21428)
  * Fix default theme-auto selector when nologin (https://github.com/go-gitea/gitea/pull/21346)
  * Fix and improve incorrect error messages (https://github.com/go-gitea/gitea/pull/21342)
  * Fix formatted link for PR review notifications to matrix (https://github.com/go-gitea/gitea/pull/21319)
  * Center-aligning content of WebAuthN page (https://github.com/go-gitea/gitea/pull/21127)
  * Remove follow from commits by file (https://github.com/go-gitea/gitea/pull/20765)
  * Fix commit status popup (https://github.com/go-gitea/gitea/pull/20737)
  * Fix init mail render logic (https://github.com/go-gitea/gitea/pull/20704)
  * Use correct page size for link header pagination (https://github.com/go-gitea/gitea/pull/20546)
  * Preserve unix socket file (https://github.com/go-gitea/gitea/pull/20499)
  * Use tippy.js for context popup (https://github.com/go-gitea/gitea/pull/20393)
  * Add missing parameter for error in log message (https://github.com/go-gitea/gitea/pull/20144)
  * Do not allow organisation owners add themselves as collaborator (https://github.com/go-gitea/gitea/pull/20043)
  * Rework file highlight rendering and fix yaml copy-paste (https://github.com/go-gitea/gitea/pull/19967)
  * Improve code diff highlight, fix incorrect rendered diff result (https://github.com/go-gitea/gitea/pull/19958)
* TESTING
  * Improve OAuth integration tests (https://github.com/go-gitea/gitea/pull/21390)
  * Add playwright tests (https://github.com/go-gitea/gitea/pull/20123)
* BUILD
  * Switch to building with go1.19 (https://github.com/go-gitea/gitea/pull/20695)
  * Update JS dependencies, adjust eslint (https://github.com/go-gitea/gitea/pull/20659)
  * Add more linters to improve code readability (https://github.com/go-gitea/gitea/pull/19989)

## 1.18.0-0

This release was replaced by 1.18.0-1 a few hours after being published because the release process [was interrupted](https://codeberg.org/forgejo/forgejo/issues/180).

## 1.18.0-rc1-2

This is the first Forgejo release candidate.
