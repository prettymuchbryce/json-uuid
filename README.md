Extends satori/go.uuid.NullUUID with better JSON Marshalling.

Until https://github.com/satori/go.uuid/pull/44 is merged, the NullUUID will always be marshalled to JSON as object with `UUID` and `Valid` properties, rather than a single field with a value of null or a string. 
