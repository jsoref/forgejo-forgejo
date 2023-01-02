# Release Notes

A Forgejo release is published shortly after a Gitea release is published and they have [matching release numbers](https://codeberg.org/forgejo/forgejo/src/branch/forgejo/CONTRIBUTING/RELEASE.md#release-numbering). Additional Forgejo releases may be published to address urgent security issues or bug fixes. Forgejo release notes include all Gitea release notes.

The Forgejo admin should carefully read the required manual actions before upgrading. A point release (e.g. v1.18.1 or v1.18.2) does not require manual actions but others might (e.g. v1.18.0, v1.19.0).

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
