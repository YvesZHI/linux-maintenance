# Type deduction

### template ###
##### case 1 #####
```
template<typename T>
void f(T param);

int x = 27;
const int cx = x;
const int& rx = x;

f(x);  // T's and param's types are both int
f(cx); // T's and param's types are both int
f(rx); // T's and param's types are both int
f(27); // T's and param's types are both int
```
##### case 2 #####
```
template<typename T>
void f(T& param);

int x = 27;
const int cx = x;
const int& rx = x;

f(x);  // T is int,       param's type is int&
f(cx); // T is const int, param's type is const int&
f(rx); // T is const int, param's type is const int&
f(27); // ERROR
```
##### case 3 #####
```
template<typename T>
void f(const T& param);

int x = 27;
const int cx = x;
const int& rx = x;

f(x);  // T is int, param's type is const int&
f(cx); // T is int, param's type is const int&
f(rx); // T is int, param's type is const int&
f(27); // T is int, param's type is const int&
```
##### case 4 #####
```
template<typename T>
void f(T&& param);

int x = 27;
const int cx = x;
const int& rx = x;

f(x);   // x is lvalue,  so T is int&,       param's type is int&
f(cx);  // cx is lvalue, so T is const int&, param's type is const int&
f(rx);  // rx is lvalue, so T is const int&, param's type is const int&
f(27);  // 27 is rvalue, so T is int,        param's type is int&&
```

### auto ###
`auto` type deduction is almost the same as template type deduction.<br>

### ATTENTION ###
Type deduction can't work with implicit direct-list-initialization.<br>
##### template #####
```
// f is the template function as above
f(std::initializer_list{3}); // OK
f({3});                      // ERROR
```
##### auto #####
```
auto x1 = {3}; // x1 is std::initializer_list<int>
auto x2 = std::initializer_list{1, 2}; // x2 is std::initializer_list<int>
auto x3{3};    // x3 is int --- EXCEPTION!
auto x4{1, 2}; // ERROR
auto x5{std::initializer_list{3}}; // ERROR
```
