# Stack vs Heap Allocation in Go

## Understanding Memory Allocation in Go

In Go, memory allocation happens in two primary locations:
- **Stack**: For local variables with known size at compile time
- **Heap**: For dynamically allocated memory or when escape analysis determines variables must outlive their function scope

## The Examples in Your Code

### 1. Stack Allocation Example: `bufferExample()`

```go
func bufferExample() {
    var buf bytes.Buffer // stack
    buf.WriteString("Hello Buffer!")
    fmt.Println(buf.String())
}
```

- `var buf bytes.Buffer` allocates the buffer struct on the stack
- This is a **stack allocation** because:
  - The size of `bytes.Buffer` struct is known at compile time
  - The variable has a limited scope (function scope)
  - The compiler can determine that the variable won't escape the function
  - No pointers to this buffer are returned or passed outside the function
  - Memory is automatically freed when the function returns

Stack allocation characteristics:
- **Fast allocation**: No memory manager involved
- **Automatic deallocation**: When the function returns, stack space is reclaimed
- **Predictable performance**: No garbage collection needed for stack variables
- **Limited size**: Stack size is limited (typically 1-8MB per goroutine)

### 2. Heap Allocation Example: `multiReaderExample()`

```go
func multiReaderExample() {
    r1 := strings.NewReader("Hello ")
    r2 := strings.NewReader("World!")
    mr := io.MultiReader(r1, r2)
    buf := new(bytes.Buffer) // heap
    _, err := buf.ReadFrom(mr)
    if err != nil {
        log.Fatalln("Error reading from multi reader:", err)
    }
    fmt.Println(buf.String())
}
```

- `buf := new(bytes.Buffer)` allocates memory for a `bytes.Buffer` on the heap
- This is a **heap allocation** because:
  - `new()` always allocates on the heap
  - `new()` returns a pointer (`*bytes.Buffer`)
  - The memory location must persist beyond the function's local scope
  - The allocation happens at runtime, not compile time

## Detailed Technical Explanation

### Escape Analysis

Go's compiler performs "escape analysis" to determine whether a variable should be allocated on the stack or heap:

```go
// Stack allocation - variable doesn't escape
func stackExample() {
    x := 42  // allocated on stack
    // x is used only within function scope
    fmt.Println(x)
}

// Heap allocation - variable escapes the function
func heapExample() *int {
    x := 42  // must be allocated on heap
    return &x  // reference to x escapes the function
}
```

### Why Stack vs Heap Matters in Your Code

#### Stack Allocation (`var buf bytes.Buffer`)
```go
func bufferExample() {
    var buf bytes.Buffer // stack
    buf.WriteString("Hello Buffer!")
    fmt.Println(buf.String())
}
```

- The `bytes.Buffer` struct itself is allocated on the stack
- The internal data slice inside the buffer still lives on the heap (since slices are dynamic)
- However, the struct containing the slice header is stack-allocated
- More efficient for local use with limited scope

#### Heap Allocation (`new(bytes.Buffer)`)
```go
func multiReaderExample() {
    buf := new(bytes.Buffer) // heap
    // buf is a *bytes.Buffer pointer to heap-allocated memory
    _, err := buf.ReadFrom(mr)
    // ...
}
```

- `new(bytes.Buffer)` always allocates on the heap
- Returns a pointer (`*bytes.Buffer`)
- Memory remains accessible even after function returns
- Requires garbage collection to clean up

## Performance Implications

### Stack Allocation Advantages:
- **Speed**: Stack allocation is very fast (just adjusting stack pointer)
- **No GC pressure**: Stack variables don't create garbage collection work
- **Locality**: Stack access has better cache locality
- **Automatic cleanup**: No need to track memory manually

### Heap Allocation Advantages:
- **Flexibility**: Can outlive function scope
- **Large data**: Can handle data larger than stack limits
- **Dynamic size**: Useful for runtime-determined sizes

## Real-world Example: Demonstrating the Difference

```go
// Stack-allocated buffer - efficient for small, temporary operations
func processSmallData() {
    var buf bytes.Buffer // stack allocated (struct)
    buf.WriteString("small data")
    // buf is automatically freed when function returns
}

// Heap-allocated buffer - necessary when buffer needs to outlive function
func createBuffer() *bytes.Buffer {
    return new(bytes.Buffer) // must be heap allocated
    // because we return a pointer to it
}

// Large buffer - may force heap allocation due to size
func processLargeData() {
    largeBuf := make([]byte, 1024*1024) // 1MB slice
    // Large allocations often go to heap anyway
    // due to stack size limitations
}
```

## Key Takeaways for Course Students

1. **Stack allocation** (`var buf bytes.Buffer`) - Efficient for local variables with limited scope
2. **Heap allocation** (`new(bytes.Buffer)`) - Necessary when memory must outlive function scope
3. **Escape analysis** - Go compiler automatically decides allocation location
4. **Performance matters** - Stack allocation is generally faster, use it when possible
5. **Memory management** - No need to manually free; handled by compiler (stack) or GC (heap)

## Compiler Escape Analysis in Practice

You can verify escape analysis using:
```bash
go build -gcflags="-m" main.go
```

This will show you which variables escape to the heap and which stay on the stack.

The distinction between stack and heap in your code demonstrates Go's automatic memory management while still allowing you to write efficient code by choosing the right allocation approach for your use case.