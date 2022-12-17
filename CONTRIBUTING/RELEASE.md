# Release management

## Release numbering

The Forgejo release numbers are composed of the Gitea release number followed by a dash and a serial number. For instance:

* Gitea **v1.18.0** will be Forgejo **v1.18.0-0**, **v1.18.0-1**, etc

The Gitea release candidates are suffixed with **-rcN** which is handled as a special case for packaging: although **X.Y.Z** is lexicographically lower than **X.Y.Z-rc1** is is considered greater. The Forgejo serial number must therefore be inserted before the **-rcN** suffix to preserve the expected version ordering.

* Gitea **v1.18.0-rc0** will be Forgejo **v1.18.0-0-rc0**, **v1.18.0-1-rc0**
* Gitea **v1.18.0-rc1** will be Forgejo **v1.18.0-2-rc1**, **v1.18.0-3-rc1**, **v1.18.0-4-rc1**
* Gitea **v1.18.0** will be Forgejo **v1.18.0-5**, **v1.18.0-6**, **v1.18.0-7**
* etc.

Because Forgejo is a soft fork of Gitea, it must retain the same release numbering scheme to be compatible with libraries and tools that depend on it. For instance, the tea CLI or the Gitea SDK will behave differently depending on the server version they connect to. If Forgejo had a different numbering scheme, it would no longer be compatible with the Gitea ecosystem.

