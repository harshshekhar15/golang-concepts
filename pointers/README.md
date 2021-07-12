## Pointers in Golang

### Nil pointer dereferences

**Dereference** - Obtain the address of a data item held in another location from a pointer.

**When does go runtime throw nil pointer dereferences?**
When a pointer declared in a go program points to `nil` (i.e. the address stored in the pointer is nil or the zero value of a pointer) the go runtime throws a nil pointer dereferences error.

### Pros of using pointers
**Memory efficient:** For a large struct passing a pointer to that struct to a function instead of passing a copy of the actual struct can save memory. Since Go utilizes pass by value semantics for function calls, all arguments to a function are copied at the time the function is invoked. If a struct passed to a function occupies many bytes or kilobytes (or more) of memory, all of that is copied. By contrast, if a pointer to the same struct is passed to a function, when that pointer is copied (assuming you’re using a 64 bit operating system), only eight bytes are copied. And these eight bytes reference the location in memory where the original struct is stored.

### Cons of using pointers
**Garbage collection:** Increased use of pointers equates to increased use of heap, leading in turn to more CPU cycles spent on garbage collection, and fewer CPU cycles spent executing application logic


### Blogs
- https://krancour.medium.com/go-pointers-when-to-use-pointers-4f29256ddff3