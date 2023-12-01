package main

/*
🔴 : Très important
🟠 : Important
🟡 : Bon pour la compréhension

	🔴 Marshaling is the process of transforming the memory representation
of an object to a data format suitable for storage or transmission,
and it is typically used when data must be moved between different
parts of a computer program or from one program to another.

The inverse of marshaling is called unmarshaling

This Marshaling and UnMarshaling is Golang trying to convert struct
into JSON objects and JSON objects into Golang structs.
This section of the course is about how you can transfer to JSON and
from JSON(string literal byte slice) back into Golang struct. 🔴

	❗ Remember Golang is a web backend language ❗

Eventually we'll learn how to get JSON objects via HTTP request.
While we don't yet know how to do that, Tod is preparing us for that future
where we'll have put into a byte slice a JSON object as a string literal.
Or the reverse take a JSON object that is a byte slice and convert it to a Golang struct.

	🟡 Recall that in Golang when you put a string in backtick ``, it is treated as raw string literal.
Meaning that it is read exactly as is a UTF-8 string with no escape characters only runes(utf-8 characters).
Vs the "" double quotes which allows for escaped characters.

	🟡 In quotes "" you need to escape new lines,
tabs and other characters (in Golang these are called verbs) that do not need to be escaped in backticks ``.
If you put a line break(\n) in a backtick string, it is interpreted as a '\n'.
Meaning, it's not converted into the actual line break but read as the full characters contained within the ticks.


So now we have this JSON object that we got from HTTP somehow.
How do we get it into a struct or how do we get our nice Golang struct into JSON?
That is what the JSON package is for.

	🟡 As with all structs in Go, it’s important to remember that only fields with a capital first letter
are visible to external programs/packages like the JSON Marshaller. Meaning that if you don't
capitalize the first letter it won't get exposed when you convert it to JSON using the json.Marshall() function.
This is because you cannot access that field outside of the struct,
remember if it's not a capital first letter it cannot be accessed outside of the struct.

	🟡 Every type in Go implements an empty interface without you having to do anything.
A empty interface is, any type that implements zero methods.
By explicit design of the Go lang specifications every type at least implements a zero method interface called a empty interface.
*/

/*
	🟠 func Marshal(v interface{}) ([]byte, error) 🟠

it takes a v interface{}, which can be 'any' go type.
Basically everything from a struct to a primitive it will take it and try to convert it to a JSON object.
It returns two things:

	1️⃣ A slice of byte []byte, containing the literal string that is the JSON object.
	2️⃣ An error, letting you know if anything went wrong

	🟡 Typically you give the JSON marshal function a pre-filled struct, or a raw string literal formatted to JSON

Only the exported (public) fields of a struct will be present in the JSON output.
A field with a json: tag is stored with its tag name instead of its variable name.
Pointers will be encoded as the values they point to, or null if the pointer is nil.

Important: When you tag a struct field with a 'json: tag' you are telling the Marshall package how to interpret the Go Struct.
Meaning you can tag something in Json to have a different name than what you are storing in the Go struct.
This is really powerful, because you have the ability to define what the json field will be stored as name wise,
which can help simplify things to a great deal!

	🟠 func Unmarshal(data []byte, v interface{}) error 🟠

unmarshal takes the following parameters:

	1️⃣ A slice of bytes (This a raw string, this is the JSON object that you want to parse)
	2️⃣ A pointer to a struct to parse the JSON into

Unmarshal returns an error if anything went wrong with parsing.


 ❓ How does Unmarshal decide which fields to try and parse ❓

For any key found in the JSON, Unmarshal will try to match it to a key found in the struct with the following logic.

For explanation sakes, I'm using "FieldName" to represent any member of a struct.

	1️⃣ It will first look for an exported(field member with a capital letter) with a tag json:"FieldName"
	2️⃣ An exported (field member with a capital letter) with the name FieldName
	3️⃣ Any exported field name, that matches the fieldname if casesensitivity is not an issue, e,g fIeLdNaMe, FIELDNAME, feildname.

		🔴 ONLY FIELDS FOUNDS IN THE DESTINATION type(struct) WILL BE DECODED 🔴

Meaning if there is a field in the JSON that isn't in the destination it will be ignored.

This is useful when you wish to pick only a few specific fields.
In particular, any unexported fields in the destination struct will be unaffected.
*/

/*
EXTRA:

	🟡 The escaped character  %+V, in the printf statement does the following :

If it's a struct it will print the value of that structure with %V,
when you have the +, %+V will print the members of the struct in the print statement.

	🟡 Relationship between strings and byte slices :

Recall that a string in Go is just a sequence of bytes.

	A byte is just an alias(type byte uint8) for a a uint8.

So the underpinnings of a string is a sequence of bytes.
This is why we can easily do a conversion on a sequence of bytes and turn it into a string and vice versa.
So when the marshal function returns a sequence of bytes, remember it's just a raw UTF-8 string.
*/
