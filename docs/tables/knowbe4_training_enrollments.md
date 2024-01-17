# Table: knowbe4_training_enrollments

This table shows data for Knowbe4 Training Enrollments.

The primary key for this table is **enrollment_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|enrollment_id (PK)|`int64`|
|content_type|`utf8`|
|module_name|`utf8`|
|user|`json`|
|campaign_name|`utf8`|
|enrollment_date|`timestamp[us, tz=UTC]`|
|start_date|`timestamp[us, tz=UTC]`|
|completion_date|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|time_spent|`int64`|
|policy_acknowledged|`bool`|