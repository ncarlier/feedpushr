# Expression language definition

The internal expression engine is [Expr](https://antonmedv.github.io/expr/).
You can find exhaustive documentation on the possibilities of the language at this address.

## Supported Literals

* **strings** - single and double quotes (e.g. `"hello"`, `'hello'`)
* **numbers** - e.g. `103`, `2.5`
* **arrays** - e.g. `[1, 2, 3]`
* **maps** - e.g. `{foo: "bar"}`
* **booleans** - `true` and `false`
* **nil** - `nil`

## Accessing article properties

Article properties can be accessed by using their name:

- `Title`: Article title
- `Text`: Article text description
- `Content`: Article HTML content
- `Link`: Article URL
- `Meta["key"]`: Metadata value
- `Tags`: Array of tags

## Supported Operators

### Comparison Operators

* `==` (equal)
* `!=` (not equal)
* `<` (less than)
* `>` (greater than)
* `<=` (less than or equal to)
* `>=` (greater than or equal to)

### Logical Operators

* `not` or `!`
* `and` or `&&`
* `or` or `||`

### String Operators

* `+` (concatenation)
* `matches` (regex match)
* `contains` (string contains)
* `startsWith` (has prefix)
* `endsWith` (has suffix)

To test if a string does *not* match a regex, use the logical `not` operator in combination with the `matches` operator:

```coffeescript
not (Title matches "^b.+")
```

You must use parenthesis because the unary operator `not` has precedence over the binary operator `matches`.

Example:

```coffeescript
'Arthur' + ' ' + 'Dent'
```

Result will be set to `Arthur Dent`.

### Membership Operators

* `in` (contain)
* `not in` (does not contain)

Example:

```coffeescript
"test" in Tags
```

```coffeescript
meta.author in ['me', 'other']
```

## Builtin functions

* `len` (length of array or string)
* `all` (will return `true` if all element satisfies the predicate)
* `none` (will return `true` if all element does NOT satisfies the predicate)
* `any` (will return `true` if any element satisfies the predicate)
* `one` (will return `true` if exactly ONE element satisfies the predicate)
* `filter` (filter array by the predicate)
* `map` (map all items with the closure)

Example:

```go
len(Text) > 200
```

## Closures

* `{...}` (closure)

Closures allowed only with builtin functions. To access current item use `#` symbol.

```
one(tags, {# == 'test'})
```
