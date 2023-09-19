# Go struct tags

A struct is a user-defined type that contains a collection of fields. It is used to group related data to form a single unit. A Go struct can be compared to a lightweight class without the inheritance feature.

## What is struct tag
A struct tag is additional meta data information inserted into struct fields. The meta data can be acquired through reflection. Struct tags usually provide instructions on how a struct field is encoded to or decoded from a format.

### Struct tags are used in popular packages including:
* encoding/json
* encoding/xml
* gopkg.in/mgo.v2/bson
* gorm.io/gorm
* github.com/gocarina/gocsv
* gopkg.in/yaml.v2
[See example](./struct_sample.go)

```shell
 go mod init struct_tags
 go mod tidy
 go run struct_sample.go 
```
The example uses struct tags to configure how JSON data is encoded.
```go
type User struct {
    Id         int    `json:"id"`
    Name       string `json:"name"`
    Occupation string `json:"occupation,omitempty"`
}
```
With `json:"id"` struct tag, we encode the Id field in lowercase. In addition, the omitempty omits the Occupation field if it is empty.

```go
go run main.go

/* output
{
    "id": 1,
    "name": "John Doe",
    "occupation": "gardener"
    }
{
    "id": 1,
    "name": "John Doe"
}
 */
```
>more.. https://zetcode.com/golang/struct-tag/

