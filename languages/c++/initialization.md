# Initialization
There are two kinds of initializations: Static initialization and Dynamic initialization.<br>
Static initialization is either Zero initialization or Constant initialization, any other initialization is Dynamic initialization.

### Static initialization
Static initialization happens at compile time. It happens on the global, static, static local and thread_local (abbr. GSST) variables. Initial values of GSST variables are evaluated during compilation and burned into the data section of the executable.

At compile time, if the initial value of a GSST variable can be evaluated, Constant initialization will happen; Otherwise, Zero initialization will happen, so it will be burned into the bss section of the executable.<br>
If Zero initialization happen at compile time, Dynamic initialization will happen at runtime.

##### Constant initialization
```
struct MyStruct
{
    static int a;
};
int MyStruct::a = 67; // Constant initialization
int i = 1;            // Constant initialization
```

##### Zero initialization & Dynamic initialization
```
std::string s; // compile time: zero-initialized to indeterminate value
               // runtime: then default-initialized to ""
```

##### Force Constant initialization: constexpr
It isn't quite easy to distinguish if a GSST is constant-initialization or zero-initialization & dynamic initialization. So we introduce the keyword `constexpr`. BTW, it implies `const`.

`constexpr` will force the compiler to do Constant initialization for a GSST variable. If Constant initialization can't happen, a compile error will be generated.

```
int i = 1;           // Constant initialization
constexpr int j = 1; // Constant initialization
constexpr int k = i; // ERROR: the value of ‘i’ is not usable in a constant expression

int func()
{
    return 1;
}

constexpr int func2()
{
    return 1;
}

int a = func();            // zero initialization & dynamic initialization
constexpr int b = func();  // ERROR: b must be initialized by a constant expression, func cannot be used in a constant expression
int c = func2();           // Constant initialization
constexpr int d = func2(); // Constant initialization
```

##### constinit & consteval
_to be continued..._


### Dynamic initialization





```
struct A { int m; };
struct B { ~B(); int m; };
struct C { C() : m(){}; ~C(); int m; };
struct D { D(){}; int m; };
struct E { E() = default; int m;};  
struct F {F(); int m;};  F::F() = default;
```
