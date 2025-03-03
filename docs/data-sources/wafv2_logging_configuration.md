---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_wafv2_logging_configuration Data Source - terraform-provider-awscc"
subcategory: ""
description: |-
  Data Source schema for AWS::WAFv2::LoggingConfiguration
---

# awscc_wafv2_logging_configuration (Data Source)

Data Source schema for AWS::WAFv2::LoggingConfiguration



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) Uniquely identifies the resource.

### Read-Only

- `log_destination_configs` (List of String) The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.
- `logging_filter` (Attributes) Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation. (see [below for nested schema](#nestedatt--logging_filter))
- `managed_by_firewall_manager` (Boolean) Indicates whether the logging configuration was created by AWS Firewall Manager, as part of an AWS WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration.
- `redacted_fields` (Attributes List) The parts of the request that you want to keep out of the logs. For example, if you redact the HEADER field, the HEADER field in the firehose will be xxx. (see [below for nested schema](#nestedatt--redacted_fields))
- `resource_arn` (String) The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.

<a id="nestedatt--logging_filter"></a>
### Nested Schema for `logging_filter`

Read-Only:

- `default_behavior` (String) Default handling for logs that don't match any of the specified filtering conditions.
- `filters` (Attributes List) The filters that you want to apply to the logs. (see [below for nested schema](#nestedatt--logging_filter--filters))

<a id="nestedatt--logging_filter--filters"></a>
### Nested Schema for `logging_filter.filters`

Read-Only:

- `behavior` (String) How to handle logs that satisfy the filter's conditions and requirement.
- `conditions` (Attributes List) Match conditions for the filter. (see [below for nested schema](#nestedatt--logging_filter--filters--conditions))
- `requirement` (String) Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.

<a id="nestedatt--logging_filter--filters--conditions"></a>
### Nested Schema for `logging_filter.filters.conditions`

Read-Only:

- `action_condition` (Attributes) A single action condition. (see [below for nested schema](#nestedatt--logging_filter--filters--conditions--action_condition))
- `label_name_condition` (Attributes) A single label name condition. (see [below for nested schema](#nestedatt--logging_filter--filters--conditions--label_name_condition))

<a id="nestedatt--logging_filter--filters--conditions--action_condition"></a>
### Nested Schema for `logging_filter.filters.conditions.label_name_condition`

Read-Only:

- `action` (String) Logic to apply to the filtering conditions. You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.


<a id="nestedatt--logging_filter--filters--conditions--label_name_condition"></a>
### Nested Schema for `logging_filter.filters.conditions.label_name_condition`

Read-Only:

- `label_name` (String) The label name that a log record must contain in order to meet the condition. This must be a fully qualified label name. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label.





<a id="nestedatt--redacted_fields"></a>
### Nested Schema for `redacted_fields`

Read-Only:

- `json_body` (Attributes) Inspect the request body as JSON. The request body immediately follows the request headers. This is the part of a request that contains any additional data that you want to send to your web server as the HTTP request body, such as data from a form. (see [below for nested schema](#nestedatt--redacted_fields--json_body))
- `method` (Map of String) Inspect the HTTP method. The method indicates the type of operation that the request is asking the origin to perform.
- `query_string` (Map of String) Inspect the query string. This is the part of a URL that appears after a ? character, if any.
- `single_header` (Attributes) Inspect a single header. Provide the name of the header to inspect, for example, User-Agent or Referer. This setting isn't case sensitive. (see [below for nested schema](#nestedatt--redacted_fields--single_header))
- `uri_path` (Map of String) Inspect the request URI path. This is the part of a web request that identifies a resource, for example, /images/daily-ad.jpg.

<a id="nestedatt--redacted_fields--json_body"></a>
### Nested Schema for `redacted_fields.json_body`

Read-Only:

- `invalid_fallback_behavior` (String) What AWS WAF should do if it fails to completely parse the JSON body.
- `match_pattern` (Attributes) The patterns to look for in the JSON body. AWS WAF inspects the results of these pattern matches against the rule inspection criteria. (see [below for nested schema](#nestedatt--redacted_fields--json_body--match_pattern))
- `match_scope` (String) The parts of the JSON to match against using the MatchPattern. If you specify All, AWS WAF matches against keys and values.

<a id="nestedatt--redacted_fields--json_body--match_pattern"></a>
### Nested Schema for `redacted_fields.json_body.match_pattern`

Read-Only:

- `all` (Map of String) Match all of the elements. See also MatchScope in JsonBody. You must specify either this setting or the IncludedPaths setting, but not both.
- `included_paths` (List of String) Match only the specified include paths. See also MatchScope in JsonBody.



<a id="nestedatt--redacted_fields--single_header"></a>
### Nested Schema for `redacted_fields.single_header`

Read-Only:

- `name` (String) The name of the query header to inspect.
