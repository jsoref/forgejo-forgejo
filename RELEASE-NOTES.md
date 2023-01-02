# Release Notes

A Forgejo release is published shortly after a Gitea release is published and they have [matching release numbers](https://codeberg.org/forgejo/forgejo/src/branch/forgejo/CONTRIBUTING/RELEASE.md#release-numbering). Additional Forgejo releases may be published to address urgent security issues or bug fixes. Forgejo release notes include all Gitea release notes.

The Forgejo admin should carefully read the required manual actions before upgrading. A point release (e.g. v1.19.1 or v1.19.2) does not require manual actions but others might (e.g. v1.18.0, v1.19.0).

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
