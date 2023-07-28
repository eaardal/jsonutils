# JSON Utils

Convenience utilities for working with JSON on the CLI.

## Commands

### `jsonutils array new`

Returns an empty array as a string: `"[]"`

### `jsonutils array append --array <array> --value <value> [--type <type>]`

Adds the `<value>` to the end of the `<array>`.

The `--type` argument can be either `number` or `bool`.

Example:

```shell
$arr=$(jsonutils array new)
# arr: []
$arr=$(jsonutils array append --array "$arr" --value "hello")
# arr: ["hello"]
$arr=$(jsonutils array append --array "$arr" --value "1" --type number)
# arr: ["hello", 1]
$arr=$(jsonutils array append --array "$arr" --value "true" --type bool)
# arr: ["hello", 1, true]
```

### `jsonutils array spread --array <array> --value <value>`

Spreads the `<value>`, which should be another array, into the `<array>`. This equivalent to this javascript:

```javascript
const greetings = ["hello", "world"]
const farewells = ["goodbye"]
const arr = [...arr, ...farewells]
// arr: ["hello", "world", "goodbye"]
```

Example:

```shell
$arr="[\"hello\", \"world\"]"
$arr=$(jsonutils array spread --array "$arr" --value "[\"goodbye\"]")
# arr: ["hello", "world", "goodbye"]
```
