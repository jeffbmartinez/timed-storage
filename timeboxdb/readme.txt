A key value storage system which can store multiple values per key, along with a timebox for which each value is considered active.

The Get call, when supplied a key, returns all values for which associated values are active. In other words, the values for which the associated time box is surrounding the current time (or another time, if provided).

~~~ API - json ~~~

~~ Get - (GET from /get) ~~

Request Fields:
- api (required): Integer specifying API version
- key (required): A value to use as a key, can be of any comparable type.
- activeTime (optional): Specify a time used to determine whether or not a value is active. This is in seconds since January 1, 1970 in UTC, also known as unix time or epoch time. If it is omitted, the server will use the current time.

Examples:
{
  "api": 1,
  "key": "jeff"
}

{
  "api": 1,
  "key": "jeff",
  "activeTime": 1393179574
}

Response Fields:
- values: List of active values

Examples:
{
  "values": []
}

{
  "values": [33, "martinez"]
}

~~ Put (POST to /put) ~~

Request Fields:
- api (required): Integer specifying API version
- key (required): Key with which to associate the value being stored
- startTime (required): A time to determine when this value is active. It is specified in seconds since January 1, 1970 UTC, also known as unix time or epoch time.
- endTime (optional: required if 'duration' is omitted, illegal if 'duration' is present): A time to determine when this value is no longer active. It is specified in seconds since January 1, 1970 UTC, also known as unix time or epoch time.
- duration (optional: required if 'endTime' is omitted, illegal if 'endTime' is present): The number of seconds after startTime to keep the value active.

Examples:
{
  "api": 1,
  "key": "jeff",
  "startTime": 1393179514
  "endTime": 1393179634
}

{
  "api": 1,
  "key": "jeff",
  "startTime": 1393179514
  "duration": 120
}

Response Fields:
- There is no body in the response. The HTTP status of 200 OK indicates success.
