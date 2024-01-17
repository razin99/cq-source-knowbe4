# Table: knowbe4_users

This table shows data for Knowbe4 Users.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`int64`|
|first_name|`utf8`|
|last_name|`utf8`|
|email|`utf8`|
|phish_prone_percentage|`float64`|
|groups|`list<item: int64, nullable>`|
|current_risk_score|`float64`|
|joined_on|`timestamp[us, tz=UTC]`|
|last_sign_in|`timestamp[us, tz=UTC]`|