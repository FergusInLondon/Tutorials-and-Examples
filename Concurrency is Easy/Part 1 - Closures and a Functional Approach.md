# Refactoring for Concurrency: The humble closure

Let me start by giving one bit of advice: **concurrency *should* to be baked in and architected from the beginning**. However if *should* was *would* then the world would be a far less interesting place! Real life software development involves rough proof-of-concepts, corners cut in the heat of the moment, misunderstandings over requirements, and many other unforeseen hiccups.

So what happens if you inherit a project that *simply isn't concurrent*? Having recently been in that situation, I've decided to share some of the tips I used to introduce concurrency with minor difficulties.

A word of warning though: as with any refactoring and changes, *test cases are incredibly useful*. If you're inheriting a project like the one I described, these test cases may not be available - so this is a good time to introduce them as you go on. You'll thank yourself later!

## Reducing globals and state via *closures*

One of the biggest bugbears with introducing a concurrent architecture in to an existing project is *global state*. Fortunately though, reducing global state can be achieved via the creative use of the humble "*closure*".

### What's a closure?

*If you understand what a closure is, then you can likely ignore this section.. perhaps even the whole post - as the benefits of a closure in capturing state is likely clear. Otherwise, listen on!*

This is what [Wikipedia](https://en.wikipedia.org/wiki/Closure_(computer_programming)) has to say about a closure:

> In programming languages, a closure, also lexical closure or function closure, is a technique for implementing lexically scoped name binding in a language with first-class functions. Operationally, a closure is a record storing a function[a] together with an environment.

Lexical scoping is - to be blunt, *the kind of scoping most people think is normal*. Whereby a block of code inherits access to the scope of it's parent, a classic example is an `if` statement. For example:

```js
let a = 3

if( true ) {
    console.log(a) // prints '3', variables in the parent scope is available.

    let b = 4
    console.log(b) // prints '4', the variable is in the current scope
}

console.log(a) // Yep;  'a' was declared in this scope, therefore is obviously accessible
console.log(b) // Nope; 'b' doesn't exist in this scope, it was declared in a child scope
```

This means that the following would also work:

```js
let a = 3
let b = 2

let added = (() => {
    return a + b // these are available, they're in the parent scope
})()

console.log(added) // '5'
```

Of course, we can also redeclare variables and "*shadow*" them:

```js
let a = 3
let b = 2

let added = (() => {
    let a = 3
    return a + b
})()

console.log(a)      // '3' - 'a' was *shadowed* in the child scope, it had it's own version of 'a'
console.log(added)  // '6' - the addition in the anonymous function used the local scope (i.e it's own declaration) of 'a'
```

This is *simple scoping*: but it's important for understanding the power behind a closure. Now with the basic *lexical scoping* rules out of the way, what happens if we return a function?

```js
function Counter() {
    let counter = 0
    return () => {
        counter++
        return counter
    }
}

let c1 = Counter()  // c1 = a function
let c2 = Counter()  // c2 = another function
console.log(c())    // '1'
console.log(c())    // '2'
console.log(d())    // '1' - c2 is using the lexical scope of the second call to `Counter()`, therefore has it's own isolated state
```

Notice how the calls to c() result in the state lasting longer than the actual function invokation? *That's the beauty of a closure.*

#### Yo, I'm not writing Javascript though!

This isn't a Javascript specific tip, I just find javascript as the most concise and readable language to demonstrate closures. To use a closure you can use *any language which supports [first-class functions](https://en.wikipedia.org/wiki/First-class_function). This means languages like `Rust`, `Go`, `PHP`, or `Python` - in addition to the obvious functional languages like various Lisp dialests or `Scala` .

### Using Closures to reduce global state

Given the above, it should be easy to see how a closure can be a useful tool for reducing global state. Let's look at a specific, albeit fictional, example. This example is a little different, as it demonstrates *multiple functions* that rely upon that state.

```js

var filename
var filesize

getFileDetails = () => {
    // Read file stats
    filesize = ...
}

connectionHandler = () => {
    getFileDetails()

    return {
        size: filesize
    }
}

```

In this fictional example, we define an endpoint that can be used for HTTP requests; it takes a filename and returns a JSON payload with the file size. It works, but there's one flaw: what happens if there are multiple concurrent requests? `filename` or `filesize` will get overwritten, and it's likely that some requests will return incorrect results. That's not good.

Let's remedy that, and isolate *both the `connectionHandler` and `getFileDetails`* functions, as well as the state that they rely upon.

```js
connectionHandler = ((){
    var filename
    var filesize

    getFileDetails => {
        //
        filesize = ...
    }

    return () => {
        getFileDetails()

        return {
            size: filesize
        }
    }
})()

```

And our call to 



By now it should be quite clear how the use of a closure can *help* reduce global state; however this relies upon specific objects of global state *only being used by one function*. So if you have a function that should be isolated to thread/routine, but relies upon persistance even when the value goes out of state (i.e `static` in C), then closures could well be a saving grace.

## Next time!

This has been an incredibly gentle tip; one that may be obvious to some, or a new lesson to others. Personally I've found it's served me well as a "*quickfix*" to isolate global state, bringing it down to a functional level without the need for refactoring.

Unfortunately though, there *are* times when global state is not only desired - but required. An example of this would be a service which broadcasts updates regarding a single entity to multiple clients/connections. For times where global state is required, we need to introduce mechanisms like a `mutex` or a `semaphore`. Next time we'll have a look at those, where they can be employed, and the limitations of them.