From a [Semantic Versioning](https://semver.org/) standpoint, all Forgejo releases are [pre-releases](https://semver.org/#spec-item-9) because they are suffixed with a dash. They are syntactically correct but do not comply with the Semantic Versioning recommendations. Gitea is not compliant either and as long as Forgejo is a soft fork, it inherits this problem.

## Release process

### Merge all feature branches

* Reset the vX.Y/forgejo branch to the Gitea tag vX.Y.Z
* Merge all feature branches into the vX.Y/forgejo branch

### Release Notes

* Add an entry in RELEASE-NOTES.md
* Copy/paste the matching entry from CHANGELOG.md
* Update the PR references prefixing them with https://github.com/go-gitea/gitea/pull/

### Testing

When Forgejo is released, artefacts (packages, binaries, etc.) are first published by the CI/CD pipelines in the https://codeberg.org/forgejo-experimental organization, to be downloaded and verified to work.

* Push the vX.Y.Z-N tag to https://codeberg.org/forgejo-integration (if it fails for whatever reason, the tag and the release can be removed manually)
  * Binaries are built and uploaded to https://codeberg.org/forgejo/forgejo-integration/releases
  * Container images are built and uploaded to https://codeberg.org/forgejo-integration/-/packages/container/forgejo/versions
* Push the vX.Y.Z-N tag to https://codeberg.org/forgejo/experimental
  * Binaries are downloaded from https://codeberg.org/forgejo-integration, signed and copied to https://codeberg.org/forgejo-experimental
  * Container images are copied from https://codeberg.org/forgejo-integration to https://codeberg.org/forgejo-experimental
* Fetch the Forgejo release as part of the [forgejo-ci](https://codeberg.org/Codeberg-Infrastructure/scripted-configuration/src/branch/main/hosts/forgejo-ci) test suite. Push the change to a branch of a repository enabled in https://woodpecker-local.forgejo.org/ ([read more...](https://codeberg.org/forgejo/forgejo/issues/208)). It will deploy the release and run high level integration tests.
* Reach out to packagers and users to manually verify the release works as expected

### Publication

* Push the vX.Y.Z-N tag to https://codeberg.org/forgejo/release
  * Binaries are downloaded from https://codeberg.org/forgejo-integration, signed and copied to https://codeberg.org/forgejo
  * Container images are copied from https://codeberg.org/forgejo-integration to https://codeberg.org/forgejo

### Website update

* Restart the last CI build at https://codeberg.org/forgejo/website/src/branch/main/
* Verify https://forgejo.org/download/ points to the expected release
* Manually try the instructions to work

## Release signing keys management

A GPG master key with no expiration date is created and shared with members of the Owners team via encrypted email. A subkey with a one year expiration date is created and stored in the secrets repository, to be used by the CI pipeline. The public master key is stored in the secrets repository and published where relevant.

### Master key creation

* gpg --expert --full-generate-key
* key type: ECC and ECC option with Curve 25519 as curve
* no expiration
* id: Forgejo Releases <contact@forgejo.org>
* gpg --export-secret-keys --armor EB114F5E6C0DC2BCDD183550A4B61A2DC5923710 and send via encrypted email to Owners
* gpg --export --armor EB114F5E6C0DC2BCDD183550A4B61A2DC5923710 > release-team-gpg.pub
* commit to the secret repository

### Subkey creation and renewal

* gpg --expert --edit-key EB114F5E6C0DC2BCDD183550A4B61A2DC5923710
* addkey
* key type: ECC (signature only)
* key validity: one year
* create [an issue](https://codeberg.org/forgejo/forgejo/issues) to schedule the renewal

#### 2023

* gpg --export --armor F7CBF02094E7665E17ED6C44E381BF3E50D53707 > 2023-release-team-gpg.pub
* gpg --export-secret-keys --armor F7CBF02094E7665E17ED6C44E381BF3E50D53707 > 2023-release-team-gpg
* commit to the secrets repository
* renewal issue https://codeberg.org/forgejo/forgejo/issues/58

### CI configuration

In the Woodpecker CI configuration the following secrets must be set:

* `releaseteamgpg` is the secret GPG key used to sign the releases
* `releaseteamuser` is the user name to authenticate with the Forgejo API and publish the releases
* `releaseteamtoken` is the token to authenticate `releaseteamuser` with the Forgejo API and publish the releases
* `domain` is `codeberg.org`

## Users, organizations and repositories

### Shared user: release-team

The [release-team](https://codeberg.org/release-team) user publishes and signs all releases. The associated email is mailto:release@forgejo.org.

The public GPG key used to sign the releases is [EB114F5E6C0DC2BCDD183550A4B61A2DC5923710](https://codeberg.org/release-team.gpg) `Forgejo Releases <release@forgejo.org>`

### Shared user: forgejo-ci

The [forgejo-ci](https://codeberg.org/forgejo-ci) user is dedicated to https://forgejo-ci.codeberg.org/ and provides it with OAuth2 credentials it uses to run.

### Shared user: forgejo-experimental-ci

The [forgejo-experimental-ci](https://codeberg.org/forgejo-experimental-ci) user is dedicated to provide the application tokens used by Woodpecker CI repositories to build releases and publish them to https://codeberg.org/forgejo-experimental. It does not (and must not) have permission to publish releases at https://codeberg.org/forgejo.

### Integration and experimental organization

The https://codeberg.org/forgejo-integration organization is dedicated to integration testing. Its purpose is to ensure all artefacts can effectively be published and retrieved by the CI/CD pipelines. 

The https://codeberg.org/forgejo-experimental organization is dedicated to publishing experimental Forgejo releases. They are copied from the https://codeberg.org/forgejo-integration organization.

The `forgejo-experimental-ci` user as well as all Forgejo contributors working on the CI/CD pipeline should be owners of both organizations.

The https://codeberg.org/forgejo-integration/forgejo repository is coupled with a Woodpecker CI repository configured with the credentials provided by the https://codeberg.org/forgejo-experimental-ci user. It runs the pipelines found in `releases/woodpecker-build/*.yml` which builds and publishes an unsigned release in https://codeberg.org/forgejo-integration.

### Experimental and release repositories

The https://codeberg.org/forgejo/experimental private repository is coupled with a Woodpecker CI repository configured with the credentials provided by the https://codeberg.org/forgejo-experimental-ci user. It runs the pipelines found in `releases/woodpecker-publish/*.yml` which signs and copies a release from https://codeberg.org/forgejo-integration into https://codeberg.org/forgejo-experimental.

The https://codeberg.org/forgejo/release private repository is coupled with a Woodpecker CI repository configured with the credentials provided by the https://codeberg.org/release-team user. It runs the pipelines found in `releases/woodpecker-publish/*.yml` which signs and copies a release from https://codeberg.org/forgejo-integration into https://codeberg.org/forgejo.
