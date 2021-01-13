# EntityFilter
The EntityFilter project aims to provide a generic search syntax for golang objects. Frontend filter parsers are provided along with backend searchers for different types of datastores.

## Search Syntax

### Filter Components
  * Attribute to query
  * Type of match to perform
  * Target value

#### Example
`name:=myname`
  * `name` is the attribute to query
  * `:=` specifies an equals comparison
  * `myname` is the target value

### Supported Operations
  * Equals `:=`
  * Not Equals `:!=`
  * Contains `:~`
  * Not Contains `:!~`

### Using Multiple Filters
  * Filters can be chained together with commas
  * Each filter condition must be true for the entity to match
  * There is currently no limit on the number of filters that can be chained, however you can expect that more conditions will result in a longer query depending on the backend datastore

#### Example
`name:=myname,age:=12`

In the above example, name must be "myname" AND age must be 12
