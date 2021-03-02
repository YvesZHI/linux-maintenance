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

##### aggregate type
The aggregate type is special because it can be initialized by List initialization.

An aggregate is an array or a class with no user-provided constructors, no brace-or-equal-initializers for non-static data members, no private or protected non-static data members, no base classes, and no virtual functions.<br>
In a word, an aggregate class is nothing more than the sum of its members. An array is an aggregate even if it is an array of non-aggregate class type.

```
class NotAggregate1
{
    virtual void f() {} //no virtual functions
};

class NotAggregate2
{
    int x; //x is private by default and non-static 
};

class NotAggregate3
{
public:
    NotAggregate3(int) {} //oops, user-defined constructor
};

class NotAggregate4
{
public:
    NotAggregate4() {} //oops, user-defined constructor
};

class NotAggregate5
{
public:
    NotAggregate5();
};
NotAggregate5::NotAggregate5() = default; //oops, user-defined constructor

class Aggregate1
{
public:
    NotAggregate1 member1;   //ok, public member
    Aggregate1& operator=(Aggregate1 const & rhs) {/* */} //ok, copy-assignment  
private:
    void f() {} // ok, just a private function
};

class Aggregate2
{
public:
    Aggregate2() = default; // ok, non-user-defined constructor
};
```
An array is aggregate too.
```
Type array_name[n] = {a1, a2, …, am};

if(m == n)
    the ith element of the array is initialized with ai
else if(m < n)
    the first m elements of the array are initialized with a1, a2, …, am and the other n - m elements are, if possible, value-initialized
else if(m > n)
    the compiler will issue an error
else // this is the case when n isn't specified at all like int a[] = {1, 2, 3};
    the size of the array (n) is assumed to be equal to m, so int a[] = {1, 2, 3}; is equivalent to int a[3] = {1, 2, 3};
```

##### trivial type
The trivial type is special because it can be copied by `memcpy`.

A trivially copyable class is a class that:
- has no non-trivial copy constructors,
- has no non-trivial move constructors,
- has no non-trivial copy assignment operators,
- has no non-trivial move assignment operators,
- has a trivial destructor.

A trivial class is a class that has a trivial default constructor and is trivially copyable.<br>
Note: In particular, a trivially copyable or trivial class does not have virtual functions or virtual base classes.

A copy/move constructor for class X is trivial if it is not user-provided and if
- class X has no virtual functions (10.3) and no virtual base classes, and
- the constructor selected to copy/move each direct base class subobject is trivial, and
- for each non-static data member of X that is of class type (or array thereof), the constructor selected to copy/move that member is trivial;

otherwise the copy/move constructor is non-trivial.

```
// empty classes are trivial
struct Trivial1 {};

// all special members are implicit
struct Trivial2 {
    int x;
};

struct Trivial3 : Trivial2 { // base class is trivial
    Trivial3() = default; // not a user-provided ctor
    int y;
};

struct Trivial4 {
public:
    int a;
private: // no restrictions on access modifiers
    int b;
};

struct Trivial5 {
    Trivial1 a;
    Trivial2 b;
    Trivial3 c;
    Trivial4 d;
};

struct Trivial6 {
    Trivial2 a[23];
};

struct Trivial7 {
    Trivial6 c;
    void f(); // it's okay to have non-virtual functions
};

struct Trivial8 {
     int x;
     static NonTrivial1 y; // no restrictions on static members
};

struct Trivial9 {
     Trivial9() = default; // not user-provided
      // a regular constructor is okay because we still have default ctor
     Trivial9(int x) : x(x) {};
     int x;
};

struct NonTrivial1 : Trivial3 {
    virtual void f(); // virtual members make non-trivial ctors
};

struct NonTrivial2 {
    NonTrivial2() : z(42) {} // user-provided ctor
    int z;
};

struct NonTrivial3 {
    NonTrivial3(); // user-provided ctor
    int w;
};
NonTrivial3::NonTrivial3() = default; // defaulted but not on first declaration
                                      // still counts as user-provided
struct NonTrivial5 {
    virtual ~NonTrivial5(); // virtual destructors are not trivial
};
```

##### standard-layout type
The standard-layout type is special because it has the same memory layout of the equivalent C struct or union.

A standard-layout class is a class that:
- has no non-static data members of type non-standard-layout class (or array of such types) or reference,
- has no virtual functions and no virtual base classes,
- has the same access control for all non-static data members,
- has no non-standard-layout base classes,
- either has no non-static data members in the most derived class and at most one base class with non-static data members, or has no base classes with non-static data members, 
- has no base classes of the same type as the first non-static data member.

A standard-layout struct is a standard-layout class defined with the class-key struct or the class-key class.<br>
A standard-layout union is a standard-layout class defined with the class-key union.

```
// empty classes have standard-layout
struct StandardLayout1 {};

struct StandardLayout2 {
    int x;
};

struct StandardLayout3 {
private: // both are private, so it's ok
    int x;
    int y;
};

struct StandardLayout4 : StandardLayout1 {
    int x;
    int y;

    void f(); // perfectly fine to have non-virtual functions
};

struct StandardLayout5 : StandardLayout1 {
    int x;
    StandardLayout1 y; // can have members of base type if they're not the first
};

struct StandardLayout6 : StandardLayout1, StandardLayout5 {
    // can use multiple inheritance as long only
    // one class in the hierarchy has non-static data members
};

struct StandardLayout7 {
    int x;
    int y;
    StandardLayout7(int x, int y) : x(x), y(y) {} // user-provided ctors are ok
};

struct StandardLayout8 {
public:
    StandardLayout8(int x) : x(x) {} // user-provided ctors are ok
// ok to have non-static data members and other members with different access
private:
    int x;
};

struct StandardLayout9 {
    int x;
    static NonStandardLayout1 y; // no restrictions on static members
};

struct NonStandardLayout1 {
    virtual f(); // cannot have virtual functions
};

struct NonStandardLayout2 {
    NonStandardLayout1 X; // has non-standard-layout member
};

struct NonStandardLayout3 : StandardLayout1 {
    StandardLayout1 x; // first member cannot be of the same type as base
};

struct NonStandardLayout4 : StandardLayout3 {
    int z; // more than one class has non-static data members
};

struct NonStandardLayout5 : NonStandardLayout3 {}; // has a non-standard-layout base class
```


##### POD
POD = trivial type + standard-layout type




```
struct A { int m; };
struct B { ~B(); int m; };
struct C { C() : m(){}; ~C(); int m; };
struct D { D(){}; int m; };
struct E { E() = default; int m;};  
struct F {F(); int m;};  F::F() = default;
```


