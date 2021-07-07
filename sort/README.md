## Sort package in golang:

- Search, SearchInts and SearchStrings all take sorted slice as input and panics if the input is not a slice.
- SearchInts and SearchStrings needs slice sorted in ascending order.
- Search can used for strings sorted in descending order (just reverse the condition in the func passed in the parameter to <=).
- Search uses binary search to search the data.
