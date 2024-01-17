# CloudQuery knowbe4 Source Plugin

[![test](https://github.com/razin99/cq-source-knowbe4/actions/workflows/test.yaml/badge.svg)](https://github.com/razin99/cq-source-knowbe4/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/razin99/cq-source-knowbe4)](https://goreportcard.com/report/github.com/razin99/cq-source-knowbe4)

A knowbe4 source plugin for CloudQuery that loads data from knowbe4 to any database, data warehouse or data lake supported by [CloudQuery](https://www.cloudquery.io/), such as PostgreSQL, BigQuery, Athena, and many more.

## Links

- [CloudQuery Quickstart Guide](https://www.cloudquery.io/docs/quickstart)
- [Supported Tables](docs/tables/README.md)

## Configuration

The following source configuration file will sync to a PostgreSQL database. See [the CloudQuery Quickstart](https://www.cloudquery.io/docs/quickstart) for more information on how to configure the source and destination.

```yaml
kind: source
spec:
  name: knowbe4
  registry: local
  path: ./cq-source-knowbe4
  destinations:
    - csv
  tables:
    - knowbe4_training_enrollments
    - knowbe4_account_risk_score_history
    - knowbe4_users
  spec:
    base_url: ${KB4_BASE_URL}
    token: ${KB4_TOKEN}
```

## Development

### Run tests

```bash
make test
```

### Run linter

```bash
make lint
```

### Generate docs

```bash
make gen-docs
```

### Release a new version

1. Run `git tag v1.0.0` to create a new tag for the release (replace `v1.0.0` with the new version number)
2. Run `git push origin v1.0.0` to push the tag to GitHub

Once the tag is pushed, a new GitHub Actions workflow will be triggered to build the release binaries and create the new release on GitHub.
To customize the release notes, see the Go releaser [changelog configuration docs](https://goreleaser.com/customization/changelog/#changelog).
